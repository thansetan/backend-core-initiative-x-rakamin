package response

type Meta struct {
	Total  int64 `json:"total"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
