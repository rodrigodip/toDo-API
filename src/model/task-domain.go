package model

type TaskData struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

func (td *TaskData) SetId(id int) {

	td.ID = id
}

func (td *TaskData) GetId() int {
	return td.ID
}

func (td *TaskData) SetTitle(title string) {

	td.Title = title
}

func (td *TaskData) SetDescription(description string) {
	td.Description = description
}

func (td *TaskData) SetCompleted(completed bool) {
	td.Completed = completed
}
