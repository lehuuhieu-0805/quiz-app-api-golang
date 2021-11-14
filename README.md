# QUIZ APP RESTFUL API

#### It is a simple RESTful API with GO using:
- [gin framework](https://github.com/gin-gonic/gin)
- [gorm](https://gorm.io/) is ORM library for golang
- Database: using Sql Server


## Install
Download [golang](https://golang.org/)

Clone project and install dependencies
```
git clone https://github.com/lehuuhieu-1011/quiz-app-api-golang.git
cd quiz-app-api-golang
```

## Usage

### Create a file .env in root directory:
```
SERVER=server-name-of-sqlserver
PORT_DB=1433
USER=username-of-sqlserver
PASSWORD=password-of-sqlserver
DATABASE=name-of-database

PORT=8080
```

### Run the application
> go run main.go

`API Endpoint: http://localhost:8080`
 
## API

#### /api/course-quiz
- `GET`: Get all course
- `POST`: Add new course
#### /api/course-quiz/:id
- `GET`: Get a course
- `PUT`: Update a course
- `DELETE`: Delete a course

#### /api/question-quiz
- `GET`: Get all question of quiz
- `POST`: Add new question of quiz
#### /api/question-quiz/:id
- `GET`: Get a question of quiz
- `PUT`: Update a question of quiz
- `DELETE`: Delete a question of quiz