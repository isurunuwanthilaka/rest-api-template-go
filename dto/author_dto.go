package dto

type AuthorResDto struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	GenderId int8   `json:"gender"`
}

type AuthorReqDto struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	GenderId int8   `json:"gender_id"`
}
