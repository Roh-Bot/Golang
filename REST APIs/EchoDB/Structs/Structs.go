package Structs

type User struct {
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Fname    string `json:"fname" validate:"required,min=4"`
	Lname    string `json:"lname" validate:"required,min=4"`
	Password string `json:"password"`
}

type ErrorStruct struct {
	Error map[string]string
}
