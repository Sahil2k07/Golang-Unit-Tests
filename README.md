# Golang-Unit-Tests

This project demonstrates how to write **unit tests in Golang** using the `stretchr/testify` package. It simulates a simple blog-like structure with users and posts, and tests the application logic in isolation using **mocks**.

## 🛠️ Technologies Used

- **Golang** – The main programming language
- **Echo** – A fast, minimalistic web framework for building REST APIs
- **GORM** – An ORM library for Golang to interact with MySQL
- **MySQL** – The relational database used for persistent storage
- **Testify** – Provides easy-to-use assertion and mocking capabilities for unit tests

## ✅ Unit Testing

We use the `github.com/stretchr/testify` package to write clean and expressive unit tests.

### Key features:

- Table-driven tests
- Mocking interfaces (e.g., repository methods)
- Assertions like `require.NoError()`, `require.Len()`, etc.

### Example Test

```go
func TestGetUserPosts(t *testing.T) {
    mockRepo := new(mocks.PostRepoMock)
    mockRepo.On("GetUserPosts", "test@example.com").Return([]models.Post{{Title: "Post 1"}}, nil)

    posts, err := mockRepo.GetUserPosts("test@example.com")

    require.NoError(t, err)
    require.Len(t, posts, 1)

    mockRepo.AssertExpectations(t)
}


```

## 🧪 Running Tests

Make sure dependencies are installed:

```bash
go mod tidy
go mod vendor
```

```bash
go test ./...
```

## 🧰 Setting Up MySQL (For Real Integration)

- Create a database for the project

  ```sql
  CREATE DATABASE GOLANG_UNIT_TESTS;
  ```

- Update the `DSN` string for right user-name and password in `configs/gorm.go`

  ```go
  func InitDatabase() {
      db, err := gorm.Open(mysql.New(mysql.Config{
          DSN: "root:shahil@tcp(127.0.0.1:3306)/GOLANG_UNIT_TESTS?charset=utf8&parseTime=True&loc=Local",
      }), &gorm.Config{})
      if err != nil {
          panic("Error connecting to the Database")
      }

      err = db.AutoMigrate(&models.User{}, &models.Post{})
      if err != nil {
          log.Fatal("Error auto migrating schemas", err)
      }

      DB = db
  }
  ```

## 🚀 Running the App

```bash
go run .
```
