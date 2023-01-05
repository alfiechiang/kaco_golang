package model

type Product struct {
	Name        string `gorm:"type:varchar(30)" json:"name"`
	Price       int    `gorm:"type:int(11)" json:"price"`
	Category    string `gorm:"type:varchar(30)" json:"category"`
	ChCategory  string `gorm:"type:varchar(30)" json:"ch_category"`
	CreatedAt   string `gorm:"type:varchar(30)" json:"created_at,omitempty"`
	HotSpot     string `gorm:"type:varchar(30)" json:"hot_spot,omitempty"`
	Description string `gorm:"type:varchar(30)" json:"description,omitempty"`
}


func (Product) TableName() string {
	return "products"
}
