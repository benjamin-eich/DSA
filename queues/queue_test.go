package queues

import (
	"testing"
)

func getTestQueue() *Queue[int] {
	q := NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	return q
}

func TestQueue_Empty(t *testing.T) {
	q := getTestQueue()
	want := false
	got := q.Empty()

	if got != want {
		t.Errorf("Empty() = %v, want %v", got, want)
	}

	q = NewQueue[int]()
	want = true
	got = q.Empty()

	if got != want {
		t.Errorf("Empty() = %v, want %v", got, want)
	}
}

func TestQueue_PopLeft(t *testing.T) {
	q := getTestQueue()

	want := 1
	wantErr := false
	got, err := q.PopLeft()
	if got != want {
		t.Errorf("PopLeft() = %v, want %v", got, want)
	}
	if (err != nil) != wantErr {
		t.Errorf("PopLeft() error = %v, wantErr %v", err, false)
	}

	want = 2
	wantErr = false
	got, err = q.PopLeft()
	if got != want {
		t.Errorf("PopLeft() = %v, want %v", got, want)
	}
	if (err != nil) != wantErr {
		t.Errorf("PopLeft() error = %v, wantErr %v", err, false)
	}

	want = 3
	wantErr = false
	got, err = q.PopLeft()
	if got != want {
		t.Errorf("PopLeft() = %v, want %v", got, want)
	}
	if (err != nil) != wantErr {
		t.Errorf("PopLeft() error = %v, wantErr %v", err, false)
	}

	want = 0
	wantErr = true
	got, err = q.PopLeft()
	if got != want {
		t.Errorf("PopLeft() = %v, want %v", got, want)
	}
	if (err != nil) != wantErr {
		t.Errorf("PopLeft() error = %v, wantErr %v", err, false)
	}
}

func TestQueue_PopRight(t *testing.T) {
	q := getTestQueue()

	want := 3
	wantErr := false
	got, err := q.PopRight()
	if got != want {
		t.Errorf("PopRight() = %v, want %v", got, want)
	}
	if (err != nil) != wantErr {
		t.Errorf("PopRight() error = %v, wantErr %v", err, false)
	}

	want = 2
	wantErr = false
	got, err = q.PopRight()
	if got != want {
		t.Errorf("PopRight() = %v, want %v", got, want)
	}
	if (err != nil) != wantErr {
		t.Errorf("PopRight() error = %v, wantErr %v", err, false)
	}

	want = 1
	wantErr = false
	got, err = q.PopRight()
	if got != want {
		t.Errorf("PopRight() = %v, want %v", got, want)
	}
	if (err != nil) != wantErr {
		t.Errorf("PopRight() error = %v, wantErr %v", err, false)
	}

	want = 0
	wantErr = true
	got, err = q.PopRight()
	if got != want {
		t.Errorf("PopRight() = %v, want %v", got, want)
	}
	if (err != nil) != wantErr {
		t.Errorf("PopRight() error = %v, wantErr %v", err, false)
	}
}
