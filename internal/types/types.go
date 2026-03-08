package types

type ProjectPayload struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Picture     *string `json:"picture"`
}
