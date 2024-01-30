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