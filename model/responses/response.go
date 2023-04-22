package responses

type Response[T interface{}] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
