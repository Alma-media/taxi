package spheric

import (
	"context"
	"errors"
	"testing"

	"github.com/Alma-media/taxi/model"
)

type sourceMock chan string

func (m sourceMock) Generate(context.Context) (string, error) { return <-m, nil }

type sourceFailMock struct{ error }

func (m sourceFailMock) Generate(context.Context) (string, error) { return "", m.error }

type repositoryFailMock struct{ error }

func (m repositoryFailMock) Save(context.Context, string) (*model.Order, error) { return nil, m.error }

func (m repositoryFailMock) List(context.Context) ([]*model.Order, error) { return nil, m.error }

func Test_OrderWrapper(t *testing.T) {
	t.Run("test if source error is propagated", func(t *testing.T) {
		errExpected := errors.New("source failure")

		source := sourceFailMock{error: errExpected}

		wrapper := NewOrderWrapper(source, nil)

		if _, err := wrapper.Order(context.Background()); err != errExpected {
			t.Errorf(`error "%v" was expected to be "%v"`, err, errExpected)
		}
	})

	t.Run("test if error returned by original repository is propagated", func(t *testing.T) {
		errExpected := errors.New("repository failure")

		repo := repositoryFailMock{error: errExpected}

		source := make(sourceMock)
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
		repo := NewOrderRepository()

		inputSequence := []string{"BB", "AA", "CC", "BB", "AA"}

		outputSequence := []string{"BB - 1", "AA - 1", "CC - 1", "BB - 2", "AA - 2"}

		source := make(sourceMock)
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
