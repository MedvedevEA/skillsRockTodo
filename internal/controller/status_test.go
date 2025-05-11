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

func TestStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := storemock.NewMockRepository(mockCtrl)

	lg := slog.New(slog.DiscardHandler)
	s := service.New(mockStore, lg)
	c := controller.New(s, lg)

	app := fiber.New()
	app.Post("/statuses", c.AddStatus)
	app.Get("/statuses/:statusId", c.GetStatus)
	app.Get("/statuses", c.GetStatuses)
	app.Patch("/statuses/:statusId", c.UpdateStatus)
	app.Delete("/statuses/:statusId", c.RemoveStatus)

	//Вспомогательные переменные
	statusId := uuid.New()
	statusName := "statusName"

	// POST /statuses
	t.Run("POST /statuses. Добавление нового статуса. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			AddStatus(statusName).
			Return(&entity.Status{StatusId: &statusId, Name: statusName}, nil)
		body := fmt.Sprintf(`{"name":"%s"}`, statusName)

		url := "/statuses"
		httpRequest, err := http.NewRequest("POST", url, bytes.NewReader([]byte(body)))
		assert.NoError(t, err)
		httpRequest.Header.Set("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 201, httpResponse.StatusCode)

		var response *entity.Status
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, &entity.Status{StatusId: &statusId, Name: statusName}, response)

	})
	// GET /statuses/:statusId
	t.Run("GET /statuses/:statusId. Получение статуса по идентификатору. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			GetStatus(&statusId).
			Return(&entity.Status{StatusId: &statusId, Name: statusName}, nil)

		url := fmt.Sprintf("/statuses/%v", statusId)
		httpRequest, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 200, httpResponse.StatusCode)

		var response *entity.Status
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, &entity.Status{StatusId: &statusId, Name: statusName}, response)
	})
	// GET /statuses
	t.Run("GET /statuses. Получение списка статусов. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			GetStatuses().
			Return([]*entity.Status{{StatusId: &statusId, Name: statusName}}, nil)

		url := "/statuses"
		httpRequest, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 200, httpResponse.StatusCode)

		var response []*entity.Status
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, []*entity.Status{{StatusId: &statusId, Name: statusName}}, response)
	})
	// PATCH /statuses/:statusId
	t.Run("PATCH /statuses/:statusId. Изменение статуса. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			UpdateStatus(&repoStoreDto.UpdateStatus{StatusId: &statusId, Name: &statusName}).
			Return(&entity.Status{StatusId: &statusId, Name: statusName}, nil)
		body := fmt.Sprintf(`{"name":"%s"}`, statusName)

		url := fmt.Sprintf("/statuses/%v", statusId)
		httpRequest, err := http.NewRequest("PATCH", url, bytes.NewReader([]byte(body)))
		assert.NoError(t, err)
		httpRequest.Header.Set("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 200, httpResponse.StatusCode)

		var response *entity.Status
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, &entity.Status{StatusId: &statusId, Name: statusName}, response)
	})
	// DELETE /statuses/:statusId
	t.Run("DELETE /statuses/:statusId.Удаление статуса по идентификатору. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			RemoveStatus(&statusId).
			Return(nil)

		url := fmt.Sprintf("/statuses/%v", statusId)
		httpRequest, err := http.NewRequest("DELETE", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 204, httpResponse.StatusCode)
	})
}
