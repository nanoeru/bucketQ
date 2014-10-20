package main

import (
	"fmt"
)

func NewBucketQ(step int, value int, log string, containers Containers) BucketQ {
	return BucketQ{
		step,
		value,
		log,
		containers,
	}
}

//	バケツ問題結果保存用構造体
type BucketQ struct {
	Step  int
	Value int
	Log   string
	Containers
}

func (b BucketQ) String() string {
	str := fmt.Sprintf("Sum:%d\n MinStep:%d\n", b.Value, b.Step)
	str += fmt.Sprintf("%s\n", b.Containers)
	str += b.Log
	return str
}

func NewBlankContainers(a ...int) (c Containers) {
	size := len(a)
	c = Containers(make([]Container, size))
	for i, v := range a {
		c[i] = &Bucket{0, v}
	}
	c = append(c, &ImaginaryBucket{})
	return
}

type Containers []Container

func (c Containers) SumValue() (sum int) {
	for _, v := range c {
		sum += v.GetValue()
	}
	return
}

func (c Containers) HasValue(aimValue int) bool {
	for _, v := range c {
		if v.GetValue() == aimValue {
			return true
		}
	}
	return false
}

func (c Containers) Copy() Containers {
	newC := Containers(make([]Container, len(c)))
	for i, v := range c {
		newC[i] = v.Copy()
	}
	return newC
}

func (c Containers) String() string {
	str := ""
	for i, v := range c {
		if v.IsBucket() {
			if str != "" {
				str += "\n"
			}
			str += fmt.Sprintf("[%d]: %d/ %d", i+1, v.GetValue(), v.GetCapacity())
		}
	}
	return str
}
