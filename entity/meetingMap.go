package entity

type Meeting struct {
	Host         *User
	Title        string
	Participants Usermap
	Start        string
	End          string
}

//创建之前要检查合法性，包括host，participants是否为已注册用户
func NewMeeting(n_title, n_start, n_end string, n_host *User, n_participants Usermap) *Meeting {
	n_meeting := Meeting{
		Host:         n_host,
		Title:        n_title,
		Start:        n_start,
		End:          n_end,
		Participants: n_participants,
	}
	return &n_meeting
}

type Meetingmap map[string]*Meeting

func (meetingmap Meetingmap) AddMeeting(meeting *Meeting) bool {
	_, ok := meetingmap[meeting.Title]
	if ok {
		return false
	}
	meetingmap[meeting.Title] = meeting
	return true
}

//添加参与者，调用之前需检测参与者是否是注册用户
func (meetingmap Meetingmap) AddParticipants(title string, a_participants Usermap) bool {
	//检测会议是否存在
	_, ok := meetingmap[title]
	if !ok {
		return false
	}
	//遍历a_participants，检测当前会议的参与者是否已经有该人，有则不用重复添加，没有则添加
	for name, a_user := range a_participants {
		_, exist := meetingmap[title].Participants[name]
		if exist {
			return false
		} else {
			meetingmap[title].Participants[name] = a_user
		}
	}
	return true
}

//删除参与者，调用之前需检测参与者是否是注册用户
func (meetingmap Meetingmap) DeleteParticipants(title string, d_participants Usermap) bool {
	//检测会议是否存在
	_, ok := meetingmap[title]
	if !ok {
		return false
	}
	//如果要删除的参与者不在会议中，则返回错误，否则删除该会议者
	for name, _ := range d_participants {
		_, exist := meetingmap[title].Participants[name]
		if !exist {
			return false
		} else {
			delete(meetingmap[title].Participants, name)
			if len(meetingmap[title].Participants) == 0 {
				delete(meetingmap, title)
			}
		}
	}
	return true
}

//返回某一用户所发起的会议
// func (meetingmap Meetingmap) MeetingsHosted(host *User) Meetingmap {
// 	hostmeetings := make(Meetingmap)
// 	for title, meeting := range meetingmap {
// 		if meeting.Host == host {
// 			hostmeetings[title] = meeting
// 		}
// 	}
// 	return hostmeetings
// }

//返回某一用户作为参与者的会议
// func (meetingmap Meetingmap) MeetingParticipated(participant *User) Meetingmap {
// 	username := participant.Username
// 	participatemeetings := make(Meetingmap)
// 	for title, meeting := range meetingmap {
// 		_, exist := meetingmap[title].Participants[username]
// 		if exist {
// 			participatemeetings[title] = meeting
// 		}
// 	}
// 	return participatemeetings
// }

// func (meetingmap Meetingmap) QueryMeeting(start, end string) Meetingmap {
//
// }

//取消用户发起的某一个会议
func (meetingmap Meetingmap) CancelMeeting(title string, host string) bool {
	meeting, ok := meetingmap[title]
	if !ok {
		return false
	}
	if meeting.Host.Username == host {
		delete(meetingmap, title)
		return true
	}
	return false
}

//退出用户参与的某一个会议
func (meetingmap Meetingmap) QuitMeeting(title string, participant string) bool {
	meeting, ok := meetingmap[title]
	if !ok {
		return false
	}
	_, exist := meeting.Participants[participant]
	if exist {
		delete(meetingmap[title].Participants, participant)
		if len(meetingmap[title].Participants) == 0 {
			delete(meetingmap, title)
		}
		return true
	}
	return false
}

//清空用户发起的所有会议
func (meetingmap Meetingmap) ClearMeeting(host string) bool {
	for title, meeting := range meetingmap {
		if meeting.Host.Username == host {
			delete(meetingmap, title)
		}
	}
	return true
}
