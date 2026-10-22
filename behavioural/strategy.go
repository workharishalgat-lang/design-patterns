package behavioural

import "fmt"

type StrategyBase interface {
	Execute()
}

type ApplicationPlaneStratagy struct {
}

func (ApplicationPlaneStratagy) Execute() {
	fmt.Println("application plane specific strategy")
}

type DataPlaneSpecificStratage struct {
}

func (DataPlaneSpecificStratage) Execute() {

}

type StrategyInvoker struct {
	base StrategyBase
}

func (sb StrategyInvoker) Invoker() {
	sb.base.Execute()
}
