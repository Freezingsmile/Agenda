package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(u)
	return err
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	everyone := make([]UserInfo, 0)
	err := engine.Find(&everyone)
	checkErr(err)
	return everyone
}

// FindByUserName .
func (*UserInfoAtomicService) FindByUserName(un string) *UserInfo {
	u := new(UserInfo)
	u.UserName = un
	has, err := engine.Get(u)
	checkErr(err)
	if !has {
		u.UserName = ""
	}
	return u
}
