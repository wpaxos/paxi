package paxi

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrStateMachineExecution = errors.New("StateMachine execution error")
)

type Key int64
type Value int64

const NIL Value = 0

type Operation uint8

const (
	NOOP Operation = iota
	PUT
	GET
	DELETE
	RLOCK
	WLOCK
)

type Command struct {
	Operation Operation
	Key       Key
	Value     Value
}

func (c Command) String() string {
	if c.Operation == GET {
		return fmt.Sprintf("Get{key=%v}", c.Key)
	}
	return fmt.Sprintf("Put{key=%v, val=%v}", c.Key, c.Value)
}

func (c *Command) IsRead() bool {
	return c.Operation == GET
}

type StateMachine struct {
	lock  *sync.Mutex
	Store map[Key]Value
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		lock:  new(sync.Mutex),
		Store: make(map[Key]Value),
	}
}

func (s *StateMachine) Execute(commands ...Command) (Value, error) {
	for _, c := range commands {
		switch c.Operation {
		case PUT:
			s.Store[c.Key] = c.Value
			return c.Value, nil
		case GET:
			if value, present := s.Store[c.Key]; present {
				return value, nil
			}
		}
	}
	return NIL, ErrStateMachineExecution
}

func Conflict(gamma *Command, delta *Command) bool {
	if gamma.Key == delta.Key {
		if gamma.Operation == PUT || delta.Operation == PUT {
			return true
		}
	}
	return false
}

func ConflictBatch(batch1 []Command, batch2 []Command) bool {
	for i := 0; i < len(batch1); i++ {
		for j := 0; j < len(batch2); j++ {
			if Conflict(&batch1[i], &batch2[j]) {
				return true
			}
		}
	}
	return false
}
