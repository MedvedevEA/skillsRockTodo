package store

import (
	"context"
	"fmt"
	"skillsRockTodo/internal/entity"
	repoStoreDto "skillsRockTodo/internal/repository/repostore/dto"
	repoStoreErr "skillsRockTodo/internal/repository/repostore/err"

	"github.com/google/uuid"
)

const (
	addUserQuery = `
INSERT INTO "user" (user_id,name) 
VALUES ($1,$2) 
RETURNING *;`
	getUsersQuery = `
SELECT * FROM "user" 
WHERE name ILIKE '%$3%'
OFFSET $1 LIMIT $2;`
	removeUserQuery = `
DELETE FROM "user" 
WHERE user_id=$1;`
)

func (s *Store) AddUserWithUserId(dto *repoStoreDto.AddUser) (*entity.User, error) {
	const op = "store.AddUser"
	user := new(entity.User)
	err := s.pool.QueryRow(context.Background(), addUserQuery, dto.UserId, dto.Name).Scan(&user.UserId, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrInternalServerError, err)
	}
	return user, nil
}
func (s *Store) GetUsers(dto *repoStoreDto.GetUsers) ([]*entity.User, error) {
	const op = "store.GetUsers"
	rows, err := s.pool.Query(context.Background(), getUsersQuery, dto.Offset, dto.Limit, dto.Name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrInternalServerError, err)
	}
	defer rows.Close()
	var users []*entity.User
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.UserId, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrInternalServerError, err)
		}
		users = append(users, user)
	}
	return users, nil
}
func (s *Store) RemoveUser(userId *uuid.UUID) error {
	const op = "store.RemoveUser"
	result, err := s.pool.Exec(context.Background(), removeUserQuery, userId)
	if err != nil {
		return fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrInternalServerError, err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrRecordNotFound, err)
	}
	return nil

}
