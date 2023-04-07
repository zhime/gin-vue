package model

type FadeBackModel struct {
	Model
	Email        string `json:"email"`
	Content      string `json:"content"`
	ApplyContent string `json:"apply_content"`
	IsApply      bool   `json:"is_apply"`
}
