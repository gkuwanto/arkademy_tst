package main

import (
	"fmt"
	"regexp"
)

func is_username_valid(user string) bool {
	r, _ := regexp.Compile("([a-z.]){8}")
	return r.MatchString(user)
}
func is_email_valid(email string) bool {
	r, _ := regexp.Compile("([a-zA-Z1234567890.]){4,}@([a-zA-Z1234567890]).([a-z])")
	return r.MatchString(email)
}

func main() {
	fmt.Println(is_username_valid("zeronull"))
	fmt.Println(is_email_valid("aku@kamu.com"))
	fmt.Println(is_email_valid("kamu@aku.com"))
}
