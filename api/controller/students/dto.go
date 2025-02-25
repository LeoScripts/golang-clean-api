package students

import "github.com/google/uuid"

type InputStudentDto struct {
	FullName string `json:"full_name" validate:"required|min_len:3|max_len:150|string"`
	Age      int    `json:"age" validate:"required|int|min:3|max:80"`
}

type OutputStudentDto struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age" `
}

type OutputStudentsDto struct {
	Students []OutputStudentDto `json:"students"`
}
