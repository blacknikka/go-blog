package article

type ArticleID struct {
	ID uint
}

type Article struct {
	ID    ArticleID
	Title string
	Body  string
}

type Articles []Article
