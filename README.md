# learn-go-with-tests

Notes and code examples created whilst following the [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) 

## VS Code Refactoring Quick Fix commands

| Action                   | Quick Fix Keys     |
| --- | --- |
| Run Code                 | Ctrl + F5       |
| Code Complete            | Ctrl + Space       |
| Rename                   | F2                 |
| Extract/Inline variable  | Ctrl + Shift + R   |
| Extract method/function  | Ctrl + Shift + R   |
| go fmt                   | Runs on Save       |
| View function signature  | hover over symbol  |
| View function definition | Ctrl + Shift + F10 |
| Find usages of a symbol  | Alt + Shift + F12  |

### Notes

#### Hello, World

- `t.Errorf` - For tests `%q` is very useful as it wraps your values in double
quotes.
- for helper test functions, it's a good idea to accept a `testing.TB` which is
an interface that `*testing.T` and `*testing.B` both satisfy, so you can call 
helper functions from a test, or a benchmark
- `t.Helper()` is needed to tell the test suite that this method is a helper. 
By doing this when it fails the line number reported will be in our function 
call rather than inside our test helper.

- [hello.go](/hello/hello.go)
- [hello_test.go](/hello/hello_test.go)

#### Integers

- [Examples](https://go.dev/blog/examples) not only document your code but
are checked for correctness so will be maintained as the code changes
- run tests including examples using `go test -v`
- require to include commented output value in order to be executable

- [adder.go](/integers/adder.go)
- [adder_test.go](/integers/adder_test.go)

### TODO

Investigate other VS Code configurations details for vS Code Go

## References

- [VS Code Go Docs](https://github.com/golang/vscode-go/blob/master/docs/features.md)