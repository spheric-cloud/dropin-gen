// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

// Package dropin demonstrates the usage of the dropin-gen tool.
package dropin

import (
	"fmt"

	orig1 "spheric.cloud/dropin-gen/testdata/orig"
)

type MyAlias[V any] = orig1.MyAlias[V]

func MyFunc[S fmt.Stringer](s S) string {
	return orig1.MyFunc(s)
}

var MyVar int = orig1.MyVar

// Predicate is a function that matches a value of any type.
type Predicate[V any] = orig1.Predicate[V]

// StringPredicate matches a given string.
//
// Matching should be case-insensitive.
type StringPredicate = orig1.StringPredicate

// VarargFunc is a function with variadic arguments.
func VarargFunc(orig int, bs ...int) int {
	return orig1.VarargFunc(orig, bs...)
}
