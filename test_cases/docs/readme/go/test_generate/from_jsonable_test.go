package address

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

import "fmt"

func ExampleDurationFromString_invalid() {
	_, err := durationFromString("some wrong text")
	if err == nil {
		panic("Unexpected nil error");
	}
	fmt.Println(err.Error())
	// Output: failed to match the duration pattern on: some wrong text
}

func ExampleDurationFromString_oneYear() {
	d, err := durationFromString("P1Y")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: 8765h49m12s
}

func ExampleDurationFromString_oneMonth() {
	d, err := durationFromString("P1M")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: 730h29m6s
}

func ExampleDurationFromString_oneWeek() {
	d, err := durationFromString("P1W")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: 168h0m0s
}

func ExampleDurationFromString_oneDay() {
	d, err := durationFromString("P1D")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: 24h0m0s
}

func ExampleDurationFromString_hoursMinutesSeconds() {
	d, err := durationFromString("PT1.1H2.2M3.3S")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: 1h8m15.3s
}

func ExampleDurationFromString_preciseNanoseconds() {
	d, err := durationFromString("PT0.000000001S")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: 1ns
}

func ExampleDurationFromString_secondsWithTrailingZeros() {
	d, err := durationFromString("PT1.000S")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: 1s
}

func ExampleDurationFromString_allSpecified() {
	d, err := durationFromString("P1.1Y2.1M3.1W4.1DT5.1H6.1M7.1S")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: 11800h49m26.900000007s
}

func ExampleDurationFromString_negative() {
	d, err := durationFromString("-P1D")
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()));
	}
	fmt.Println(fmt.Sprintf("%s", d))
	// Output: -24h0m0s
}

func ExampleDurationFromString_overflow() {
	_, err := durationFromString("P300Y")
	fmt.Println(err.Error())
	// Output: overflows in nanoseconds: P300Y
}

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
