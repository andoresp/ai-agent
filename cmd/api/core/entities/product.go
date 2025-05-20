package entities

type Product struct {
	Id        int32   `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Quantity  float32 `json:"quantity"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`

	CategoryId int32 `json:"categoryId"`
}
