package model

// Transactor to add a task to storage.
func (t *Task) AddTask() error {

	// Retrieve the model from the JSON file.
	li, err := RetrieveListInfo()
	if err != nil {
		return err
	}

	// Append the new task to the task list.
	li.TasksList = append(li.TasksList, *t)

	// Commit changes.
	return CommitListInfo(li)
}

// Transactor to delete task from storage.
func (t *Task) DeleteTask() error {
	// Retrieve the model from the JSON file.
	li, err := RetrieveListInfo()
	if err != nil {
		return err
	}

	var newTaskList []Task 
	for _, task := range li.TasksList {
		if (task.ID != t.ID) {
			newTaskList = append(newTaskList, task)
		}
	}
	li.TasksList = newTaskList

	// Commit changes.
	return CommitListInfo(li)

}