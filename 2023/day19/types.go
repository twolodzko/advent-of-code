package main

import "fmt"

type Part map[rune]int

func (this Part) Sum() int {
	result := 0
	for _, k := range [4]rune{'x', 'm', 'a', 's'} {
		result += this[k]
	}
	return result
}

func (this Part) String() string {
	return fmt.Sprintf("{x=%d,m=%d,a=%d,s=%d}", this['x'], this['m'], this['a'], this['s'])
}

type Check interface {
	Check(Part) bool
}

type Greater struct {
	key   rune
	value int
}

func (this Greater) String() string {
	return fmt.Sprintf("%s>%d", string(this.key), this.value)
}

func (this Greater) Check(part Part) bool {
	return part[this.key] > this.value
}

type Less struct {
	key   rune
	value int
}

func (this Less) String() string {
	return fmt.Sprintf("%s<%d", string(this.key), this.value)
}

func (this Less) Check(part Part) bool {
	return part[this.key] < this.value
}

type Rule interface {
	Applies(Part) bool
	Destination() string
}

type Conditional struct {
	destination string
	condition   Check
}

func (this Conditional) String() string {
	return fmt.Sprintf("%s:%s", this.condition, this.destination)
}

func (this Conditional) Applies(part Part) bool {
	return this.condition.Check(part)
}

func (this Conditional) Destination() string {
	return this.destination
}

type SendTo struct {
	destination string
}

func (this SendTo) String() string {
	return this.destination
}

func (this SendTo) Applies(Part) bool {
	return true
}

func (this SendTo) Destination() string {
	return this.destination
}

type Workflows map[string][]Rule
