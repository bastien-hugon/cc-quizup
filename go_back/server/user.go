package server

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	adminPseudo   = "admin"
	adminPassword = "admin"

	hoursSessionExpiration = 5

	routineRestTimeSecond = 20
)

// User struct defines the whole user, its name and its session deadline
type User struct {
	ID     uint
	name   string
	pass   string
	expire time.Time
}

// GetInfoUser will return the new user and the cookie linked to the user.
func GetInfoUser(srv *Server, w http.ResponseWriter, r *http.Request) (User, *http.Cookie) {

	c, err := r.Cookie("Quiz-CC")

	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{Name: "Quiz-CC", Value: sID.String()}
		http.SetCookie(w, c)
	}
	var u User
	if ID, ok := srv.sessions[c.Value]; ok {
		u = srv.users[ID]
	}
	return u, c
}

func (u *User) watchExpire(srv *Server, w http.ResponseWriter, r *http.Request) {

	for {
		t := time.Now()
		*u = srv.users[u.ID]
		if t.After(u.expire) == true {
			srv.users[u.ID] = User{}
			return
		}
		time.Sleep(time.Second * routineRestTimeSecond)
	}
}

// LogUserIn will log the user in, redirect to the login url if error occured.
func (u *User) LogUserIn(srv *Server, w http.ResponseWriter, r *http.Request, c *http.Cookie) {

	if u.name != adminPseudo && u.pass != adminPassword {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	u.ID = 1
	u.expire = time.Now().Add(time.Hour * hoursSessionExpiration)
	srv.sessions[c.Value] = u.ID
	srv.users[u.ID] = *u

	go u.watchExpire(srv, w, r)
}
