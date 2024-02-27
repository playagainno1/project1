package migration

import "github.com/jmoiron/sqlx"

func Migrate(db *sqlx.DB) error {
	for _, sql := range SQLs {
		_, err := db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

var (
	SQLs = []string{}
)
