package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Profile   Profile   `json:"profile"`
	Roles     []string  `json:"roles"`
}

type Profile struct {
	Bio      string            `json:"bio"`
	Avatar   string            `json:"avatar"`
	Location string            `json:"location"`
	Website  string            `json:"website"`
	Metadata map[string]string `json:"metadata"`
}

type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     string                 `json:"level"`
	Service   string                 `json:"service"`
	Message   string                 `json:"message"`
	Context   map[string]interface{} `json:"context"`
	TraceID   string                 `json:"trace_id"`
}
