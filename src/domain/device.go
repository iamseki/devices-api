package domain

import (
	"errors"
	"time"
)

type UpdateDeviceParams struct {
	Name         string
	Brand        string
	State        string
	CreationTime time.Time
}

func ValidateUpdateDevice(params *UpdateDeviceParams, currentDeviceState string) error {
	if currentDeviceState == "IN_USE" && (params.Name != "" || params.Brand != "") {
		return errors.New("device in use")
	}

	if !params.CreationTime.IsZero() {
		return errors.New("cannot update creation time")
	}

	return nil
}

func ValidateDeleteDevice(state string) error {
	if state == "IN_USE" {
		return errors.New("device in use")
	}

	return nil
}
