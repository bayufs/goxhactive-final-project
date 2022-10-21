package resources

type InputComments struct {
	UserID  uint   `json:"user_id" validate:"required"`
	PhotoId uint   `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
