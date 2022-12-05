// https://go.dev/play/p/gZ5AK8-o4zA

package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

var ranOnce bool

const (
	UnknownCode     = 0
	UnreachableCode = 1
	AuthFailureCode = 2
)

// Custome network error with error codes and a message.
type ErrNetwork struct {
	Code int
	Msg  string
}

// Error method on type ErrNetwork
func (e ErrNetwork) Error() string {
	return fmt.Sprintf("network error(%d): %s", e.Code, e.Msg)
}

// someFunc is a stand-in for some function that takes in some data.
// On the first call it will return an ErrNetwork and on the second
// just a standard error.
func someFunc(data string) error {
	if !ranOnce {
		ranOnce = true
		return ErrNetwork{
			Code: AuthFailureCode,
			Msg:  "user unrecognized",
		}
	}
	return fmt.Errorf("another error")
}

func main() {
	// This loop will continue as long as we have network errors
	// that are not code AuthFailureCode and will terminate on
	// any other error or success.
	for {
		if err := someFunc("data"); err != nil {
			var netErr ErrNetwork
			if errors.As(err, &netErr) {
				if netErr.Code == AuthFailureCode {
					log.Println("autn failure! Danger Will Robinson!")
					return
				}
				log.Println("network error: ", err)
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("unrecoverable: ", err)
		}
		break
	}
}
