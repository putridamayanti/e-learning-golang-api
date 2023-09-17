package models

import "time"

const (
	AdminRole	= iota
	StudentRole
	InstructorRole
)

type User struct {
	Id 			string			`json:"id"`
	Name 		string			`json:"name"`
	Email 		string			`json:"email"`
	Password 	string			`json:"password"`
	Role		int 			`json:"role"`
	CreatedAt	time.Time		`json:"createdAt"`
}

type Course struct {
	Id 				string			`json:"id"`
	Title 			string			`json:"title"`
	InstructorId	string			`json:"instructorId"`
	Description		string			`json:"description"`
	Content			[]string		`json:"content"`
	EnrolledUsers	[]string		`json:"enrolledUsers"`
	CreatedAt		time.Time		`json:"createdAt"`
}

type Enrollment struct {
	Id 				string			`json:"id"`
	InstructorId	string			`json:"instructorId"` // User Instructor Id
	CourseId		string			`json:"courseId"`
	StudentId		string			`json:"studentId"` // User Student Id
	CreatedAt		time.Time		`json:"createdAt"`
}