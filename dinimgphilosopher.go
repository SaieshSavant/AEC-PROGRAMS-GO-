package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numphilosophers = 3
	maxdiningcycles = 2
)

type philosopher struct {
	id                  int
	leftfork, rightfork *sync.Mutex
	diningcycles        int
}

type diningtable struct {
	philosophers []*philosopher
	waiter       *sync.Mutex
}

func (p *philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Second)
}

func (p *philosopher) eat() {
	p.leftfork.Lock()
	p.rightfork.Lock()

	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Second)
	p.rightfork.Unlock()
	p.leftfork.Unlock()
	p.diningcycles++
}

func (p *philosopher) dine(table *diningtable){
	for p.diningcycles < maxdiningcycles{
		p.think()
		table.waiter.Lock()
		p.eat()
		table.waiter.Unlock()
	}
}


func main(){
	table :=&diningtable{
		philosophers: make([]*philosopher, numphilosophers),
        waiter: &sync.Mutex{},
	}

	forks :=make([]*sync.Mutex,numphilosophers)
	for i:=0;i<numphilosophers;i++{
		forks[i]=&sync.Mutex{}
	}

	for i:=0;i<numphilosophers;i++{
		table.philosophers[i]=&philosopher{
			id: i,
			leftfork: forks[i],
			rightfork: forks[(i+1)%numphilosophers],
		}
	}

    var wg sync.WaitGroup
	wg.Add(numphilosophers)
	for i:=0;i<numphilosophers;i++{
		go func (p*philosopher)  {
			defer wg.Done()
			p.dine(table)

			
		}(table.philosophers[i])
	}
	wg.Wait()
}

/*
Philosopher 2 is thinking
Philosopher 0 is thinking
Philosopher 1 is thinking
Philosopher 2 is thinking
Philosopher 2 is thinking
Philosopher 1 is thinking
Philosopher 1 is thinking
Philosopher 0 is thinking
Philosopher 0 is thinking
Philosopher 2 is thinking
Philosopher 1 is thinking
Philosopher 0 is thinking

*/