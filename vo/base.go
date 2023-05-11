package vo

type Base struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
