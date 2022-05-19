package psql

import (
	"time"

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

func (p passPsql) AddPassage(passage domain.AddPassage) error {
	query := `
	insert into passage(pass_id, passage_status, is_exit) VALUES
($1,$2,$3) returning passage_id
	`
	var passage_id int32
	err := p.Conn.QueryRow(query, passage.PassID, passage.Status, passage.Exit).Scan(&passage_id)
	if err != nil {
		return err
	}
	return nil
}

func (p passPsql) AllPassages() ([]domain.Passage, error) {
	var passages []domain.Passage

	query := `
	select passage_datetime,is_exit,account_fullname from passage join pass p on p.pass_id = passage.pass_id join account a on p.pass_account_id = a.account_id
	`

	rows, err := p.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.Passage{}
		err := rows.Scan(&obj.Time, &obj.Exit, &obj.Fullname)
		if err != nil {
			return nil, err
		}
		passages = append(passages, obj)
	}
	return passages, nil
}
func (p passPsql) UpdatePassTime(data time.Time, passID int32) error {
	query := `UPDATE pass
				SET pass_expiration_date = $1
				WHERE pass_id = $2 returning pass_id`

	var passIDCheck int32
	err := p.Conn.QueryRow(query, data, passID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}

func (p passPsql) CheckPassByData(data string) (*domain.CheckResult, error) {
	checkRes := new(domain.CheckResult)

	query := `
	select pass_id, pass_expiration_date, pass_issue_date, 
	account_role_id, pass_name, user_email, user_created_date, 
	account_fullname from pass join account ON account.account_id=pass.pass_account_id 
	join users u on account.account_user_id = u.user_id 
	where pass_secure_data=$1;
	`

	err := p.Conn.QueryRow(query, data).Scan(&checkRes.PassID,
		&checkRes.ExpirationDate, &checkRes.IssueDate, &checkRes.UserRoleID,
		&checkRes.PassName, &checkRes.UserEmail, &checkRes.UserCreated, &checkRes.UserName)
	if err != nil {
		return nil, err
	}
	return checkRes, nil
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

func (p passPsql) FindPassByID(passID int32) (*domain.Pass, error) {
	obj := new(domain.Pass)

	query := `select * from pass where pass_id = $1`

	err := p.Conn.QueryRow(query, passID).Scan(&obj.PassID, &obj.AccountID,
		&obj.DynamicQR, &obj.ExpirationDate, &obj.IssueDate,
		&obj.Name, &obj.SecureData, &obj.Active, &obj.Disabled)
	if err != nil {
		return nil, err
	}

	return obj, nil
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
