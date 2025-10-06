package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Delete(id int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)

	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	for _, product := range r.productList {
		if id == product.ID {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if product.ID == p.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

func (r *productRepo) Delete(id int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if id != p.ID {
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList

	return nil
}

func generateInitialProducts(r *productRepo) {
	pd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Photo of Orange",
		Price:       20,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	pd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Photo of Apple",
		Price:       2,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	pd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Photo of Banana",
		Price:       60,
		ImgUrl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcQCWfkZS0Zx7Hz7LBbW1-VggJsj1vDD_2bJnBaezT3YpRDljzGkASfYN8iI3wNFCHM59cwSWAHAlBZzQmwUg_n1I-WI8loYzTib-Xs40lM",
	}

	r.productList = append(r.productList, pd1)
	r.productList = append(r.productList, pd2)
	r.productList = append(r.productList, pd3)
}

/*
Repository Design Pattern
--------------------------
-
*/
