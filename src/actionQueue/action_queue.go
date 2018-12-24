package actionQueue

import "fmt"
import "sync"

type (
	ActionQueue struct {
		que       chan Action
		QueActive bool
		wg        sync.WaitGroup
	}

	// Action interface {
	// 	Action() (Object, error)
	// }
	Action func() (Object, error)

	Object interface{}
	// ActionGetPlayer struct{} // This can be move to other file

	// ActionGetMatch struct{}
)

// type Action func() (*Summoner, error)

func (q *ActionQueue) Add(a Action) {
	if q.QueActive {
		q.wg.Add(1)
		q.que <- a
	} else {
		fmt.Println("Queue is closed")
	}

}

func (q *ActionQueue) Start() {
	q.QueActive = true
	// q.wg = sync.WaitGroup
	q.que = make(chan Action, 10)
	go consume(q.que, &q.wg)

}

func consume(que <-chan Action, wg *sync.WaitGroup) {
	// defer close(que) // Close when channel is empty
	for {
		action, ok := <-que
		// fmt.Println(wg)
		if ok {
			action()  // Perform Action
			wg.Done() // Decrease Semaphore
		} else {
			return
		}
	}
	// for a, ok := <- que; ok{ // Pull action from channel
	// 	// select {
	// 	// 	case action <- que:

	// 	// }
	// 	a.Action() // Perform the action
	// }

}

func (q *ActionQueue) Wait() {
	q.wg.Wait()
}

func (q *ActionQueue) Stop() {
	q.QueActive = false // Stop receiving Action
	close(q.que)
}
