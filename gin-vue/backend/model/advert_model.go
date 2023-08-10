package model

type AdvertModel struct {
	Model
	Title  string `json:"title"`
	Href   string `json:"href"`
	Images string `json:"images"`
	IsShow bool   `json:"is_show"`
}
