package controller

import ()

//var users [...]User
//var meetings [...]Meeting
//var currentuser User

func init() {
	//users = ...
	//meetings = ...
	//currentuser = ...
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
	if !a {
		//todo: output error info
		return
	}
	b, err = isPasswordValid(password)
	if !b {
		//todo: output error info
		return
	}
	c, err = isEmailValid(email)
	if !c {
		//todo: output error info
		return
	}
	d, err = isTelNumValid(telphone)
	if !d {
		//todo: output error info
		return
	}
	init()
	for _, u := range users {
		if a == u.username {
			//todu: output error info
			return
		}
	}
	//todo: output successed
	//write to file
}

func isUserNameValid(userName string) (bool, error) {
	m, err := regexp.MatchString("^[a-zA-Z]{4-30}$", userName)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isPasswordVaild(password string) (bool, error) {
	m, err := regexp.MatchString("^[0-9a-zA-Z@.]{6-30}$", password)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isEmailVaild(email string) (bool, error) {
	m, err := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isTelNumberVaild(telNum string) (bool, error) {
	m, err := regexp.MatchString("^[0-9]{11}$", telNum)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func Login(username, password string) {
	//todo:
}
