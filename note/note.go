package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	tittle    string
	content   string
	createtAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", note.tittle, note.content)
}

func New(tittle, content string) (Note, error) {
	if tittle == "" || content != "" {
		return Note{}, errors.New("Invalit input")
	}

	return Note{
		tittle:    tittle,
		content:   content,
		createtAt: time.Now(),
	}, nil
}
