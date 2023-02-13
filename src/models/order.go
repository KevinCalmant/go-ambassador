package models

type Order struct {
	Model
	TransactionId   string      `json:"trasanctionId" gorm:"null"`
	UserId          uint        `json:"userId"`
	Code            string      `json:"code"`
	AmbassadorEmail string      `json:"ambassadorEmail"`
	FirstName       string      `json:"firstname"`
	LastName        string      `json:"lastname"`
	Email           string      `json:"email"`
	Address         string      `json:"address" gorm:"null"`
	City            string      `json:"city" gorm:"null"`
	Country         string      `json:"country" gorm:"null"`
	Zip             string      `json:"zip" gorm:"null"`
	Complete        bool        `json:"complete" gorm:"default:false"`
	OrderItem       []OrderItem `json:"orderItem" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Model
	OrderId           string  `json:"orderId"`
	ProductTitle      string  `json:"productTitle"`
	Price             float64 `json:"price"`
	Quantity          uint    `json:"quantity"`
	AdminRevenue      float64 `json:"adminRevenue"`
	AmbassadorRevenue float64 `json:"ambassadorRevenu"`
}
