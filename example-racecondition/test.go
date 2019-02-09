package main

import (
	"encoding/json"
	"fmt"
)

type UserType int

const (
	Personal UserType = iota
	Business
	Institution
)

func (u *UserType) MarshalJSON() ([]byte, error) {
	switch *u {
	case Personal:
		return []byte(`"PERSONAL"`), nil
	case Business:
		return []byte(`"BUSINESS"`), nil
	case Institution:
		return []byte(`"INSTITUTION"`), nil
	default:
		err := fmt.Errorf("unknown type %d", u)
		return nil, err
	}
}

func (u *UserType) UnmarshalJSON(i []byte) error {
	switch string(i) {
	case `"PERSONAL"`:
		*u = Personal
		return nil
	case `"BUSINESS"`:
		*u = Business
		return nil
	case `"INSTITUTION"`:
		*u = Institution
		return nil
	default:
		err := fmt.Errorf("unknown type %s", i)
		return err
	}
}

type User struct {
	Name string   `json:"name"`
	Type UserType `json:"type"`
}

func main() {
	user1 := User{
		Name: "Martin",
		Type: Business,
	}

	j, err := json.Marshal(&user1)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))

	j2 := []byte(`{"name": "Martin", "type": "BUSINESS"}`)
	user2 := User{}

	err = json.Unmarshal(j2, &user2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", user2)
}
