package sshtool

import (
	"os"
	"fmt"
)

func ce(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}