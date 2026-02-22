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
