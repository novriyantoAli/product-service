package domain

type Radgroupreply struct {
	ID        uint   `json:"id"`
	Groupname string `json:"groupname"`
	Attribute string `json:"attribute"`
	OP        string `json:"op"`
	Value     string `json:"value"`
}

func (Radgroupreply) TableName() string {
	return "radgroupreply"
}
