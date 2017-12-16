package service

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

/*/v1/app/userinfo/{id}
/v1/app/userinfo/{id}/participated
/v1/app/userinfo/{id}/participated/{title}
/v1/app/userinfo/{id}/owned
/v1/app/allmeetings
/v1/app/userinfo/{id}/post/participated/{title}
/v1/app/userinfo/{id}/post/own/{title}
/v1/app/login
/v1/app/reg
*/
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/v1/app/login", loginHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/app/reg", regHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/app/userinfo/{id}", getInfoHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/app/userinfo/{id}/participated", allParticipatedHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/app/userinfo/{id}/participated/{title}", participatedOrNotHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/app/userinfo/{id}/owned", ownedHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/app/allmeetings", allmeetingsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/app/userinfo/{id}/post/participated/{title}", postParticipatedHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/app/userinfo/{id}/post/own/{title}", postMeetingsHandler(formatter)).Methods("POST")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
