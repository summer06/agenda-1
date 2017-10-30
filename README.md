# Agenda - Go Version

##### requirement：Implement a agenda CLI program
##### for more detail：
file:///Users/dengxiajun/Downloads/cli-agenda(1).html

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
