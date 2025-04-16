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
	addTaskQuery    = `INSERT INTO task (title, description) VALUES ($1, $2) RETURNING *;`
	getTaskQuery    = `SELECT * FROM task WHERE task_id=$1`
	getTasksQuery   = `SELECT * FROM task OFFSET $1 LIMIT $2`
	updateTaskQuery = `
		UPDATE task SET 
		title = CASE WHEN $2::character varying IS NULL THEN title ELSE $2 END,
		description = CASE WHEN $3::character varying IS NULL THEN description ELSE $3 END,
		status = CASE WHEN $4::character varying IS NULL THEN status ELSE $4 END,
		update_at = now()
		WHERE task_id=$1
		RETURNING *`
	removeTaskQuery = `DELETE FROM task WHERE task_id=$1 RETURNING task_id`
)

func (p *PostgreSql) AddTask(dto *dto.AddTask) (*entity.Task, error) {
	const op = "postgresql.AddTask"
	task := new(entity.Task)
	err := p.pool.QueryRow(context.Background(), addTaskQuery, dto.Title, dto.Description).Scan(&task.TaskId, &task.Title, &task.Description, &task.Status, &task.CreateAt, &task.UpdateAt)
	if err != nil {
		p.lg.Error("failed to add task", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.InternalServerError
	}
	return task, nil
}

func (p *PostgreSql) GetTask(taskId *uuid.UUID) (*entity.Task, error) {
	const op = "postgresql.GetTask"
	task := new(entity.Task)
	err := p.pool.QueryRow(context.Background(), getTaskQuery, taskId).Scan(&task.TaskId, &task.Title, &task.Description, &task.Status, &task.CreateAt, &task.UpdateAt)
	if errors.Is(err, sql.ErrNoRows) {
		p.lg.Error("failed to get task", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.RecordNotFound
	}
	if err != nil {
		p.lg.Error("failed to get task", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.InternalServerError
	}
	return task, nil
}

func (p *PostgreSql) GetTasks(dto *dto.GetTasks) ([]*entity.Task, error) {
	const op = "postgresql.GetTasks"
	rows, err := p.pool.Query(context.Background(), getTasksQuery, dto.Offset, dto.Limit)
	if err != nil {
		p.lg.Error("failed to get tasks", slog.String("op", op), slog.Any("error", err))
	}
	defer rows.Close()
	var tasks []*entity.Task
	for rows.Next() {
		task := new(entity.Task)
		err := rows.Scan(&task.TaskId, &task.Title, &task.Description, &task.Status, &task.CreateAt, &task.UpdateAt)
		if err != nil {
			p.lg.Error("failed to get tasks", slog.String("op", op), slog.Any("error", err))
			return nil, servererrors.InternalServerError
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (p *PostgreSql) UpdateTask(dto *dto.UpdateTask) (*entity.Task, error) {
	const op = "postgresql.UpdateTask"
	task := new(entity.Task)
	err := p.pool.QueryRow(context.Background(), updateTaskQuery, dto.TaskId, dto.Title, dto.Description, dto.Status).Scan(&task.TaskId, &task.Title, &task.Description, &task.Status, &task.CreateAt, &task.UpdateAt)
	if errors.Is(err, sql.ErrNoRows) {
		p.lg.Error("failed to update task", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.RecordNotFound
	}
	if err != nil {
		p.lg.Error("failed to update task", slog.String("op", op), slog.Any("error", err))
		return nil, servererrors.InternalServerError
	}
	return task, nil
}

func (p *PostgreSql) RemoveTask(taskId *uuid.UUID) error {
	const op = "postgresql.RemoveTask"
	err := p.pool.QueryRow(context.Background(), removeTaskQuery, taskId).Scan(&taskId)
	if errors.Is(err, sql.ErrNoRows) {
		p.lg.Error("failed to remove task", slog.String("op", op), slog.Any("error", err))
		return servererrors.RecordNotFound
	}
	if err != nil {
		p.lg.Error("failed to remove task", slog.String("op", op), slog.Any("error", err))
		return servererrors.InternalServerError
	}
	return nil

}
