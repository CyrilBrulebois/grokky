# grokky

[![GoDoc](https://godoc.org/github.com/logrusorgru/grokky?status.svg)](https://godoc.org/github.com/logrusorgru/grokky)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/grokky.svg)](https://travis-ci.org/logrusorgru/grokky)
[![Coverage Status](https://coveralls.io/repos/logrusorgru/grokky/badge.svg?branch=master)](https://coveralls.io/r/logrusorgru/grokky?branch=master)
[![GoReportCard](http://goreportcard.com/badge/logrusorgru/grokky)](http://goreportcard.com/report/logrusorgru/grokky)

Package grokky is a pure Golang Grok-like patterns library. This can
help you to parse log files and other. This is based on
[RE2](https://en.wikipedia.org/wiki/RE2_%28software%29)
regexp that
[much more faster](https://swtch.com/~rsc/regexp/regexp1.html)
then
[Oniguruma](https://en.wikipedia.org/wiki/Oniguruma).
The library disigned for creating
many patterns and using it many times. The behavior and capabilities
are slightly different from the original library. The golas of the
library are: (1) simplicity, (2) performance, (3) ease of use.

# Also

See also another golang implementation: [vjeantet/grok](https://github.com/vjeantet/grok). This implementation is closer to the original library.

The difference:

1. The grokky allows named captures only. Any name of a pattern is just
  name of a pattern and nothing more. You can treat is as an alias
  for regexp. It's impossible to use a name of a pattern as a capture group.
  In some cases the grooky is similar to the grok that created as
  `g, err := grok.NewWithConfig(&grok.Config{NamedCapturesOnly:   true})`
  But.

2. The grokky prefered top named group. Unfortunately it is very
  difficult to explain what it means. If you have two patterns. The second
  pattern has same named group and nested into first. Then the named group of
  the first pattern will be used. The grok uses last (closer to tail) group
  in any cases. But the grok also has `ParseToMultiMap` method. To see the
  difference explanation get the package and run following command
  `go test -v -run the_difference github.com/logrusorgru/grokky`

3. The grokky designed as a factory of patterns.

# Get it

```
go get -u -t github.com/logrusorgru/grokky
```

Run test case

```
go test github.com/logrusorgru/grokky
```

Run benchmark comparsion with vjeantet/grok

```
go test -bench=.* github.com/logrusorgru/grokky
```


# Example


```go

package main

import (
	"github.com/logrusorgru/grokky"
	"fmt"
	"log"
	"time"
)

func createHost() grokky.Host {
	h := grokky.New()
	// add patterns to the Host
	h.Must("YEAR", `(?:\d\d){1,2}`)
	h.Must("MONTHNUM2", `0[1-9]|1[0-2]`)
	h.Must("MONTHDAY", `(?:0[1-9])|(?:[12][0-9])|(?:3[01])|[1-9]`)
	h.Must("HOUR", `2[0123]|[01]?[0-9]`)
	h.Must("MINUTE", `[0-5][0-9]`)
	h.Must("SECOND", `(?:[0-5]?[0-9]|60)(?:[:.,][0-9]+)?`)
	h.Must("TIMEZONE", `Z%{HOUR}:%{MINUTE}`)
	h.Must("DATE", "%{YEAR:year}-%{MONTHNUM2:month}-%{MONTHDAY:day}")
	h.Must("TIME", "%{HOUR:hour}:%{MINUTE:min}:%{SECOND:sec}")
	return h
}

func main() {
	h := createHost()
	// compile the pattern for RFC3339 time
	p, err := h.Compile("%{DATE:date}T%{TIME:time}%{TIMEZONE:tz}")
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range p.Parse(time.Now().Format(time.RFC3339)) {
		fmt.Printf("%s: %v\n", k, v)
	}
	//
	// Yes, it's better to use time.Parse for time values
	// but this is just example.
	//
}

```

# Licensing

Copyright © 2015 Konstantin Ivanov <kostyarin.ivanov@gmail.com>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE.md file for more details.