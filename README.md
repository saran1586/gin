# Student API with Gin

First steps with the **Gin** web framework in Go: small programs that each expose part of a
student CRUD API, backed by **GORM** on **PostgreSQL**.

| File | What it shows |
|---|---|
| `student/student/create.go` | `POST` a student and save it |
| `student/student/insert.go` | `POST /students` and `GET /students` |
| `student/student/update.go` | `PUT /students/:id` |
| `student/student/delete.go` | `DELETE /students/:id` |

## Run

Each file is a separate program, so run one at a time:

```sh
cd student
export DATABASE_URL="host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
go run ./student/insert.go
```

The server starts on `:8080`.
