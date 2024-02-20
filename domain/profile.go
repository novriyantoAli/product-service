package domain

type Profile struct {
	Username  string `json:"username"`
	Groupname string `json:"groupname"`
	Priority  uint   `json:"priority"`
}

func (Profile) TableName() string {
	return "radusergroup"
}

type ProfileRepository interface {
	Find(profile *Profile) (res []Profile, err error)
	FindRadcheck(param *Radgroupcheck) (res []Radgroupcheck, err error)
	FindRadreply(param *Radgroupreply) (res []Radgroupreply, err error)
}

type ProfileUsecase interface {
	Find(profile *Profile) (res []Profile, err error)
	FindRadcheck(param *Radgroupcheck) (res []Radgroupcheck, err error)
	FindRadreply(param *Radgroupreply) (res []Radgroupreply, err error)
}
