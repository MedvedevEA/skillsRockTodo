package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"skillsRockTodo/internal/controller"
	"skillsRockTodo/internal/entity"
	"skillsRockTodo/internal/infrastructure/storemock"
	repoStoreDto "skillsRockTodo/internal/repository/repostore/dto"
	"skillsRockTodo/internal/service"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTaskUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := storemock.NewMockRepository(mockCtrl)

	lg := slog.New(slog.DiscardHandler)
	s := service.New(mockStore, lg)
	c := controller.New(s, lg)

	app := fiber.New()
	app.Post("/taskusers", c.AddTaskUser)
	app.Get("/taskusers", c.GetTaskUsers)
	app.Delete("/taskusers/:TaskUserId", c.RemoveTaskUser)

	//Вспомогательные переменные
	taskUserId := uuid.New()
	userId := uuid.New()
	taskId := uuid.New()

	// POST /taskusers
	t.Run("POST /taskusers. Добавление новой записи. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			AddTaskUser(&repoStoreDto.AddTaskUser{TaskId: &taskId, UserId: &userId}).
			Return(&entity.TaskUser{TaskUserId: &taskUserId, TaskId: &taskId, UserId: &userId}, nil)
		body := fmt.Sprintf(`{"taskId":"%v","userId":"%v"}`, taskId, userId)

		url := "/taskusers"
		httpRequest, err := http.NewRequest("POST", url, bytes.NewReader([]byte(body)))
		assert.NoError(t, err)
		httpRequest.Header.Set("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 201, httpResponse.StatusCode)

		var response *entity.TaskUser
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, &entity.TaskUser{TaskUserId: &taskUserId, TaskId: &taskId, UserId: &userId}, response)
	})
	// GET /taskUsers
	t.Run("GET /TaskUsers. Получение списка записей. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			GetTaskUsers(&repoStoreDto.GetTaskUsers{Offset: 0, Limit: 1, TaskId: nil, UserId: nil}).
			Return([]*entity.TaskUser{{TaskUserId: &taskUserId, TaskId: &taskId, UserId: &userId}}, nil)

		url := "/taskusers?offset=0&limit=1"
		httpRequest, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 200, httpResponse.StatusCode)

		var response []*entity.TaskUser
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, []*entity.TaskUser{{TaskUserId: &taskUserId, TaskId: &taskId, UserId: &userId}}, response)
	})
	// DELETE /taskusers/:taskUserId
	t.Run("DELETE /taskuseres/:taskUserId.Удаление записи по идентификатору. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			RemoveTaskUser(&taskUserId).
			Return(nil)

		url := fmt.Sprintf("/taskusers/%v", taskUserId)
		httpRequest, err := http.NewRequest("DELETE", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 204, httpResponse.StatusCode)
	})
}
