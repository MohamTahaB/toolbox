package model

// Model structure, consists of a list of all tasks, and the index of the selected one.
type Model struct {
	TasksList []Task
	Selected  int
}

// Task structure, consists of a title and a boolean that describes its state.
type Task struct {
	Title string
	Done  bool
}
