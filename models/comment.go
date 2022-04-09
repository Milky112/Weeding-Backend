package models

import "time"

type PostData struct {
	Name       string    `json:"name"`
	Attendance int       `json:"attendance"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at`
}
