package entities

type Category struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
