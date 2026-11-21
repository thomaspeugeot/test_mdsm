package basemodel2

type IF_to_system_weawing_model interface {
	GetName() string // or whatever interface is needed to display the interface content in a base model diagram/semantic rule
}

// gong:export
type System struct {
	Name string

	// IF_to_system_weawing_model is populated  at runtime at the end to probe if there is the interface is fulfilled
	IF_to_system_weawing_model IF_to_system_weawing_model
}
