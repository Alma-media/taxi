package spheric

import (
	"reflect"
	"testing"

	"github.com/Alma-media/taxi/model"
)

var (
	aa = model.NewOrder("AA")
	bb = model.NewOrder("BB")
	cc = model.NewOrder("CC")
	dd = model.NewOrder("DD")
)

func Test_OrderRepository(t *testing.T) {
	toBeSaved := []string{"DD", "AA", "CC", "BB", "AA", "AA"}

	t.Run("test save order", func(t *testing.T) {
		repo := NewOrderRepository()

		for _, order := range toBeSaved {
			repo.Save(nil, order)
		}

		expected := map[string]*model.Order{"AA": aa, "BB": bb, "CC": cc, "DD": dd}

		if !reflect.DeepEqual(repo.orders, expected) {
			t.Errorf("repository was expected to contain %v but has %v", expected, repo.orders)
		}
	})

	t.Run("test order list", func(t *testing.T) {
		repo := NewOrderRepository()

		for _, order := range toBeSaved {
			repo.Save(nil, order)
		}

		expected := []*model.Order{aa, bb, cc, dd}

		actual, _ := repo.List(nil)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("list was expected to contain %v but has %v", expected, actual)
		}
	})
}
