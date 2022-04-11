package se_panics

import "fmt"

//DataDuplicate throw a panic to duplicate data
func DataDuplicate(message string) string {
	messageFill := fmt.Sprintf(`%s email já cadastrado`, message)
	return messageFill
}
