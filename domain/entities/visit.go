package entities

type Visit struct {
	Id        int    `json:"id"`
	Referrer  string `json:"referrer"`
	UserAgent string `json:"userAgent"`
	Ip        string `json:"ip"`
	LinkId    int    `json:"linkId"`
	CreatedAt string `json:"createdAt"`
}
