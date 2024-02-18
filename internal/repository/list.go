package repository

import (
	"context"
	"github.com/Victor90001/prod/internal/entity"
	"github.com/Victor90001/prod/internal/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresListRepository struct {
	db *pgxpool.Pool
}

func NewPostgresListRepository(db *pgxpool.Pool) (interfaces.ListRepository, error) {
	return &PostgresListRepository{
		db: db,
	}, nil
}

func (r *PostgresListRepository) GetLists() ([]entity.List, error) {
	var lists []entity.List
	q := "select row_to_json(x) from (select * from Product) as x"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return lists, err
	}
	for rows.Next() {
		var list entity.List
		if err := rows.Scan(&list); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil

}

func (r *PostgresListRepository) InsertList(list entity.List) error {
	q := "INSERT INTO Product (groupID, name, price, code, prodDate, describe, size, country, addParam) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	if _, err := r.db.Exec(context.Background(), q, list.DealerId, list.Name, list.Price, list.Amount, list.CreatedAt, list.Info, list.Carrier, list.ContactPerson, list.Note); err != nil {
		return err
	}
	return nil
}

func (r *PostgresListRepository) UpdateList(list entity.List) error {
	q := "UPDATE Product SET (name, price, code, prodDate, describe, size, country, addParam) = ($1, $2, $3, $4, $5, $6, $7, $8) WHERE id=$9"
	if _, err := r.db.Exec(context.Background(), q, list.Name, list.Price, list.Amount, list.CreatedAt, list.Info, list.Carrier, list.ContactPerson, list.Note, list.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresListRepository) DeleteList(id int) error {
	q := "DELETE FROM Product WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}
