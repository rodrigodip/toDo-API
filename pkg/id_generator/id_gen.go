package IDgenerator

import (
	"strconv"
)

var lastID int

// TODO: Função para receber o ultimo ID persistido no DB ou no FS
func NewID() string {

	var newID int = lastID + 1
	lastID = newID

	return strconv.Itoa(newID)
}
