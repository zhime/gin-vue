package model

type BannerModel struct {
	Model
	Path string `json:"path"`
	Hash string `json:"hash"`
	Name string `json:"name"`
}
