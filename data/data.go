package data

import (
	"sync"
)

// User .
type User struct {
	ID   int
	Name string
}

// State .
type State struct {
	User User
}

// Store .
var Store *store

type store struct {
	state       State
	mu          sync.RWMutex
	subscribers []chan State
}

func init() {
	Store = &store{
		state: State{
			User: User{ID: 5, Name: "Gopher"},
		},
	}
}

// Subscribe .
func (s *store) Subscribe() chan State {
	ch := make(chan State)
	s.mu.Lock()
	s.subscribers = append(s.subscribers, ch)
	state := s.state
	s.mu.Unlock()
	go func() { ch <- state }()
	return ch
}

// Unsubscribe .
func (s *store) Unsubscribe(ch chan State) {
	s.mu.Lock()
	for i, subscriber := range s.subscribers {
		if subscriber == ch {
			s.subscribers[i] = s.subscribers[len(s.subscribers)-1]
			s.subscribers = s.subscribers[:len(s.subscribers)-1]
			close(ch)
			break
		}
	}
	s.mu.Unlock()
}

func (s *store) dispatch() {
	for _, subscriber := range s.subscribers {
		select {
		case subscriber <- s.state:
		default:
		}
	}
}

// SetUser .
func (s *store) SetUser(user User) {
	s.mu.Lock()
	s.state.User = user
	s.dispatch()
	s.mu.Unlock()
}
