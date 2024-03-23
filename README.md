# `demp`

`demp` is a `Dollar Templating` library.

## What is `Dollar Templating` ?

Go templates use `{{ }}` as the template delimeter.

But it might not be the best choice for a end-user facing templating language.

Meaning

```toml
someVar = "$pwd"

# is better than

someVar = "{{ .pwd }}"
```

## Usage

`demp` exposes a single function `ResolveTemplate`. 

Which takes 2 arguments.
- `template`
    The template string
- `variables`
    The variables to the substituted.

It returns the resolved template string.

```go
import "github.com/pspiagicw/demp"

func main() {
    input := "$name is $age years old"
    variables := map[string]string{
        "name": "John",
        "age": "25",
    }

    fmt.Println(demp.ResolveTemplate(input, variables)) // => "John is 25 years old"
}
```

## Templating

- `demp` uses `$` as the template delimeter.
- You can use `${someVar}` for convenience.
    Example `Plural: ${noun}s`, should print `Plural: apples` for `noun=apple`.
- Names can only contain `a-z`, `A-Z`, `0-9`, `_` and `-`.
- `$$` is used to escape `$`.
    Example: `$$name` should print `$name`. `$$$name` should print `$John` for `name=John`.

## Contribution

Feel free to open issues and PRs.

## License

MIT
