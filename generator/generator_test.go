package generator

import (
	"context"
	"testing"
	"time"

	"github.com/Alma-media/taxi/config"
)

func Test_Generator(t *testing.T) {
	config := config.Generator{
		KeySize:  10,
		PoolSize: 5000,
		KeyBytes: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		// make interval as short as possible to trigger the timer at least once
		ReplaceInterval: time.Nanosecond,
	}

	t.Run("test getting order from generator", func(t *testing.T) {
		gen := New(config)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		gen.Run(ctx)

		for i := 0; i < config.PoolSize*10; i++ {
			key := gen.Generate()

			if size := len(key); size != config.KeySize {
				t.Errorf("generated key has invalid size: %d", size)
			}
		}
	})
}
