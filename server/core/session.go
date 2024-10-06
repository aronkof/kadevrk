package core

import (
	"fmt"
)

func (s *Rk) AddClient() error {
	if s.currClients.Load() == s.maxClients {
		return fmt.Errorf("reached max number of clients %d", s.maxClients)
	}

	s.currClients.Add(1)

	return nil
}

func (s *Rk) RemoveClient() {
	if s.currClients.Load() == 0 {
		return
	}

	s.currClients.Add(-1)
}
