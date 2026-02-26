package types

type ProjectPayload struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Picture     *string `json:"picture"`
}

type Project struct {
	ID          *string `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Picture     *string `json:"picture" db:"picture"`
}
