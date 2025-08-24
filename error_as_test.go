package gotchas_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Mikhalevich/gotchas"
)

func TestErrorAsWithPointerReceiver(t *testing.T) {
	t.Parallel()

	var (
		errFunc = func() error {
			return &gotchas.CustomPointerReceiverError{}
		}
	)

	t.Run("error.As with pointer receiver value", func(t *testing.T) {
		t.Parallel()

		var (
			err = errFunc()
			p   *gotchas.CustomPointerReceiverError
		)

		require.ErrorAs(t, err, &p)
	})

	t.Run("error.As with value receiver value", func(t *testing.T) {
		t.Parallel()

		var (
			err = errFunc()
			p   gotchas.CustomPointerReceiverError
		)

		require.Panics(t, func() {
			//nolint:govet
			errors.As(err, &p)
		})
	})
}

func TestErrorAsWithValueReceiver(t *testing.T) {
	t.Parallel()

	t.Run("error value and error.As with value receiver value", func(t *testing.T) {
		t.Parallel()

		var (
			err = func() error {
				return gotchas.CustomValueReceiverError{}
			}()
			v gotchas.CustomValueReceiverError
		)

		require.ErrorAs(t, err, &v)
	})

	t.Run("error value error.As with pointer receiver value", func(t *testing.T) {
		t.Parallel()

		var (
			err = func() error {
				return gotchas.CustomValueReceiverError{}
			}()
			v *gotchas.CustomValueReceiverError
		)

		require.NotErrorAs(t, err, &v)
	})

	t.Run("error pointer and error.As with value receiver value", func(t *testing.T) {
		t.Parallel()

		var (
			err = func() error {
				return &gotchas.CustomValueReceiverError{}
			}()
			v gotchas.CustomValueReceiverError
		)

		require.NotErrorAs(t, err, &v)
	})

	t.Run("error pointer and error.As with pointer receiver value", func(t *testing.T) {
		t.Parallel()

		var (
			err = func() error {
				return &gotchas.CustomValueReceiverError{}
			}()
			v *gotchas.CustomValueReceiverError
		)

		require.ErrorAs(t, err, &v)
	})
}
