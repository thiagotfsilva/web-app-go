package models

import "time"

type User struct {
	ID           uint64        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Nick         string        `json:"nick,omitempty"`
	Email        string        `json:"email,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}
