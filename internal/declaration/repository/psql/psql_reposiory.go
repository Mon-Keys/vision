package psql

import (
	"github.com/jackc/pgx"
	"github.com/patrickmn/go-cache"
	"github.com/perlinleo/vision/internal/domain"
)

type declarationPsql struct {
	Conn  *pgx.ConnPool
	Cache *cache.Cache
}

func NewDeclarationPSQLRepository(ConnectionPool *pgx.ConnPool, Cache *cache.Cache) domain.DeclarationRepository {
	return &declarationPsql{
		ConnectionPool,
		Cache,
	}
}

func (u declarationPsql) RoleDeclarationByID(id int32) (*domain.AskRoleDeclaration, error) {
	declaration := new(domain.AskRoleDeclaration)

	query := `select role_request_account_id, role_request_comment, 
	role_request_wanted_role_id, account_role_id from role_request 
	join account a on a.account_id = role_request.role_request_account_id where role_request_id=$1`

	err := u.Conn.QueryRow(
		query, id).Scan(&declaration.CreatorID, &declaration.Comment, &declaration.RoleID, &declaration.CurrentRole)
	if err != nil {
		return nil, err
	}

	return declaration, nil
}

func (u declarationPsql) TimeDeclarationByID(id int32) (*domain.AskTimeDeclaration, error) {
	declaration := new(domain.AskTimeDeclaration)

	query := `select time_request_account_id, time_request_comment, 
	time_request_pass_new_expiration_date, time_request_pass_id from time_request 
	join account a on a.account_id = time_request.time_request_account_id where time_request_id=$1`

	err := u.Conn.QueryRow(
		query, id).Scan(&declaration.CreatorID, &declaration.Comment, &declaration.TimeExtended, &declaration.PassID)
	if err != nil {
		return nil, err
	}

	return declaration, nil
}

func (u declarationPsql) CreateRoleDeclaration(declaration domain.AskRoleDeclaration) error {
	query := `insert into role_request 
	(role_request_account_id, role_request_wanted_role_id, role_request_approved, role_request_comment)
 	VALUES
 	($1,$2,false,$3)  RETURNING role_request_id`

	var roleReqID int32

	err := u.Conn.QueryRow(
		query, declaration.CreatorID, declaration.RoleID, declaration.Comment).Scan(&roleReqID)
	if err != nil {
		return err
	}

	return nil
}
func (u declarationPsql) CreatePassDeclaration(declaration domain.AskPassDeclaration) error {
	query := `insert into pass_request (pass_request_account_id, pass_request_pass_id, 
		pass_request_approved, pass_request_comment) VALUES ($1,$2,$3,$4) RETURNING pass_request_id`

	var passReqID int32

	err := u.Conn.QueryRow(query, declaration.CreatorID, declaration.NewPassID, declaration.Approved, declaration.Comment).Scan(&passReqID)

	if err != nil {
		return err
	}

	return nil
}
func (u declarationPsql) CreateTimeDeclaration(declaration domain.AskTimeDeclaration) error {
	query := `insert into time_request (time_request_account_id, time_request_pass_id, 
		time_request_approved, time_request_comment, time_request_pass_new_expiration_date) VALUES ($1,$2,$3,$4,$5) RETURNING time_request_id`

	var passReqID int32

	err := u.Conn.QueryRow(query, declaration.CreatorID, declaration.PassID, false, declaration.Comment, declaration.TimeExtended).Scan(&passReqID)

	if err != nil {
		return err
	}

	return nil
}
func (u declarationPsql) AllDeclarations() ([]domain.DeclarationCommon, error) {
	var declarations []domain.DeclarationCommon
	passDeclarations, err := u.PassRequestDeclarationsAll()
	if err != nil {
		return nil, err
	}
	roleDeclarations, err := u.RoleRequestDeclarationsAll()
	if err != nil {
		return nil, err
	}

	timeDeclarations, err := u.TimeRequestDeclarationsAll()
	if err != nil {
		return nil, err
	}

	declarations = append(declarations, passDeclarations...)
	declarations = append(declarations, roleDeclarations...)
	declarations = append(declarations, timeDeclarations...)

	return declarations, nil
}

func (u declarationPsql) PassRequestDeclarationsAll() ([]domain.DeclarationCommon, error) {
	var declarations []domain.DeclarationCommon

	query := `select pass_request_id, pass_request_account_id,pass_request_created,account_fullname,
		pass_request_approved, pass_request_denied
		 from pass_request join account a on a.account_id = pass_request.pass_request_account_id`

	rows, err := u.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.DeclarationCommon{}
		obj.DeclarationType = 0
		err := rows.Scan(&obj.DeclarationInnerID, &obj.AuthorID,
			&obj.CreatedDate, &obj.AuthorName, &obj.Accepted, &obj.Denied)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, obj)
	}
	return declarations, nil
}

func (u declarationPsql) RoleRequestDeclarationsAll() ([]domain.DeclarationCommon, error) {
	var declarations []domain.DeclarationCommon

	query := `select role_request_id, role_request_account_id,role_request_created,account_fullname,
	role_request_approved, role_request_denied
		 from role_request join account a on a.account_id = role_request.role_request_account_id`

	rows, err := u.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.DeclarationCommon{}
		obj.DeclarationType = 2
		err := rows.Scan(&obj.DeclarationInnerID, &obj.AuthorID,
			&obj.CreatedDate, &obj.AuthorName, &obj.Accepted, &obj.Denied)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, obj)
	}
	return declarations, nil
}

