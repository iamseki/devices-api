package domain_test

import (
	"testing"
	"time"

	"github.com/iamseki/devices-api/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestValidateDeleteDevice(t *testing.T) {
	t.Run("Valid state", func(t *testing.T) {
		err := domain.ValidateDeleteDevice("AVAILABLE")
		assert.NoError(t, err)
	})

	t.Run("Invalid state", func(t *testing.T) {
		err := domain.ValidateDeleteDevice("IN_USE")
		assert.Error(t, err)
	})
}

func TestValidateUpdateDevice(t *testing.T) {
	t.Run("Valid update", func(t *testing.T) {
		params := &domain.UpdateDeviceParams{Name: "test"}
		err := domain.ValidateUpdateDevice(params, "AVAILABLE")
		assert.NoError(t, err)
	})

	t.Run("Cannot be update for IN_USE", func(t *testing.T) {
		params := &domain.UpdateDeviceParams{Name: "foo", Brand: "bar"}
		err := domain.ValidateUpdateDevice(params, "IN_USE")
		assert.Error(t, err)
	})

	t.Run("Creation time cannot be updated", func(t *testing.T) {
		params := &domain.UpdateDeviceParams{CreationTime: time.Now()}
		err := domain.ValidateUpdateDevice(params, "")
		assert.Error(t, err)
	})
}
