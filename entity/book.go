package entity

type Book struct {
	ID          uint64 `gorm:"primari_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"uniqueIndex;type:varchar(255)" json:"description"`
	UserId      uint64 `gorm:"not null" json:"-"`
	User        string `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
