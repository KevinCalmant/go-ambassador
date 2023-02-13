package models

type Order struct {
	Model
	TransactionId   string      `json:"trasanctionId" gorm:"null"`
	UserId          uint        `json:"userId"`
	Code            string      `json:"code"`
	AmbassadorEmail string      `json:"ambassadorEmail"`
	FirstName       string      `json:"-"`
	LastName        string      `json:"-"`
	Name            string      `json:"name" gorm:"-"`
	Email           string      `json:"email"`
	Address         string      `json:"address" gorm:"null"`
	City            string      `json:"city" gorm:"null"`
	Country         string      `json:"country" gorm:"null"`
	Zip             string      `json:"zip" gorm:"null"`
	Complete        bool        `json:"-" gorm:"default:false"`
	Total           float64     `json:"total" gorm:"-""`
	OrderItems      []OrderItem `json:"orderItems" gorm:"foreignKey:OrderId"`
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

func (order *Order) FullName() string {
	return order.FirstName + " " + order.LastName
}

func (order *Order) GetTotal() float64 {
	var total float64 = 0
	for _, orderItem := range order.OrderItems {
		total += orderItem.Price * float64(orderItem.Quantity)
	}
	return total
}
