package main

import (
	"fmt"
)

type Container interface {
	Pour(c Container) bool
	Poured(v int) bool
	CanPoured(v int) bool
	GetFreeSpace() int
	GetCapacity() int
	GetValue() int
	Full()
	Empty()
	Copy() Container
	IsBucket() bool
	fmt.Stringer
}
