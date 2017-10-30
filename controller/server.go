package controller

import (
	"fmt"
	//"log"
	. "agenda/entity"
	//fileio "agenda/fileio"
	"regexp"
)

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
	update()
	return
}

func Login(username, password string) {
	initialization()
	//check if current user alright exist
	//if exist then suggest logout
	if currentUser.Username != "NULL" {
		fmt.Println("Login failed! Error : already Logined. Please logout first")
		return
	} else {
		//if not do follows
		user := users.QueryUser(username)
		if user != nil {
			if user.Password != password {
				fmt.Println("Login failed : wrong password!")
			} else {
				//change the current user and write to file
				currentUser = user
				fmt.Println("Login!")
			}
		} else {
			fmt.Println("Login failed : wrong user!")
		}
	}
	update()
	return
}

func Logout() {
	initialization()
	currentUser = NULLUSER
	update()
	return
}

func ListUser() {
	initialization()
	if currentUser.Username != "NULL" {
		for user := range users {
			fmt.Println(user)
		}
	} else {
		fmt.Println("Please login first!")
	}
}

func DeleteUser() {

	initialization()
	update()
	return
}

func CreateMeeting(title string, participators []string, starttime string, endtime string) {

	initialization()
	update()
	return
}

func ModifyMeeting(title string, addedparticipators []string, deletedparticipators []string) {

	initialization()
	update()
	return
}

func QueryMeeting(starttime string, endtime string) {

	initialization()
	update()
	return
}

func QuitMeeting(title string) {
	initialization()
	update()
	return
}

func CancelMeeting(title string) {

	initialization()
	update()
	return
}

func ClearMeeting() {

	initialization()
	update()
	return
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
