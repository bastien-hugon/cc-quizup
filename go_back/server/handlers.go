package server

import (
	"net/http"
	"strconv"
)

func (srv *Server) homeHandler(w http.ResponseWriter, r *http.Request) {

	srv.tmpl.ExecuteTemplate(w, "home.tmpl", nil)
}

func (srv *Server) adminHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := GetInfoUser(srv, w, r)

	if u.ID == 0 {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	srv.tmpl.ExecuteTemplate(w, "admin.tmpl", u)
}

func (srv *Server) adminLoginHandler(w http.ResponseWriter, r *http.Request) {

	_, c := GetInfoUser(srv, w, r)

	if r.Method == http.MethodPost {
		u := User{
			name: r.FormValue("name"),
			pass: r.FormValue("pass"),
		}
		u.LogUserIn(srv, w, r, c)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
	srv.tmpl.ExecuteTemplate(w, "adminLogin.tmpl", nil)
}

func (srv *Server) adminSeeQuizHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := GetInfoUser(srv, w, r)

	if u.ID == 0 {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
	quiz := getAllQuiz(srv.db)
	srv.tmpl.ExecuteTemplate(w, "adminSeeQuiz.tmpl", quiz)
}

func (srv *Server) adminCreateQuizHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := GetInfoUser(srv, w, r)

	if u.ID == 0 {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
	if r.Method == http.MethodPost {
		q := quiz{
			Question: r.FormValue("question"),
			Answer:   r.FormValue("answer"),
			Choice1:  r.FormValue("choice1"),
			Choice2:  r.FormValue("choice2"),
			Choice3:  r.FormValue("choice3"),
			Choice4:  r.FormValue("choice4"),
		}
		q.addNewQuiz(srv.db)
		http.Redirect(w, r, "/admin/see/quiz", http.StatusSeeOther)
	}
	srv.tmpl.ExecuteTemplate(w, "adminNewQuiz.tmpl", nil)
}

func (srv *Server) adminDeleteQuizHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := GetInfoUser(srv, w, r)

	if u.ID == 0 {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}

	// Getting quiz's ID, with an error handling in case of bad path.
	id, err := strconv.Atoi(r.URL.Path[len("/admin/delete/quiz/"):])
	if err != nil {
		http.Redirect(w, r, "/admin/see/quiz", http.StatusSeeOther)
	}

	deleteQuiz(srv.db, id)
	http.Redirect(w, r, "/admin/see/quiz", http.StatusSeeOther)
}

func (srv *Server) adminSeeTeamsHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := GetInfoUser(srv, w, r)

	if u.ID == 0 {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
	teams := getAllTeam(srv.db)
	srv.tmpl.ExecuteTemplate(w, "adminSeeTeams.tmpl", teams)
}

func (srv *Server) adminDeleteTeamHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := GetInfoUser(srv, w, r)

	if u.ID == 0 {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}

	// Getting quiz's ID, with an error handling in case of bad path.
	id, err := strconv.Atoi(r.URL.Path[len("/admin/delete/team/"):])
	if err != nil {
		http.Redirect(w, r, "/admin/see/team", http.StatusSeeOther)
	}

	deleteTeam(srv.db, id)
	http.Redirect(w, r, "/admin/see/team", http.StatusSeeOther)
}

/////////////////////
// Gameplay handlers
/////////////////////
func (srv *Server) playNewGameHandler(w http.ResponseWriter, r *http.Request) {

	_, c := GetInfoUser(srv, w, r)

	if r.Method == http.MethodPost {
		t := team{
			Name: r.FormValue("name"),
		}
		t.LogTeamIn(srv, w, r, c)
		http.Redirect(w, r, "/play/start", http.StatusSeeOther)
	}
	srv.tmpl.ExecuteTemplate(w, "playNewGame.tmpl", nil)
}

func (srv *Server) playHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := GetInfoTeam(srv, w, r)

	if t.ID == 0 {
		http.Redirect(w, r, "/play/new-game", http.StatusSeeOther)
	}
	srv.tmpl.ExecuteTemplate(w, "play.tmpl", nil)
}
