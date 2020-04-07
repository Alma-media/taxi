package repository

import (
	"context"
	"errors"
	"testing"

	genMock "github.com/Alma-media/taxi/generator/mock"
	repoMock "github.com/Alma-media/taxi/repository/mock"
)

func Test_OrderWrapper(t *testing.T) {
	t.Run("test if source error is propagated", func(t *testing.T) {
		errExpected := errors.New("source failure")

		source := genMock.GeneratorFailure{Err: errExpected}

		wrapper := NewOrderWrapper(source, nil)

		if _, err := wrapper.Order(context.Background()); err != errExpected {
			t.Errorf(`error "%v" was expected to be "%v"`, err, errExpected)
		}
	})

	t.Run("test if error returned by original repository is propagated", func(t *testing.T) {
		errExpected := errors.New("repository failure")

		repo := repoMock.OrderFailure{Err: errExpected}

		source := make(genMock.GeneratorSuccess)
		go func() {
			defer close(source)
			source <- "AA"
		}()

		wrapper := NewOrderWrapper(source, repo)

		if _, err := wrapper.Order(context.Background()); err != errExpected {
			t.Errorf(`error "%v" was expected to be "%v"`, err, errExpected)
		}
	})

	t.Run("test save order happy flow", func(t *testing.T) {
		repo := repoMock.OrderSuccess{}

		inputSequence := []string{"AA", "BB", "CC"}

		outputSequence := []string{"AA - 1", "BB - 1", "CC - 1"}

		source := make(genMock.GeneratorSuccess)
		go func() {
			defer close(source)
			for _, id := range inputSequence {
				source <- id
			}
		}()

		wrapper := NewOrderWrapper(source, repo)

		for _, expected := range outputSequence {
			order, _ := wrapper.Order(context.Background())
			if order.String() != expected {
				t.Errorf("order %q was expected to be %q", order, expected)
			}
		}
	})
}
