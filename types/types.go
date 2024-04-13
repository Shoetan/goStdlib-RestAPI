package types

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

type User struct {
	Id uint `json:"id" autoincr:"true" primary:"true"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string
	Token string
}