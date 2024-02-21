package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/novriyantoAli/product-service/domain"
)

type handler struct {
	UC domain.VoucherUsecase
}

func NewHandler(e *echo.Echo, uc domain.VoucherUsecase) {
	h := &handler{UC: uc}

	g := e.Group("/api/v1/voucher")
	g.GET("", h.Fetch)
	g.POST("", h.Save)
}

func (h *handler) Fetch(c echo.Context) error {
	res, err := h.UC.Find(&domain.Voucher{})
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

func (h *handler) Save(c echo.Context) error {
	cvr := new(domain.CreateVoucherRequest)
	err := c.Bind(cvr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: "permintaan tidak lengkap",
			Data:    err,
		})
	}

	if err := c.Validate(cvr); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorMessage{
			Success: false,
			Message: fmt.Sprintf("validasi gagal: %s", err.Error()),
			Data:    err,
		})
	}

	err = h.UC.Create(cvr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Success: false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, domain.ErrorMessage{
		Success: true,
		Message: "berhasil membuat data",
	})
}
