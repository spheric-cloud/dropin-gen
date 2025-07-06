// SPDX-FileCopyrightText: 2024 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

// Package dropin demonstrates the usage of the dropin-gen tool.
package dropin

import (
	"dropin-gen/testdata/orig"
	"fmt"
)

type MyAlias[V any] = orig.MyAlias[V]

func MyFunc[S fmt.Stringer](s S) string {
	return orig.MyFunc(s)
}

var MyVar int = orig.MyVar

// Predicate is a function that matches a value of any type.
type Predicate[V any] = orig.Predicate[V]

// StringPredicate matches a given string.
//
// Matching should be case-insensitive.
type StringPredicate = orig.StringPredicate
