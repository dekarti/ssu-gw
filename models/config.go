package models

type (
	Config struct {
		Command  []string `json:"command"`
		Language string   `json:"language"`
		Version  string   `json:"version"`
	}
)
