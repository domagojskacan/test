package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string
	Username string
	Id       int
}

type Claims struct {
	Ident int
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

var data Credentials
var db *sql.DB
var allUsers string

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("../testhtml/signup.gtpl")
		temp.Execute(w, nil)
		return
	}
	r.ParseForm()
	name := r.Form["username"]
	trueName := strings.Join(name, ", ")
	data.Username = trueName
	sifra := r.Form["password"]
	trueSifra := strings.Join(sifra, ", ")
	data.Password = trueSifra

	users, err := db.Query("select username from users")
	if err != nil {
		fmt.Println(err)
	}

	for users.Next() {
		err := users.Scan(&allUsers)
		if err != nil {
			fmt.Println(err)
			return
		}
		if allUsers == data.Username {
			tempSame, _ := template.ParseFiles("../testhtml/isti.gtpl")
			tempSame.Execute(w, nil)
			return
		}

	}

	// username not found

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	if _, err := db.Query("insert into users values ($1, $2)", data.Username, string(hashedPassword)); err != nil {
		fmt.Println(err)
		return
	}
	tempDobro, _ := template.ParseFiles("../testhtml/dobro.gtpl")
	tempDobro.Execute(w, nil)

}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp, _ := template.ParseFiles("../testhtml/login.gtpl")
		temp.Execute(w, nil)
	} else {
		r.ParseForm()

		name := r.Form["username"]
		trueName := name[0] //strings.Join(name, ", ")
		data.Username = trueName
		sifra := r.Form["password"]
		trueSifra := strings.Join(sifra, ", ")
		data.Password = trueSifra

		result := db.QueryRow("select password from users where username=$1", data.Username)

		storedData := &Credentials{}
		err := result.Scan(&storedData.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(storedData.Password), []byte(data.Password)); err != nil {
			tempLose, _ := template.ParseFiles("../testhtml/lose.gtpl")
			tempLose.Execute(w, nil)
		} else {
			tempDobro, _ := template.ParseFiles("../testhtml/dobro.gtpl")
			tempDobro.Execute(w, nil)
			uid := db.QueryRow("select id from users where username=$1", data.Username)

			err := uid.Scan(&storedData.Id)
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			expirationTime := time.Now().Add(5 * time.Minute)
			claims := &Claims{
				Ident:          storedData.Id,
				StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(tokenString)

		}

	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("../testhtml/head.gtpl")
	temp.Execute(w, nil)
}
func main() {
	var err error

	const (
		host     = "192.168.8.45"
		port     = 5432
		user     = "postgres"
		password = "example"
		dbname   = "mydb"
	)
	//client := redis.NewClient(&redis.Options{
	//	Addr: "192.168.8.45:6379",
	//		Password: "",
	//		DB: 10,
	//	})

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
