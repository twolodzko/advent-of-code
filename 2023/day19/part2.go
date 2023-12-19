package main

import "fmt"

type Range struct {
	min, max uint64
}

func (this Range) Size() uint64 {
	if this.max < this.min {
		return 0
	}
	return this.max - this.min + 1
}

func (this Range) Intersection(other Range) Range {
	lower := max(this.min, other.min)
	upper := min(this.max, other.max)
	return Range{lower, upper}
}

func (this Range) String() string {
	return fmt.Sprintf("%d:%d", this.min, this.max)
}

func (this Range) TrimMax(val uint64) Range {
	return Range{this.min, min(this.max, val)}
}

func (this Range) TrimMin(val uint64) Range {
	return Range{max(this.min, val), this.max}
}

type PartRanges map[rune]Range

func NewPartRanges() PartRanges {
	ranges := make(PartRanges)
	for _, k := range [4]rune{'x', 'm', 'a', 's'} {
		ranges[k] = Range{1, 4000}
	}
	return ranges
}

func (this PartRanges) Copy() PartRanges {
	copy := make(PartRanges)
	for k, v := range this {
		copy[k] = v
	}
	return copy
}

func (this PartRanges) String() string {
	return fmt.Sprintf("{x=%s,m=%s,a=%s,s=%s}", this['x'], this['m'], this['a'], this['s'])
}

func (this PartRanges) Size() uint64 {
	var result uint64 = 1
	for _, r := range this {
		result *= r.Size()
	}
	return result
}

func (this PartRanges) Intersection(other PartRanges) PartRanges {
	result := make(PartRanges)
	for k := range this {
		result[k] = this[k].Intersection(other[k])
	}
	return result
}

func (this PartRanges) TrimMax(key rune, val uint64) PartRanges {
	copy := make(PartRanges)
	for k, v := range this {
		copy[k] = v
	}
	copy[key] = copy[key].TrimMax(val)
	return copy
}

func (this PartRanges) TrimMin(key rune, val uint64) PartRanges {
	copy := make(PartRanges)
	for k, v := range this {
		copy[k] = v
	}
	copy[key] = copy[key].TrimMin(val)
	return copy
}

func explore(workflow string, workflows Workflows, part PartRanges, final string) []PartRanges {
	if workflow == final {
		return []PartRanges{part}
	}

	var paths []PartRanges
	rules := workflows[workflow]
	for _, rule := range rules {
		switch rule := rule.(type) {
		case Conditional:
			switch cond := rule.condition.(type) {
			case Less:
				value := uint64(cond.value)
				tmp := explore(rule.Destination(), workflows, part.TrimMax(cond.key, value-1), final)
				part = part.TrimMin(cond.key, value)
				paths = append(paths, tmp...)
			case Greater:
				value := uint64(cond.value)
				tmp := explore(rule.Destination(), workflows, part.TrimMin(cond.key, value+1), final)
				part = part.TrimMax(cond.key, value)
				paths = append(paths, tmp...)
			}
		case SendTo:
			paths = append(paths, explore(rule.Destination(), workflows, part, final)...)
		}
	}
	return paths
}

func part2(workflows Workflows) {
	paths := explore("in", workflows, NewPartRanges(), "A")

	var size, intersection, result uint64
	for i := 0; i < len(paths); i++ {
		size = paths[i].Size()
		result += size
		for j := i + 1; j < len(paths); j++ {
			intersection = paths[i].Intersection(paths[j]).Size()
			result -= intersection
		}
	}

	fmt.Println(result)

	// const whole uint64 = 4000 * 4000 * 4000 * 4000
	// fmt.Println(whole - result)
}
