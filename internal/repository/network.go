package repository

import (
	"context"
	"github.com/Victor90001/prod/internal/entity"
	"github.com/Victor90001/prod/internal/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresNetworkRepository struct {
	db *pgxpool.Pool
}

func NewPostgresNetworkRepository(db *pgxpool.Pool) (interfaces.NetworkRepository, error) {
	return &PostgresNetworkRepository{
		db: db,
	}, nil
}

func (r *PostgresNetworkRepository) GetNetworks() ([]entity.Network, error) {
	var networks []entity.Network
	q := "select row_to_json(x) from (select * from Section) as x"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return networks, err
	}
	defer rows.Close()
	for rows.Next() {
		var network entity.Network
		if err := rows.Scan(&network); err != nil {
			return nil, err
		}
		networks = append(networks, network)
	}
	return networks, nil
}

func (r *PostgresNetworkRepository) InsertNetwork(network entity.Network) error {
	q := "INSERT INTO Section (name) VALUES ($1)"
	if _, err := r.db.Exec(context.Background(), q, network.Name); err != nil {
		return err
	}
	return nil
}

func (r *PostgresNetworkRepository) UpdateNetwork(network entity.Network) error {
	q := "UPDATE Section SET name=$1 WHERE id=$2"
	if _, err := r.db.Exec(context.Background(), q, network.Name, network.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresNetworkRepository) DeleteNetwork(id int) error {
	q := "DELETE FROM Section WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}
