package somegraph

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

import (
	"fmt"
	"strings"
)

// SomeGraphFromJSONable parses SomeGraph from a JSONable value.
//
// If there are any errors, the state of target is undefined.
//
// SomeGraphFromJSONable requires:
//  * target != nil
//  * errors != nil
//  * errors.Empty()
func SomeGraphFromJSONable(
	value interface{},
	ref string,
	target *SomeGraph,
	errors *Errors) {

	if target == nil {
		panic("unexpected nil target")
	}

	if errors == nil {
		panic("unexpected nil errors")
	}

	if !errors.Empty() {
		panic("unexpected non-empty errors")
	}

	cast, ok := value.(map[string]interface{})
	if !ok {
		errors.Add(
			ref,
			fmt.Sprintf(
				"expected a map[string]interface{}, but got: %T",
				value))
		return
	}

	////
	// Parse SomeBool
	////

	value0, ok0 := cast[
		"some_bool"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: some_bool")
	} else {
		cast1, ok1 := value0.(bool)
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "some_bool"},
					"/"),
				fmt.Sprintf(
					"expected a bool, but got: %T",
					value0))
		} else {
			target.SomeBool = cast1
		}
	}

	if errors.Full() {
		return
	}

	return
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!