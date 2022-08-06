# gmap
[![GoDoc](https://godoc.org/github.com/kkHAIKE/gmap?status.svg)](https://godoc.org/github.com/kkHAIKE/gmap)

simple library for get value from map[string]interface{} or []interface{} use path.

It's use like [objx](https://github.com/stretchr/objx) but more simple and add case-insensitive key support.

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
