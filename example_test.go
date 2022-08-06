package gmap_test

import (
	"encoding/json"
	"fmt"

	"github.com/kkHAIKE/gmap"
)

func Example() {
	data := []byte(`{
		"Foo": [{
			"Bar": {
				"hello.world": [1,2]
			}
		}]
	}`)

	var m map[string]interface{}
	_ = json.Unmarshal(data, &m)

	m2 := gmap.GetIgnoreCase(m, "foo.0.bar")
	fmt.Println(m2.MustMap().Get(`hello\.world.1`).MustInt())
	// Output: 2
}
