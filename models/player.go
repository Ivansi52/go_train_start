package models

type Player struct {
	Nickname string `json:"nickname"`
	Level    int    `json:"level"`
	Gold     int    `json:"gold"`
	Online   bool   `json:"online"`
}
