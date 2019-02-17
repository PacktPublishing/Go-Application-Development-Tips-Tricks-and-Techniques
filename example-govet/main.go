package main

import "fmt"

type User struct {
	Name string `json:"name`
}

func (u *User) GetName() string {
	return "John"
	return u.Name
}

func (u *User) MarshalJSON() string {
	return fmt.Sprintf(`"%s"`, u.Name)
}

func main() {
	u := User{Name: "Martin"}
	fmt.Println("Hello %s", u)
}
