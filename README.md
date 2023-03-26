# gmap
[![GoDoc](https://godoc.org/github.com/kkHAIKE/gmap?status.svg)](https://godoc.org/github.com/kkHAIKE/gmap)

This is a simple library that allows you to retrieve values from a `map[string]interface{}` or `[]interface{}` using a path

It is similar to objx but simpler, and it includes support for case-insensitive keys.

# usage
```go
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
```

# link
- [objx](https://github.com/stretchr/objx)
- [gjson](https://github.com/tidwall/gjson)
