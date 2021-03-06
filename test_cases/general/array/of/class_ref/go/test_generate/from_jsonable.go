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
	id string,
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

	target.ID = id

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
	// Pre-allocate Empties
	////

	emptiesRef := ref+"/empties";
	var emptiesOk bool
	var emptiesValue interface{}
	var emptiesMap map[string]interface{}

	emptiesValue, emptiesOk = cast[
		"empties"]
	if emptiesOk {
		emptiesMap, ok = emptiesValue.(map[string]interface{})
		if !ok {
			errors.Add(
				emptiesRef,
				fmt.Sprintf(
					"expected a map[string]interface{}, but got: %T",
					emptiesValue));
		} else {
			target.Empties = make(
				map[string]*Empty)

			for id := range emptiesMap {
				target.Empties[id] = &Empty{}
			}
		}
	}

	// Pre-allocating class instances is critical.
	// If the pre-allocation failed, we can not continue to parse the instances.
	if !errors.Empty() {
		return
	}

	////
	// Parse Empties
	////

	if emptiesOk {
		for id, value := range emptiesMap {
			EmptyFromJSONable(
				value,
				id,
				strings.Join([]string{
					emptiesRef, id}, "/"),
				target.Empties[id],
				errors)

			if errors.Full() {
				break
			}
		}
	}

	if errors.Full() {
		return
	}

	////
	// Parse ArrayOfClassRefs
	////

	value0, ok0 := cast[
		"array_of_class_refs"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: array_of_class_refs")
	} else {
		cast1, ok1 := value0.([]interface{})
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "array_of_class_refs"},
					"/"),
				fmt.Sprintf(
					"expected a []interface{}, but got: %T",
					value0))
		} else {
			target1 := make(
				[]*Empty,
				len(cast1))
			for i1 := range cast1 {
				cast2, ok2 := (cast1[i1]).(string)
				if !ok2 {
					errors.Add(
						strings.Join(
							[]string{
								ref, "array_of_class_refs", strconv.Itoa(i1)},
							"/"),
						fmt.Sprintf(
							"expected a string, but got: %T",
							cast1[i1]))
				} else {
					target2, ok2 := target.Empties[cast2]
					if !ok2 {
						errors.Add(
							strings.Join(
								[]string{
									ref, "array_of_class_refs", strconv.Itoa(i1)},
								"/"),
							fmt.Sprintf(
								"reference to an instance of class Empty not found: %s",
								cast1[i1]))
					} else {
						target1[i1] = target2
					}
				}

				if errors.Full() {
					break;
				}
			}

			target.ArrayOfClassRefs = target1
		}
	}

	if errors.Full() {
		return
	}

	return
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
