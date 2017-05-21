package models

type Task struct {
	Name           string
	Number         int
	Description    string
	DefaultInput   string
	ExpectedOutput string
}

var (
	Task1 *Task = &Task{
		Name:           "task 1",
		Description:    "Task 1",
		DefaultInput:   "1e2\nf\n1e-3",
		ExpectedOutput: "True\nFalse\nTrue",
	}
)
