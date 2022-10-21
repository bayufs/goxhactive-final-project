package helpers

type Response struct {
	Payload any    `json:"payload"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func JWTResponse[D any](message string, status int, tokenize D) Response {
	jsonResponse := Response{
		Payload: tokenize,
		Message: message,
		Status:  status,
	}

	return jsonResponse
}
