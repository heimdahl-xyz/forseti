package repositories

import (
	"fmt"
	"github.com/heimdahl-xyz/forseti/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	DB *sqlx.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s, due to %s", url, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping connection to %s due to %s", url, err)
	}
	return &PostgresRepository{
		DB: db,
	}, nil
}

func (p *PostgresRepository) ProcessTransfer(ft *types.FTMessage) error {
	query := `
                INSERT INTO transfers (
                        timestamp, 
					    from_address, 
					    from_owner, 
					    to_address, 
					    to_owner,
                        amount, 
					    token_address, 
					    symbol,
					    chain, 
					    network,
					    tx_hash,
					    decimals, 
					    position
                ) VALUES (
                          $1, 
                          $2, 
                          $3,
                          $4,
                          $5,
                          $6,
                          $7,
                          $8, 
                          $9, 
                          $10, 
                          $11, 
                          $12,
                          $13)`

	_, err := p.DB.Exec(query,
		ft.Timestamp,
		ft.FromAddress,
		ft.FromOwner,
		ft.ToAddress,
		ft.ToOwner,
		ft.Amount.String(),
		ft.TokenAddress,
		ft.Symbol,
		ft.Chain,
		ft.Network,
		ft.TxHash,
		ft.Decimals,
		ft.Position)
	if err != nil {
		return fmt.Errorf("failed to insert transfer message %s", err)
	}
	return nil
}
