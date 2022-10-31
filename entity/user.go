package entity

type User struct {
	ID       uint `gorm:"primaryKey;column:we"`
	Username string
	Age      int
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
}
