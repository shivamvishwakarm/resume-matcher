package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/qdrant/go-client/qdrant"
	"github.com/shivamvishwakarm/resume-matcher/internal/config"
	"github.com/shivamvishwakarm/resume-matcher/internal/controler"
	"github.com/shivamvishwakarm/resume-matcher/internal/middleware"
	"github.com/shivamvishwakarm/resume-matcher/internal/models"
	"github.com/shivamvishwakarm/resume-matcher/internal/services"
)

func main() {

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10 MB limit for file uploads
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/register", createuser)
	v1.Post("/login", login)
	v1.Post("/upload-resume", middleware.Auth(), uploadResume)
	v1.Get("/me", middleware.Auth(), controler.GetUser())
	v1.Post("/match-resume", middleware.Auth(), getBestResume)
	v1.Get("/resumes", middleware.Auth(), controler.GetAllResume)

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("health")
	})

	log.Println("Server starting on port:8080")
	log.Fatal(app.Listen(":8080"))
}

func login(c *fiber.Ctx) error {

	var user controler.LoginReq

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := controler.LoginUser(user)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON((fiber.Map{
			"error": err.Error(),
		}))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    result.Token,
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(result)
}

func createuser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	createdUser, err := controler.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdUser.Password = ""

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func uploadResume(c *fiber.Ctx) error {

	// Get the logged-in user's email
	userClaims, _ := c.Locals("user").(map[string]interface{})
	userEmail, _ := userClaims["email"].(string)

	// 1. Get the PDF file from the multipart form
	file, err := c.FormFile("resume")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded. Use form field 'resume'",
		})
	}

	// Validate that it's a PDF
	if filepath.Ext(file.Filename) != ".pdf" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Only PDF files are accepted",
		})
	}

	// 2. Save the PDF to the server
	uploadsDir := "./uploads"
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create uploads directory",
		})
	}

	uniqueFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	savePath := filepath.Join(uploadsDir, uniqueFilename)

	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	// Read the file bytes for PDF text extraction
	pdfBytes, err := os.ReadFile(savePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read saved file",
		})
	}

	// 3. Extract text from the PDF
	text, err := services.ExtractTextFromPDF(pdfBytes)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to extract text from PDF: %v", err),
		})
	}

	// 4. Generate embedding for the extracted text
	embedding, err := services.GenerateEmbedding(text)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to generate embedding: %v", err),
		})
	}

	// 5. Store the embedding in Qdrant vector DB
	pointID := uuid.New().String()

	_, err = config.QdrantClient.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: config.ResumeCollectionName,
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewID(pointID),
				Vectors: qdrant.NewVectors(embedding...),
				Payload: qdrant.NewValueMap(map[string]any{
					"user_email":  userEmail,
					"filename":    file.Filename,
					"filepath":    savePath,
					"uploaded_at": time.Now().Format(time.RFC3339),
					"text":        text,
				}),
			},
		},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to store embedding in vector DB: %v", err),
		})
	}

	// 6. Save resume metadata to MongoDB
	resumeDoc := models.Resume{
		UserEmail:  userEmail,
		Filename:   file.Filename,
		Filepath:   savePath,
		QdrantID:   pointID,
		UploadedAt: time.Now(),
	}

	coll := config.Db.Collection("resumes")
	_, err = coll.InsertOne(context.TODO(), resumeDoc)
	if err != nil {
		log.Println("Error saving resume metadata to MongoDB:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save resume metadata",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Resume uploaded and processed successfully",
		"point_id": pointID,
		"filename": file.Filename,
	})
}

type matchRequest struct {
	JobDescription string `json:"job_description"`
}

func getBestResume(c *fiber.Ctx) error {

	// Get the logged-in user's email
	userClaims, _ := c.Locals("user").(map[string]interface{})
	userEmail, _ := userClaims["email"].(string)

	// 1. Parse the job description from the request body
	var req matchRequest
	if err := c.BodyParser(&req); err != nil || req.JobDescription == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "job_description is required in request body",
		})
	}

	// 2. Generate embedding for the job description
	embedding, err := services.GenerateEmbedding(req.JobDescription)
	if err != nil {
		log.Println("Error generating embedding for JD:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate embedding for job description",
		})
	}

	// 3. Search Qdrant for the most similar resumes — only this user's resumes
	limit := uint64(10)
	searchResult, err := config.QdrantClient.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: config.ResumeCollectionName,
		Query:          qdrant.NewQuery(embedding...),
		Filter: &qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("user_email", userEmail),
			},
		},
		WithPayload: qdrant.NewWithPayload(true),
		Limit:       &limit,
	})
	if err != nil {
		log.Println("Error querying Qdrant:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search resumes",
		})
	}

	// 4. Return the best matching resumes with match percentage
	results := make([]fiber.Map, 0, len(searchResult))
	for _, point := range searchResult {
		matchPercent := point.Score * 100 // cosine similarity as percentage

		filename := ""
		if fn, ok := point.Payload["filename"]; ok {
			filename = fn.GetStringValue()
		}

		results = append(results, fiber.Map{
			"id":            point.Id.GetUuid(),
			"filename":      filename,
			"match_percent": fmt.Sprintf("%.2f%%", matchPercent),
			"score":         point.Score,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_results": len(results),
		"results":       results,
	})
}
