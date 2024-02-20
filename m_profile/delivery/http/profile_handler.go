package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/novriyantoAli/product-service/domain"
)

type handler struct {
	UC domain.ProfileUsecase
}

func NewHandler(e *echo.Echo, uc domain.ProfileUsecase) {
	h := &handler{UC: uc}

	g := e.Group("/api/v1/profile")
	g.GET("", h.Find)
	g.GET("/radcheck/:groupname", h.FindRadcheck)
	g.GET("/radreply/:groupname", h.FindRadreply)
}

func (h *handler) FindRadcheck(c echo.Context) error {
	groupname := c.Param("groupname")
	res, err := h.UC.FindRadcheck(&domain.Radgroupcheck{Groupname: groupname})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, domain.ErrorMessage{
		Success: true,
		Message: "berhasil mengambil data",
		Data:    res,
	})
}

func (h *handler) FindRadreply(c echo.Context) error {
	groupname := c.Param("groupname")
	res, err := h.UC.FindRadreply(&domain.Radgroupreply{Groupname: groupname})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, domain.ErrorMessage{
		Success: true,
		Message: "berhasil mengambil data",
		Data:    res,
	})
}

func (h *handler) Find(c echo.Context) error {
	res, err := h.UC.Find(&domain.Profile{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, domain.ErrorMessage{
		Success: true,
		Message: "berhasil mengambil data",
		Data:    res,
	})
}
