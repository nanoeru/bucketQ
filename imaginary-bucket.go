package main

import (
	"fmt"
	"math"
)

//	無制限の水入出力
type ImaginaryBucket struct {
}

func (c ImaginaryBucket) Empty() {
}

func (c ImaginaryBucket) Full() {
}

func (c ImaginaryBucket) GetValue() int {
	return 0
}

func (c ImaginaryBucket) GetCapacity() int {
	return math.MaxInt64
}

func (c ImaginaryBucket) GetFreeSpace() int {
	return math.MaxInt64
}

func (c ImaginaryBucket) Pour(toC Container) bool {
	if toC.GetFreeSpace() > 0 {
		toC.Full()
		return true
	}
	return false
}

func (c ImaginaryBucket) CanPoured(v int) bool {
	return true
}

func (c ImaginaryBucket) Poured(v int) bool {
	return true
}

func (c ImaginaryBucket) Copy() Container {
	return &c
}

func (c ImaginaryBucket) String() string {
	return fmt.Sprintf("ImaginaryBucket")
}

func (c ImaginaryBucket) IsBucket() bool {
	return false
}
