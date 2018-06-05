package server

import (
	"database/sql"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type team struct {
	ID    uint
	Name  string
	Score int
}

func getAllTeam(db *sql.DB) []team {

	selDB, err := db.Query("SELECT * FROM team ORDER BY score")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer selDB.Close()

	ts := make([]team, 0)

	for selDB.Next() {
		var t team
		err = selDB.Scan(&t.ID, &t.Name, &t.Score)
		if err != nil {
			log.Fatal(err.Error())
		}
		ts = append(ts, t)
	}
	return ts
}

func (t *team) addNewTeam(db *sql.DB) {

	ins, err := db.Prepare("INSERT INTO team (name) VALUES(?)")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ins.Close()

	_, err = ins.Exec(t.Name)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func deleteTeam(db *sql.DB, id int) {

	ins, err := db.Prepare("DELETE FROM team WHERE id=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ins.Close()

	ins.Exec(id)
}

func (t *team) LogTeamIn(srv *Server, w http.ResponseWriter, r *http.Request, c *http.Cookie) {

	t.ID = 1
	srv.game[c.Value] = t.ID
	srv.teams[t.ID] = *t
}

// GetInfoTeam will return a cookie and the team for this browser
func GetInfoTeam(srv *Server, w http.ResponseWriter, r *http.Request) (team, *http.Cookie) {

	c, err := r.Cookie("Quiz-CC")

	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{Name: "Quiz-CC", Value: sID.String()}
	}
	var t team
	if ID, ok := srv.game[c.Value]; ok {
		t = srv.teams[ID]
	}
	return t, c
}
