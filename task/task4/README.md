# Self Blog Golang Backend Server

This is a golang backend server for self blog.

## Requirements
- Golang
  - gin `go get -u github.com/gin-gonic/gin`
  - gorm `go get -u gorm.io/gorm`
  - mysql `go get -u gorm.io/driver/mysql`
- Swagger
  - swaggo/swag `go install github.com/swaggo/swag/cmd/swag@latest`
  - swaggo/gin-swagger `go get -u github.com/swaggo/gin-swagger`
  - swaggo/files `go get -u github.com/swaggo/files`


## Project Structure

### facade
Facade layer. That's about the interface logic

### middleware
Middleware layer. For defining the gin middleware that's using at router.
Such as logging, auth, etc.

### dao
Data access layer. For initializing the database.

## How to run
1. Run `go run main.go`
2. Open `http://localhost:8080/swagger/index.html` to check the endpoints
3. Register an account via `POST /user/register`
4. Login the account via `POST /user/login` and get the token
5. Create a post via `POST /post/create` with the token
6. Read the post via `GET /post/read/:id`
7. Update the post via `PUT /post/update/`
8. List the posts via `GET /post/list`
9. Delete the post via `DELETE /post/delete?postId=id`
10. Create a comment via `POST /comment/create`
11. Get the comments via `GET /comment/list?postId=id`

