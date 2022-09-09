// https://www.youtube.com/watch?v=Kf5LLi7Ti9A&t=12s

package main

import "errors"

type Password string

func (p Password) Validate() error {
	if len(p) < 5 {
		return errors.New("password is too short")
	}
	return nil
}

func main () {
	password := Password("123")
	if err := password.Validate(); err != nil {
		panic(err)
	} 
}