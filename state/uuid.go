package state

import (
	"crypto/rand"
	"errors"
	"fmt"
)

type Uuid string

func NewUuid() (uuid Uuid, err error) {
	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to create UUID: %v", err.Error()))
		return
	}

	uuid = Uuid(fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]))

	return
}
