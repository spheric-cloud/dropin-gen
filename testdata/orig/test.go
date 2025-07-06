// SPDX-FileCopyrightText: 2024 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

// Package orig demonstrates the usage of the dropin-gen tool.
package orig

import "fmt"

// Predicate is a function that matches a value of any type.
type Predicate[V any] func(V) bool

type MyAlias[V any] = Predicate[V]

// StringPredicate matches a given string.
//
// Matching should be case-insensitive.
type StringPredicate = Predicate[string]

func MyFunc[S fmt.Stringer](s S) string {
	return "foo" + s.String()
}

var MyVar = 1
