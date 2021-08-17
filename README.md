**RestN** ( **rest** **n**ormalization )

Purpose of this library is more control under responses.

Install it:<br>
`go get https://github.com/TrueGameover/RestN`

<br>
<br>
IResponseNormalizer is main interface that should be implemented by your normalizer. 
Normalizer allows you to control how objects are converted to "raw" format.

For example:

```
type Test struct {
    time time.Time,
    another struct { gg int }
}
```

How "time" field would be serialized to json? Usually this behavior needs explicit conversion to string in every place.
Normalizers resolve this pain.

<br>
Let's implement simple normalizer:

```
import "RestN/rest"

type TimeNormalizer struct {
    rest.IResponseNormalizer
}

func (n TimeNormalizer) Normalize(object interface{}, normalize rest.NormalizeMethod, options rest.Options, depth int) interface{} {
    // depth is used for control of normalization's depth
    // options - your params that passed to normalizers 
    
    if test, ok := object.(Test); ok {
            dict := map[string]interface{}{
                "Time": test.time.Format(time.RFC3339),
                "Another": normalize(test.another, options, depth), // if need normalize another struct deeper
            }

	    return dict
	}

    return object
}

func (n TimeNormalizer) Support(object interface{}) (ok bool) {
    _, ok = object.(rest.Locale)
    return
}

```

First time initialize normalizers :<br>
`normalizers.Init()`
<br>
<br>
Then register it:<br>
`rest.RegisterNormalizer(TimeNormalizer{})`
<br>
<br>
Create response:<br>

```
r := rest.RestResponse{}
r.SetBody(Test{ time: time.Now(), another: {gg: 5} })
resp.SetNormalizationOption("custom_key", "custom_value") // your customization
println(r.NormalizeResponse())
```

It will be normalized\customized by your normalizer every time.



