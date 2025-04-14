package dto

type AddTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type UpdateTask struct {
	Id          int     `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}
