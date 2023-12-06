package main

type Range struct {
	min, max int
}

func NewRange(min, max int) Range {
	return Range{min, max}
}

func (this Range) Add(value int) Range {
	return Range{this.min + value, this.max + value}
}

func (this Range) Before(other Range) (Range, bool) {
	if this.max < other.min {
		return this, true
	}
	if this.min > other.min {
		return Range{}, false
	}
	r := Range{this.min, min(this.max, other.min-1)}
	return r, r.min <= r.max
}

func (this Range) Intersection(other Range) (Range, bool) {
	if this.max < other.min || this.min > other.max {
		return Range{}, false
	}
	r := Range{max(this.min, other.min), min(this.max, other.max)}
	return r, r.min <= r.max
}

func (this Range) After(other Range) (Range, bool) {
	if this.min > other.max {
		return this, true
	}
	if this.max < other.min {
		return Range{}, false
	}
	r := Range{max(this.min, other.max+1), this.max}
	return r, r.min <= r.max
}

func (this Range) Equal(other Range) bool {
	return this.min == other.min && this.max == other.max
}

func (this Range) IsValid() bool {
	return this.min <= this.max
}

func (this Range) Min(other Range) Range {
	if this.min < other.min {
		return this
	}
	return other
}
