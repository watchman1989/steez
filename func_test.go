package main

import (
	"testing"
	"github.com/watchman1989/steez/comm"
	"github.com/watchman1989/steez/data/modle"
	"math/rand"
	"fmt"
)


func TestMysql(t *testing.T) {
	db, err := comm.InitMysqlClient(comm.MysqlConfig{
		Host: "127.0.0.1",
		Port: 3306,
		User: "local_test",
		Password: "Mysql?3.1415926",
		Dbname: "steez",
	})
	if err != nil {
		t.Fatalf("init mysql client failed: %s", err.Error())
	}

	var accounts []modle.TestAccounts
	result := db.Table("tb_test_accounts").Find(&accounts)
	if result.Error != nil {
		t.Fatalf("query accounts failed: %s", result.Error.Error())
		return
	}

	ids := make([]int64, 0)
	for _, account := range accounts {
		ids = append(ids, account.ID)
	}

	for i := 0; i < 300; i++ {
		fromIdx := rand.Intn(len(ids) - 1)
		toIdx := rand.Intn(len(ids) - 1) 

		if fromIdx == toIdx {
			continue
		}

		transferRecord := modle.TestTransferRecords{
			TransferNo:    comm.GetSha256([]byte(fmt.Sprintf("%d-%d", ids[fromIdx], ids[toIdx]))),
			FromAccountID: ids[fromIdx],
			ToAccountID:   ids[toIdx],
		}
		result = db.Create(&transferRecord)
		if result.Error != nil {
			t.Fatalf("create transfer record failed: %s", result.Error.Error())
			return
		}
		t.Logf("%d transferRecord: %+v\n", i, transferRecord)
	}
}