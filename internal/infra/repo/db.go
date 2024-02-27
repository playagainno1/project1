package repo

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

const (
	dbDialet = "mysql"
	prepared = true
)

type (
	selectorFunc func() *goqu.SelectDataset
	inserterFunc func() *goqu.InsertDataset
	updaterFunc  func() *goqu.UpdateDataset
	deleterFunc  func() *goqu.DeleteDataset
)

type operators struct {
	selector selectorFunc
	inserter inserterFunc
	updater  updaterFunc
	deleter  deleterFunc
}

func newOperators(db *sqlx.DB, table string) *operators {
	return &operators{
		selector: func() *goqu.SelectDataset {
			return selector(db, table)
		},
		inserter: func() *goqu.InsertDataset {
			return inserter(db, table)
		},
		updater: func() *goqu.UpdateDataset {
			return updater(db, table)
		},
		deleter: func() *goqu.DeleteDataset {
			return deleter(db, table)
		},
	}
}

func selector(db *sqlx.DB, table string) *goqu.SelectDataset {
	return goqu.New(dbDialet, db).From(table).Prepared(prepared)
}

func inserter(db *sqlx.DB, table string) *goqu.InsertDataset {
	return goqu.New(dbDialet, db).Insert(table).Prepared(prepared)
}

func updater(db *sqlx.DB, table string) *goqu.UpdateDataset {
	return goqu.New(dbDialet, db).Update(table).Prepared(prepared)
}

func deleter(db *sqlx.DB, table string) *goqu.DeleteDataset {
	return goqu.New(dbDialet, db).Delete(table).Prepared(prepared)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func intToBool(i int) bool {
	if i != 0 {
		return true
	}
	return false
}
