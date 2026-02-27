package config

import (
	"context"
	"fmt"
	"log"

	"github.com/qdrant/go-client/qdrant"
)

const (
	QdrantHost           = "localhost"
	QdrantPort           = 6334
	ResumeCollectionName = "resumes"
	EmbeddingDimension   = 768
)

var QdrantClient *qdrant.Client

func init() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: QdrantHost,
		Port: QdrantPort,
	})
	if err != nil {
		log.Fatal("Failed to connect to Qdrant:", err)
	}

	QdrantClient = client

	// Ensure the "resumes" collection exists
	err = ensureCollection(client)
	if err != nil {
		log.Fatal("Failed to ensure Qdrant collection:", err)
	}

	log.Println("Qdrant db is running...")
}

func ensureCollection(client *qdrant.Client) error {
	ctx := context.Background()

	// Check if collection already exists
	collections, err := client.ListCollections(ctx)
	if err != nil {
		return fmt.Errorf("failed to list collections: %w", err)
	}

	for _, col := range collections {
		if col == ResumeCollectionName {
			log.Println("Qdrant collection 'resumes' already exists")
			return nil
		}
	}

	// Create the collection
	err = client.CreateCollection(ctx, &qdrant.CreateCollection{
		CollectionName: ResumeCollectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     uint64(EmbeddingDimension),
			Distance: qdrant.Distance_Cosine,
		}),
	})
	if err != nil {
		return fmt.Errorf("failed to create collection: %w", err)
	}

	log.Println("Created Qdrant collection 'resumes'")
	return nil
}
