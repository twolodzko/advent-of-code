package main_test

// import (
// 	m "day20"
// 	"testing"
// )

// func test_module(module m.Module, expected m.Pulse, should_send bool, t *testing.T) {
// 	result, ok := module.Send()
// 	if ok != should_send {
// 		t.Errorf("expected send status to be %v, got %v", should_send, ok)
// 	}
// 	if result != expected {
// 		t.Errorf("expected pulse %v, got %v", expected, result)
// 	}
// }

// func Test_Broadcaster(t *testing.T) {
// 	module := &m.Broadcaster{}

// 	// no input
// 	test_module(module, false, true, t)

// 	module.Receive(true)
// 	test_module(module, true, true, t)

// 	module.Receive(false)
// 	test_module(module, false, true, t)
// }

// func Test_FlipFlop(t *testing.T) {
// 	module := m.NewFlipFlop("test")
// 	test_module(module, false, false, t)

// 	module.Receive(true)
// 	test_module(module, false, false, t)

// 	module.Receive(false)
// 	test_module(module, true, true, t)

// 	module.Receive(false)
// 	test_module(module, false, true, t)

// 	module.Receive(true)
// 	test_module(module, false, false, t)

// 	module.Receive(false)
// 	test_module(module, true, true, t)

// 	module.Receive(true)
// 	test_module(module, true, false, t)
// }

// func Test_Conjunction(t *testing.T) {
// 	module := m.NewConjunction("test")
// 	test_module(module, true, true, t)

// 	module.Receive(true)
// 	test_module(module, false, true, t)

// 	// fresh state
// 	test_module(module, true, true, t)

// 	module.Receive(true)
// 	module.Receive(true)
// 	module.Receive(true)
// 	test_module(module, false, true, t)

// 	// fresh state
// 	test_module(module, true, true, t)

// 	module.Receive(true)
// 	module.Receive(false)
// 	module.Receive(true)
// 	test_module(module, true, true, t)
// }
