package persistance

import basemodel2 "github.com/fullstack-lang/test_mdsm/base_model2"

// hook to the exported basemodel2 instance, pre-allocated so links can be wired in any Restore order
var E__basemodel2__System_000000 = &basemodel2.System{}

func RestoreBaseModel2() {
	__model__System_000000_ := E__basemodel2__System_000000
	__model__System_000000_.Name = "System2"
}
