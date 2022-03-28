package model

import "github.com/lib/pq"

type User struct {
	Name     string     `json:"name"`
	Age      int        `json:"age"`
	Id       int        `json:"id"`
	Friends  []string 	`json:"friends"`
	SourceId int        `json:"source_id"`
	TargetId int        `json:"target_id"`
	NewAge   int        `json:"new_age"`
}

type UserGet struct {
	Name     string     `json:"name"`
	Age      int        `json:"age"`
	Id       int        `json:"id"`
	Friends  pq.StringArray 	`json:"friends"`
}

type UserUpdate struct {
	NewAge   int        `json:"new_age"`
}