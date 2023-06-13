package todo_webapp

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Done        bool   `json:"done"`
}
