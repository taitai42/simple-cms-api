package model

import "github.com/ASWATFZLLC/VirtualCard/db"

type Home struct {
	Title string `json:"title" db:"title"`
	ImageURL string `json:"image_url" db:"image_url"`
	Description string `json:"description" db:"description"`
}

func (h *Home) GetHome () error {
	err := db.GetDB().Get(h, "select * from home")
	if err != nil {
		return err
	}

	return nil
}

func (h *Home) UpdateHome() error {
	_, err := db.GetDB().NamedExec("update home set title = :title, image_url = :image_url, description = :description", h)
	if err != nil {
		return err
	}

	return nil
}


func (h *Home) InsertHome() error {
	_, err := db.GetDB().NamedExec("insert into home set title = :title, image_url = :image_url, description = :description", h)
	if err != nil {
		return err
	}

	return nil
}

