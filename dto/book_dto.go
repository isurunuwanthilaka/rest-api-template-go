package dto

type BookResDto struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Genre   string `json:"genre"`
	Summary string `json:"summary"`
}

type BookReqDto struct {
	Title    string `json:"title"`
	AuthorId uint   `json:"author_id"`
	GenreId  uint   `json:"genre_id"`
	Summary  string `json:"summary"`
}
