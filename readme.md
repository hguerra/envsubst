# envsubst

`envsubst` is a Go package for expanding variables in a string using `${var}` syntax.
Includes support for bash string replacement functions. **This fork allows you to search Google Secret Manager variables**.

## Documentation

[Documentation can be found on GoDoc][doc].


```
envsubst -no-empty < input.tmpl > output.txt
```

**Example with Google Secret Manager**

file: `input.tmpl`
```
Hello ${USER}
My secret from Google: ${gcp:secretmanager:projects/xxx/secrets/mykey/versions/1}
```

#### Imposing restrictions

The flags and their restrictions are:

|__Option__     | __Meaning__    | __Type__ | __Default__  |
| ------------| -------------- | ------------ | ------------ |
|`-no-empty`  | fail if a variable is set but empty | `flag` | `false`

## Supported Functions

| __Expression__                | __Meaning__                                                     |
| -----------------             | --------------                                                  |
| `${var}`                      | Value of `$var`
| `${#var}`                     | String length of `$var`
| `${var^}`                     | Uppercase first character of `$var`
| `${var^^}`                    | Uppercase all characters in `$var`
| `${var,}`                     | Lowercase first character of `$var`
| `${var,,}`                    | Lowercase all characters in `$var`
| `${var:n}`                    | Offset `$var` `n` characters from start
| `${var:n:len}`                | Offset `$var` `n` characters with max length of `len`
| `${var#pattern}`              | Strip shortest `pattern` match from start
| `${var##pattern}`             | Strip longest `pattern` match from start
| `${var%pattern}`              | Strip shortest `pattern` match from end
| `${var%%pattern}`             | Strip longest `pattern` match from end
| `${var-default`               | If `$var` is not set, evaluate expression as `$default`
| `${var:-default`              | If `$var` is not set or is empty, evaluate expression as `$default`
| `${var=default`               | If `$var` is not set, evaluate expression as `$default`
| `${var:=default`              | If `$var` is not set or is empty, evaluate expression as `$default`
| `${var/pattern/replacement}`  | Replace as few `pattern` matches as possible with `replacement`
| `${var//pattern/replacement}` | Replace as many `pattern` matches as possible with `replacement`
| `${var/#pattern/replacement}` | Replace `pattern` match with `replacement` from `$var` start
| `${var/%pattern/replacement}` | Replace `pattern` match with `replacement` from `$var` end

For a deeper reference, see [bash-hackers](https://wiki.bash-hackers.org/syntax/pe#case_modification) or [gnu pattern matching](https://www.gnu.org/software/bash/manual/html_node/Pattern-Matching.html).

## Unsupported Functions

* `${var-default}`
* `${var+default}`
* `${var:?default}`
* `${var:+default}`

[doc]: http://godoc.org/github.com/hguerra/envsubst
