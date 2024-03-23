package database

import "database/sql"

type BaseClient struct {
	DB *sql.DB
}

func (b *BaseClient) Close() error {
	if b.DB != nil {
		return b.DB.Close()
	}
	return nil
}
