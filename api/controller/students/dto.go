package students

type InputStudentDto struct {
	FullName string `json:"full_name" validate:"required|min_len:3|max_len:150|string"`
	Age      int    `json:"age" validate:"required|int|min:3|max:80"`
}
