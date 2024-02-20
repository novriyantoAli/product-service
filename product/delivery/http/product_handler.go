package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/novriyantoAli/product-service/domain"
	"github.com/novriyantoAli/product-service/model"
)

type productHandler struct {
	UC domain.ProductUsecase
}

func NewHandler(e *echo.Echo, uc domain.ProductUsecase) {
	handler := &productHandler{UC: uc}

	group := e.Group("/api/v1/product")
	group.GET("", handler.Find)
	group.GET("/:id/detail", handler.Detail)
	group.POST("", handler.Save)
	group.PUT("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)
}

func (h *productHandler) Find(c echo.Context) error {
	res, err := h.UC.Find(&model.Product{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.ErrorMessage{
		Success: true,
		Message: "fetched...",
		Data:    res.List,
	})
}

func (h *productHandler) Detail(c echo.Context) error {
	idString := c.Param("id")
	idUint, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "id tidak dapat di proses",
			Data:    err.Error(),
		})
	}

	res, err := h.UC.FindProduct(&domain.Product{ID: uint(idUint)})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.ErrorMessage{
		Success: true,
		Message: "berhasil mengabil data",
		Data:    res,
	})
}

func (h *productHandler) Save(c echo.Context) error {
	product := new(domain.Product)
	err := c.Bind(product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "permintaan tidak lengkap",
			Data:    err,
		})
	}

	if err := c.Validate(product); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: fmt.Sprintf("validasi gagal: %s", err.Error()),
			Data:    err,
		})
	}

	httpStatus := http.StatusCreated
	success := true
	message := "berhasil membuat data"

	err = h.UC.Create(product)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		success = false
		message = "internal server error"
	}

	return c.JSON(httpStatus, domain.ErrorMessage{
		Success: success,
		Message: message,
	})
}

func (h *productHandler) Update(c echo.Context) error {
	idString := c.Param("id")
	idUint64, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "id tidak valid",
			Data:    err.Error(),
		})
	}

	product := new(domain.Product)
	err = c.Bind(product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "permintaan tidak lengkap",
			Data:    err,
		})
	}

	if err := c.Validate(product); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: fmt.Sprintf("validasi gagal: %s", err.Error()),
			Data:    err,
		})
	}

	product.ID = uint(idUint64)
	err = h.UC.Save(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, domain.ErrorMessage{
		Success: true,
		Message: "berhasil memperbaharui data",
	})
}

func (h *productHandler) Delete(c echo.Context) error {
	idString := c.Param("id")
	idUint64, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "id tidak valid",
			Data:    err.Error(),
		})
	}

	err = h.UC.Delete(uint(idUint64))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusAccepted, domain.ErrorMessage{
		Success: true,
		Message: "berhasil menghapus data",
	})
}
