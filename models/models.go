package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Nickname  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Phone     string `gorm:"not null"`
	Position  string
	Graduated string
	Country   string
	City      string
	Bio       string
	AboutMe   string
	Profile   UserProfile `gorm:"foreignKey:UserID"`
	Hobbies   []Hobby     `gorm:"foreignKey:UserID"`
	Links     []Link      `gorm:"foreignKey:UserID"`
}

type UserProfile struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Company   string
}

type Hobby struct {
	ID     uint   `gorm:"primaryKey"`
	Hobby  string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}

type Link struct {
	ID     uint   `gorm:"primaryKey"`
	Link   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
