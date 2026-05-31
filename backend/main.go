package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
}

type OrderItem struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type Order struct {
	ID              string      `json:"id"`
	UserID          string      `json:"user_id"`
	CustomerName    string      `json:"customer_name"`
	CustomerEmail   string      `json:"customer_email"`
	CustomerPhone   string      `json:"customer_phone"`
	CustomerAddress string      `json:"customer_address"`
	Items           []OrderItem `json:"items"`
	TotalAmount     float64     `json:"total_amount"`
	Status          string      `json:"status"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type Session struct {
	UserID    string    `json:"user_id"`
	Role      string    `json:"role"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Database struct {
	Users    map[string]User    `json:"users"`
	Products map[string]Product `json:"products"`
	Orders   map[string]Order   `json:"orders"`
}

var (
	db       *Database
	dbMutex  sync.RWMutex
	sessions = make(map[string]Session)
	sessMux  sync.RWMutex
	dbPath   = "db.json"
)

func loadDatabase() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	db = &Database{
		Users:    make(map[string]User),
		Products: make(map[string]Product),
		Orders:   make(map[string]Order),
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Println("Database file not found, initializing seed data...")
		seedData()
		saveDatabaseLocked()
		return
	}

	data, err := os.ReadFile(dbPath)
	if err != nil {
		log.Printf("Error reading database file: %v. Re-initializing...", err)
		seedData()
		saveDatabaseLocked()
		return
	}

	err = json.Unmarshal(data, db)
	if err != nil {
		log.Printf("Error unmarshaling database: %v. Re-initializing...", err)
		seedData()
		saveDatabaseLocked()
	}
}

func saveDatabase() {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	saveDatabaseLocked()
}

func saveDatabaseLocked() {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		log.Printf("Error marshaling database: %v", err)
		return
	}
	err = os.WriteFile(dbPath, data, 0644)
	if err != nil {
		log.Printf("Error writing database file: %v", err)
	}
}

func generateID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func seedData() {
	adminHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	customerHash, _ := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)

	db.Users["admin-id"] = User{
		ID:       "admin-id",
		Name:     "Administrator FourMart",
		Email:    "admin@fourmart.com",
		Password: string(adminHash),
		Role:     "admin",
		Address:  "Gedung Rektorat Lt. 2, Jakarta",
		Phone:    "081234567890",
	}

	db.Users["customer-id"] = User{
		ID:       "customer-id",
		Name:     "Ahmad Budiman",
		Email:    "user@fourmart.com",
		Password: string(customerHash),
		Role:     "customer",
		Address:  "Jl. Kebon Jeruk No. 12, Jakarta Barat",
		Phone:    "089876543210",
	}

	products := []Product{
		{
			ID:          "p1",
			Name:        "Tas Ransel Sekolah Oxford Premium",
			Description: "Tas ransel sekolah awet dengan bahan Oxford tahan air, kompartemen laptop 14 inch, dan busa punggung tebal yang nyaman dipakai seharian.",
			Price:       185000,
			Image:       "https://images.unsplash.com/photo-1553062407-98eeb64c6a62?w=500&auto=format&fit=crop&q=60",
			Stock:       45,
			Category:    "Tas & Kotak Pensil",
		},
		{
			ID:          "p2",
			Name:        "Buku Tulis Spiral A5 Joyko (Isi 5 Pcs)",
			Description: "Paket 5 buku tulis spiral ukuran A5 dengan kertas 80gsm tebal, isi 60 lembar. Sangat cocok untuk mencatat materi pelajaran sekolah maupun kuliah.",
			Price:       45000,
			Image:       "https://images.unsplash.com/photo-1531346878377-a5be20888e57?w=500&auto=format&fit=crop&q=60",
			Stock:       120,
			Category:    "Tulis-Menulis",
		},
		{
			ID:          "p3",
			Name:        "Pulpen Gel Hitam Joyko 0.5mm (1 Lusin)",
			Description: "Satu lusin (12 pcs) pulpen gel tinta hitam ukuran mata pena 0.5mm. Aliran tinta lancar, cepat kering, dan tidak luntur di kertas.",
			Price:       28000,
			Image:       "https://images.unsplash.com/photo-1583485088034-697b5bc54ccd?w=500&auto=format&fit=crop&q=60",
			Stock:       80,
			Category:    "Tulis-Menulis",
		},
		{
			ID:          "p4",
			Name:        "Pensil Warna Faber-Castell Classic 24 Warna",
			Description: "Pensil warna klasik berkualitas tinggi isi 24 warna cerah. Kayu ramah lingkungan, tidak mudah patah saat diraut, dan pigmen warna halus.",
			Price:       62000,
			Image:       "https://images.unsplash.com/photo-1513364776144-60967b0f800f?w=500&auto=format&fit=crop&q=60",
			Stock:       35,
			Category:    "Tulis-Menulis",
		},
		{
			ID:          "p5",
			Name:        "Kalkulator Scientific Casio FX-991EX",
			Description: "Kalkulator ilmiah dengan 552 fungsi matematika, statistik, matriks, dan persamaan. Dilengkapi layar LCD beresolusi tinggi dan dual power (solar & baterai).",
			Price:       320000,
			Image:       "https://images.unsplash.com/photo-1629909613654-28e377c37b09?w=500&auto=format&fit=crop&q=60",
			Stock:       15,
			Category:    "Elektronik & Belajar",
		},
		{
			ID:          "p6",
			Name:        "Kotak Pensil Canvas Estetik",
			Description: "Kotak pensil berbahan kanvas tebal dengan ritsleting kokoh dan kapasitas besar untuk menyimpan pulpen, pensil, penghapus, dan penggaris kecil.",
			Price:       22000,
			Image:       "https://images.unsplash.com/photo-1578844251758-2f71da64c96f?w=500&auto=format&fit=crop&q=60",
			Stock:       60,
			Category:    "Tas & Kotak Pensil",
		},
		{
			ID:          "p7",
			Name:        "Penggaris Stainless Steel Kenko 30cm",
			Description: "Penggaris besi stainless tahan karat dengan ukuran presisi cm dan inch. Tepi penggaris halus dan tidak tajam sehingga aman untuk anak sekolah.",
			Price:       12000,
			Image:       "https://images.unsplash.com/photo-1501504905252-473c47e087f8?w=500&auto=format&fit=crop&q=60",
			Stock:       150,
			Category:    "Tulis-Menulis",
		},
		{
			ID:          "p8",
			Name:        "Sepatu Sekolah Hitam Warrior Classic",
			Description: "Sepatu sekolah hitam polos berbahan kanvas tebal dengan sol karet anti licin. Sesuai dengan standar peraturan seragam sekolah di Indonesia.",
			Price:       135000,
			Image:       "https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500&auto=format&fit=crop&q=60",
			Stock:       25,
			Category:    "Seragam & Aksesoris",
		},
	}

	for _, p := range products {
		db.Products[p.ID] = p
	}
}

