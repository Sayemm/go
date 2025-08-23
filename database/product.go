package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Photo of Orange",
		Price:       20,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Photo of Apple",
		Price:       2,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	pd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Photo of Banana",
		Price:       60,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	ProductList = append(ProductList, pd1)
	ProductList = append(ProductList, pd2)
	ProductList = append(ProductList, pd3)
}
