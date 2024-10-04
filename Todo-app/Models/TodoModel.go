package Models

import "time"

type Todos struct {
	NoteId      string    `json:"note_id" db:"id"`
	UserId      string    `json:"user_id" db:"userId"`
	Name        string    `json:"title" db:"name"`
	Note        string    `json:"user_note" db:"note"`
	IsCompleted bool      `json:"is_completed" db:"iscompleted"`
	CreatedAt   time.Time `json:"created_at" db:"createdAt"`
	UpdatedAt   time.Time `json:"updated_at" db:"updatedAt"`
}
