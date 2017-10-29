package entity

type meeting struct {
	host         *User
	title        string
	participants Usermap
	start        string
	end          string
}

//创建之前要检查合法性，包括host，participants是否为已注册用户
func NewMeeting(n_title, n_start, n_end string, n_host *User, n_participants Usermap) *meeting {
	n_meeting := meeting{
		host:         n_host,
		title:        n_title,
		start:        n_start,
		end:          n_end,
		participants: n_participants,
	}
	return &n_meeting
}
