package service

import (
	"cloudgo-data-orm/entities"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// login ,if info correct return a struct struct { CorrectPW bool}
func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["UserName"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		u := entities.UserInfoService.FindByUserName(req.Form["UserName"][0])
		if u.PassWord != req.Form["PassWord"][0] {
			formatter.JSON(w, http.StatusOK, struct {
				CorrectPW bool `json:"CorrectPW"`
			}{true})
		} else {
			formatter.JSON(w, http.StatusOK, struct {
				CorrectPW bool `json:"CorrectPW"`
			}{false})
		}
	}
}

//register, if success ,return { Success bool}
func regHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		u := entities.UserInfoService.FindByUserName(req.Form["UserName"][0])
		if u.UserName == "" {
			u = &entities.UserInfo{UserName: req.Form["UserName"][0], PassWord: req.Form["PassWord"][0], Email: req.Form["Email"][0]}
			entities.UserInfoService.Save(u)
			formatter.JSON(w, 201, struct {
				Success bool `json:"Success"`
			}{true})
		} else {
			formatter.JSON(w, 200, struct {
				Success bool `json:"Success"`
			}{false})
		}
	}
}

// get userinfo
func getInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		u := entities.UserInfoService.FindByUserName(id)
		if u.UserName == "" {
			formatter.JSON(w, 200, struct {
				Success  bool      `json:"Success"`
				UserName string    `json:"UserName"`
				Email    string    `json:"Email"`
				CreateAt time.Time `json:"CreateAt"`
			}{false, "", "", time.Now()})
		} else {
			formatter.JSON(w, 200, struct {
				Success  bool      `json:"Success"`
				UserName string    `json:"UserName"`
				Email    string    `json:"Email"`
				CreateAt time.Time `json:"CreateAt"`
			}{true, u.UserName, u.PassWord, u.CreateAt})
		}
	}
}

// get all meetings that user participates
func allParticipatedHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		p := entities.ParticipatedService.GetUserAllParticipated(id)
		formatter.JSON(w, 200, p)
	}
}

// check a user participates certain meeting or not.return a struct{ Exist bool }
func participatedOrNotHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		title := vars["title"]
		p := &entities.Participated{UserName: id, Title: title}
		has := entities.ParticipatedService.GetUserParticipated(p)
		formatter.JSON(w, 200, struct{ Exist bool }{has})
	}
}

//return all meetings that a user created
func ownedHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		m := entities.MeetingInfoService.GetUserAllMeetings(id)
		formatter.JSON(w, 200, m)
	}
}

//return all meetings
func allmeetingsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		m := entities.MeetingInfoService.GetAllMeetings()
		formatter.JSON(w, 200, m)
	}
}

//post new meeting
//you MUST post a form which has string start and string end, their format is time.ANSIC
//the form must contain password too.
//we need req.Form["start"],req.Form["end"],req.Form["password"]
//return struct{Success bool}
func postMeetingsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		vars := mux.Vars(req)
		m := entities.MeetingInfoService.GetMeetingByTitle(vars["title"])
		if m.Title == "" {
			m = new(entities.MeetingInfo)
			m.Title = vars["title"]
			m.Owned = vars["id"]
			var err1, err2 error
			m.Start, err1 = time.Parse(time.ANSIC, req.Form["start"][0])
			checkErr(err1)
			m.End, err2 = time.Parse(time.ANSIC, req.Form["end"][0])
			checkErr(err2)
			u := entities.UserInfoService.FindByUserName(m.Owned)
			if u.UserName == "" || req.Form["password"][0] != u.PassWord {
				formatter.JSON(w, 200, struct {
					Success bool `json:"Success"`
				}{false})
			} else {
				err := entities.MeetingInfoService.Save(m)
				checkErr(err)
				formatter.JSON(w, 201, struct {
					Success bool `json:"Success"`
				}{true})
			}
		} else {
			formatter.JSON(w, 200, struct {
				Success bool `json:"Success"`
			}{false})
		}
	}
}

//delcare a user participates a meetings
//you MUST post a form contain password,that is  req.Form["password"]
//title must exsit, password must be correct, or it will fail.
func postParticipatedHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		vars := mux.Vars(req)
		m := entities.MeetingInfoService.GetMeetingByTitle(vars["title"])
		p := &entities.Participated{UserName: vars["id"], Title: vars["title"]}
		phas := entities.ParticipatedService.GetUserParticipated(p)
		if m.Title != "" && (!phas) {
			p.UserName = vars["id"]
			p.Title = vars["title"]
			u := entities.UserInfoService.FindByUserName(p.UserName)
			if u.UserName == "" || req.Form["password"][0] != u.PassWord {
				formatter.JSON(w, 200, struct {
					Success bool `json:"Success"`
				}{false})
			} else {
				err := entities.ParticipatedService.Save(p)
				checkErr(err)
				formatter.JSON(w, 201, struct {
					Success bool `json:"Success"`
				}{true})
			}
		} else {
			formatter.JSON(w, 200, struct {
				Success bool `json:"Success"`
			}{false})
		}
	}
}
