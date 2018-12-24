package actionQueue

import (
	"fmt"
	"testing"
)

type TestObject struct{}
type TestRequest struct {
	Object Object
	err    error
	action func() (Object, error)
}

func (t *TestRequest) Run() {
	summoner, err := t.action()
	fmt.Println(summoner, err)
}

func TestActionQueue(t *testing.T) {
	que := &ActionQueue{}
	que.Start()
	action := func() (Object, error) {
		fmt.Println("Test")
		return &TestObject{}, nil
	}

	newRequest := &TestRequest{
		action: action,
	}

	que.Add(newRequest)
	que.wg.Wait()
	que.Stop()
}

// type (
// 	testAction struct {
// 		Name string
// 	}
// 	testObject struct{}
// )

// func (t *testAction) Action() (Object, error) {
// 	fmt.Println(t.Name)
// 	return &testObject{}, nil
// }

// func TestQueue(t *testing.T) {
// 	t.Log("Given I created a ActionQueue")
// 	{
// 		que := &ActionQueue{}

// 		que.Start()
// 		// Create Action
// 		action := &testAction{}
// 		que.Add(action)
// 		que.wg.Wait()
// 		que.Stop()
// 	}

// }

// func TestConcurrenyQue(t *testing.T) {
// 	que := &ActionQueue{}
// 	que.Start()
// 	go func() {
// 		action := &testAction{
// 			Name: "Test1",
// 		}
// 		que.Add(action)
// 	}()

// 	go func() {
// 		action := &testAction{
// 			Name: "Test2",
// 		}
// 		que.Add(action)
// 	}()

// 	que.wg.Wait()
// 	que.Stop()
// }
