package Models

import "time"

type Todos struct {
	NoteId      string    `json:"noteid" db:"id"`
	UserId      string    `json:"userid" db:"userId"`
	Name        string    `json:"title" db:"name"`
	Note        string    `json:"user_note" db:"note"`
	IsCompleted bool      `json:"iscompleted" db:"iscompleted"`
	CreatedAt   time.Time `json:"createdat" db:"createdAt"`
	UpdatedAt   time.Time `json:"updatedat" db:"updatedAt"`
}
