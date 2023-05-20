This project just contains tests, code snippets, etc. all written as Go tests.

# Table-Driven Tests

In the future I will set up more map-based, table-driven tests. For example, see TestMarshal and TestUnmarshal in StandardLibrary/JSON_YAML_marshal_etc/json_test.go.

# IDE Support For Running Tests

If you have an IDE like GoLand it is easy to run a single test from a table. In GoLand you to click on the green arrow in the gutter next to a row of the table of tests to run a single test.  Moreover, table-driven tests make it simple to add a new test, or a variation of an existing test, without adding or modifying any code and immediately try it.

Note that they have been set up carefully. E.g., you must use a struct for the data of each test and .  See https://www.jetbrains.com/go/guide/tips/run-test-in-table-test/?playlist=new-features-in-goland-2020-3).

See my presentation at  [Using Tests for Snippets](https://docs.google.com/presentation/d/1xY6NBQWX--jTtCL125OoKtid8BhUzBkEj1O8wV6FjkA/edit#slide=id.p) for more explanation.

# Organisation

Some tests are just "snippets" that I put here just in case I need them.  They can then be used later to assist memory, or even as a source of code to be copied and pasted elsewhere.

Some tests started as experiments to assist understanding of Go behaviour.  Note that I don't advocate learning how to use a language by experiment (though this is a lot safer in Go than C/C++); however, sometimes an experiment is useful to confirm something and solidify understanding.

Because it's all tests, most or all of the files have names of the form *_test.go.  Since it only has test functions (the func name beginning with "Test") the package name does not mean anything, so I generally use __ (two underscores) for the name.  (I tried using a single underscore but that had problems.)  These "tests" are not nec. used to verify the correctness of any code -- in other words, they typically need *not* (but may) call testing.T.Fail() etc.

