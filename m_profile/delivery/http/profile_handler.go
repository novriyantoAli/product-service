package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/novriyantoAli/product-service/domain"
	"github.com/sirupsen/logrus"
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
	g.POST("", h.Save)
	g.POST("/radcheck/:groupname", h.SaveRadcheck)
	g.POST("/radreply/:groupname", h.SaveRadreply)
	g.DELETE("/radcheck/:groupname/:id", h.DeleteRadcheck)
	g.DELETE("/radreply/:groupname/:id", h.DeleteRadreply)
	g.DELETE("/:id", h.Delete)
}

func (h *handler) DeleteRadcheck(c echo.Context) error {
	groupname := c.Param("groupname")
	res, err := h.UC.Find(&domain.Profile{Groupname: groupname})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	if len(res) != 1 {
		logrus.Error("anomali behaviour groupname in profile(radusergroup)")
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	idString := c.Param("id")
	idUint64, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "id tidak valid",
		})
	}

	err = h.UC.DeleteRadcheck(uint(idUint64))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusAccepted, domain.ErrorMessage{
		Success: true,
		Message: "berhasil menghapus data...",
	})
}

func (h *handler) DeleteRadreply(c echo.Context) error {
	groupname := c.Param("groupname")
	res, err := h.UC.Find(&domain.Profile{Groupname: groupname})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	if len(res) != 1 {
		logrus.Error("anomali behaviour groupname in profile(radusergroup)")
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	idString := c.Param("id")
	idUint64, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "id tidak valid",
		})
	}

	err = h.UC.DeleteRadreply(uint(idUint64))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusAccepted, domain.ErrorMessage{
		Success: true,
		Message: "berhasil menghapus data...",
	})
}

func (h *handler) Delete(c echo.Context) error {
	idString := c.Param("id")
	err := h.UC.Delete(idString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusAccepted, domain.ErrorMessage{
		Success: true,
		Message: "berhasil menghapus data...",
	})
}

func (h *handler) SaveRadcheck(c echo.Context) error {
	groupname := c.Param("groupname")
	res, err := h.UC.Find(&domain.Profile{Groupname: groupname})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	if len(res) != 1 {
		logrus.Error("anomali behaviour groupname in profile(radusergroup)")
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	radcheck := new(domain.Radgroupcheck)
	err = c.Bind(radcheck)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "permintaan tidak lengkap",
			Data:    err,
		})
	}

	if err := c.Validate(radcheck); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: fmt.Sprintf("validasi gagal: %s", err.Error()),
			Data:    err,
		})
	}

	err = h.UC.SaveRadcheck(radcheck)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, domain.ErrorMessage{
		Success: true,
		Message: "berhasil menyimpan data...",
	})
}

func (h *handler) SaveRadreply(c echo.Context) error {
	groupname := c.Param("groupname")
	res, err := h.UC.Find(&domain.Profile{Groupname: groupname})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	if len(res) != 1 {
		logrus.Error("anomali behaviour groupname in profile(radusergroup)")
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	radreply := new(domain.Radgroupreply)
	err = c.Bind(radreply)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "permintaan tidak lengkap",
			Data:    err,
		})
	}

	if err := c.Validate(radreply); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: fmt.Sprintf("validasi gagal: %s", err.Error()),
			Data:    err,
		})
	}

	err = h.UC.SaveRadreply(radreply)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, domain.ErrorMessage{
		Success: true,
		Message: "berhasil menyimpan data...",
	})
}

func (h *handler) Save(c echo.Context) error {
	profile := new(domain.Profile)
	err := c.Bind(profile)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "permintaan tidak lengkap",
			Data:    err,
		})
	}

	if err := c.Validate(profile); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: fmt.Sprintf("validasi gagal: %s", err.Error()),
			Data:    err,
		})
	}

	profile.Groupname = profile.Username

	err = h.UC.Save(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, domain.ErrorMessage{
		Success: true,
		Message: "data berhasil dibuat...",
	})
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
