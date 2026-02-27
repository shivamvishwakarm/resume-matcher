- login
  - email , password
    return JWT - store in cookie, refresh token and access token

- Register
  - name, email, password, confirm password

- upload resume
- verity JWT
- PDF only
- store in server then create embeding and store the embed in the vector db

get resume based on JD:

    - Score of maching the resume ( how well fit to that JD)
    -

approach:

    - JD embeding then we'll match against the resume embeding of the user

run docker

```
docker run -d -p 27017:27017 \
 -e MONGO_INITDB_ROOT_USERNAME=admin \
 -e MONGO_INITDB_ROOT_PASSWORD=password \
 --name mongodb \
 mongodb/mongodb-community-server:latest
```

qudrant db:

docker run -p 6333:6333 -p 6334:6334 \
 -v "$(pwd)/qdrant_storage:/qdrant/storage:z" \
 qdrant/qdrant

using this as a embeding model:
https://ollama.com/library/nomic-embed-text
