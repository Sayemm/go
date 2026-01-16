package cart

import (
	"ecommerce/domain"
	"fmt"
	"sync"
)

type service struct {
	cartRepo    CartRepo
	productRepo ProductRepo
}

func NewService(cartRepo CartRepo, productRepo ProductRepo) Service {
	return &service{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *service) AddItem(userID, productID, quantity int) (*domain.CartItemDetail, error) {
	var wg sync.WaitGroup
	errCh := make(chan error, 3)

	var mu sync.Mutex
	var cart *domain.Cart
	var product *domain.Product

	// GOROUTINE 1: Get or create cart
	wg.Add(1)
	go func() {
		defer wg.Done()

		c, err := s.cartRepo.GetOrCreate(userID)
		if err != nil {
			fmt.Println("Error Getting Cart!")
			errCh <- fmt.Errorf("failed to get cart: %w", err)
			return
		}

		mu.Lock()
		cart = c // need mutex because another goroutine might be reading
		mu.Unlock()
	}()

	// GOROUTINE 2: Validate product exists
	wg.Add(1)
	go func() {
		defer wg.Done()

		p, err := s.productRepo.Get(productID)
		if err != nil {
			fmt.Println("Error fetching product")
			errCh <- fmt.Errorf("failed to validate product: %w", err)
			return
		}

		if p == nil {
			fmt.Println("Product not found")
			errCh <- fmt.Errorf("product not found")
			return
		}

		mu.Lock()
		product = p
		mu.Unlock()
	}()

	// GOROUTINE 3: Validate quantity and stock
	wg.Add(1)
	go func() {
		defer wg.Done()

		if quantity <= 0 {
			fmt.Println("Invalid quantity")
			errCh <- fmt.Errorf("quantity must be positive")
			return
		}
	}()

	wg.Wait()
	close(errCh)

	fmt.Println("Checking for errors...")
	for err := range errCh {
		if err != nil {
			fmt.Printf("Validation failed: %v\n", err)
			return nil, err
		}
	}
	fmt.Println("No errors - all validations passed!")

	fmt.Println("Adding item to cart...")
	item, err := s.cartRepo.AddItem(cart.ID, productID, quantity)
	if err != nil {
		return nil, fmt.Errorf("failed to add item: %w", err)
	}

	return &domain.CartItemDetail{
		ID:          item.ID,
		ProductID:   product.ID,
		ProductName: product.Title,
		Price:       product.Price,
		ImgUrl:      product.ImgUrl,
		Quantity:    item.Quantity,
		Subtotal:    product.Price * float64(item.Quantity),
	}, nil
}
func (s *service) GetCart(userID int) (*domain.CartDetails, error) {
	cart, err := s.cartRepo.GetOrCreate(userID)
	if err != nil {
		return nil, err
	}

	items, err := s.cartRepo.GetItems(cart.ID)
	if err != nil {
		return nil, err
	}

	var total float64
	for _, item := range items {
		total += item.Subtotal
	}

	return &domain.CartDetails{
		CartID: cart.ID,
		UserID: userID,
		Items:  items,
		Total:  total,
	}, nil
}

func (s *service) UpdateItem(itemID, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity must be positive")
	}

	item, err := s.cartRepo.GetItem(itemID)
	if err != nil {
		return err
	}
	if item == nil {
		return fmt.Errorf("cart item not found")
	}

	return s.cartRepo.UpdateItemQuantity(itemID, quantity)
}

func (s *service) RemoveItem(itemID int) error {
	return s.cartRepo.RemoveItem(itemID)
}

func (s *service) ClearCart(userID int) error {
	cart, err := s.cartRepo.GetOrCreate(userID)
	if err != nil {
		return nil
	}
	return s.cartRepo.Clear(cart.ID)
}
