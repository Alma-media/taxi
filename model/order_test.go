package model

import (
	"sync"
	"sync/atomic"
	"testing"
)

func Test_Order(t *testing.T) {
	t.Run("test order to string conversion", func(t *testing.T) {
		order := NewOrder("ID")

		if str := order.String(); str != "ID - 0" {
			t.Errorf("unexpected output: %s", str)
		}
	})

	t.Run("test order counter (concurrent)", func(t *testing.T) {
		order := NewOrder("ID")

		const expected = 50

		var wg sync.WaitGroup
		for i := 0; i < expected; i++ {
			wg.Add(1)
			go func() {
				order.Increment()
				wg.Done()
			}()
		}

		wg.Wait()

		if actual := atomic.LoadUint64(&order.counter); actual != expected {
			t.Errorf("counter was expected to equal %d instead of %d", expected, actual)
		}
	})
}
