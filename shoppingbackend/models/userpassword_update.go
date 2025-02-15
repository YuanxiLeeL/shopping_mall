package models

type OldPassword struct {
	OldPassword string ` json:"old_password" binding:"required"`
}

type NewPassword struct {
	NewPassword string ` json:"new_password" binding:"required"`
}
