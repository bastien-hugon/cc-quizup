package server

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	// importing mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	templatesPath = "./templates/*.tmpl"

	databaseDriver = "mysql"
	databaseUser   = "root"
	databasePass   = "mysql;root123"
	databaseName   = "quiz"
)

// Server defines the main server struct
type Server struct {
	peer     string
	db       *sql.DB
	tmpl     *template.Template
	sessions map[string]uint
	users    map[uint]User
	game     map[string]uint
	teams    map[uint]team
}

func connectDatabase() *sql.DB {

	db, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

// InitServer returns the initialized server with the given peer
func InitServer(peer string) *Server {

	srv := &Server{
		peer:     peer,
		tmpl:     template.Must(template.ParseGlob(templatesPath)),
		db:       connectDatabase(),
		sessions: map[string]uint{},
		users:    map[uint]User{},
		game:     map[string]uint{},
		teams:    map[uint]team{},
	}
	return srv
}

// InitHandler will init all the Handler for the URLs
func (srv *Server) InitHandler() {

	http.HandleFunc("/", srv.homeHandler)

	http.HandleFunc("/admin", srv.adminHandler)
	http.HandleFunc("/admin/login", srv.adminLoginHandler)

	http.HandleFunc("/admin/see/quiz", srv.adminSeeQuizHandler)
	http.HandleFunc("/admin/create/new/quiz", srv.adminCreateQuizHandler)
	http.HandleFunc("/admin/delete/quiz/", srv.adminDeleteQuizHandler)

	http.HandleFunc("/admin/see/team", srv.adminSeeTeamsHandler)
	http.HandleFunc("/admin/delete/team/", srv.adminDeleteTeamHandler)

	http.HandleFunc("/play", srv.playHandler)
	http.HandleFunc("/play/new-game", srv.playNewGameHandler)
}

// RunServer will run the server and log if there is an error
func (srv *Server) RunServer() {

	err := http.ListenAndServe(":"+srv.peer, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}
