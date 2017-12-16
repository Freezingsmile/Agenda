package entities

import (
	"time"
)

type MeetingInfo struct {
	Title string `xorm:"notnull pk"`
	Owned string
	Start time.Time
	End   time.Time
}

func NewMeetingInfo(m MeetingInfo) *MeetingInfo {
	if len(m.Title) == 0 {
		panic("MeetingTitle shold not null!")
	}
	return &m
}

type Participated struct {
	UserName string
	Title    string
}
