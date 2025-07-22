package dto

type ProductPatchBody struct {
	Sku      *string `json:"sku,omitempty"`
	Name     *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
	Price    *int    `json:"price,omitempty"`
}
