package inspect_test

import "github.com/resyahrial/go-commerce/pkg/inspect"

func ExampleDo() {
	var1 := "string"
	var2 := true
	var3 := []int{42}
	inspect.Do(var1, var2, var3)
}
