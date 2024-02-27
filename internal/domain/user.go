package domain

type User struct {
	ID             string `json:"id,omitempty"`
	Username       string `json:"username,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	DeviceID       string `json:"deviceId,omitempty"`
	Password       string `json:"password,omitempty"`
	Email          string `json:"email,omitempty"`
	Avatar         string `json:"avatar,omitempty"`
	AppToken       string `json:"appToken,omitempty"`
	WebToken       string `json:"webToken,omitempty"`
	RegisterTime   int64  `json:"registerTime,omitempty"`
	RegisterIP     string `json:"registerIp,omitempty"`
	Lang           string `json:"lang,omitempty"`
	LoginTime      int64  `json:"loginTime,omitempty"`
	IsPro          bool   `json:"isPro,omitempty"`
	ProType        int    `json:"proType,omitempty"`
	ProIsAutoRenew bool   `json:"oroIsAutoRenew,omitempty"`
	ProStart       int64  `json:"proStart,omitempty"`
	ProEnd         int64  `json:"proEnd,omitempty"`
	CreateTime     int64  `json:"createTime,omitempty"`
	UpdateTime     int64  `json:"updateTime,omitempty"`
	DeleteTime     int64  `json:"deleteTime,omitempty"`
}

type Users []User
