package entities

import (
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	u := &UserInfo{"test1", "passwd", "qq@qq", time.Now()}
	err := UserInfoService.Save(u)
	if err != nil {
		t.Error("Error happen when saving userinfo")
	}
	defer engine.Delete(&UserInfo{UserName: "test1"})
	check := UserInfoService.FindByUserName("test1")
	if check.UserName == "" {
		t.Error("not found it after saving")
	}
	t.Log("test1 pass")
}
func TestSaveMeeting(t *testing.T) {
	m := &MeetingInfo{"title1", "test1", time.Now(), time.Now().Add(time.Hour)}
	err := MeetingInfoService.Save(m)
	if err != nil {
		t.Error("Error happen when saving meetinginfo")
	}
	defer engine.Delete(&MeetingInfo{Owned: "test1"})
	check := MeetingInfoService.GetMeetingByTitle("title1")
	if check.Title != "title1" {
		t.Error("not found this meeting")
	}
	t.Log("test2 pass")
}

func TestSaveParticipated(t *testing.T) {
	p := &Participated{"test1", "title2"}
	err := ParticipatedService.Save(p)
	if err != nil {
		t.Error("Error happen when saving participated")
	}
	defer engine.Delete(p)
	check := new(Participated)
	check.Title = "title2"
	check.UserName = "test1"
	has, err2 := engine.Get(check)
	if err2 != nil {
		t.Error("mistake happened")
	}
	if !has {
		t.Error("not found it in participated table")
	}
	t.Log("test3 pass")
}

func TestGetMeetings(t *testing.T) {
	m1 := &MeetingInfo{"title3", "test3", time.Now(), time.Now().Add(time.Hour)}
	m2 := &MeetingInfo{"title4", "test3", time.Now(), time.Now().Add(time.Hour)}
	m3 := &MeetingInfo{"title5", "test3", time.Now(), time.Now().Add(time.Hour)}
	err := MeetingInfoService.Save(m1)
	if err != nil {
		t.Error("saveError")
	}
	err = MeetingInfoService.Save(m2)
	if err != nil {
		t.Error("saveError")
	}
	err = MeetingInfoService.Save(m3)
	if err != nil {
		t.Error("saveError")
	}
	mlist := MeetingInfoService.GetUserAllMeetings("test3")
	if len(mlist) != 3 {
		t.Error("GetUserAllMeeting Error")
	}
	defer engine.Delete(&MeetingInfo{Owned: "test3"})
	check := MeetingInfoService.GetMeetingByTitle("title4")
	if check.Owned != "test3" {
		t.Error("GetMeetingByTitle Error")
	}
}

func TestGetParticipated(t *testing.T) {
	p1 := &Participated{"test4", "title4"}
	p2 := &Participated{"test4", "title5"}
	err := ParticipatedService.Save(p1)
	if err != nil {
		t.Error("saveError")
	}
	err = ParticipatedService.Save(p2)
	if err != nil {
		t.Error("saveError")
	}
	defer engine.Delete(&Participated{UserName: "test4"})
	plist := ParticipatedService.GetUserAllParticipated("test4")
	check := ParticipatedService.GetUserParticipated(p1)
	if len(plist) != 2 {
		t.Error("GetUserAllParticipated Error")
	}
	if check == false {
		t.Error("GetUserParticipated Error")
	}
}
