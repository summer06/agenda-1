package entity

type User struct {
	username  string
	password  string
	email     string
	telephone string
}

//用于外部可以构造新的user实例
func NewUser(u_name, p_word, e_mail, t_phone string) *User {
	newUser := User{
		username:  u_name,
		password:  p_word,
		email:     e_mail,
		telephone: t_phone,
	}
	return &newUser
}

//增删查操作默认合法性检测在controller里面做好了

//装user实例的map
type Usermap map[string]*User

//将一个user实例加入map
func (usermap Usermap) AddUser(user *User) bool {
	usermap[user.username] = user
	return true
}

//从map中删除一个user实例
func (usermap Usermap) DeleteUser(username string) bool {
	delete(usermap, username)
	return true
}

//在map中查找一个user实例
func (usermap Usermap) QueryUser(username string) *User {
	_, ok := usermap[username]
	if ok {
		return usermap[username]
	} else {
		return nil
	}
}

//todo 用文件读写初始化Usermap实例
