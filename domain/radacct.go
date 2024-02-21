package domain

import "time"

type Radacct struct {
	Radacctid          uint      `json:"radacctid"`
	Acctsessionid      string    `json:"acctsessionid"`
	Acctuniqueid       string    `json:"acctuniqueid"`
	Username           string    `json:"username"`
	Realm              string    `json:"realm"`
	Nasipaddress       string    `json:"nasipaddress"`
	Nasportid          string    `json:"nasportid"`
	Nasporttype        string    `json:"nasporttype"`
	Acctstarttime      time.Time `json:"acctstarttime"`
	Acctupdatetime     time.Time `json:"acctupdatetime"`
	Acctstoptime       time.Time `json:"acctstoptime"`
	Acctinterval       uint      `json:"acctinterval"`
	Acctsessiontime    uint      `json:"acctsessiontime"`
	Acctauthentic      string    `json:"acctauthentic"`
	ConnectinfoStart   string    `json:"connectinfo_start"`
	ConnectinfoStop    string    `json:"connectinfo_stop"`
	Acctinputoctets    uint      `json:"acctinputoctets"`
	Acctoutputoctets   uint      `json:"acctoutputoctets"`
	Calledstationid    string    `json:"calledstationid"`
	Callingstationid   string    `json:"callingstationid"`
	Acctterminatecause string    `json:"acctterminatecause"`
	Servicetype        string    `json:"servicetype"`
	Framedprotocol     string    `json:"framedprotocol"`
	Framedipaddress    string    `json:"framedipaddress"`
	Secret             string    `json:"secret"`
}

func (Radacct) TableName() string {
	return "radacct"
}

type RadacctRepository interface {
	FirstUsername(param string) (res Radacct, err error)
}
