package do

import (
	"fmt"
	"strings"
	"github.com/watchman1989/steez/comm"
	"context"
)


const (
	recursiveQueryFmt = `WITH RECURSIVE transfer_chain AS (
    SELECT
        tr.id,
        tr.transfer_no,
        tr.from_account_id,
        tr.to_account_id,
        a1.account_no AS from_account_no,
        a2.account_no AS to_account_no,
        1 AS level,
        CAST(CONCAT(a1.account_no, '->', a2.account_no) AS CHAR(1000)) AS path
    FROM tb_test_transfer_records tr
    JOIN tb_test_accounts a1 ON tr.from_account_id = a1.id
    JOIN tb_test_accounts a2 ON tr.to_account_id = a2.id
    WHERE a1.account_no = 'ACCOUNT_NO'
      AND tr.deleted_at IS NULL
      AND a1.deleted_at IS NULL
      AND a2.deleted_at IS NULL

    UNION ALL

    SELECT
        tr.id,
        tr.transfer_no,
        tr.from_account_id,
        tr.to_account_id,
        a1.account_no AS from_account_no,
        a2.account_no AS to_account_no,
        tc.level + 1 AS level,
        CAST(CONCAT(tc.path, '->', a2.account_no) AS CHAR(1000)) AS path
    FROM tb_test_transfer_records tr
    JOIN tb_test_accounts a1 ON tr.from_account_id = a1.id
    JOIN tb_test_accounts a2 ON tr.to_account_id = a2.id
    JOIN transfer_chain tc ON tr.from_account_id = tc.to_account_id
    WHERE tc.level < LEVEL
      AND tr.deleted_at IS NULL
      AND a1.deleted_at IS NULL
      AND a2.deleted_at IS NULL
)
SELECT
    level,
    from_account_no,
    to_account_no,
    transfer_no,
    path
FROM transfer_chain
ORDER BY level, transfer_no;`
    
)

type TransferRecord struct {
	Level int
	FromAccountNo string
	ToAccountNo string
	TransferNo string
	Path string
}

func RecursiveQuery(ctx context.Context, accountNo string, level int) (records []TransferRecord, err error) {
	sql := strings.Replace(recursiveQueryFmt, "ACCOUNT_NO", accountNo, 1)
	sql = strings.Replace(sql, "LEVEL", fmt.Sprintf("%d", level), 1)
	records = make([]TransferRecord, 0)
	rows, err := comm.GContext.Mysql.Raw(sql).Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record TransferRecord
		err = rows.Scan(&record.Level, &record.FromAccountNo, &record.ToAccountNo, &record.TransferNo, &record.Path)
		if err != nil {
			return
		}
		records = append(records, record)
	}
	return
}