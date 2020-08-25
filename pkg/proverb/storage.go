package proverb

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stevedesilva/my-proverb-app/pkg/proverb/generated/silvade/table"
)

//go:generate jet -source=MySQL -host=localhost -port=3306 -user=root -password=rootpassword -dbname=silvade -path=generated

type Storage interface {
	Create(ctx context.Context, proverbData *Data) error
}

type mysqlStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &mysqlStorage{db: db}
}

func (m *mysqlStorage) Create(ctx context.Context, proverbData *Data) error {
	// context
	if err := ctx.Err(); err != nil {
		return err
	}

	stmt := table.Proverbs.INSERT(table.Proverbs.MutableColumns).MODEL(proverbData)

	if _, err := stmt.ExecContext(ctx, m.db); err != nil {
		return fmt.Errorf("insert proverb data failed %w", err)
	}

	return nil
}
