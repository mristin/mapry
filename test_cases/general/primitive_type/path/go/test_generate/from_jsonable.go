package somegraph

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

import (
	"fmt"
	"regexp"
	"strings"
)

var pattern0 = regexp.MustCompile(
	`^/[a-zA-Z]+-[0-9]+$`)

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
	// Parse SomePath
	////

	value0, ok0 := cast[
		"some_path"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: some_path")
	} else {
		cast1, ok1 := value0.(string)
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "some_path"},
					"/"),
				fmt.Sprintf(
					"expected a string, but got: %T",
					value0))
		} else {
			if !pattern0.MatchString(cast1) {
				errors.Add(
					strings.Join(
						[]string{
							ref, "some_path"},
						"/"),
					fmt.Sprintf(
						"expected to match ^/[a-zA-Z]+-[0-9]+$, but got: %s",
						cast1))
			} else {
				target.SomePath = cast1
			}
		}
	}

	if errors.Full() {
		return
	}

	////
	// Parse UnconstrainedPath
	////

	value2, ok2 := cast[
		"unconstrained_path"]

	if !ok2 {
		errors.Add(
			ref,
			"property is missing: unconstrained_path")
	} else {
		cast3, ok3 := value2.(string)
		if !ok3 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "unconstrained_path"},
					"/"),
				fmt.Sprintf(
					"expected a string, but got: %T",
					value2))
		} else {
			target.UnconstrainedPath = cast3
		}
	}

	if errors.Full() {
		return
	}

	return
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
