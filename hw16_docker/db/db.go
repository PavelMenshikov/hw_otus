package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitDB(cfg Config) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка открытия подключения: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("ошибка подключения: %w", err)
	}

	log.Println("Успешное подключение к БД")
	return nil
}

func ExecTx(txFunc func(*sql.Tx) error) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	if err := txFunc(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %w, rollback err: %w", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID          int     `json:"id"`
	UserID      int     `json:"userId"`
	OrderDate   string  `json:"orderDate"`
	TotalAmount float64 `json:"totalAmount"`
}

type UserStats struct {
	UserID          int     `json:"userId"`
	TotalSpent      float64 `json:"totalSpent"`
	AvgProductPrice float64 `json:"avgProductPrice"`
}

func GetAllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetAllProducts() ([]Product, error) {
	rows, err := DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetOrdersByUser(userID int) ([]Order, error) {
	rows, err := DB.Query("SELECT id, user_id, order_date, total_amount FROM orders WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.OrderDate, &o.TotalAmount); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func GetUserStats(userID int) (*UserStats, error) {
	query := `
	SELECT u.id, 
	       COALESCE(SUM(o.total_amount), 0) AS total_spent,
	       COALESCE(AVG(p.price), 0) AS avg_product_price
	FROM users u
	LEFT JOIN orders o ON u.id = o.user_id
	LEFT JOIN order_products op ON o.id = op.order_id
	LEFT JOIN products p ON op.product_id = p.id
	WHERE u.id = $1
	GROUP BY u.id;
	`
	var stats UserStats
	err := DB.QueryRow(query, userID).Scan(&stats.UserID, &stats.TotalSpent, &stats.AvgProductPrice)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

func CreateUser(u User) (int, error) {
	var newID int
	err := DB.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		u.Name, u.Email, u.Password,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}
