// A synchornized queue consists of elements that need to be processed in 
// a particula sequence. Passenger queue and ticket processing queues 
// ar types of syncrhonized queues as follows 
package main 

import (
	"fmt" 
	"time" 
	"math/rand"
)

// constants 
const (
	messagePassStart = iota 
	messageTicketStart 
	messagePassEnd 
	messageTicketEnd 
)

// Queue class 
type Queue struct{
	waitPass int 
	waitTicket int 
	playPass bool 
	playTicket bool 
	queuePass chan int 
	queueTicket chan int 
	message chan int 
}

// THe New method 
// the new method on Queue initializes message, queuePass and queueTicket with nil 
// values. the make method creates a Queue with a chan integer parameter as follows 
func (queue *Queue) New(){
	queue.message = make(chan int) 
	queue.queuePass = make(chan int) 
	queue.queueTicket = make(chan int)
}


// StartTicketIssue method 
// the StartTicketIssue method on Qeueu sends messageTicketStart 
// to the message queue and queueTicket receives the message. 
// the ticket issue is started by sending messages to the queue, 
// as follows 
func (Queue *Queue) StartTicketIssue(){
	Queue.message <- messageTicketStart 
	<-Queue.queueTicket
}

// EndTicketIssue method 
// the EndTicketIssue method finishes the issuing of a ticket 
// to a passenger standing in the queue. In the following 
// code, the EndTicketIssue method on Queue sends 
// messageTicketEnd to the message queeu. 

func (queue *Queue) EndTicketIssue(){
	queue.message <- messageTicketEnd
}

// ticketIssue method 
// this method starts and finishes the issuing of a 
// ticket to the passenger. 
// The ticketIssue method invokes the StartTicketIssue 
// and EndTicketIssue methods after Sleep calls for 
// 10 seconds and two seconds 
func ticketIssue(queue *Queue){
	for{
		// sleep up to 10 seconds 
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond) 
		fmt.Println("Ticket issue starts") 
		// Sleep up to 2 seconds 
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) 
		fmt.Println("Ticket issue ends") 
		queue.EndTicketIssue()
	}
}

// the StartPass method 
// The StartPass method starts the pasenger queue moving toward the ticket counter 
// the StartPass method on Queue sends messagePassStart to the message queue 
// and queuePass receives the message. Passengers are moved into the queue 
// as follows 

func (q *Queue) StartPass(){
	q.message <- messagePassStart 
	<-q.queuePass
}

// The EndPass method 
// the EndPass method stops the passenger queue moving toward the ticket 
// counter. The EndPass method on Queue sends 
// messagePassEnd to the message queue in the following code 

// EndPass ends the Pass Queue 
func (q *Queue) EndPass(){
	q.message <- messagePassEnd
}

// Passenger method 
// THe passenger method starts and ends the passenger movement 
// to the queue. The passenger method invokes the StartPass 
// method, and the EndPass method ends after sleep calls for 10 
// seconds and two seconds. 

func passenger(q *Queue){
	fmt.Println("Starting the Passenger Queue...") 
	for{
		fmt.Println("starting the processing...") 
		// sleep up to 10 seconds 
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond) 
		q.StartPass() 
		fmt.Println(" Passenger starts") 
		// Sleep up to 2 seconds 
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) 
		fmt.Println(" Passenger ends") 
		q.EndPass()
	}
}

// the main method calls the passgner and ticketIssue methods 
// after creating a queue. the passenger enters into the queue 
// and a ticket is issued at the counter in the processing queue 
func main(){
	var q *Queue = &Queue{} 
	q.New() 
	fmt.Println(q) 
	
	for i:= 0; i < 10; i++{
		go passenger(q)
	}

	for j:= 0; j < 5; j++{
		go ticketIssue(q)
	}

	select{}
}
