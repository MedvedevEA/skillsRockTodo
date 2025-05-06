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
	addTaskUserQuery = `
INSERT INTO user_task (task_id, user_id) 
VALUES ($1, $2) RETURNING *;`
	getTaskUsersQuery = `
SELECT * 
FROM task_user 
WHERE ($3::uuid IS NULL OR task_id=$3) AND ($4::uuid IS null OR user_id=$4) 
OFFSET $1 LIMIT $2;`
	removeTaskUserQuery = `
DELETE FROM task_user
WHERE task_user_id=$1;`
)

func (s *Store) AddTaskUser(dto *repoStoreDto.AddTaskUser) (*entity.TaskUser, error) {
	const op = "store.AddTaskUser"
	TaskUser := new(entity.TaskUser)
	err := s.pool.QueryRow(context.Background(), addTaskUserQuery, dto.TaskId, dto.UserId).Scan(&TaskUser.TaskUserId, &TaskUser.UserId, &TaskUser.TaskId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrInternalServerError, err)
	}
	return TaskUser, nil
}
func (s *Store) GetTaskUsers(dto *repoStoreDto.GetTaskUsers) ([]*entity.TaskUser, error) {
	const op = "store.GetTaskUsers"
	rows, err := s.pool.Query(context.Background(), getTaskUsersQuery, dto.Offset, dto.Limit, dto.TaskId, dto.UserId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrInternalServerError, err)
	}
	defer rows.Close()
	var TaskUsers []*entity.TaskUser
	for rows.Next() {
		TaskUser := new(entity.TaskUser)
		err := rows.Scan(&TaskUser.TaskUserId, &TaskUser.TaskId, &TaskUser.UserId)
		if err != nil {
			return nil, fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrInternalServerError, err)
		}
		TaskUsers = append(TaskUsers, TaskUser)
	}
	return TaskUsers, nil
}

func (s *Store) RemoveTaskUser(TaskUserId *uuid.UUID) error {
	const op = "store.RemoveTaskUser"
	result, err := s.pool.Exec(context.Background(), removeTaskUserQuery, TaskUserId)
	if err != nil {
		return fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrInternalServerError, err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w (%v)", op, repoStoreErr.ErrRecordNotFound, err)
	}
	return nil
}
