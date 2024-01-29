package model

// Model structure, consists of a list of all tasks, and the index of the selected one.
type Model struct {
	TasksList []Task `json:"tasksList"`
	Selected  int    `json:"selected"`
}

// Task structure, consists of a title and a boolean that describes its state.
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
