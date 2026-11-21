package persistance

import "github.com/fullstack-lang/test_mdsm/weaving_model"

// for debug only, a Stage would make it available
var E__weaving_model__System_000000 *weaving_model.System

func RestoreWeavingModel() {
	__model__System_000000_ := &weaving_model.System{Name: "WeavingModel"}

	__model__System_000000_.System1 = E__basemodel1__System_000000
	E__basemodel1__System_000000.IF_to_system_weawing_model = __model__System_000000_

	__model__System_000000_.System2 = E__basemodel2__System_000000
	E__basemodel2__System_000000.IF_to_system_weawing_model = __model__System_000000_

	// for overcoming lack of stage
	E__weaving_model__System_000000 = __model__System_000000_
}
