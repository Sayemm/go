package repo

import (
	"database/sql"
	"ecommerce/cart"
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type CartRepo interface {
	cart.CartRepo
}

type cartRepo struct {
	db *sqlx.DB
}

func NewCartRepo(db *sqlx.DB) CartRepo {
	return &cartRepo{
		db: db,
	}
}

// Get user's cart, or create one if it doesn't exist
func (r *cartRepo) GetOrCreate(userID int) (*domain.Cart, error) {
	var cart domain.Cart
	query := `SELECT id, user_id, created_at, updated_at FROM carts WHERE user_id = $1`

	err := r.db.Get(&cart, query, userID)
	if err == nil {
		return &cart, nil
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	insertQuery := `
		INSERT INTO carts (user_id)
		VALUES ($1)
		RETURNING id, user_id, created_at, updated_at
	`

	err = r.db.QueryRow(insertQuery, userID).Scan(
		&cart.ID,
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

// Add item or update quantity if already exists
func (r *cartRepo) AddItem(cartID, productID, quantity int) (*domain.CartItem, error) {
	query := `
		INSERT INTO cart_items (cart_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (cart_id, product_id)
		DO UPDATE SET 
			quantity = cart_items.quantity + $3,
			updated_at = CURRENT_TIMESTAMP
		RETURNING id, cart_id, product_id, quantity, created_at, updated_at
	`
	var item domain.CartItem
	err := r.db.QueryRow(query, cartID, productID, quantity).Scan(
		&item.ID,
		&item.CartID,
		&item.ProductID,
		&item.Quantity,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *cartRepo) GetItems(cartID int) ([]domain.CartItemDetail, error) {
	query := `
		SELECT 
			ci.id,
			ci.product_id,
			p.title as product_name,
			p.price,
			p.img_url,
			ci.quantity,
			(p.price * ci.quantity) as subtotal
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		WHERE ci.cart_id = $1
		ORDER BY ci.created_at DESC
	`

	var items []domain.CartItemDetail
	err := r.db.Select(&items, query, cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return []domain.CartItemDetail{}, nil
		}
		return nil, err
	}
	return items, nil
}
func (r *cartRepo) UpdateItemQuantity(itemID, quantity int) error {
	query := `
		UPDATE cart_items
		SET quantity = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	result, err := r.db.Exec(query, quantity, itemID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
func (r *cartRepo) RemoveItem(itemID int) error {
	query := `DELETE FROM cart_items WHERE id = $1`

	result, err := r.db.Exec(query, itemID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *cartRepo) Clear(cartID int) error {
	query := `DELETE FROM cart_items WHERE cart_id = $1`
	_, err := r.db.Exec(query, cartID)
	return err
}

func (r *cartRepo) GetItem(itemID int) (*domain.CartItem, error) {
	var item domain.CartItem
	query := `
		SELECT id, cart_id, product_id, quantity, created_at, updated_at
		FROM cart_items
		WHERE id = $1
	`

	err := r.db.Get(&item, query, itemID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}
