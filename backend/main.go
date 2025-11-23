package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db, _ = gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	db.AutoMigrate(&User{}, &Item{}, &Cart{}, &Order{})

	// default items
	db.Create(&Item{Name: "T-Shirt", Price: 299})
	db.Create(&Item{Name: "Cap", Price: 199})

	r := gin.Default()

	r.POST("/users", signup)
	r.POST("/users/login", login)
	r.GET("/items", getItems)
	r.POST("/carts", addToCart)
	r.GET("/carts", listCarts)
	r.POST("/orders", placeOrder)
	r.GET("/orders", listOrders)

	r.Run(":8080")
}

func signup(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	db.Create(&u)
	c.JSON(200, gin.H{"msg": "user created"})
}

func login(c *gin.Context) {
	var body User
	c.BindJSON(&body)

	var u User
	err := db.Where("username=? AND password=?", body.Username, body.Password).First(&u).Error
	if err != nil {
		c.JSON(400, gin.H{"msg": "Invalid username/password"})
		return
	}

	u.Token = "token123"
	db.Save(&u)

	c.JSON(200, gin.H{
		"token":  u.Token,
		"userID": u.ID,
	})
}

func getItems(c *gin.Context) {
	var items []Item
	db.Find(&items)
	c.JSON(200, items)
}

func addToCart(c *gin.Context) {
	var ct Cart
	c.BindJSON(&ct)
	db.Create(&ct)
	c.JSON(200, gin.H{"msg": "added to cart"})
}

func listCarts(c *gin.Context) {
	var ct []Cart
	db.Find(&ct)
	c.JSON(200, ct)
}

func placeOrder(c *gin.Context) {
	var order Order
	c.BindJSON(&order)
	db.Create(&order)
	c.JSON(200, gin.H{"msg": "Order successful"})
}

func listOrders(c *gin.Context) {
	var o []Order
	db.Find(&o)
	c.JSON(200, o)
}
