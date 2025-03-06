# Golang Clean API

API developed using good software development practices with the intention of making the concept of clean architecture visible.

With that in mind, I left the database in memory (to make it easier for users, especially beginners), simulating the flow of how it would be with a real database, whether PostgreSQL, MySQL, MongoDB, or any other you prefer, as the flow is quite similar.

Treat this project as a laboratory where you can modify and play around with it, haha. I hope it is very useful in your development.

Other important points are:

- The folder structure is not the most common one for Golang, but I really liked it, haha.
- Environment variables – I didn’t add them to simplify things (I have other projects with this configuration ready, and I’ll upload them here to GitHub soon).
- [Air](https://github.com/air-verse/air) (automatically restarts the server after modifications) – I’ve already configured it if you want to use it.
  - If you don’t have it installed, just run the command below in the terminal.

```bash
go install github.com/air-verse/air@latest
```
```diff
- Attention

+ For security reasons, occasionally use the command "go run".
```

## Techs

* [Golang](https://go.dev/) - Programing Lenguage
* [Gin](https://github.com/gin-gonic/gin) - Gin Web Framework
* [Clean Arch](https://dev.to/booscaaa/implementando-clean-architecture-com-golang-4n0a) - Clean Architecture
* [Solid](https://aprendagolang.com.br/o-que-e-solid/) - SOLID Principles

## **Endpoints**

- **GET** `/api/v1/heart` :  Route index
- **GET** `/api/v1/students` : Get all students
- **GET** `/api/v1/students/:id` : Get a student By ID
- **POST** `/api/v1/students` : Create a new student
- **PUT** `/api/v1/students/:id` : Update a student
- **DELETE** `/api/v1/students/:id` : Remove a student

## Running this project

1 **Clone the repository**

 ```bash
 git clone https://github.com/LeoScripts/golang-clean-api
 cd golang-clean-api
```

2 **Install dependencies**

```bash
go mod tidy
```

3 **Run the application**

```bash
go run main.go
# or
air
```

## Http Client collection

Import the file `_other/Insomnia_2025-03-05.json` into your HTTP client (Insomnia, Postman, etc......)
