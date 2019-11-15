package controller

import "regexp"

func ValidateName(name string)(bool){
	matched,_:=regexp.MatchString(`^[a-zA-Z ]+$`,name)
	return matched
}


func ValidateEmail(email string)(bool){
	match,_:=regexp.MatchString(`^.+@[a-zA-Z0-9-.]+.([a-zA-Z]{2,3}|[0-9]{1,3}$)`,email)
	return match
}
