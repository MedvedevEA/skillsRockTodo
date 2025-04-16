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
	addUserTaskQuery    = `INSERT INTO user_task (user_id, task_id) VALUES ($1, $2) RETURNING *;`
	getUserTasksQuery   = `SELECT * FROM user_task WHERE ($3::uuid IS null OR user_id=$3) AND ($4::uuid IS NULL OR task_id=$4) OFFSET $1 LIMIT $2`
	removeUserTaskQuery = `DELETE FROM user_task WHERE user_task_id=$1 RETURNING user_task_id`
)

func (p *PostgreSql) AddUserTask(dto *dto.AddUserTask) (*entity.UserTask, error) {
	const op = "postgresql.AddUserTask"
	userTask := new(entity.UserTask)
	err := p.pool.QueryRow(context.Background(), addUserTaskQuery, dto.UserId, dto.TaskId).Scan(&userTask.UserTaskId, &userTask.UserId, &userTask.TaskId, &userTask.CreateAt, &userTask.UpdateAt)
	if err != nil {
		p.lg.Error("failed to add User", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.InternalServerError
	}
	return userTask, nil
}

func (p *PostgreSql) GetUserTasks(dto *dto.GetUserTasks) ([]*entity.UserTask, error) {
	const op = "postgresql.GetUserTasks"
	rows, err := p.pool.Query(context.Background(), getUserTasksQuery, dto.Offset, dto.Limit, dto.UserId, dto.TaskId)
	if err != nil {
		p.lg.Error("failed to get UserTasks", slog.String("op", op), slog.Any("error", err))
	}
	defer rows.Close()
	var userTasks []*entity.UserTask
	for rows.Next() {
		userTask := new(entity.UserTask)
		err := rows.Scan(&userTask.UserTaskId, &userTask.UserId, &userTask.TaskId, &userTask.CreateAt, &userTask.UpdateAt)
		if err != nil {
			p.lg.Error("failed to get userTasks", slog.String("op", op), slog.Any("error", err))
			return nil, servererrors.InternalServerError
		}
		userTasks = append(userTasks, userTask)
	}
	return userTasks, nil
}

func (p *PostgreSql) RemoveUserTask(userTaskId *uuid.UUID) error {
	const op = "postgresql.RemoveUserTask"
	err := p.pool.QueryRow(context.Background(), removeUserTaskQuery, userTaskId).Scan(&userTaskId)
	if errors.Is(err, sql.ErrNoRows) {
		p.lg.Error("failed to remove userTask", slog.String("op", op), slog.Any("error", err))
		return servererrors.RecordNotFound
	}
	if err != nil {
		p.lg.Error("failed to remove userTask", slog.String("op", op), slog.Any("error", err))
		return servererrors.InternalServerError
	}
	return nil

}
