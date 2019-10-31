package singleton

import (
	"sync"
)

var (
	p *person
	once sync.Once
)

func init() {
	once.Do(func() {
		p = &person{}
	})
}

func GetInstance() *person {
	return p
}

type person struct {
	name string
	age int
	mux sync.Mutex
}

func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}