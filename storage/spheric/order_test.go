package spheric

import (
	"reflect"
	"testing"

	"github.com/Alma-media/taxi/model"
	"github.com/Alma-media/taxi/storage"
)

var _ storage.Order = (*OrderStorage)(nil)

var (
	aa = model.NewOrder("AA")
	bb = model.NewOrder("BB")
	cc = model.NewOrder("CC")
	dd = model.NewOrder("DD")
)

func Test_OrderStorage(t *testing.T) {
	toBeSaved := []string{"DD", "AA", "CC", "BB", "AA", "AA"}

	t.Run("test save order", func(t *testing.T) {
		repo := NewOrderStorage()

		for _, order := range toBeSaved {
			repo.Save(nil, order)
		}

		expected := map[string]*model.Order{"AA": aa, "BB": bb, "CC": cc, "DD": dd}

		if !reflect.DeepEqual(repo.orders, expected) {
			t.Errorf("repository %v was expected to equal %v", repo.orders, expected)
		}
	})

	t.Run("test order list", func(t *testing.T) {
		repo := NewOrderStorage()

		for _, order := range toBeSaved {
			repo.Save(nil, order)
		}

		expected := []*model.Order{aa, bb, cc, dd}

		actual, _ := repo.List(nil)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("list %v was expected to contain %v", actual, expected)
		}
	})
}
