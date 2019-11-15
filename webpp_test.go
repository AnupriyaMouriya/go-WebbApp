package main

import "testing"


func TestValidateName(t *testing.T) {
	name:=ValidateName("anu")
	if name==false{t.Errorf("invalid name")}
}


func TestValidateEmail(t *testing.T) {
	email:=ValidateEmail("anu@gmail.com")
	if email==false{t.Errorf("invalid email")}
}


