package domain

// model/entity - existance
type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imgUrl" db:"img_url"`
}
