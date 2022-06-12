# learn-go-with-tests

Notes and code examples created whilst following the [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) 

## VS Code Refactoring Quick Fix commands

| Action                      | Key Binding          |
| --- | --- |
| Run code                    | Ctrl + F5            |
| Debug code                  | F5                   |
| Debug previous              | Alt + Shift + D    * |
| Debug - Continue/Pause      | F5                   |
| Debug - Step over           | F10                  |
| Debug - Step into           | F11                  |
| Debug - Step out            | Shift + F11          |
| Debug - Restart             | Ctrl + Shift + F5    |
| Debug - Stop                | Shift + F5           |
| Run test at cursor          | Alt + Shift + T    * |
| Run previous test           | Alt + Shift + P    * |
| Run benchmark at cursor     | Alt + Shift + B    * |
| Run subtest at cursor       | Alt + Shift + S    * |
| Debug test at cursor        | F5                   |
| Run tests in file           | Alt + Shift + Y    * |
| Run benchmarks in file      | Alt + Shift + N    * |
| Run tests in package        | Alt + Shift + U    * |
| Run benchmarks in package   | Alt + Shift + M    * |
| Run tests in workspace      | Alt + Shift + W    * |
| Cancel running tests        | Alt + Shift + C    * |
| Lint package                | Alt + Shift + K    * |
| Lint workspace              | Alt + Shift + L    * |
| Vet package                 | Alt + Shift + V    * |
| Vet workspace               | Alt + Shift + G    * |
| Toggle test file            | Alt + Shift + X    * |
| Code complete               | Ctrl + Space         |
| Signature Help              | Ctrl + Shift + Space |
| Add import                  | Alt + Shift + A    * |
| Rename                      | F2                   |
| Extract/Inline variable     | Ctrl + Shift + R     |
| Extract method/function     | Ctrl + Shift + R     |
| Extract to function         | Alt + Shift + E    * |
| Extract to variable         | Alt + Shift + R    * |
| go fmt                      | Runs on Save         |
| View function signature     | hover over symbol    |
| Go to function definition   | F12                  |
| View function definition    | Ctrl + Shift + F10   |
| Find usages of a symbol     | Alt + Shift + F12    |
| Show call hierachy          | Shift + Alt + H      |
| Go to symbol in file        | Ctrl + Shift + O     |
| Go to symbol in workspace   | Ctrl + T             |

\* - cutstom key binding

### Notes

