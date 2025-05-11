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
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := storemock.NewMockRepository(mockCtrl)

	lg := slog.New(slog.DiscardHandler)
	s := service.New(mockStore, lg)
	c := controller.New(s, lg)

	app := fiber.New()
	app.Post("/messages", c.AddMessage)
	app.Get("/messages/:messageId", c.GetMessage)
	app.Get("/messages", c.GetMessages)
	app.Patch("/messages/:messageId", c.UpdateMessage)
	app.Delete("/messages/:messageId", c.RemoveMessage)

	//Вспомогательные переменные
	messageId := uuid.New()
	taskId := uuid.New()
	userId := uuid.New()
	messageText := "messageText"
	createAt := time.Now().Round(0)
	updateAt := time.Now().Round(0)

	// POST /messages
	t.Run("POST /messages. Добавление нового сообщения. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			AddMessage(&repoStoreDto.AddMessage{TaskId: &taskId, UserId: &userId, Text: messageText}).
			Return(&entity.Message{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt}, nil)
		body := fmt.Sprintf(`{"taskId":"%v","userId":"%v","text":"%s"}`, taskId, userId, messageText)

		url := "/messages"
		httpRequest, err := http.NewRequest("POST", url, bytes.NewReader([]byte(body)))
		assert.NoError(t, err)
		httpRequest.Header.Set("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 201, httpResponse.StatusCode)

		var response *entity.Message
		json.NewDecoder(httpResponse.Body).Decode(&response)
		fmt.Printf("%+v\n", &entity.Message{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt})
		assert.Equal(t, &entity.Message{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt}, response)
	})
	// GET /messages/:messageId
	t.Run("GET /messages/:messageId. Получение сообщения по идентификатору. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			GetMessage(&messageId).
			Return(&entity.Message{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt}, nil)

		url := fmt.Sprintf("/Messages/%v", messageId)
		httpRequest, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 200, httpResponse.StatusCode)

		var response *entity.Message
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, &entity.Message{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt}, response)
	})
	// GET /messages
	t.Run("GET /messages. Получение списка сообщений. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			GetMessages(&repoStoreDto.GetMessages{Offset: 0, Limit: 1, TaskId: nil}).
			Return([]*entity.Message{{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt}}, nil)

		url := "/messages?offset=0&limit=1"
		httpRequest, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 200, httpResponse.StatusCode)

		var response []*entity.Message
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, []*entity.Message{{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt}}, response)
	})
	// PATCH /messages/:messageId
	t.Run("PATCH /messages/:messageId. Изменение сообщения. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			UpdateMessage(&repoStoreDto.UpdateMessage{MessageId: &messageId, Text: &messageText}).
			Return(&entity.Message{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt}, nil)
		body := fmt.Sprintf(`{"text":"%s"}`, messageText)

		url := fmt.Sprintf("/Messages/%v", messageId)
		httpRequest, err := http.NewRequest("PATCH", url, bytes.NewReader([]byte(body)))
		assert.NoError(t, err)
		httpRequest.Header.Set("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 200, httpResponse.StatusCode)

		var response *entity.Message
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, &entity.Message{MessageId: &messageId, TaskId: &taskId, UserId: &userId, Text: messageText, CreateAt: createAt, UpdateAt: &updateAt}, response)

	})
	// DELETE /Messages/:MessageId
	t.Run("DELETE /messagees/:messageId.Удаление сообщения по идентификатору. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			RemoveMessage(&messageId).
			Return(nil)

		url := fmt.Sprintf("/Messages/%v", messageId)
		httpRequest, err := http.NewRequest("DELETE", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 204, httpResponse.StatusCode)
	})
}
