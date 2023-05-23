package models

type User struct {
	ID   uint
	Name string
}

type Article struct {
	ID      uint
	Title   string
	Content string
	UserID  uint
}

type Tag struct {
	ID    uint
	Name  string
	Color string
}

type ArticleTag struct {
	ArticleID uint
	TagID     uint
}
