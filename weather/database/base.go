package database

import "database/sql"

type BaseClient struct {
	DB *sql.DB
}

func (d *BaseClient) Close() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}
