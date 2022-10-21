package resources

type InputSocialMedias struct {
	UserID         uint   `json:"user_id" validate:"required"`
	Name           string `json:"username" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}
