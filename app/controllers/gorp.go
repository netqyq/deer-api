package controllers

import (
	"database/sql"

	"github.com/netqyq/deer-api/app/models"

	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"

	"github.com/revel/modules/db/app"
	r "github.com/revel/revel"
)

var (
	Dbm *gorp.DbMap
)

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func InitDB1() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.User{}).SetKeys(true, "UserId")
	t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Email": 60,
		"Name":  100,
	})

	t1 := Dbm.AddTable(models.Product{}).SetKeys(true, "Id")
	setColumnSizes(t1, map[string]int{
		"Name": 200,
	})

	err := Dbm.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
