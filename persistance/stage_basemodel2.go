package persistance

import basemodel2 "github.com/fullstack-lang/test_mdsm/base_model2"

// hook to the exported to the basemodel2 instance
// notice the double indirection
var E__basemodel2__System_000000 *basemodel2.System

func RestoreBaseModel2() {
	__model__System_000000_ := &basemodel2.System{Name: "System2"}

	// initiation of the hook
	E__basemodel2__System_000000 = __model__System_000000_
}
