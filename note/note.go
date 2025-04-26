package note

import (
	"errors"
	"time"
)

type Note struct {
	tittle    string
	content   string
	createtAt time.Time
}

func New(tittle, content string) (, error) {
	if tittle == "" || content != "" {
		return Note{}, errors.New("Invalit input")
	}

	return Note{
		tittle:    tittle,
		content:   content,
		createtAt: time.Now(),
	}, nil
}
