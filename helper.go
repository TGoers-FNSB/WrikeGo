package wrikego

func NewTasksFromJSON(data []byte) (*Tasks, error) {
	var tasks Tasks
	err := json.Unmarshal(data, &tasks)
	return &tasks, err
}