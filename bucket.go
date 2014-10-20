package main

import (
	"fmt"
)

//	バケツ
type Bucket struct {
	Value    int
	Capacity int
}

func (c *Bucket) Empty() {
	c.Value = 0
}

func (c *Bucket) Full() {
	c.Value = c.Capacity
}

func (c Bucket) GetValue() int {
	return c.Value
}

func (c Bucket) GetCapacity() int {
	return c.Capacity
}

func (c Bucket) GetFreeSpace() int {
	return c.Capacity - c.Value
}

//	変化があるかないか
func (c *Bucket) Pour(toC Container) bool {
	toContainerFreeSpace := toC.GetFreeSpace()
	//	空
	if c.Value == 0 || toContainerFreeSpace == 0 {
		return false
	}
	//	移動可能量
	movableValue := toContainerFreeSpace
	if c.Value < movableValue {
		movableValue = c.Value
	}
	c.Value -= movableValue
	//	上記条件より必ずtrue
	return toC.Poured(movableValue)
	//	return true
}

//	注ぐ(失敗すると何もせずにfalse)
func (c *Bucket) Poured(v int) bool {
	if c.CanPoured(v) {
		c.Value += v
		return true
	}
	return false
}

func (c *Bucket) CanPoured(v int) bool {
	return c.Value+v <= c.Capacity
}

func (c Bucket) Copy() Container {
	return &c
}

func (c Bucket) String() string {
	return fmt.Sprintf("Bucket(%d/%d)", c.Value, c.Capacity)
}

func (c Bucket) IsBucket() bool {
	return true
}
