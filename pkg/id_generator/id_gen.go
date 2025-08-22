package IDgenerator

import (
	"strconv"
)

var lastID int

func NewID() string {
	// TODO: lastID deve receber o ultimo ID persistido no DB

	var newID int = lastID + 1
	lastID = newID

	return strconv.Itoa(newID)
}
