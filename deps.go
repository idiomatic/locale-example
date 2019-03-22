package main

// Trigger `go get` to fetch things needed by derivative projects.
// This file is optional.

import (
	_ "github.com/blevesearch/bleve"
	_ "github.com/idiomatic/locale"
	_ "github.com/idiomatic/middleware"
	_ "golang.org/x/text/language"
)
