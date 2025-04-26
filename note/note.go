package note

import "time"

type Note struct {
	tittle    string
	content   string
	createtAt time.Time
}

func New(tittle, content string) Note {
	return Note{
		tittle:    tittle,
		content:   content,
		createtAt: time.Now(),
	}
}
