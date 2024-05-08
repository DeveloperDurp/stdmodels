package stdmodels

type StandardMessage struct {
	Message string `json:"message"`
}

type StandardError struct {
	Message     string   `json:"message"`
	Status      int      `json:"status"`
	Description []string `json:"description"`
}
