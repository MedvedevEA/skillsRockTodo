package controller_test

import (
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

func TestUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := storemock.NewMockRepository(mockCtrl)

	lg := slog.New(slog.DiscardHandler)
	s := service.New(mockStore, lg)
	c := controller.New(s, lg)

	app := fiber.New()
	app.Get("/users", c.GetUsers)

	//Вспомогательные переменные
	userId := uuid.New()
	userName := "userName"

	// GET /users
	t.Run("GET /users. Получение списка пользователей. Успех", func(t *testing.T) {
		mockStore.EXPECT().
			GetUsers(&repoStoreDto.GetUsers{Offset: 0, Limit: 1, Name: &userName}).
			Return([]*entity.User{{UserId: &userId, Name: userName}}, nil)

		url := fmt.Sprintf("/users?offset=0&limit=1&name=%s", userName)
		httpRequest, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)
		assert.Equal(t, 200, httpResponse.StatusCode)

		var response []*entity.User
		json.NewDecoder(httpResponse.Body).Decode(&response)
		assert.Equal(t, []*entity.User{{UserId: &userId, Name: userName}}, response)
	})
}
