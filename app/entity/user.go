package entity

type Product struct {
	UserId       string `gorm:"not null;" json:"user_id"`
	ProductId    string `gorm:"size:128;not null;uniqueIndex;primary_key" json:"product_id"`
	ProductName  string `gorm:"size:128;not null" json:"product_name"`
	ProductStock int    `gorm:"not null" json:"product_stoct"`
	ProductPrice int    `gorm:"not null" json:"product_price"`
	CreatedDate  string `gorm:"size:128;null" json:"created_date"`
	LastUpdate   string `gorm:"size:128;null" json:"last_update"`
}

type User struct {
	UserId       string `gorm:"size:128;not null;uniqueIndex;primary_key" json:"user_id"`
	UserEmail    string `gorm:"size:128;not null;uniqueIndex" json:"user_email"`
	UserName     string `gorm:"size:128;not null" json:"user_name"`
	UserImage    string `gorm:"size:256;not null" json:"user_image"`
	UserPassword string `gorm:"size:128;not null" json:"user_password"`
	UserSession  string `gorm:"size:256;null" json:"user_session"`
	UserStatus   StatusUser
	UserStatusId int8      `gorm:"not null;" json:"user_status"`
	CreatedDate  string    `gorm:"size:128;null" json:"created_date"`
	LastUpdate   string    `gorm:"size:128;null" json:"last_update"`
	Products     []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StatusUser struct {
	Id     int8   `gorm:"not null;uniqueIndex;primaryKey"`
	Status string `gorm:"not null;unique"`
}
