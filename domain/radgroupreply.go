package domain

type Radgroupreply struct {
	ID        uint   `json:"id"`
	Groupname string `json:"groupname" validate:"required"`
	Attribute string `json:"attribute" validate:"required"`
	OP        string `json:"op" validate:"required"`
	Value     string `json:"value" validate:"required"`
}

func (Radgroupreply) TableName() string {
	return "radgroupreply"
}
