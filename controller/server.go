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

func isLogined() bool {
	if currentUser.Username != "NULL" {
		return true
	} else {
		return false
	}
}

func Login(username, password string) {
	initialization()
	//check if current user alright exist
	//if exist then suggest logout
	if isLogined() {
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
	if isLogined() {
		logout()
	} else {
		fmt.Println("Logout failed : no logined user!")
	}
	update()
	return
}

func logout() {
	currentUser = NULLUSER
}

func ListUser() {
	initialization()
	if isLogined() {
		for user := range users {
			fmt.Println(user)
		}
	} else {
		fmt.Println("Please login first!")
	}
}

func DeleteUser() {
	initialization()
	if isLogined() {
		users.DeleteUser(currentUser.Username)
		logout()
	} else {
		fmt.Println("delete failed! not login.")
	}
	update()
	return
}

func CreateMeeting(title string, participators []string, starttime string, endtime string) {
	initialization()
	if isLogined() {
		for _, s := range participators {
			if users.QueryUser(s) == nil {
				fmt.Println("Create Meeting failed! invalid user")
				return
			}
		}
		t, _ := isTimeValid(starttime)
		r, _ := isTimeValid(endtime)
		if t == false || r == false {
			fmt.Println("wrong time")
			return
		}
		if meetings.AddMeeting(NewMeeting(title, starttime, endtime, currentUser.Username, participators)) == false {
			fmt.Println("filed!")
			return
		}
		fmt.Println("Create Meeting successed!")
	} else {
		fmt.Println("Please login first!")
	}
	update()
	return
}

func ModifyMeeting(title string, addedparticipators []string, deletedparticipators []string) {
	initialization()
	if isLogined() {
		if len(addedparticipators) != 0 {
			for _, s := range addedparticipators {
				if users.QueryUser(s) == nil {
					fmt.Println("Modify Meeting failed! invalid user")
					return
				}
			}
			if meetings.AddParticipants(title, addedparticipators) == false {
				fmt.Println("Modify Meeting failed! invalid title or add user")
				return
			}
		}
		if len(deletedparticipators) != 0 {
			for _, s := range deletedparticipators {
				if users.QueryUser(s) == nil {
					fmt.Println("Modify Meeting failed! invalid user")
					return
				}
			}
			if meetings.DeleteParticipants(title, deletedparticipators) == false {
				fmt.Println("Modify Meeting failed! invalid title or delete user")
				return
			}
		}
		fmt.Println("Modify Meeting successed!")
	}
	update()
	return
}

func QueryMeeting(starttime string, endtime string) {
	initialization()

	if isLogined() {
		t, _ := isTimeValid(starttime)
		r, _ := isTimeValid(endtime)
		if t == false || r == false {
			fmt.Println("time wrong!")
		}
		meeting := meetings.QueryMeeting(starttime, endtime, currentUser.Username)
		for _, value := range meeting {
			fmt.Println(value)
		}
	}
	update()
	return
}

func QuitMeeting(title string) {
	initialization()
	if isLogined() {
		if meetings.QuitMeeting(title, currentUser.Username) {
			fmt.Println("quit successed!")
		} else {
			fmt.Println("title wrong or you aren't hostor!")
		}
	}
	update()
	return
}

func CancelMeeting(title string) {

	initialization()
	if isLogined() {
		if meetings.CancelMeeting(title, currentUser.Username) {
			fmt.Println("meeting cancle successed!")
		} else {
			fmt.Println("meeting title wrong or you aren't hostor!")
		}
	}
	update()
	return
}

func ClearMeeting() {

	initialization()
	if isLogined() {
		if meetings.ClearMeeting(currentUser.Username) {
			fmt.Println("clear meeting successed!")
		}
	}
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

func isTimeValid(time string) (bool, error) {
	m, err := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}\\s[0-9]{2}:[0-9]{2}:[0-9]{2}$", time)
	if m {
		return true, err
	} else {
		return false, err
	}
}
