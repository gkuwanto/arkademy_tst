package main

import (
	"net/http"
	"strconv"
	"text/template"
)

type User struct {
	ID   int
	Name string
}

type Skill struct {
	ID      int
	Name    string
	User_id int
}

var users []User
var skills []Skill
var user_idx int
var skill_idx int

func main() {
	users = append(users, User{0, "Garry"})
	http.HandleFunc("/", index)
	http.HandleFunc("/user", addUser)
	http.HandleFunc("/skill", addSkill)
	http.ListenAndServe(":8000", nil)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		users = append(users, User{ID: user_idx + 1, Name: r.FormValue("name")})
	}
}

func addSkill(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		user_id := r.FormValue("user")
		user, _ := strconv.Atoi(user_id)
		skills = append(skills, Skill{ID: skill_idx + 1, Name: name, User_id: user})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tmpl.Execute(w, struct {
		Users  []User
		Skills []Skill
	}{Users: users, Skills: skills})
}
