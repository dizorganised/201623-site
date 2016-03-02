package web

import (
	"github.com/UniversityRadioYork/2016-site/controllers"
	"github.com/UniversityRadioYork/2016-site/structs"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/UniversityRadioYork/myradio-go"
)

type Server struct {
	*negroni.Negroni
}

func NewServer(c *structs.Config) (*Server, error) {

	s := Server{negroni.Classic()}

	session, err := myradio.NewSessionFromKeyFile()

	if err != nil {
		return &s, err;
	}

	router := mux.NewRouter()

	getRouter := router.Methods("GET").Subrouter()

	// Routes go in here
	nfc := controllers.NewNotFoundController(c)
	router.NotFoundHandler = http.HandlerFunc(nfc.Get)

	ic := controllers.NewIndexController(session, c)
	getRouter.HandleFunc("/", ic.Get)

	sc := controllers.NewSearchController(session, c)
	getRouter.HandleFunc("/search", sc.Get);

	// End routes

	s.UseHandler(router)

	return &s, nil;

}
