package dto

type BookResDto struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Author  int64  `json:"author"`
	Genre   int64  `json:"genre"`
	Summary string `json:"summary"`
}

type BookReqDto struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	AuthorId int64  `json:"author_id"`
	GenreId  int64  `json:"genre_id"`
	Summary  string `json:"summary"`
}
