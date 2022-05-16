package models

type DiscordStat struct {
	Time          int64 `json:"time"`
	MemberCount   int   `json:"memberCount"`
	PresenceCount int   `json:"presenceCount"`
}
