package dto

type AddTask struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type GetTasks struct {
}
type GetTask struct {
	Id int `validate:"required,gte=1"`
}
type UpdateTask struct {
	Id          int     `validate:"required,gte=1"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status" validate:"omitempty,oneof=new in_progress done"`
}
type RemoveTask struct {
	Id int `validate:"required,gte=1"`
}
