package persistance

import (
	basemodel1 "github.com/fullstack-lang/test_mdsm/base_model1"
)

// hook to the exported basemodel1 instance, pre-allocated so links can be wired in any Restore order
var E__basemodel1__System_000000 = &basemodel1.System{}

func RestoreBaseModel1() {
	__model__System_000000_ := E__basemodel1__System_000000
	__model__System_000000_.Name = "System1"
}
