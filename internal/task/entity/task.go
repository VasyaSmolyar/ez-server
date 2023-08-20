package entity

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	UserId      string `json:"user_id"`
	ImageId     string `json:"image_id"`
}