func createSession(userID string, role string) string {
	sessMux.Lock()
	defer sessMux.Unlock()

	token := generateID()
	sessions[token] = Session{
		UserID:    userID,
		Role:      role,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	return token
}

func getSessionUser(token string) (User, error) {
	sessMux.RLock()
	sess, exists := sessions[token]
	sessMux.RUnlock()

	if !exists {
		return User{}, errors.New("invalid or expired session token")
	}

	if time.Now().After(sess.ExpiresAt) {
		sessMux.Lock()
		delete(sessions, token)
		sessMux.Unlock()
		return User{}, errors.New("session token expired")
	}

	dbMutex.RLock()
	user, exists := db.Users[sess.UserID]
	dbMutex.RUnlock()

	if !exists {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func destroySession(token string) {
	sessMux.Lock()
	delete(sessions, token)
	sessMux.Unlock()
}

func authRequired(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	user, err := getSessionUser(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Locals("user", user)
	return c.Next()
}

func adminRequired(c fiber.Ctx) error {
	userVal := c.Locals("user")
	if userVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access",
		})
	}

	user := userVal.(User)
	if user.Role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: Admin access required",
		})
	}

	return c.Next()
}
func main() {
	loadDatabase()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			log.Printf("Server Error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "healthy", "time": time.Now().Format(time.RFC3339)})
	})

	app.Post("/api/auth/register", func(c fiber.Ctx) error {
		type RegisterReq struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Address  string `json:"address"`
			Phone    string `json:"phone"`
		}

		var req RegisterReq
		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if req.Name == "" || req.Email == "" || req.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name, Email, and Password are required"})
		}

		dbMutex.Lock()
		defer dbMutex.Unlock()

		for _, u := range db.Users {
			if strings.EqualFold(u.Email, req.Email) {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already registered"})
			}
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to encrypt password"})
		}

		newUser := User{
			ID:       generateID(),
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashedPassword),
			Role:     "customer",
			Address:  req.Address,
			Phone:    req.Phone,
		}

		db.Users[newUser.ID] = newUser
		saveDatabaseLocked()

		token := createSession(newUser.ID, newUser.Role)

		userResponse := newUser
		userResponse.Password = ""

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User registered successfully",
			"token":   token,
			"user":    userResponse,
		})
	})

	app.Post("/api/auth/login", func(c fiber.Ctx) error {
		type LoginReq struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var req LoginReq
		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		dbMutex.RLock()
		var foundUser *User
		for _, u := range db.Users {
			if strings.EqualFold(u.Email, req.Email) {
				temp := u
				foundUser = &temp
				break
			}
		}
		dbMutex.RUnlock()

		if foundUser == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		token := createSession(foundUser.ID, foundUser.Role)

		userResponse := *foundUser
		userResponse.Password = ""

		return c.JSON(fiber.Map{
			"message": "Login successful",
			"token":   token,
			"user":    userResponse,
		})
	})

	app.Get("/api/auth/me", authRequired, func(c fiber.Ctx) error {
		user := c.Locals("user").(User)
		user.Password = ""
		return c.JSON(user)
	})

	app.Post("/api/auth/logout", func(c fiber.Ctx) error {
		token := c.Get("Authorization")
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}
		if token != "" {
			destroySession(token)
		}
		return c.JSON(fiber.Map{"message": "Logged out successfully"})
	})
	app.Get("/api/products", func(c fiber.Ctx) error {
		category := c.Query("category")
		search := c.Query("search")

		dbMutex.RLock()
		defer dbMutex.RUnlock()

		var result []Product = make([]Product, 0)
		for _, p := range db.Products {
			matchCategory := true
			if category != "" && !strings.EqualFold(p.Category, category) {
				matchCategory = false
			}

			matchSearch := true
			if search != "" {
				searchLower := strings.ToLower(search)
				nameLower := strings.ToLower(p.Name)
				descLower := strings.ToLower(p.Description)
				if !strings.Contains(nameLower, searchLower) && !strings.Contains(descLower, searchLower) {
					matchSearch = false
				}
			}

			if matchCategory && matchSearch {
				result = append(result, p)
			}
		}

		return c.JSON(result)
	})

	app.Get("/api/products/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		dbMutex.RLock()
		product, exists := db.Products[id]
		dbMutex.RUnlock()

		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}

		return c.JSON(product)
	})

	app.Post("/api/products", authRequired, adminRequired, func(c fiber.Ctx) error {
		var p Product
		if err := c.Bind().JSON(&p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product data"})
		}

		if p.Name == "" || p.Price <= 0 || p.Stock < 0 || p.Category == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name, Category, Stock, and Positive Price are required"})
		}

		p.ID = generateID()
		if p.Image == "" {
			p.Image = "https://images.unsplash.com/photo-1501504905252-473c47e087f8?w=500&auto=format&fit=crop&q=60"
		}

		dbMutex.Lock()
		db.Products[p.ID] = p
		saveDatabaseLocked()
		dbMutex.Unlock()

		return c.Status(fiber.StatusCreated).JSON(p)
	})

	app.Put("/api/products/:id", authRequired, adminRequired, func(c fiber.Ctx) error {
		id := c.Params("id")

		var req Product
		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product data"})
		}

		dbMutex.Lock()
		defer dbMutex.Unlock()

		p, exists := db.Products[id]
		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}

		if req.Name != "" {
			p.Name = req.Name
		}
		if req.Description != "" {
			p.Description = req.Description
		}
		if req.Price > 0 {
			p.Price = req.Price
		}
		if req.Image != "" {
			p.Image = req.Image
		}
		if req.Stock >= 0 {
			p.Stock = req.Stock
		}
		if req.Category != "" {
			p.Category = req.Category
		}

		db.Products[id] = p
		saveDatabaseLocked()

		return c.JSON(p)
	})

	app.Delete("/api/products/:id", authRequired, adminRequired, func(c fiber.Ctx) error {
		id := c.Params("id")

		dbMutex.Lock()
		defer dbMutex.Unlock()

		_, exists := db.Products[id]
		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}

		delete(db.Products, id)
		saveDatabaseLocked()

		return c.JSON(fiber.Map{"message": "Product deleted successfully"})
	})
	app.Post("/api/orders", authRequired, func(c fiber.Ctx) error {
		user := c.Locals("user").(User)

		type OrderReqItem struct {
			ProductID string `json:"product_id"`
			Quantity  int    `json:"quantity"`
		}

		type OrderReq struct {
			CustomerName    string         `json:"customer_name"`
			CustomerPhone   string         `json:"customer_phone"`
			CustomerAddress string         `json:"customer_address"`
			Items           []OrderReqItem `json:"items"`
		}

		var req OrderReq
		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if len(req.Items) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shopping cart is empty"})
		}

		if req.CustomerName == "" || req.CustomerPhone == "" || req.CustomerAddress == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Delivery details (Name, Phone, Address) are required"})
		}

		dbMutex.Lock()
		defer dbMutex.Unlock()
		var orderItems []OrderItem
		var totalAmount float64

		for _, item := range req.Items {
			p, exists := db.Products[item.ProductID]
			if !exists {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Product %s not found", item.ProductID)})
			}

			if item.Quantity <= 0 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid quantity for product %s", p.Name)})
			}

			if p.Stock < item.Quantity {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Stok produk '%s' tidak mencukupi (Tersisa %d)", p.Name, p.Stock)})
			}
			p.Stock -= item.Quantity
			db.Products[item.ProductID] = p

			orderItems = append(orderItems, OrderItem{
				ProductID:   p.ID,
				ProductName: p.Name,
				Price:       p.Price,
				Quantity:    item.Quantity,
			})

			totalAmount += p.Price * float64(item.Quantity)
		}

		newOrder := Order{
			ID:              generateID(),
			UserID:          user.ID,
			CustomerName:    req.CustomerName,
			CustomerEmail:   user.Email,
			CustomerPhone:   req.CustomerPhone,
			CustomerAddress: req.CustomerAddress,
			Items:           orderItems,
			TotalAmount:     totalAmount,
			Status:          "Pending",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		db.Orders[newOrder.ID] = newOrder
		saveDatabaseLocked()

		return c.Status(fiber.StatusCreated).JSON(newOrder)
	})

	app.Get("/api/orders", authRequired, func(c fiber.Ctx) error {
		user := c.Locals("user").(User)

		dbMutex.RLock()
		defer dbMutex.RUnlock()

		var result []Order = make([]Order, 0)
		for _, o := range db.Orders {
			if user.Role == "admin" || o.UserID == user.ID {
				result = append(result, o)
			}
		}
		for i := 0; i < len(result); i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].CreatedAt.Before(result[j].CreatedAt) {
					result[i], result[j] = result[j], result[i]
				}
			}
		}

		return c.JSON(result)
	})

	app.Get("/api/orders/:id", authRequired, func(c fiber.Ctx) error {
		id := c.Params("id")
		user := c.Locals("user").(User)

		dbMutex.RLock()
		order, exists := db.Orders[id]
		dbMutex.RUnlock()

		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
		}

		if user.Role != "admin" && order.UserID != user.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return c.JSON(order)
	})

	app.Put("/api/orders/:id/status", authRequired, adminRequired, func(c fiber.Ctx) error {
		id := c.Params("id")

		type StatusReq struct {
			Status string `json:"status"`
		}

		var req StatusReq
		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		validStatus := map[string]bool{
			"Pending":   true,
			"Paid":      true,
			"Shipped":   true,
			"Cancelled": true,
		}

		if !validStatus[req.Status] {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status value. Must be Pending, Paid, Shipped, or Cancelled."})
		}

		dbMutex.Lock()
		defer dbMutex.Unlock()

		order, exists := db.Orders[id]
		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
		}
		if req.Status == "Cancelled" && order.Status != "Cancelled" {
			for _, item := range order.Items {
				if p, ok := db.Products[item.ProductID]; ok {
					p.Stock += item.Quantity
					db.Products[item.ProductID] = p
				}
			}
		} else if order.Status == "Cancelled" && req.Status != "Cancelled" {
			for _, item := range order.Items {
				if p, ok := db.Products[item.ProductID]; ok {
					if p.Stock < item.Quantity {
						return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Gagal mengubah status: Stok '%s' tidak mencukupi untuk memulihkan pesanan", p.Name)})
					}
					p.Stock -= item.Quantity
					db.Products[item.ProductID] = p
				}
			}
		}

		order.Status = req.Status
		order.UpdatedAt = time.Now()
		db.Orders[id] = order
		saveDatabaseLocked()

		return c.JSON(order)
	})
	app.Get("/api/dashboard/stats", authRequired, adminRequired, func(c fiber.Ctx) error {
		dbMutex.RLock()
		defer dbMutex.RUnlock()

		totalProducts := len(db.Products)
		totalOrders := len(db.Orders)
		totalCustomers := 0
		for _, u := range db.Users {
			if u.Role == "customer" {
				totalCustomers++
			}
		}

		var totalRevenue float64 = 0
		var pendingSales float64 = 0
		var completedSales float64 = 0
		var cancelledSales float64 = 0

		for _, o := range db.Orders {
			if o.Status == "Paid" || o.Status == "Shipped" {
				totalRevenue += o.TotalAmount
				completedSales += o.TotalAmount
			} else if o.Status == "Pending" {
				pendingSales += o.TotalAmount
			} else if o.Status == "Cancelled" {
				cancelledSales += o.TotalAmount
			}
		}

		categoryStats := make(map[string]int)
		for _, p := range db.Products {
			categoryStats[p.Category]++
		}

		return c.JSON(fiber.Map{
			"total_products":  totalProducts,
			"total_orders":    totalOrders,
			"total_customers": totalCustomers,
			"total_revenue":   totalRevenue,
			"pending_sales":   pendingSales,
			"completed_sales": completedSales,
			"cancelled_sales": cancelledSales,
			"category_stats":  categoryStats,
		})
	})

	log.Println("FourMart School Supplies Backend starting on port 5000...")
	log.Fatal(app.Listen(":5000"))
}