Settings saved and sync via: [Settings Sync](https://code.visualstudio.com/docs/editor/settings-sync)

#### Hello, World

- `t.Errorf` - For tests `%q` is very useful as it wraps your values in double
quotes.
- for helper test functions, it's a good idea to accept a `testing.TB` which is
an interface that `*testing.T` and `*testing.B` both satisfy, so you can call 
helper functions from a test, or a benchmark
- `t.Helper()` is needed to tell the test suite that this method is a helper. 
By doing this when it fails the line number reported will be in our function 
call rather than inside our test helper.

##### code

- [hello.go](/hello/hello.go)
- [hello_test.go](/hello/hello_test.go)

#### Integers

- [Examples](https://go.dev/blog/examples) not only document your code but
are checked for correctness so will be maintained as the code changes
- run tests including examples using `go test -v`
- require to include commented output value in order to be executable

##### code

- [adder.go](/integers/adder.go)
- [adder_test.go](/integers/adder_test.go)

#### Iteration

- [repeat.go](/arrays-and-slices/repeat.go)
- [repeat_test.go](/arrays-and-slices/repeat_test.go)

#### Arrays and slices

- `range` lets you iterate over an array or slice. Each iteration returns two
values - the index and the value
- `make` allows the creation of a slice with a starting capacity
- `append` will grow a new slice from the one supplied to have a capacity to 
include the new values supplied

##### code

- [sum.go](/arrays-and-slices/sum.go)
- [sum_test.go](/arrays-and-slices/sum_test.go)

#### Structs, methods & interfaces

- [shapes.go](/structs-methods-interfaces/shapes.go)
- [shapes_test.go](/structs-methods-interfaces/shapes_test.go)

#### Pointers & errors

- if you're writing a function that needs to mutate state you'll need it to
take a pointer to the thing you want to change
- pointers can be `nil`
- errors are values, so we can reference them as a variable
- it is possible to create new types from existing ones

##### code

- [wallet.go](/pointers-errors/wallet.go)
- [wallet_test.go](/pointers-errors/wallet_test.go)

#### Maps

- [constant error](https://dave.cheney.net/2016/04/07/constant-errors) pattern 
looks worth adopting
- a `map` value is a pointer to a `runtime.hmap` structure
    - to avoid panics never initialise an empty map **variable**
    - instead initialise an empty map

##### code

- [dictionary.go](/maps/dictionary.go)
- [dictionary_test.go](/maps/dictionary_test.go)

#### Dependency Injection

- you don't need a framework for dependency injection
- it encourages writing of general-purpose functions

##### code

- [greet.go](/dependency-injection/greet.go)
- [greet_test.go](/dependency-injection/greet_test.go)

#### Mocking

- use an *iterative, test-driven approach*
- slow test ruin developer productivity

##### code

- [countdown.go](/mocking/countdown.go)
- [countdown_test.go](/mocking/countdown_test.go)

#### Concurrency

- benchmark tests can give insight into performance: `go test -bench=.`
- *goroutines* runs as a separate process - created by putting `go` in front 
of a function
- *anonymous functions* are frequently combined with goroutines as they:
    - can be executed at the same time they're declared
    - maintain access to lexical scope
- *channels* organise and control communication between different processes
- *race detector* can help spot race conditions: `go test -race`

> [Make it work, make it right, make it fast](http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast)

##### code

- [checkwebsites.go](/concurrency/checkwebsites.go)
- [checkwebsites_test.go](/concurrency/checkwebsites_test.go)

#### Select

- `select` helps you wait on multiple channels
- `time.After` can be handy as one of your `case` statements to timeout long
running code
- `httptest` is a convenient way of creating test servers

##### code

- [racer.go](/select/racer.go)
- [racer_test.go](/select/racer_test.go)

#### Reflection

- only use reflection if you really need to

##### code

- [reflection.go](/reflection/reflection.go)
- [reflection_test.go](/reflection/reflection_test.go)

#### Sync

- `Mutex` allows us to add locks to our data
- `WaitGroup` is a means of waiting for goroutines to complete
- When to use locks over channels
    - Use channels when passing ownership of data
    - Use mutexes for managing state
- Don't use embedding because it is convenient
    - Think about the effect embedding has on your public API

##### code

- [sync.go](/sync/sync.go)
- [sync_test.go](/sync/sync_test.go)

#### Context

- `context` helps to manage long running processes
- you should derive your context so that cancellations are propogated throughout
the call stack
- idiomatic go encourage the passing of a context as the first parameter of
all methods in involved on the call path between incoming and outgoing
requests
- you can use `select`, channels and goroutines when managing the 
cancellation of contexts
- using `context.Value` is a last resort, not a first 

##### code

- [context.go](/context/context.go)
- [context_test.go](/context/context_test.go)

#### Intro to property based tests

- built into standard library
- `quick.Check` runs a number of random inputs checkinf for failures

##### code

- [numeral.go](/property-based/numeral.go)
- [numeral_test.go](/property-based/numeral_test.go)

#### Maths

- developing an acceptance test alongside unit test can tell us when we are 
done
- `encoding/xml` Go package helps handle simple xml parsing
- refactoring to a public API was an interesting exercise
- namespacing go.mod made much more sense when doing this exercise

##### code

- [go.mod](/maths/go.mod)
- [main.go](/maths/clockface/main.go)
- [clockface.go](/maths/clockface.go)
- [clockface.go](/maths/clockface_test.go)
- [svg.go](/maths/svg/svg.go)
- [svg_test.go](/maths/svg/svg_test.go)

#### Reading files

- `io/fs` package introduced an abstraction of the filesystem
- this can help decouple or code from the particular filesystem details
- `testing/fstest` offers an implementation of `io/fs` for writing tests 

##### code

- [post.go](/reading-files/post.go)
- [blogposts.go](/reading-files/blogposts.go)
- [blogposts_test.go](/reading-files/blogposts_test.go)

#### Templating

- `text/template` and `html/template` are useful templating packages
- `embed` provides access to files embedded in the Go program
- Approval tests allow for easy testing of larger objects, strings, etc
    - [github.com/approvals/go-approval-tests](github.com/approvals/go-approval-tests)
- #TODO investigate rendering of Body of posts

##### code

- [renderer.go](/blogrenderer/renderer.go)
- [renderer_test.go](/blogrenderer/renderer_test.go)

#### Generics

- `interface{}` doesn't allow the compiler to help us write code
- in Go when using generics you need to provide *type parameters* e.g. `[T comparable]`
- generics can remove duplication of code and tests

##### code

- [stack.go](/generics/stack.go)
- [stack_test.go](/generics/stack_test.go)
- [generics_test.go](/generics/generics_test.go)

#### Revisited - Arrays and slices

- `Reduce` and `Find` higher-order functions illustrate how generics can result 
in simpler to read and maintain code
- be open-minded about what is and isn't idiomatic

##### code

- [sum.go](/arrays-and-slices-g/sum.go)
- [sum_test.go](/arrays-and-slices-g/sum_test.go)

#### Building an application

##### HTTP server

- `http.Handler` use this interface to create web servers
- `http.HandlerFunc` turns ordinary fuctions into `http.Handler`s
- build system in smaller chuncks

###### code

- [main.go](/1-http-server/main.go)
- [server.go](/1-http-server/server.go)
- [in_memory_player_store.go](/1-http-server/in_memory_player_store.go)
- [server_test.go](/1-http-server/server_test.go)
- [server_integration_test.go](/1-http-server/server_integration_test.go)

##### JSON, routing and embedding

- The standard library provides an easy to use to do **routing**
- **Embedding** can be helpful but need to remember the impact on the public api
- **JSON deserializing and serialzing** can be handled again using the standard library

###### code

- [main.go](/2-json/main.go)
- [server.go](/2-json/server.go)
- [in_memory_player_store.go](/2-json/in_memory_player_store.go)
- [server_test.go](/2-json/server_test.go)
- [server_integration_test.go](/2-json/server_integration_test.go)

##### IO and sorting

- covers working with files
- `sort.Slice` for sorting slices

###### code

- [main.go](/3-io-and-sorting/main.go)
- [server.go](/3-io-and-sorting/server.go)
- [league.go](/3-io-and-sorting/league.go)
- [tape.go](/3-io-and-sorting/tape.go)
- [file_system_player_store.go](/3-io-and-sorting/file_system_player_store.go)
- [server_test.go](/3-io-and-sorting/server_test.go)
- [file_system_player_store_test.go](/3-io-and-sorting/file_system_player_store_test.go)
- [tape_test.go](/3-io-and-sorting/tape_test.go)
- [server_integration_test.go](/3-io-and-sorting/server_integration_test.go)

Useful format verbs:

| Verb | Type            | Use                                         |
| --- | --- |--- |
| `%v` | General         | the value in a default format               |
| `%d` | integer         | base 10                                     |
| `%f` | floating point  | decmial point but not exponent e.g. `9.2f%` |
| `%g` | floating point  | for a precise decimal number                |
| `%q` | string          | double-quoted safely escaped with Go syntax |

### TODO

Investigate other VS Code configurations details for VS Code Go

## References

- [Go in VS Code](https://code.visualstudio.com/docs/languages/go)
- [VS Code Go Docs](https://github.com/golang/vscode-go/blob/master/docs/Home.md)