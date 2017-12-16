package entities

type MeetingInfoAtomicService struct{}
type ParticipatedAtomicService struct{}

var MeetingInfoService = MeetingInfoAtomicService{}
var ParticipatedService = ParticipatedAtomicService{}

func (*MeetingInfoAtomicService) Save(m *MeetingInfo) error {
	_, err := engine.Insert(m)
	return err
}

func (*MeetingInfoAtomicService) GetAllMeetings() []MeetingInfo {
	meetings := make([]MeetingInfo, 0)
	err := engine.Find(&meetings)
	checkErr(err)
	return meetings
}

func (*MeetingInfoAtomicService) GetUserAllMeetings(owned string) []MeetingInfo {
	meetings := make([]MeetingInfo, 0)
	err := engine.Where("Owned = ?", owned).Find(&meetings)
	checkErr(err)
	return meetings
}

func (*MeetingInfoAtomicService) GetMeetingByTitle(title string) *MeetingInfo {
	m := new(MeetingInfo) // m is a pointer,notice
	m.Title = title
	has, err := engine.Get(m)
	checkErr(err)
	if !has {
		m.Title = ""
	}
	return m
}
func (*ParticipatedAtomicService) Save(p *Participated) error {
	_, err := engine.Insert(p)
	return err
}

func (*ParticipatedAtomicService) GetUserAllParticipated(username string) []Participated {
	ps := make([]Participated, 0)
	err := engine.Where("UserName = ?", username).Find(&ps)
	checkErr(err)
	return ps
}
func (*ParticipatedAtomicService) GetUserParticipated(p *Participated) bool {
	has, err := engine.Get(p)
	checkErr(err)
	return has
}
