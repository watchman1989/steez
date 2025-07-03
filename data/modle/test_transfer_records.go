package modle

import "time"

type TestTransferRecords struct {
	ID            int64      `gorm:"column:id;primaryKey;autoIncrement"`
	TransferNo    string     `gorm:"column:transfer_no;unique;not null"`
	FromAccountID int64      `gorm:"column:from_account_id;not null"`
	ToAccountID   int64      `gorm:"column:to_account_id;not null"`
	CreatedAt     time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:"column:deleted_at"`
}

// TableName returns the table name for the TestTransferRecords model
func (t *TestTransferRecords) TableName() string {
	return "tb_test_transfer_records"
}
