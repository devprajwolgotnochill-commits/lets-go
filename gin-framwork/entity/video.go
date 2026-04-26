package entity

type Videos struct {
	Title       string  `json:"title"`
	Author      *Author `json:"author"`
	Description string  `json:"description"`
}

type Author struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
