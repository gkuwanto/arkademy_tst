package main

import (
	"encoding/json"
	"fmt"
)

type School struct {
	HighSchool string `json: "highSchool"`
	University string `json: "university"`
}
type Skill struct {
	Name  string `json: "name"`
	Score int    `json: "score"`
}

type BioData struct {
	Name       string   `json: "name"`
	Address    string   `json: "address"`
	Hobbies    []string `json: "hobbies"`
	Is_married bool     `json: "is_married"`
	School     School   `json: "school"`
	Skills     []Skill  `json: "skills"`
}

func bioData() string {
	var hobbies []string
	hobbies = append(hobbies, "Video Games")
	hobbies = append(hobbies, "Travelling")
	hobbies = append(hobbies, "Running")
	school := School{HighSchool: "Harapan Bangsa", University: "Insititut Teknologi Bandung"}
	var skills []Skill
	skills = append(skills, Skill{Name: "Python", Score: 50})
	skills = append(skills, Skill{Name: "Golang", Score: 50})
	skills = append(skills, Skill{Name: "Design", Score: 20})
	bio := &BioData{
		Name:       "Garry Kuwanto",
		Address:    "Taman Bukit Golf Barat HG5/43 Tangerang Banten",
		Hobbies:    hobbies,
		Is_married: false,
		School:     school,
		Skills:     skills,
	}
	content, _ := json.Marshal(bio)
	return string(content)
}

func main() {
	jsonText := bioData()
	fmt.Println(jsonText)
}