func (u declarationPsql) AllDeclarationsByAccountID(accountID int32) ([]domain.DeclarationCommon, error) {
	var declarations []domain.DeclarationCommon
	passDeclarations, err := u.PassRequestDeclarationsByAccountID(accountID)
	if err != nil {
		return nil, err
	}
	roleDeclarations, err := u.RoleRequestDeclarationsByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	timeDeclarations, err := u.TimeRequestDeclarationsByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	declarations = append(declarations, passDeclarations...)
	declarations = append(declarations, roleDeclarations...)
	declarations = append(declarations, timeDeclarations...)

	return declarations, nil
}

func (u declarationPsql) TimeRequestDeclarationsAll() ([]domain.DeclarationCommon, error) {
	var declarations []domain.DeclarationCommon

	query := `select time_request_id, time_request_account_id,time_request_created,account_fullname,
	time_request_approved, time_request_denied
		 from time_request join account a on a.account_id = time_request.time_request_account_id`

	rows, err := u.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.DeclarationCommon{}
		obj.DeclarationType = 1
		err := rows.Scan(&obj.DeclarationInnerID, &obj.AuthorID, &obj.CreatedDate, &obj.AuthorName, &obj.Accepted, &obj.Denied)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, obj)
	}
	return declarations, nil
}

func (u declarationPsql) PassRequestDeclarationsByAccountID(accountID int32) ([]domain.DeclarationCommon, error) {
	var declarations []domain.DeclarationCommon

	query := `select pass_request_id, pass_request_account_id,pass_request_created,account_fullname,
		pass_request_approved, pass_request_denied
		 from pass_request join account a on a.account_id = pass_request.pass_request_account_id where pass_request_account_id=$1`

	rows, err := u.Conn.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.DeclarationCommon{}
		obj.DeclarationType = 0
		err := rows.Scan(&obj.DeclarationInnerID, &obj.AuthorID, &obj.CreatedDate,
			&obj.AuthorName, &obj.Accepted, &obj.Denied)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, obj)
	}
	return declarations, nil
}
func (u declarationPsql) RoleRequestDeclarationsByAccountID(accountID int32) ([]domain.DeclarationCommon, error) {
	var declarations []domain.DeclarationCommon

	query := `select role_request_id, role_request_account_id,role_request_created,account_fullname,
	role_request_approved, role_request_denied
		 from role_request join account a on a.account_id = role_request.role_request_account_id where role_request_account_id=$1`

	rows, err := u.Conn.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.DeclarationCommon{}
		obj.DeclarationType = 2
		err := rows.Scan(&obj.DeclarationInnerID, &obj.AuthorID, &obj.CreatedDate, &obj.AuthorName, &obj.Accepted, &obj.Denied)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, obj)
	}
	return declarations, nil
}
func (u declarationPsql) TimeRequestDeclarationsByAccountID(accountID int32) ([]domain.DeclarationCommon, error) {
	var declarations []domain.DeclarationCommon

	query := `select time_request_id, time_request_account_id,time_request_created,account_fullname,
	time_request_approved, time_request_denied
		 from time_request join account a on a.account_id = time_request.time_request_account_id where time_request_account_id=$1`

	rows, err := u.Conn.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.DeclarationCommon{}
		obj.DeclarationType = 1
		err := rows.Scan(&obj.DeclarationInnerID, &obj.AuthorID, &obj.CreatedDate, &obj.AuthorName, &obj.Accepted, &obj.Denied)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, obj)
	}
	return declarations, nil
}

func (u declarationPsql) AcceptRoleDeclaration(RoleDeclarationID int32) error {
	query := `UPDATE role_request
				SET role_request_approved = true
				WHERE role_request_id = $1 returning role_request_id`

	var passIDCheck int32
	err := u.Conn.QueryRow(query, RoleDeclarationID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}
func (u declarationPsql) AcceptPassDeclaration(PassDeclarationID int32) error {
	query := `UPDATE pass_request
				SET pass_request_approved = true
				WHERE pass_request_id = $1 returning pass_request_id`

	var passIDCheck int32
	err := u.Conn.QueryRow(query, PassDeclarationID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}
func (u declarationPsql) AcceptTimeDeclaration(TimeDeclarationID int32) error {
	query := `UPDATE time_request
				SET time_request_approved = true
				WHERE time_request_id = $1 returning time_request_id`

	var passIDCheck int32
	err := u.Conn.QueryRow(query, TimeDeclarationID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}

func (u declarationPsql) DenyRoleDeclaration(RoleDeclarationID int32) error {
	query := `UPDATE role_request
				SET role_request_denied = true
				WHERE role_request_id = $1 returning role_request_id`

	var passIDCheck int32
	err := u.Conn.QueryRow(query, RoleDeclarationID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}
func (u declarationPsql) DenyPassDeclaration(PassDeclarationID int32) error {
	query := `UPDATE pass_request
				SET pass_request_denied = true
				WHERE pass_request_id = $1 returning pass_request_id`

	var passIDCheck int32
	err := u.Conn.QueryRow(query, PassDeclarationID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}
func (u declarationPsql) DenyTimeDeclaration(TimeDeclarationID int32) error {
	query := `UPDATE time_request
				SET time_request_denied = true
				WHERE time_request_id = $1 returning time_request_id`

	var passIDCheck int32
	err := u.Conn.QueryRow(query, TimeDeclarationID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
}

func (u declarationPsql) PassRequestDeclarationByID(AskPassDeclarationID int32) (*domain.AskPassDeclaration, error) {
	askPassDec := new(domain.AskPassDeclaration)

	query := `select pass_request_account_id, pass_request_pass_id, pass_request_approved,pass_request_denied, pass_request_comment
	 from pass_request where pass_request_id = $1`
	err := u.Conn.QueryRow(query, AskPassDeclarationID).Scan(&askPassDec.CreatorID, &askPassDec.NewPassID, &askPassDec.Approved, &askPassDec.Denied, &askPassDec.Comment)
	if err != nil {
		return nil, err
	}
	return askPassDec, nil
}
