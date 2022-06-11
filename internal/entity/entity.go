package entity

type Task struct {
	ID          int64
	Open        bool
	Title       string
	Description string
	AuthorID    int64
}
