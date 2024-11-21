package caseofficer

import (
	"fmt"
	"github.com/advanced-go/common/messaging"
)

type dispatchT struct{}

func newTestDispatcher() dispatcher {
	d := new(dispatchT)
	return d
}

func (d *dispatchT) setup(_ *caseOfficer, _ string) {}

func (d *dispatchT) dispatch(agent *caseOfficer, event string) {
	switch event {
	case messaging.DataChangeEvent:
		agent.handler.Trace(agent, event, "Broadcast() -> calendar data change event")
	}
}

func ExampleTestDispatcher() {
	fmt.Printf("test: TestDispatch() \n")

	//Output:
	//test: TestDispatch()

}