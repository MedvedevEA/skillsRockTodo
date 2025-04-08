package repository

type DtoCreateTaskReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
type DtoUpdateTaskReq struct {
	Id          int     `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}
