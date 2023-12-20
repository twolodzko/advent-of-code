package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pulse bool

const (
	Low  Pulse = false
	High Pulse = true
)

func (this Pulse) String() string {
	if this == High {
		return "high"
	}
	return "low"
}

type Module interface {
	Receive(string, Pulse)
	Send() (Pulse, bool)
}

type Broadcaster struct {
	state Pulse
}

func (this Broadcaster) String() string {
	return "broadcaster"
}

func (this *Broadcaster) Receive(from string, pulse Pulse) {
	this.state = pulse
}

func (this Broadcaster) Send() (Pulse, bool) {
	return this.state, true
}

type FlipFlop struct {
	name  string
	state Pulse
	send  bool
}

func (this FlipFlop) String() string {
	// return fmt.Sprintf("%%%s", this.name)
	return fmt.Sprint(this.state)
}

func (this *FlipFlop) Receive(from string, pulse Pulse) {
	if pulse == Low {
		this.state = !this.state
		this.send = true
	} else {
		this.send = false
	}
}

func (this *FlipFlop) Send() (Pulse, bool) {
	defer func() { this.send = true }()
	return this.state, this.send
}

type Conjunction struct {
	name   string
	states map[string]Pulse
}

func (this Conjunction) String() string {
	// return fmt.Sprintf("&%s", this.name)
	return fmt.Sprint(this.states)
}

func (this *Conjunction) Reset() {
	for key := range this.states {
		this.states[key] = Low
	}
}

func (this *Conjunction) Receive(from string, pulse Pulse) {
	this.states[from] = pulse
}

func (this *Conjunction) Send() (Pulse, bool) {
	for _, state := range this.states {
		if state == Low {
			return High, true
		}
	}
	return Low, true
}

type Untyped struct {
	name string
}

func (this Untyped) String() string {
	return this.name
}

func (this Untyped) Receive(from string, pulse Pulse) {}

func (this Untyped) Send() (Pulse, bool) {
	return false, false
}

func ModuleFrom(name string) (Module, string) {
	if name == "broadcaster" {
		return &Broadcaster{}, name
	}
	switch name[0] {
	case '%':
		name = name[1:]
		return NewFlipFlop(name), name
	case '&':
		name = name[1:]
		return NewConjunction(name), name
	default:
		return &Untyped{name}, name
	}
}

func NewFlipFlop(name string) Module {
	return &FlipFlop{name, false, false}
}

func NewConjunction(name string) Module {
	return &Conjunction{name, make(map[string]Pulse)}
}

type Broadcast struct {
	modules       map[string]Module
	transmissions map[string][]string
}

func NewBroadcast() Broadcast {
	return Broadcast{make(map[string]Module), make(map[string][]string)}
}

func (this *Broadcast) Transmit(transmission Transmission) (Pulse, bool) {
	transmitter := this.modules[transmission.from]
	receiver := this.modules[transmission.to]

	pulse, ok := transmitter.Send()
	if ok {
		// fmt.Printf("%s -%s-> %s\n", transmitter, pulse, receiver)
		receiver.Receive(transmission.from, pulse)
	}
	return pulse, ok
}

type Transmission struct {
	from, to string
}

func (this Broadcast) NewQueue(from string) []Transmission {
	var transmissions []Transmission
	for _, name := range this.transmissions[from] {
		transmissions = append(transmissions, Transmission{from, name})
	}
	return transmissions
}

func (this *Broadcast) Broadcast() (int, int) {
	var (
		low, high int
		from      string         = "broadcaster"
		queue     []Transmission = this.NewQueue(from)
		new_queue []Transmission
	)

	for len(queue) != 0 {
		transmission := queue[0]
		queue = queue[1:]

		pulse, ok := this.Transmit(transmission)
		if ok {
			switch pulse {
			case Low:
				low++
			case High:
				high++
			}
			new_queue = append(new_queue, this.NewQueue(transmission.to)...)
		}

		if len(queue) == 0 {
			queue = new_queue
			new_queue = nil
		}
	}

	// + 1 for the button
	return low + 1, high
}

func (this Broadcast) String() string {
	var out []string
	for from, to := range this.transmissions {
		out = append(out, fmt.Sprintf("%s -> %s", from, strings.Join(to, ", ")))
	}
	return strings.Join(out, "\n")
}

func (this *Broadcast) ParseAdd(line string) {
	fields := strings.Split(line, " -> ")
	mod, from := ModuleFrom(fields[0])
	if _, ok := this.modules[from]; !ok {
		this.modules[from] = mod
	}
	var to []string
	for _, name := range strings.Split(fields[1], ", ") {
		to = append(to, name)
	}
	this.transmissions[from] = to
}

func BroadcastFrom(scanner *bufio.Scanner) Broadcast {
	broadcast := NewBroadcast()
	for scanner.Scan() {
		line := scanner.Text()
		broadcast.ParseAdd(line)
	}

	// update module information for modules not found on lhs of the transmission list
	for from, to := range broadcast.transmissions {
		for _, name := range to {
			if _, ok := broadcast.modules[name]; !ok {
				mod, name := ModuleFrom(name)
				broadcast.modules[name] = mod
			}
			switch mod := broadcast.modules[name].(type) {
			case *Conjunction:
				mod.states[from] = Low
			}
		}
	}

	return broadcast
}

type State struct {
	iteration int
	state     map[string]Module
	low, high int
}

func part1(broadcast Broadcast) {
	var (
		i         int
		low, high int
	)
	memory := make(map[string]State)

	for i < 1000 {
		// fmt.Println(i)
		key := fmt.Sprint(broadcast.modules)
		if state, ok := memory[key]; ok {
			// fmt.Printf("it was already seen on iteration %d\n", state.iteration)
			jump := i - state.iteration
			// jump := (1000 - i) / cycle

			if jump > 0 && i+jump < 1000 {
				fmt.Printf("jump = %d\n", jump)
				i += jump
				low += (low - state.low)
				high += (high - state.high)
				// broadcast.modules = state.state
				jump = 0
				continue
			}
		} else {
			clone := make(map[string]Module)
			for k, v := range broadcast.modules {
				clone[k] = v
			}
			memory[key] = State{i, clone, low, high}
		}

		l, h := broadcast.Broadcast()
		low += l
		high += h
		i++

		// fmt.Println(low, high)
	}
	fmt.Println(low * high)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	broadcast := BroadcastFrom(scanner)

	// fmt.Println(broadcast.modules)
	part1(broadcast)

	// fmt.Println(broadcast.Broadcast())
	// fmt.Println()
	// fmt.Println(broadcast.Broadcast())
	// fmt.Println()
	// fmt.Println(broadcast.Broadcast())
}
