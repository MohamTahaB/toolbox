package model


func AddTask(t *Task) error {

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