package data

// Item model
type Item struct {
	BaseModel
	Name     string `json:"name"`
	SKU      string `json:"sku"`
	Price    string `json:"price"`
	Approved bool   `json:"approved"`
	// CreatedBy User   `gorm:"ForeignKey:CreatedID"`
	CreatedID uint `json:"createdId" binding:"-"`
}

// CreatedBy gets the user who owns this item
func (item *Item) CreatedBy(db *DB) User {
	var user User
	db.Model(&item).Related(&user, "CreatedBy")
	return user
}
