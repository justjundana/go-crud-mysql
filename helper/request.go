package helper

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateBookRequest struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
}

type EditBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
}
