package handlers

import (
	"net/http"

	"github.com/Splucheviy/codewithteeEchoLesson/cmd/api/requests"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	payload := new(requests.RegisterUserRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	c.Logger().Info(payload)
	return c.String(http.StatusBadRequest, "good request")
}
