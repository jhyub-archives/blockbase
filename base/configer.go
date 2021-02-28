package base

import (
	"crypto/sha512"
	"fmt"
)

type User struct {
	id string
	password string
	authorities string
}

func GenerateUser(name string, password string, read bool, write bool) {
	at := ""
	if read {
		at = at+"1"
	} else {
		at = at+"0"
	}
	if write {
		at = at+"1"
	} else {
		at = at+"0"
	}
	abc := User{name, password, at}
	print(abc)
}