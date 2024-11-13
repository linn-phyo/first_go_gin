package domain

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type Users struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;column:id;" json:"id"`
	UserName    string    `gorm:"column:user_name;" json:"user_name"`
	Email       string    `gorm:"column:email;" json:"email"`
	Password    string    `gorm:"column:password;" json:"password"`
	CreatedDate time.Time `gorm:"column:createddate;" json:"createddate"`
	UpdatedDate time.Time `gorm:"column:updateddate;" json:"updateddate"`
}
