package controller

import (
	"fmt"
	//"log"
	. "agenda/entity"
	"regexp"
)

var users Usermap

func writeToFile() {
	//write users
	//write meetings
	//write currentuser
}

func Register(username, password, email, telphone string) {
	var a, b, c, d bool
	var err error

	a, err = isUserNameValid(username)
	if false == a {
		fmt.Println("username fail", err)
		return
	}
	b, err = isPasswordValid(password)
	if false == b {
		fmt.Println("password fail", err)
		return
	}
	c, err = isEmailValid(email)
	if false == c {
		fmt.Println("email fail", err)
		return
	}
	d, err = isTelNumberValid(telphone)
	if false == d {
		fmt.Println("telphone fail", err)
		return
	}
	initialization()
	if users.AddUser(NewUser(username, password, email, telphone)) {
		fmt.Println(username, password, email, telphone, "register successed!")
	} else {
		fmt.Println(username, password, email, telphone, "register failed!")
	}
	return
}

func initialization() bool {
	users = make(Usermap)
	//meetings = ...
	//currentuser = ...
	return true
}

func isUserNameValid(username string) (bool, error) {
	m, err := regexp.MatchString("^[a-zA-Z]{4,30}$", username)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isPasswordValid(password string) (bool, error) {
	m, err := regexp.MatchString("^[0-9a-zA-Z@.]{6,30}$", password)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isEmailValid(email string) (bool, error) {
	m, err := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isTelNumberValid(telNum string) (bool, error) {
	m, err := regexp.MatchString("^[0-9]{11}$", telNum)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func Login(username, password string) {
	//todo: user map find this user

}

func Logout() {

}

func ListUser() {

}

func DeleteUser() {

}

func CreateMeeting(title string, participators []string, starttime string, endtime string) {

}

func ModifyMeeting(title string, addedparticipators []string, deletedparticipators []string) {

}

func QueryMeeting(starttime string, endtime string) {

}

func QuitMeeting(title string) {
}

func CancelMeeting(title string) {

}

func ClearMeeting() {

}
