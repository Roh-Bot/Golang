package Structs

type Users struct {
	User_code    string
	Display_name string `json:"display_name" validate:"required,min=3,max=14"`
	Email        string `json:"email" validate:"email"`
	Mobile       int    `json:"mobile" valdiate:"min=10"`
	Password     string `json:"password" validate:"min=8"`
	Mpin         int    `json:"mpin"`
}

type ErrorStruct struct {
	Error map[string]string
}

//func ValidateEmailOrMobile(fl validator.FieldLevel) bool {
//	email := fl.Field().String()
//
//	mobile := fl.Parent().Elem().FieldByName("mobile").Interface().(int)
//
//	// Check that either email or mobile is non-null, but not both
//	return (email != "" && mobile == 0) || (email == "" && mobile != 0)
//}
