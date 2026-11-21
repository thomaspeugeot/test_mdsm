package main

import (
	"fmt"

	"github.com/fullstack-lang/test_mdsm/persistance"
)

func main() {
	fmt.Println("Hello world")

	persistance.RestoreBaseModel1()
	persistance.RestoreBaseModel2()
	persistance.RestoreWeavingModel()
	persistance.Inspect()

}
