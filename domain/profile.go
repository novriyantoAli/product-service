package domain

type Profile struct {
	Username  string `json:"username" validate:"required"`
	Groupname string `json:"groupname"`
	Priority  uint   `json:"priority" validate:"required"`
}

func (Profile) TableName() string {
	return "radusergroup"
}

type ProfileRepository interface {
	Find(profile *Profile) (res []Profile, err error)
	FindRadcheck(param *Radgroupcheck) (res []Radgroupcheck, err error)
	FindRadreply(param *Radgroupreply) (res []Radgroupreply, err error)
	Save(param *Profile) error
	SaveRadcheck(param *Radgroupcheck) error
	SaveRadreply(param *Radgroupreply) error
	DeleteRadcheck(id uint) error
	DeleteRadreply(id uint) error
	Delete(username string) error
}

type ProfileUsecase interface {
	Find(profile *Profile) (res []Profile, err error)
	FindRadcheck(param *Radgroupcheck) (res []Radgroupcheck, err error)
	FindRadreply(param *Radgroupreply) (res []Radgroupreply, err error)
	Save(param *Profile) error
	SaveRadcheck(param *Radgroupcheck) error
	SaveRadreply(param *Radgroupreply) error
	DeleteRadcheck(id uint) error
	DeleteRadreply(id uint) error
	Delete(username string) error
}
