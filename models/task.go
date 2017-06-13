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
			Name:           "Task 1",
			Description:    "Determine if number is floating",
			DefaultInput:   "1e2\nf\n1e-3",
			ExpectedOutput: "True\nFalse\nTrue",
		},
		{
			Id:		123457,
			Name:           "Task 2",
			Description:    "Task 1",
			DefaultInput:   "1e2\nf\n1e-3",
			ExpectedOutput: "True\nFalse\nTrue",
		},
		{
			Id:		123458,
			Name:           "Task 3",
			Description:    "Task 1",
			DefaultInput:   "1e2\nf\n1e-3",
			ExpectedOutput: "True\nFalse\nTrue",
		},
	}
)
