package entity

type User struct {
	ID       uint `gorm:"primaryKey;column:uid"`
	Username string
	Age      int
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
}
