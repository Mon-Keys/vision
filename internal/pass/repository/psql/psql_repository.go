package psql

import (
	"github.com/jackc/pgx"
	"github.com/patrickmn/go-cache"
	"github.com/perlinleo/vision/internal/domain"
)

type passPsql struct {
	Conn  *pgx.ConnPool
	Cache *cache.Cache
}

func NewPassPSQLRepository(ConnectionPool *pgx.ConnPool, Cache *cache.Cache) domain.PassRepository {
	return &passPsql{
		ConnectionPool,
		Cache,
	}
}

func (p passPsql) PassesByAccountID(accountID int32) ([]domain.Pass, error) {
	var passes []domain.Pass

	query := `select * from pass where pass_account_id = $1 and pass_disabled = false`

	rows, err := p.Conn.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.Pass{}
		err := rows.Scan(&obj.PassID, &obj.AccountID,
			&obj.DynamicQR, &obj.ExpirationDate, &obj.IssueDate,
			&obj.Name, &obj.SecureData, &obj.Active, &obj.Disabled)
		if err != nil {
			return nil, err
		}
		passes = append(passes, obj)
	}
	return passes, nil
}

func (p passPsql) CreatePass(pass domain.Pass) (error, int32) {
	query := `insert into pass 
    (pass_account_id, pass_dynamic_qr, pass_expiration_date, 
     pass_name, pass_secure_data, pass_active) 
     VALUES (
             $1,$2,$3,$4,$5,$6
            ) RETURNING pass_id`

	var passID int32
	err := p.Conn.QueryRow(query, pass.AccountID, pass.DynamicQR, pass.ExpirationDate, pass.Name, pass.SecureData, pass.Active).Scan(&passID)
	if err != nil {
		return err, -1
	}

	return nil, passID
}

func (p passPsql) ActivatePass(passID int32) error {
	query := `UPDATE pass
				SET pass_active = true
				WHERE pass_id = $1 returning pass_id`

	var passIDCheck int32
	err := p.Conn.QueryRow(query, passID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}

func (p passPsql) DisablePass(passID int32) error {
	query := `UPDATE pass
				SET pass_disabled = true
				WHERE pass_id = $1 returning pass_id`

	var passIDCheck int32
	err := p.Conn.QueryRow(query, passID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}
func (p passPsql) DeletePass(passID int32) error {
	query := `DELETE from pass
			where pass_id = $1  returning pass_id`
	var passIDCheck int32
	err := p.Conn.QueryRow(query, passID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil

}
