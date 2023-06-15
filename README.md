# Tmplify

tmplify is a powerful string templating package for Go, inspired by Jinja2 in Python. It provides a convenient way to template strings by replacing tokens with values from a map.

## Features

- Simple and intuitive API for string templating.
- Supports dynamic token replacement using a map of key-value pairs.
- Flexible syntax for defining and using tokens in template strings.
- Efficient and optimized for performance.

## Installation

To use tmplify in your Go project, you need to have Go installed and set up. Then, you can install the package by running the following command:

```shell
go get github.com/Soul-Remix/tmplify
```

## Usage

```go
import "github.com/Soul-Remix/tmplify"
```

## Template a string

```go
template := "Hello, {{ .name }}!"
data := map[string]interface{}{
    "name": "John",
}

result, err := tmplify.TemplateString(template, data)
if err != nil {
    // Handle error
}

fmt.Println(result) // Output: Hello, John!
```

## Contributing

Contributions to GoTmplify are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request on GitHub. Make sure to follow the project's code of conduct.

## License

tmplify is licensed under the [MIT License](https://github.com/Soul-Remix/tmplify/blob/main/LICENSE).
