package repository

import (
	"context"
	"errors"
	"testing"

	genMock "github.com/Alma-media/taxi/generator/mock"
	storageMock "github.com/Alma-media/taxi/storage/mock"
)

func Test_Pipe(t *testing.T) {
	t.Run("test if source error is propagated", func(t *testing.T) {
		errExpected := errors.New("source failure")

		source := genMock.GeneratorFailure{Err: errExpected}

		wrapper := NewPipe(source, nil)

		if _, err := wrapper.Order(context.Background()); err != errExpected {
			t.Errorf(`error "%v" was expected to be "%v"`, err, errExpected)
		}
	})

	t.Run("test if error returned by original repository is propagated", func(t *testing.T) {
		errExpected := errors.New("repository failure")

		repo := storageMock.OrderFailure{Err: errExpected}

		source := make(genMock.GeneratorSuccess)
		go func() {
			defer close(source)
			source <- "AA"
		}()

		wrapper := NewPipe(source, repo)

		if _, err := wrapper.Order(context.Background()); err != errExpected {
			t.Errorf(`error "%v" was expected to be "%v"`, err, errExpected)
		}
	})

	t.Run("test save order happy flow", func(t *testing.T) {
		repo := storageMock.OrderSuccess{}

		inputSequence := []string{"AA", "BB", "CC"}

		outputSequence := []string{"AA - 1", "BB - 1", "CC - 1"}

		source := make(genMock.GeneratorSuccess)
		go func() {
			defer close(source)
			for _, id := range inputSequence {
				source <- id
			}
		}()

		wrapper := NewPipe(source, repo)

		for _, expected := range outputSequence {
			order, _ := wrapper.Order(context.Background())
			if order.String() != expected {
				t.Errorf("order %q was expected to be %q", order, expected)
			}
		}
	})
}
