package model

// Transactor to add a task to storage.
func (t *Task) AddTask() error {

	// Retrieve the model from the JSON file.
	m, err := RetrieveModel()
	if err != nil {
		return err
	}

	// Append the new task to the task list.
	m.TasksList = append(m.TasksList, *t)

	// Commit changes.
	return CommitModel(m)
}