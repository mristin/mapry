package somegraph

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

// SomeGraph converts the instance to a JSONable representation.
//
// SomeGraph requires:
//  * instance != nil
//
// SomeGraph ensures:
//  * (err == nil && target != nil) || (err != nil && target == nil)
func SomeGraphToJSONable(
	instance *SomeGraph) (
	target map[string]interface{}, err error) {

	if instance == nil {
		panic("unexpected nil instance")
	}

	target = make(map[string]interface{})
	defer func() {
		if err != nil {
			target = nil
		}
	}()
	////
	// Serialize SomeBool
	////

	target["some_bool"] = instance.SomeBool

	return
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
