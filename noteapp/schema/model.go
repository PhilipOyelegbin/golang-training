package schema

type UserData struct{
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type NoteData struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	UserId int `json:"userId"`
}