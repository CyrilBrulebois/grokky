//
// Copyright (c) 2016 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE.md file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package grokky

import (
	"fmt"
	"log"
)

func ExamplePattern_Parse() {
	h := New()
	h.Add("WORD", `\w+`)
	h.Add("NUMBER", `\d+`)
	p, err := h.Compile("%{WORD:name}/%{NUMBER:age}")
	if err != nil {
		log.Fatal(err)
	}
	result := p.Parse("Alice/15")
	fmt.Println("Name:", result["name"])
	fmt.Println("Age:", result["age"])
	// Output:
	// Name: Alice
	// Age: 15
}
