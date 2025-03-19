package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/iamseki/devices-api/src/repository/queries"
	"github.com/labstack/echo/v4"
)

// @Summary Insert a new device into the database
// @Description Insert a new device into the database
// @ID insert-device
// @Tags devices
// @Accept json
// @Produce json
// @Param device body DeviceResponse true "Device data"
// @Success 201 {object} queries.Device
// @Router /devices [POST]
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

	params := &queries.InsertDeviceParams{Name: device.Name, Brand: device.Brand}
	fmt.Println(params)
	err = h.Repository.Queries.InsertDevice(ctx, conn, params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, device)
}

// @Summary Retrieve a device by ID
// @Description Get details of a device by its ID
// @ID get-device
// @Tags devices
// @Accept json
// @Produce json
// @Param id path int true "Device ID"
// @Success 200 {object} queries.Device
// @Router /devices/{id} [get]
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

// @Summary List devices
// @Description Retrieve a list of devices filtered by brand, state, or name
// @ID list-devices
// @Tags devices
// @Accept json
// @Produce json
// @Param brand query string false "Brand of the device"
// @Param state query string false "State of the device"
// @Param name query string false "Name of the device"
// @Success 200 {array} queries.Device
// @Router /devices [get]
func (h *Handler) ListDevice(c echo.Context) error {
	ctx := c.Request().Context()

	conn, err := h.Repository.Pool.Acquire(ctx)
	if err != nil {
		return err
	}

	filter := &queries.ListDevicesParams{}

	brand := c.QueryParam("brand")
	state := c.QueryParam("state")
	name := c.QueryParam("name")
	if brand != "" {
		filter.Brand = brand
	}

	if state != "" {
		filter.State = state
	}

	if name != "" {
		filter.Name = name
	}

	fmt.Println(state)

	devices, err := h.Repository.Queries.ListDevices(ctx, conn, filter)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, devices)
}

// @Summary Update a device
// @Description Update details of an existing device
// @ID update-device
// @Tags devices
// @Accept json
// @Produce json
// @Param id path int true "Device ID"
// @Param device body queries.Device true "Updated device data"
// @Success 200 {object} queries.Device
// @Router /devices/{id} [patch]
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

	if currentDevice.State == "IN_USE" && (device.Name != "" || device.Brand != "") {
		return echo.NewHTTPError(http.StatusBadRequest, "Device in use")
	}
	fmt.Println(device)
	if !device.CreationTime.IsZero() {
		return echo.NewHTTPError(http.StatusBadRequest, "Cannot update creation time")
	}

	err = h.Repository.Queries.UpdateDevice(ctx, conn, &queries.UpdateDeviceParams{Name: device.Name, Brand: device.Brand, State: device.State, ID: int32(id)})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, device)
}

// @Summary Delete a device
// @Description Delete a device by its ID
// @ID delete-device
// @Tags devices
// @Accept json
// @Produce json
// @Param id path int true "Device ID"
// @Success 204 "No Content"
// @Router /devices/{id} [delete]
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
