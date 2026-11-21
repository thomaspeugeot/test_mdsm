package persistance

import "fmt"

func Inspect() {
	// inspect the model
	fmt.Println("Inspecting the model")

	// base model 1
	fmt.Println("E__basemodel1__System_000000", E__basemodel1__System_000000)
	if E__basemodel1__System_000000 != nil && E__basemodel1__System_000000.IF_to_system_weawing_model != nil {
		fmt.Println("E__basemodel1__System_000000.IF_to_system_weawing_model", E__basemodel1__System_000000.IF_to_system_weawing_model)
		fmt.Println("E__basemodel1__System_000000.IF_to_system_weawing_model.GetName()", E__basemodel1__System_000000.IF_to_system_weawing_model.GetName())
	} else {
		fmt.Println("E__basemodel1__System_000000.IF_to_system_weawing_model is nil")
	}

	// base model 2

	fmt.Println("E__basemodel2__System_000000", E__basemodel2__System_000000)
	if E__basemodel2__System_000000 != nil && E__basemodel2__System_000000.IF_to_system_weawing_model != nil {
		fmt.Println("E__basemodel2__System_000000.IF_to_system_weawing_model", E__basemodel2__System_000000.IF_to_system_weawing_model)
		fmt.Println("E__basemodel2__System_000000.IF_to_system_weawing_model.GetName()", E__basemodel2__System_000000.IF_to_system_weawing_model.GetName())
	} else {
		fmt.Println("E__basemodel2__System_000000.IF_to_system_weawing_model is nil")
	}

	fmt.Println("E__weaving_model__System_₀₀₀₀₀₀", E__weaving_model__System_000000)
}
