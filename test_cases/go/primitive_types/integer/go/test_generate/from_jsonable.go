package somegraph

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

import (
	"fmt"
	"math"
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
	// Parse SomeInt
	////

	value0, ok0 := cast[
		"some_int"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: some_int")
	} else {
		fcast1, ok1 := value0.(float64)
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "some_int"},
					"/"),
				fmt.Sprintf(
					"expected a float64, but got: %T",
					value0))
		} else if fcast1 != math.Trunc(fcast1) {
			errors.Add(
				strings.Join(
					[]string{
						ref, "some_int"},
					"/"),
				fmt.Sprintf(
					"expected a whole number, but got: %f",
					fcast1))
		// 9223372036854775808.0 == 2^63 is the first float > MaxInt64.
		// -9223372036854775808.0 == -(2^63) is the last float >= MinInt64.
		} else if fcast1 >= 9223372036854775808.0 ||
			fcast1 < -9223372036854775808.0 {

			errors.Add(
				strings.Join(
					[]string{
						ref, "some_int"},
					"/"),
				fmt.Sprintf(
					"expected the value to fit into int64, but got an overflow: %f",
					fcast1))
		} else {
			target.SomeInt = int64(fcast1)
		}
	}

	if errors.Full() {
		return
	}

	return
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
