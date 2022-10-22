package resources

type InputComments struct {
	UserID  uint   `json:"user_id"`
	PhotoId uint   `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type InputCommentsUpdate struct {
	UserID  uint   `json:"user_id"`
	PhotoId uint   `json:"photo_id"`
	Message string `json:"message" validate:"required"`
}
