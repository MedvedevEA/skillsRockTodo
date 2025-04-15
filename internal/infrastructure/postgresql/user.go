package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/repository/dto"
	"skillsRockTodo/pkg/servererrors"

	"github.com/google/uuid"
)

const (
	addUserQuery    = `INSERT INTO "user" (name, password) VALUES ($1, $2) RETURNING *;`
	getUserQuery    = `SELECT * FROM "user" WHERE user_id=$1`
	getUsersQuery   = `SELECT * FROM "user" OFFSET $1 LIMIT $2`
	updateUserQuery = `
		UPDATE "user" SET 
		name = CASE WHEN $2::character varying IS NULL THEN name ELSE $2 END,
		password = CASE WHEN $3::character varying IS NULL THEN password ELSE $3 END,
		update_at = now()
		WHERE user_id=$1
		RETURNING *`
	removeUserQuery = `DELETE FROM "user" WHERE user_id=$1 RETURNING user_id`
)

func (p *PostgreSql) AddUser(dto *dto.AddUser) (*entity.User, error) {
	const op = "postgresql.AddUser"
	user := new(entity.User)
	err := p.pool.QueryRow(context.Background(), addUserQuery, dto.Name, dto.Password).Scan(&user.UserId, &user.Name, &user.Password, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		p.lg.Error("failed to add User", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.InternalServerError
	}
	return user, nil
}

func (p *PostgreSql) GetUser(userId *uuid.UUID) (*entity.User, error) {
	const op = "postgresql.GetUser"
	user := new(entity.User)
	err := p.pool.QueryRow(context.Background(), getUserQuery, userId).Scan(&user.UserId, &user.Name, &user.Password, &user.CreateAt, &user.UpdateAt)
	if errors.Is(err, sql.ErrNoRows) {
		p.lg.Error("failed to get User", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.RecordNotFound
	}
	if err != nil {
		p.lg.Error("failed to get User", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.InternalServerError
	}
	return user, nil
}

func (p *PostgreSql) GetUsers(dto *dto.GetUsers) ([]*entity.User, error) {
	const op = "postgresql.GetUsers"
	rows, err := p.pool.Query(context.Background(), getUsersQuery, dto.Offset, dto.Limit)
	if err != nil {
		p.lg.Error("failed to get Users", slog.String("op", op), slog.Any("error", err))
	}
	defer rows.Close()
	var users []*entity.User
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.UserId, &user.Name, &user.Password, &user.CreateAt, &user.UpdateAt)
		if err != nil {
			p.lg.Error("failed to get users", slog.String("op", op), slog.Any("error", err))
			return nil, servererrors.InternalServerError
		}
		users = append(users, user)
	}
	return users, nil
}

func (p *PostgreSql) UpdateUser(dto *dto.UpdateUser) (*entity.User, error) {
	const op = "postgresql.UpdateUser"
	user := new(entity.User)
	err := p.pool.QueryRow(context.Background(), updateUserQuery, dto.UserId, dto.Name, dto.Password).Scan(&user.UserId, &user.Name, &user.Password, &user.CreateAt, &user.UpdateAt)
	if errors.Is(err, sql.ErrNoRows) {
		p.lg.Error("failed to update user", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.RecordNotFound
	}
	if err != nil {
		p.lg.Error("failed to update user", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.InternalServerError
	}
	return user, nil
}

func (p *PostgreSql) RemoveUser(userId *uuid.UUID) error {
	const op = "postgresql.RemoveUser"
	err := p.pool.QueryRow(context.Background(), removeUserQuery, userId).Scan(&userId)
	if errors.Is(err, sql.ErrNoRows) {
		p.lg.Error("failed to remove user", slog.String("op", op), slog.Any("error", err))
		return servererrors.RecordNotFound
	}
	if err != nil {
		p.lg.Error("failed to remove user", slog.String("op", op), slog.Any("error", err))
		return servererrors.InternalServerError
	}
	return nil

}
