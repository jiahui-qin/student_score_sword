package models

import "time"

type Exam struct {
	Name      string
	Id        int
	TeacherId int
	ClassId   int
	ExamTime  time.Time
	Desc      string
	FullScore int
}
