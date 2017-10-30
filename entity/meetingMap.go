package entity

import "time"

type Meeting struct {
	Host         string
	Title        string
	Participants []string
	Start        string
	End          string
}

//创建之前要检查合法性，包括host，participants是否为已注册用户
func NewMeeting(n_title, n_start, n_end string, n_host string, n_participants []string) *Meeting {
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

//下面的方法接收者为map类型，不知道是否会出现没有对原对象赋值的情况，留意！

//添加会议
func (meetingmap Meetingmap) AddMeeting(meeting *Meeting) bool {
	_, ok := meetingmap[meeting.Title]
	if ok {
		return false
	}
	meetingmap[meeting.Title] = meeting
	return true
}

//添加参与者，调用之前需检测参与者是否是注册用户
func (meetingmap Meetingmap) AddParticipants(title string, a_participants []string) bool {
	//检测会议是否存在
	_, ok := meetingmap[title]
	if !ok {
		return false
	}
	//遍历a_participants，检测当前会议的参与者是否已经有该人，有则不用重复添加，没有则添加
	for _, name := range a_participants {
		for _, existname := range meetingmap[title].Participants {
			if name == existname {
				return false
			}
		}
		meetingmap[title].Participants = append(meetingmap[title].Participants, name)
	}
	return true
}

//删除参与者，调用之前需检测参与者是否是注册用户
func (meetingmap Meetingmap) DeleteParticipants(title string, d_participants []string) bool {
	//检测会议是否存在
	_, ok := meetingmap[title]
	if !ok {
		return false
	}
	//如果要删除的参与者不在会议中，则返回错误，否则删除该会议者
	for _, name := range d_participants {
		for index, existname := range meetingmap[title].Participants {
			if name == existname {
				meetingmap[title].Participants = append(meetingmap[title].Participants[:index], meetingmap[title].Participants[index+1:]...)
				if len(meetingmap[title].Participants) == 0 {
					delete(meetingmap, title)
					return true
				}
				break
			}
			if index == len(meetingmap[title].Participants)-1 {
				return false
			}
		}
	}
	return true
}

//返回某一用户所发起的会议
func (meetingmap Meetingmap) MeetingsHosted(host string) Meetingmap {
	hostmeetings := make(Meetingmap)
	for title, meeting := range meetingmap {
		if meeting.Host == host {
			hostmeetings[title] = meeting
		}
	}
	return hostmeetings
}

//返回某一用户作为参与者的会议
func (meetingmap Meetingmap) MeetingParticipated(participant string) Meetingmap {
	username := participant
	participatemeetings := make(Meetingmap)
	for title, meeting := range meetingmap {
		for _, existname := range meeting.Participants {
			if existname == username {
				participatemeetings[title] = meeting
				break
			}
		}
	}
	return participatemeetings
}

//查找某一时间段内的会议
//默认在调用前检查start和end的前后关系
//检查方法：用time.Parse转化成time格式后，用end_t.After(start_t)返回真假判断
func (meetingmap Meetingmap) QueryMeeting(start, end, username string) Meetingmap {
	resultMeeting := make(Meetingmap)
	start_t, _ := time.Parse("2006-01-02 15:04:05", start)
	end_t, _ := time.Parse("2006-01-02 15:04:05", end)
	meetingHost := meetingmap.MeetingsHosted(username)
	meetingParticipate := meetingmap.MeetingParticipated(username)
	for title, meeting := range meetingHost {
		s, _ := time.Parse("2006-01-02 15:04:05", meeting.Start)
		e, _ := time.Parse("2006-01-02 15:04:05", meeting.End)
		if (s == start_t || s.After(start_t)) && (e == end_t || end_t.After(e)) {
			resultMeeting[title] = meeting
		}
	}
	for title, meeting := range meetingParticipate {
		s, _ := time.Parse("2006-01-02 15:04:05", meeting.Start)
		e, _ := time.Parse("2006-01-02 15:04:05", meeting.End)
		if (s == start_t || s.After(start_t)) && (e == end_t || end_t.After(e)) {
			resultMeeting[title] = meeting
		}
	}
	return resultMeeting
}

//取消用户发起的某一个会议
func (meetingmap Meetingmap) CancelMeeting(title string, host string) bool {
	meeting, ok := meetingmap[title]
	if !ok {
		return false
	}
	if meeting.Host == host {
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
	for index, existname := range meeting.Participants {
		if participant == existname {
			meeting.Participants = append(meeting.Participants[:index], meeting.Participants[index+1:]...)
			if len(meeting.Participants) == 0 {
				delete(meetingmap, title)
				return true
			}
			return true
		}
	}
	return false
}

//清空用户发起的所有会议
func (meetingmap Meetingmap) ClearMeeting(host string) bool {
	for title, meeting := range meetingmap {
		if meeting.Host == host {
			delete(meetingmap, title)
		}
	}
	return true
}
