package somegraph

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

import (
	"fmt"
	"strconv"
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
	// Parse ArrayOfMaps
	////

	value0, ok0 := cast[
		"array_of_maps"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: array_of_maps")
	} else {
		cast1, ok1 := value0.([]interface{})
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "array_of_maps"},
					"/"),
				fmt.Sprintf(
					"expected a []interface{}, but got: %T",
					value0))
		} else {
			target1 := make(
				[]map[string]bool,
				len(cast1))
			for i1 := range cast1 {
				cast2, ok2 := (cast1[i1]).(map[string]interface{})
				if !ok2 {
					errors.Add(
						strings.Join(
							[]string{
								ref, "array_of_maps", strconv.Itoa(i1)},
							"/"),
						fmt.Sprintf(
							"expected a map[string]interface{}, but got: %T",
							cast1[i1]))
				} else {
					target2 := make(map[string]bool)
					for k2 := range cast2 {
						cast3, ok3 := (cast2[k2]).(bool)
						if !ok3 {
							errors.Add(
								strings.Join(
									[]string{
										ref, "array_of_maps", strconv.Itoa(i1), k2},
									"/"),
								fmt.Sprintf(
									"expected a bool, but got: %T",
									cast2[k2]))
						} else {
							target2[k2] = cast3
						}

						if errors.Full() {
							break;
						}
					}

					target1[i1] = target2
				}

				if errors.Full() {
					break;
				}
			}

			target.ArrayOfMaps = target1
		}
	}

	if errors.Full() {
		return
	}

	return
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!