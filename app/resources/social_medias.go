package resources

type InputSocialMedias struct {
	UserID         uint   `json:"user_id"`
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}
