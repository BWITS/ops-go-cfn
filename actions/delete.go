package actions

import (
	"github.com/MYOB-Technology/ops-go-cfn/stack"
)

// Delete ...
func Delete(svc *stack.Service, name string) (status string, err error) {
	ctrl := stack.Controller(svc)
	waiter := stack.Waiter(svc)

	stackOutput, err := ctrl.Delete(&name)
	if err != nil {
		return
	}

	return stackOutput.String(), waiter.WaitDelete(&name)
}
