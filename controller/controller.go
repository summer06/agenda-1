package controller

import (
	"fmt"
	//"log"
	. "agenda/entity"
	fileio "agenda/fileio"
	"regexp"
)

var users Usermap

//初始化所有的数据结构
func initialization() {
	users = make(Usermap)
	readFromFile()
	//meetings ....
}

func readFromFile() {
	//read users
	t, _ := fileio.ReadFile("user.json")
	fmt.Println("t: ", t)
	for _, u := range t {
		//fmt.Println(u)
		// 这一句很长很丑陋的代码，里面的interface{}要断言才能使用，这里没有写断言的检查，可能引发panic
		users.AddUser(NewUser(u["Username"].(string), u["Password"].(string), u["Email"].(string), u["Telephone"].(string)))
	}
	//todo: read meetings and current user
}

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
	//todo : call the writetofile func
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

func Login(username, password string) {
	initialization()
	//todo : check if current user alright exist
	//if exist then suggest logout
	//if not do follows
	user := users.QueryUser(username)
	if user != nil {
		if user.Password != password {
			fmt.Println("Login failed : wrong password!")
		} else {
			// todo : change the current user and write to file
			fmt.Println("Login!")
		}
	} else {
		fmt.Println("Login failed : wrong user!")
	}
}

func Logout() {
	//todo: clear current user
}

func ListUser() {
	initialization()
	for user := range users {
		fmt.Println(user)
	}
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
