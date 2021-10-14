# Blog API
Blog CRUD implementation using Golang and MongoDB

# Stack of the project:
- Golang with the standard library
- MongoDB
- MVC pattern
  <br>

# Installation

To start the project, simply run the following command inside the project's folder:

```sh
docker-compose up
```

You should be able to test the API at localhost:3000

## To-do API specification

| Task | URL | Method | Response code | Response |
|:----:|:---:|:------:|:-------------:|:--------:|
| Check API health | /| GET | 200 | API is up |
| Read all users | /user | GET | 200 | All users |
| Create a user | /user | POST | 201 | Created user |
| Get single user by id | /user/profile?id=value | GET | 200 | User belonging to that id |
| Update user by id | /user/profile?id=value | PUT | 204 | None |
| Delete user by id | /user/profile?id=value | DELETE | 204 | None |
| Get all posts from a user by its id | /user/posts?id=value | GET | 200 | Posts from user |
| Read all posts | /post | GET | 200 | All posts |
| Create a post | /post | POST | 201 | Created post |
| Read post by id | /post/filter?id=value | GET | 200 | Post belonging to that id |
| Update post by id | /post/filter?id=value | PUT | 204 | None |
| Delete post by id | /post/filter?id=value | DELETE | 204 | None |

## Create a new user
```
http://localhost:3000/user
```
Example body of the post:
```json
{
    "name":"John Doe",
    "email":"jdoe@gmail.com",
    "password":"Golang.21"
}
```
Example response:
```json
{
    "InsertedID": "e5866552-9ae3-4afb-95c2-b95b5c1b2b97"
}
```

## Create a new post
```
http://localhost:3000/post
```
Example body of the post:
```json
{
    "Content":"test post",
    "Title":"this is a title",
    "author": "e5866552-9ae3-4afb-95c2-b95b5c1b2b97"
}
```
Example response:
```json
{
    "InsertedID": "9f12030a-ea35-46d3-a893-991faaff9378"
}
```
