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
		PoolSize: 50,
		KeyBytes: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		// make interval as short as possible to trigger the timer at least once
		ReplaceInterval: time.Nanosecond,
	}

	t.Run("test generator initialization", func(t *testing.T) {
		gen := New(config)

		if err := gen.Init(); err != nil {
			t.Error("should not return an error")
		}
		defer gen.Close()
	})

	t.Run("test if error is returned trying to close inactive generator", func(t *testing.T) {
		gen := New(config)

		err := gen.Close()
		if err == nil {
			t.Error("error was expected")
		}
		if err != errNotInitialized {
			t.Errorf("%q error was expected but got %q", errNotInitialized, err)
		}
	})

	t.Run("test if error is returned trying to initialize active generator", func(t *testing.T) {
		gen := New(config)

		if err := gen.Init(); err != nil {
			t.Fatal(err)
		}
		defer gen.Close()

		err := gen.Init()
		if err == nil {
			t.Error("error was expected")
		}
		if err != errAlreadyInitialized {
			t.Errorf("%q error was expected but got %q", errAlreadyInitialized, err)
		}
	})

	t.Run("test getting order from generator", func(t *testing.T) {
		gen := New(config)

		if err := gen.Init(); err != nil {
			t.Fatal(err)
		}
		defer gen.Close()

		for i := 0; i < config.PoolSize*10; i++ {
			key, err := gen.Generate(context.Background())
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if size := len(key); size != config.KeySize {
				t.Errorf("generated key has invalid size: %d", size)
			}
		}
	})

	t.Run("test if error is returned generating the key by inactive generator", func(t *testing.T) {
		gen := New(config)

		// "fake" context with timeout/deadline (any canceled context is suitable)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := gen.Generate(ctx)
		if err == nil {
			t.Error("error was expected")
		}
		if err != context.Canceled {
			t.Errorf("%q error was expected but got %q", context.Canceled, err)
		}
	})
}
