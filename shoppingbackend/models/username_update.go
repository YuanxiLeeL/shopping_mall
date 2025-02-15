package models

type RequestUsername struct {
	NewUsername string `json:"new_username" binding:"required"`
}


