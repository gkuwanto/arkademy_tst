package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID     int
	Name   string
	Skills []Skill
}

type Skill struct {
	ID      int
	Name    string
	User_id int
}

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_DATABASENAME = "arkademy"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user", addUser)
	http.HandleFunc("/skill", addSkill)
	http.ListenAndServe(":8000", nil)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp(localhost:3306)/"+DB_DATABASENAME)
		if err != nil {
			log.Fatal(err)
		}
		db.Exec("INSERT INTO `users` (`Name`) VALUES (?)", r.FormValue("name"))
	}
	http.Redirect(w, r, "/", 303)
}

func addSkill(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		user_id := r.FormValue("user")
		user, _ := strconv.Atoi(user_id)
		db, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp(localhost:3306)/"+DB_DATABASENAME)
		if err != nil {
			log.Fatal(err)
		}
		db.Exec("INSERT INTO `skills` (`NAME`, `user_id`) VALUES (?, ?)", name, user)
	}
	http.Redirect(w, r, "/", 303)
}

func index(w http.ResponseWriter, r *http.Request) {
	var users []User
	var user User
	db, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp(localhost:3306)/"+DB_DATABASENAME)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		scanErr := rows.Scan(&user.ID, &user.Name)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		users = append(users, user)
	}
	for i := range users {
		rows, err := db.Query("SELECT * FROM skills WHERE user_id = ?", users[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var skill Skill
			scanErr := rows.Scan(&skill.ID, &skill.Name, &skill.User_id)
			if scanErr != nil {
				log.Fatal(scanErr)
			}
			users[i].Skills = append(users[i].Skills, skill)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tmpl.Execute(w, struct {
		Users []User
	}{Users: users})
}
