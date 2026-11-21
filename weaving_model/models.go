package weaving_model

import (
	basemodel1 "github.com/fullstack-lang/test_mdsm/base_model1"
	basemodel2 "github.com/fullstack-lang/test_mdsm/base_model2"
)

type System struct {
	Name string

	System1 *basemodel1.System
	System2 *basemodel2.System
}

func (s *System) GetName() string {
	return s.Name
}
