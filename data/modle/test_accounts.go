package modle

import "time"

type TestAccounts struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement"`
	AccountNo string     `gorm:"column:account_no;unique;not null"`
	Property  int8       `gorm:"column:property;default:1"`
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

// TableName returns the table name for the TestAccounts model
func (t *TestAccounts) TableName() string {
	return "tb_test_accounts"
}
