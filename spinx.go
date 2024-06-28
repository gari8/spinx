package spinx

import (
	"fmt"
	"time"
)

var (
	defaultSpinSpeed = 100 * time.Millisecond
	DefaultRune      = []rune{'⠧', '⠇', '⠏', '⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦'}
	DefaultStr       = []string{"*.........", ".*........", "..*.......", "...*......", "....*.....", ".....*....", "......*...", ".......*..", "........*.", ".........*"}
)

type chars interface {
	rune | string
}

type Spinner[T chars] struct {
	chars     []T
	done      chan struct{}
	spinSpeed time.Duration
}

type Option[T chars] struct {
	Chars     []T
	SpinSpeed *time.Duration
}

type options[T chars] []Option[T]

func (opts options[T]) validate() *Option[T] {
	option := &Option[T]{
		SpinSpeed: &defaultSpinSpeed,
	}
	switch any(*new(T)).(type) {
	case rune:
		option.Chars = any(DefaultRune).([]T)
	case string:
		option.Chars = any(DefaultStr).([]T)
	}
	if len(opts) > 0 {
		if opts[0].Chars != nil {
			option.Chars = opts[0].Chars
		}
		if opts[0].SpinSpeed != nil {
			option.SpinSpeed = opts[0].SpinSpeed
		}
	}
	return option
}

func NewSpinner[T chars](opts ...Option[T]) *Spinner[T] {
	option := options[T](opts).validate()
	return &Spinner[T]{
		chars:     option.Chars,
		done:      make(chan struct{}),
		spinSpeed: *(option.SpinSpeed),
	}
}

func (s *Spinner[T]) Spin() {
	go func(s *Spinner[T]) {
		for {
			for _, r := range s.chars {
				select {
				case <-s.done:
					fmt.Printf("\r\r")
					return
				default:
					switch any(s.chars).(type) {
					case []rune:
						fmt.Printf("\r%c", r)
					case []string:
						fmt.Printf("\r%s", r)
					}
					time.Sleep(s.spinSpeed)
				}
			}
		}
	}(s)
}

func (s *Spinner[T]) Stop() {
	s.done <- struct{}{}
}

func (s *Spinner[T]) Shift(speed time.Duration) {
	s.spinSpeed = speed
}
