package somegraph

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

import (
	"fmt"
	"strconv"
	"strings"
)

// EmptyFromJSONable parses Empty from a JSONable value.
//
// If there are any errors, the state of the target is undefined.
//
// EmptyFromJSONable requires:
//  * target != nil
//  * errors != nil
//  * errors.Empty()
func EmptyFromJSONable(
	value interface{},
	ref string,
	target *Empty,
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

	_, ok := value.(map[string]interface{})
	if !ok {
		errors.Add(
			ref,
			fmt.Sprintf(
				"expected a map[string]interface{}, but got: %T",
				value))
		return
	}

	return
}

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
	// Parse ArrayOfEmbeds
	////

	value0, ok0 := cast[
		"array_of_embeds"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: array_of_embeds")
	} else {
		cast1, ok1 := value0.([]interface{})
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "array_of_embeds"},
					"/"),
				fmt.Sprintf(
					"expected a []interface{}, but got: %T",
					value0))
		} else {
			target1 := make(
				[]Empty,
				len(cast1))
			for i1 := range cast1 {
				EmptyFromJSONable(
					cast1[i1],
					strings.Join(
						[]string{
							ref, "array_of_embeds", strconv.Itoa(i1)},
						"/"),
					&(target1[i1]),
					errors)

				if errors.Full() {
					break;
				}
			}

			target.ArrayOfEmbeds = target1
		}
	}

	if errors.Full() {
		return
	}

	return
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
