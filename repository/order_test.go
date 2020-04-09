package repository

import (
	"context"
	"errors"
	"testing"

	genMock "github.com/Alma-media/taxi/generator/mock"
	storageMock "github.com/Alma-media/taxi/storage/mock"
)

func Test_OrderRepository(t *testing.T) {
	t.Run("test if error returned by original repository is propagated", func(t *testing.T) {
		errExpected := errors.New("repository failure")

		storage := storageMock.OrderFailure{Err: errExpected}

		source := make(genMock.Generator)
		go func() {
			defer close(source)
			source <- "AA"
		}()

		repository := NewOrderRepository(source, storage)

		if _, err := repository.Order(context.Background()); err != errExpected {
			t.Errorf(`error "%v" was expected to be "%v"`, err, errExpected)
		}
	})

	t.Run("test save order happy flow", func(t *testing.T) {
		storage := storageMock.OrderSuccess{}

		inputSequence := []string{"AA", "BB", "CC"}

		outputSequence := []string{"AA - 1", "BB - 1", "CC - 1"}

		source := make(genMock.Generator)
		go func() {
			defer close(source)
			for _, id := range inputSequence {
				source <- id
			}
		}()

		repository := NewOrderRepository(source, storage)

		for _, expected := range outputSequence {
			order, _ := repository.Order(context.Background())
			if order.String() != expected {
				t.Errorf("order %q was expected to be %q", order, expected)
			}
		}
	})

	t.Run("test if original repository is invoked calling List()", func(t *testing.T) {
		errExpected := errors.New("source failure")

		storage := storageMock.OrderFailure{Err: errExpected}

		source := make(genMock.Generator)
		go func() {
			defer close(source)
			source <- "AA"
		}()

		repository := NewOrderRepository(source, storage)

		if _, err := repository.List(context.Background()); err != errExpected {
			t.Errorf(`error "%v" was expected to be "%v"`, err, errExpected)
		}
	})
}
