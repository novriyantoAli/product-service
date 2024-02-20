package domain

type Radgroupcheck struct {
	ID        uint   `json:"id"`
	Groupname string `json:"groupname"`
	Attribute string `json:"attribute"`
	OP        string `json:"op"`
	Value     string `json:"value"`
}

func (Radgroupcheck) TableName() string {
	return "radgroupcheck"
}
