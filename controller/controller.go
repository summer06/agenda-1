package controller

import (
	"fmt"
	//"log"
	. "agenda/entity"
	fileio "agenda/fileio"
	//"regexp"
)

var users Usermap
var currentUser *User
var meetings Meetingmap

var NULLUSER = NewUser("NULL", "", "", "")

//初始化所有的数据结构
func initialization() {
	users = make(Usermap)
	meetings = make(Meetingmap)
	readFromFile()
	//meetings ....
}

func readFromFile() {
	//read users
	t, _ := fileio.ReadFile("json/user.json")
	//fmt.Println("t: ", t)
	for _, u := range t {
		//fmt.Println("user:", u)
		// 这一句很长很丑陋的代码，里面的interface{}要断言才能使用，这里没有写断言的检查，可能引发panic

		users.AddUser(NewUser(u["Username"].(string), u["Password"].(string), u["Email"].(string), u["Telephone"].(string)))
	}
	//read current user
	d, _ := fileio.ReadFile("json/current.json")
	for _, c := range d {
		currentUser = NewUser(c["Username"].(string), c["Password"].(string), c["Email"].(string), c["Telephone"].(string))
	}

	// read meetings
	m, _ := fileio.ReadFile("json/meeting.json")
	for _, meeting := range m {
		fmt.Println("meeting:\n", meeting)

	}
}

func writeToFile() {
	//write users
	var alluser []User
	for _, v := range users {
		alluser = append(alluser, *v)
	}
	fileio.WriteFile("json/user.json", alluser)
	//write meetings
	var current []User
	current = append(current, *currentUser)
	//fmt.Println("###", current, "###")
	fileio.WriteFile("json/current.json", current)
	//write currentuser
}

func update() {
	writeToFile()
}
