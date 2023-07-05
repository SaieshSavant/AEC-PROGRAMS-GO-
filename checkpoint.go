package main

import (
	"fmt"
	"sync"
	"time"
)
func worker(id int,checkpoint chan bool,resume chan bool,wg*sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("worker %d:starting\n",id)
	time.Sleep(time.Duration(id)*time.Second)
    
	fmt.Printf("worker %d:checkout reached\n",id)
	checkpoint <- true
	<-resume
	fmt.Printf("worker %d: Resuming\n",id)



}

func main(){
	numworkers:=5
	checkpoint:=make(chan bool)
	resume:=make(chan bool)
	var wg sync.WaitGroup

	for i:=1;i<numworkers;i++{
		wg.Add(1)
		go worker(i,checkpoint,resume,&wg)
	}
   
	for i:=1;i<numworkers;i++{
		<-checkpoint
	}

	fmt.Println("All workers reached checkpoint")
	fmt.Println("Resuming all workers")

	for i:=1;i<numworkers;i++{
		resume<-true

	}

	wg.Wait()
    fmt.Println("All workers completed their work")
}
/*
worker 4:starting
worker 1:starting
worker 2:starting
worker 3:starting
worker 1:checkout reached
worker 2:checkout reached
worker 3:checkout reached
worker 4:checkout reached
All workers reached checkpoint
Resuming all workers
worker 4: Resuming
worker 1: Resuming
worker 2: Resuming
worker 3: Resuming
All workers completed their work
*/
