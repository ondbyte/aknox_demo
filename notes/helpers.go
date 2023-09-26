package notes

import "fmt"

var _id = -1

func newId() string {
	_id++
	return fmt.Sprint(_id)
}
