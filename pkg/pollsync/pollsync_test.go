package pollsync

import (
	"context"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestPollSyncBasic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	i := New(ctx)

	goroutineCount := runtime.NumGoroutine()

	if err := i.Error(); err != nil {
		t.Fatalf("Unknown error received early in test suite: %v", err)
	}

	if data := i.Data("test"); data != nil {
		t.Fatal("data received early was non-nil:", data)
	}

	var x int

	i.Register(100*time.Millisecond, "test", func(context.Context) (interface{}, error) {
		x++
		return x, nil
	})

	time.Sleep(time.Second)

	if x < 10 {
		t.Fatal("Count did not increment properly:", x)
	}

	var y int

	i.Register(time.Second, "test2", func(context.Context) (interface{}, error) {
		y++
		return y, nil
	})

	time.Sleep(2 * time.Second)

	if x < 30 {
		t.Fatal("x couldn't possibly be this small right now:", x)
	}

	if y < 1 {
		t.Fatal("y is a little behind", y)
	}

	i.Stop() // stop the tape. We have to sleep a second for our goroutines to cleanly exit.
	time.Sleep((100 * time.Millisecond) + time.Second)

	if runtime.NumGoroutine() != goroutineCount {
		t.Fatalf("Stop did not shutter all goroutines: orig: %v, now: %v", goroutineCount, runtime.NumGoroutine())
	}

	if !reflect.DeepEqual(i.Data("test"), x) {
		t.Fatal("we are not getting the right data back from x")
	}

	if !reflect.DeepEqual(i.Data("test2"), y) {
		t.Fatal("we are not getting the right data back from y")
	}

}
