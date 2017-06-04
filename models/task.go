package models

type (
	Task struct {
		Id             int	`json:"id"`
		Name           string	`json:"name"`
		Description    string	`json:"description"`
		DefaultInput   string	`json:"defaultInput"`
		ExpectedOutput string	`json:"expectedOutput"`
	}
)

var (
	Tasks = []*Task{
		{
			Id:		123456,
			Name:           "task 1",
			Description:    "Task 1",
			DefaultInput:   "1e2\nf\n1e-3",
			ExpectedOutput: "True\nFalse\nTrue",
		},
	}
)
