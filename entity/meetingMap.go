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

func (meetingmap Meetingmap) AddParticipants(title string, a_participants Usermap) bool {
	_, ok := meetingmap[title]
	if !ok {
		return false
	}
	return true
}
