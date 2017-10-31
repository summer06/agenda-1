# Agenda - Go Version

##### requirement：Implement a agenda CLI program
##### for more detail：
the cli-agenda.html in qq

##### Introduction: This program allow users to register, login, logout, list user, delete user, create meeting, modify meeting, query meeting, cancel meeting, quit meeting and clear meeting. Briefly, it is a meeting management system.
---
### Design
We use **top-down design method**, firstly design the commands which are shown to user directly. Then we design the functions for the commands to call. After that we design the structure to store users and meetings information which are needed by the processing functions. Finally we design the method for file IO, to make data permanent.

#### command design
Here we design 11 commands for users to use.

- user register

 `$ ./agenda register -u username -p password -e email -t teltephone`
- user login

  `$ ./agenda login -u username -p password`
- user logout

 `$ ./agenda logout`
- list all user

  `$ ./agenda listUser`
- delete user

  `$ ./agenda deleteUser`
- create meeting

  `$ ./agenda createMeeting -t title -p 'participant1 participant2 ...' -s startTime -e endTime`
- modify meeting's participants(-a for add, -d for delete)

  `$ ./agenda modifyMeeting -a 'participants1 participant2 ...' -d 'participant1 participant2 ...'`
- query meetings

  `$ ./agenda queryMeeting -s startTime -e endTime`
- cancel meeting

  `$ ./agenda cancelMeeting -t title`
- quit meeting

  `$ ./agenda quitMeeting -t title`
- clear all meetings

  `$ ./agenda clearMeeting`


#### Data structure
we use a User struct to record information of each user, and we use Meeting struct to record information of a meeting. Then we use a map named Usermap to collect each user and a Meetingmap to collect each meeting.
```
type User struct {
	Username  string
	Password  string
	Email     string
	Telephone string
}
type Usermap map[string]*User
```
```
type Meeting struct {
	Host         string
	Title        string
	Participants []string
	Start        string
	End          string
}
type Meetingmap map[string]*Meeting
```
---
### Implement Detail
#### cobra
Package cobra is a commander providing a simple interface to create powerful modern CLI interfaces. In addition to providing an interface, Cobra simultaneously provides a controller to organize your application code.

In this project, we use cobra to construct a framework, and some cobra functions really help.

for example, we use `cobra.Command` to create a command and use `RootCmd.AddCommand` to conveniently add the command to parent command.

here is one of our commands.

```
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "user register",
	Long: `register command is used to register a new user, you are required to offer
	username, password, email and telephone number.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		fmt.Println(username)
		password, _ := cmd.Flags().GetString("password")
		fmt.Println(password)
		email, _ := cmd.Flags().GetString("email")
		fmt.Println(email)
		telephone, _ := cmd.Flags().GetString("telephone")
		fmt.Println(telephone)

		controller.Register(username, password, email, telephone)
	},
}
```
It is easy to build command line program with cobra.
For more about cobra, see this https://godoc.org/github.com/spf13/cobra

#### json

we use the encode/json to encode and decode the json file to the struct

in our package fileio , we implement the file write/read and marshal/unmarshal function, and provide two func for controller to use it. for that these two func can be used universally, it just return the type []interface{}, witch is the slice of the universal type interface{}. interface{} means empty interface that have nothing need to implement, so any type implement the interface{}.

``` lang=golang
func ReadFile(filename string) ([]map[string]interface{}, error) {
    logsome(filename)
    if checkFileIsExist(filename) {
        bytes, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Println("ReadFile: ", err.Error())
            return nil, err
        }
        var xxx []map[string]interface{}
        if err := json.Unmarshal(bytes, &xxx); err != nil {
            fmt.Println("Unmarshal: ", err.Error())
            return nil, err
        }
        return xxx, nil
    } else {
        file, _ := os.Create(filename)
        defer file.Close()
        return nil, nil
    }
}


```
