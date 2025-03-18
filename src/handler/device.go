package handler

import (
	"net/http"
	"strconv"

	"github.com/iamseki/devices-api/src/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func (h *Handler) InsertDevice(c echo.Context) error {
	device := &queries.Device{}
	if err := c.Bind(device); err != nil {
		return err
	}

	ctx := c.Request().Context()
	conn, err := h.Repository.Pool.Acquire(ctx)
	if err != nil {
		return err
	}

	err = h.Repository.Queries.InsertDevice(ctx, conn, &queries.InsertDeviceParams{Name: device.Name, Brand: device.Brand, State: device.State})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, device)
}

func (h *Handler) GetDevice(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	conn, err := h.Repository.Pool.Acquire(ctx)
	if err != nil {
		return err
	}

	device, err := h.Repository.Queries.GetDevice(ctx, conn, int32(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, device)
}

func (h *Handler) ListDevice(c echo.Context) error {
	ctx := c.Request().Context()

	conn, err := h.Repository.Pool.Acquire(ctx)
	if err != nil {
		return err
	}

	filter := &queries.ListDevicesParams{
		Brand: pgtype.Text{String: c.QueryParam("brand"), Valid: true},
		Name:  pgtype.Text{String: c.QueryParam("name"), Valid: true},
		State: c.QueryParam("state"),
	}

	devices, err := h.Repository.Queries.ListDevices(ctx, conn, filter)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, devices)
}

func (h *Handler) UpdateDevice(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	device := &queries.Device{}
	if err := c.Bind(device); err != nil {
		return err
	}

	conn, err := h.Repository.Pool.Acquire(ctx)
	if err != nil {
		return err
	}

	currentDevice, err := h.Repository.Queries.GetDevice(ctx, conn, int32(id))
	if err != nil {
		return err
	}

	if currentDevice.State == "IN_USE" && (device.Name.String != "" || device.Brand.String != "") {
		return echo.NewHTTPError(http.StatusBadRequest, "Device in use")
	}

	if device.CreationTime.Time.String() != "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Cannot update creation time")
	}

	err = h.Repository.Queries.UpdateDevice(ctx, conn, &queries.UpdateDeviceParams{Name: device.Name, Brand: device.Brand, State: device.State, ID: int32(id)})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, device)
}

func (h *Handler) DeleteDevice(c echo.Context) error {
	ctx := c.Request().Context()

	conn, err := h.Repository.Pool.Acquire(ctx)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	device, err := h.Repository.Queries.GetDevice(ctx, conn, int32(id))
	if err != nil {
		return err
	}

	if device.State == "IN_USE" {
		return echo.NewHTTPError(http.StatusBadRequest, "Device in use")
	}

	err = h.Repository.Queries.DeleteDevice(ctx, conn, int32(id))
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
