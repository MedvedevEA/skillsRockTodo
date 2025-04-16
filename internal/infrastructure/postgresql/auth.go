package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/pkg/servererrors"
)

const (
	//GetAccessPermissionsQuery =``
	LoginQuery = `SELECT password FROM user WHERE "name"=$1`
)

func (p *PostgreSql) GetAccessPermissions() []*entity.AccessPermission {
	return nil
}
func (p *PostgreSql) Login(userName string) (string, error) {
	const op = "postgresql.Login"
	var password string
	err := p.pool.QueryRow(context.Background(), LoginQuery, userName).Scan(&password)
	if errors.Is(err, sql.ErrNoRows) {
		return "", servererrors.RecordNotFound
	}
	if err != nil {
		return "", servererrors.InternalServerError
	}
	return password, nil

}
