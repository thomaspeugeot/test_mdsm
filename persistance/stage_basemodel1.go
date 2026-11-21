package persistance

import (
	basemodel1 "github.com/fullstack-lang/test_mdsm/base_model1"
)

// hook to the exported to the basemodel1 instance
// notice the double indirection
var E__basemodel1__System_000000 *basemodel1.System

func RestoreBaseModel1() {
	__model__System_000000_ := &basemodel1.System{Name: "System1"}

	// initiation of the hook
	E__basemodel1__System_000000 = __model__System_000000_
}
