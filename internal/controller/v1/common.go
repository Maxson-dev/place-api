package v1

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
