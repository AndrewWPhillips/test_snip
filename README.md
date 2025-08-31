This project just contains tests, code snippets, etc. all written as Go tests (except files in the main and private directories).

# Organisation

This project is called "test_snip" as it consists of tests (mainly) and code snippets.  But it's all tests so (most) files have names of the form *_test.go.

Since it's all tests package names don't really mean anything, so I generally use __ (two underscores) for the name.  (I tried using a single underscore but that had problems.)

## Tests

Some tests started as experiments to assist understanding of Go behaviour.  Note that I don't advocate learning how to use a language by experiment (though this is a lot safer in Go than C/C++); however, sometimes an experiment is useful to confirm something and solidify understanding.

These "tests" are not nec. used to verify the correctness of code (as Test*() functions are intended to do).  They need *not* (but may) call testing.T.Fail() etc.

## Snippets

Apart from tests there are also some code snippets, that I have added to remind me of something or even copy and paste.  These are mainly in the utils directory - eg see TestGoroutineID() which allows you to get an ID of a goroutine (something that is discouraged but can be useful for debugging).

# Table-Driven Tests

In the future I will set up more map-based, table-driven tests. For example, see TestMarshal and TestUnmarshal in StandardLibrary/JSON_YAML_marshal_etc/json_test.go.

## IDE Support For Running Tests

If you have an IDE that supports it (like GoLand), it is easy to run a single test from a table. In GoLand you to click on the green arrow in the gutter next to a row of the table of tests to run a single test.  Moreover, table-driven tests make it simple to add a new test, or a variation of an existing test, without adding or modifying any code, then immediately try it.

Note that they have been set up carefully. E.g., you must use a struct for the data of each test and .  See https://www.jetbrains.com/go/guide/tips/run-test-in-table-test/?playlist=new-features-in-goland-2020-3).

See my presentation at  [Using Tests for Snippets](https://docs.google.com/presentation/d/1xY6NBQWX--jTtCL125OoKtid8BhUzBkEj1O8wV6FjkA/edit#slide=id.p) for more explanation.

