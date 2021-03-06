package address

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

import (
	"fmt"
	"strings"
	"time"
)

// PersonFromJSONable parses Person from a JSONable value.
//
// If there are any errors, the state of the target is undefined.
//
// PersonFromJSONable requires:
//  * target != nil
//  * errors != nil
//  * errors.Empty()
func PersonFromJSONable(
	value interface{},
	id string,
	ref string,
	target *Person,
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

	target.ID = id

	////
	// Parse FullName
	////

	value0, ok0 := cast[
		"full_name"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: full_name")
	} else {
		cast1, ok1 := value0.(string)
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "full_name"},
					"/"),
				fmt.Sprintf(
					"expected a string, but got: %T",
					value0))
		} else {
			target.FullName = cast1
		}
	}

	if errors.Full() {
		return
	}

	////
	// Parse Birthday
	////

	value2, ok2 := cast[
		"birthday"]

	if !ok2 {
		errors.Add(
			ref,
			"property is missing: birthday")
	} else {
		cast3, ok3 := value2.(string)
		if !ok3 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "birthday"},
					"/"),
				fmt.Sprintf(
					"expected a string, but got: %T",
					value2))
		} else {
			target3, err3 := time.Parse(
				"2006-01-02",
				cast3)
			if err3 != nil {
				errors.Add(
					strings.Join(
						[]string{
							ref, "birthday"},
						"/"),
					fmt.Sprintf(
						"expected layout 2006-01-02, got: %s",
						cast3))
			} else {
				target.Birthday = target3
			}
		}
	}

	if errors.Full() {
		return
	}

	////
	// Parse Address
	////

	value4, ok4 := cast[
		"address"]

	if !ok4 {
		errors.Add(
			ref,
			"property is missing: address")
	} else {
		AddressFromJSONable(
			value4,
			strings.Join(
				[]string{
					ref, "address"},
				"/"),
			&(target.Address),
			errors)
	}

	if errors.Full() {
		return
	}

	return
}

// AddressFromJSONable parses Address from a JSONable value.
//
// If there are any errors, the state of the target is undefined.
//
// AddressFromJSONable requires:
//  * target != nil
//  * errors != nil
//  * errors.Empty()
func AddressFromJSONable(
	value interface{},
	ref string,
	target *Address,
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
	// Parse Text
	////

	value0, ok0 := cast[
		"text"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: text")
	} else {
		cast1, ok1 := value0.(string)
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "text"},
					"/"),
				fmt.Sprintf(
					"expected a string, but got: %T",
					value0))
		} else {
			target.Text = cast1
		}
	}

	if errors.Full() {
		return
	}

	return
}

// PipelineFromJSONable parses Pipeline from a JSONable value.
//
// If there are any errors, the state of target is undefined.
//
// PipelineFromJSONable requires:
//  * target != nil
//  * errors != nil
//  * errors.Empty()
func PipelineFromJSONable(
	value interface{},
	ref string,
	target *Pipeline,
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
	// Pre-allocate Persons
	////

	personsRef := ref+"/persons";
	var personsOk bool
	var personsValue interface{}
	var personsMap map[string]interface{}

	personsValue, personsOk = cast[
		"persons"]
	if personsOk {
		personsMap, ok = personsValue.(map[string]interface{})
		if !ok {
			errors.Add(
				personsRef,
				fmt.Sprintf(
					"expected a map[string]interface{}, but got: %T",
					personsValue));
		} else {
			target.Persons = make(
				map[string]*Person)

			for id := range personsMap {
				target.Persons[id] = &Person{}
			}
		}
	}

	// Pre-allocating class instances is critical.
	// If the pre-allocation failed, we can not continue to parse the instances.
	if !errors.Empty() {
		return
	}

	////
	// Parse Persons
	////

	if personsOk {
		for id, value := range personsMap {
			PersonFromJSONable(
				value,
				id,
				strings.Join([]string{
					personsRef, id}, "/"),
				target.Persons[id],
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
	// Parse Maintainer
	////

	value0, ok0 := cast[
		"maintainer"]

	if !ok0 {
		errors.Add(
			ref,
			"property is missing: maintainer")
	} else {
		cast1, ok1 := value0.(string)
		if !ok1 {
			errors.Add(
				strings.Join(
					[]string{
						ref, "maintainer"},
					"/"),
				fmt.Sprintf(
					"expected a string, but got: %T",
					value0))
		} else {
			target1, ok1 := target.Persons[cast1]
			if !ok1 {
				errors.Add(
					strings.Join(
						[]string{
							ref, "maintainer"},
						"/"),
					fmt.Sprintf(
						"reference to an instance of class Person not found: %s",
						value0))
			} else {
				target.Maintainer = target1
			}
		}
	}

	if errors.Full() {
		return
	}

	return
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
