package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Credentials struct {
	Password string
	Username string
}

var data Credentials

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../testhtml/signup.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		name := r.Form["username"]
		trueName := strings.Join(name, ", ")
		data.Username = trueName
		sifra := r.Form["password"]
		trueSifra := strings.Join(sifra, ", ")
		fmt.Println(trueSifra)
		data.Password = trueSifra

		fmt.Println(data)
		//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
		//if _, err = db.Query("insert into users values ($1, $2)", data.Username, string(hashedPassword)); err != nil {

		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}

	}
}

func main() {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "example"
		dbname   = "mydb"
	)

	http.HandleFunc("/signup", signup)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()

	//err = db.Ping()
	//if err != nil {
	//	panic(err)
//	}

//	fmt.Println("Successfully connected!")
//}
