package entity

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Assigned *User  `json:"user"`
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
