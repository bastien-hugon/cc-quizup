package server

import (
	"database/sql"
	"log"
)

type quiz struct {
	ID       int
	Question string
	Answer   string
	Choice1  string
	Choice2  string
	Choice3  string
	Choice4  string
}

func getAllQuiz(db *sql.DB) []quiz {

	selDB, err := db.Query("SELECT * FROM quiz")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer selDB.Close()

	qs := make([]quiz, 0)

	for selDB.Next() {
		var q quiz
		err = selDB.Scan(&q.ID, &q.Question, &q.Answer, &q.Choice1, &q.Choice2, &q.Choice3, &q.Choice4)
		if err != nil {
			log.Fatal(err.Error())
		}
		qs = append(qs, q)
	}
	return qs
}

func (q *quiz) addNewQuiz(db *sql.DB) {

	ins, err := db.Prepare("INSERT INTO quiz (question,answer,choice1,choice2,choice3,choice4) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ins.Close()

	_, err = ins.Exec(q.Question, q.Answer, q.Choice1, q.Choice2, q.Choice3, q.Choice4)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func deleteQuiz(db *sql.DB, id int) {

	ins, err := db.Prepare("DELETE FROM quiz WHERE id=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ins.Close()

	ins.Exec(id)
}
