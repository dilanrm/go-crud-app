package main

import "sync"

type Store struct {
	mu sync.Mutex
	people map[int]Person
	nextID int
}

var store = Store{
	people: make(map[int]Person),
	nextID: 1,
}

func (s *Store) CreatePerson(p Person) Person {
	s.mu.Lock()
    defer s.mu.Unlock()
    p.ID = s.nextID
    s.people[p.ID] = p
    s.nextID++
    return p
}

func (s *Store) UpdatePerson(id int, p Person) bool {
	s.mu.Lock()
    defer s.mu.Unlock()
	_, exists := s.people[id]
    if exists {
		p.ID = id
        s.people[id] = p
    }
    return exists
}

func (s *Store) DeletePerson(id int) bool {
	s.mu.Lock()
    defer s.mu.Unlock()
    _, exists := s.people[id]
    if exists {
        delete(s.people, id)
    }
    return exists
}

func (s *Store) GetPerson(id int) (Person) {
	s.mu.Lock()
    defer s.mu.Unlock()
	p := s.people[id]
	return p
}

func (s *Store) GetAllPeople() []Person {
	s.mu.Lock()
    defer s.mu.Unlock()
    people := make([]Person, 0, len(s.people))
    
    for _, p := range s.people {
        people = append(people, p)
    }
    return people
}