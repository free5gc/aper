package aper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeSingleBitString(t *testing.T) {

	for i, test := range singleBitStringTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		perTestTrace(2, fmt.Sprintf("	in.Bytes : %0x, in.BitsLength: %d", reflect.ValueOf(test.Out).Field(0).Field(0).Bytes(), reflect.ValueOf(test.Out).Field(0).Field(1).Uint()))
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeStructBitString(t *testing.T) {

	for i, test := range structBitStringTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		for j := 0; j < reflect.TypeOf(test.Out).NumField(); j++ {
			perTestTrace(2, fmt.Sprintf("	in.Bytes : %0x, in.BitsLength: %d", reflect.ValueOf(test.Out).Field(j).Field(0).Bytes(), reflect.ValueOf(test.Out).Field(j).Field(1).Uint()))
		}
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeSingleOctetString(t *testing.T) {

	for i, test := range singleOctetStringTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		perTestTrace(2, fmt.Sprintf("	in : %0x, len: %d", reflect.ValueOf(test.Out).Field(0).Bytes(), len(reflect.ValueOf(test.Out).Field(0).Bytes())))
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeSingleInt(t *testing.T) {

	for i, test := range integerTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		perTestTrace(2, fmt.Sprintf("	in : %d", reflect.ValueOf(test.Out).Field(0).Int()))
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeStructInt(t *testing.T) {

	for i, test := range integerStructTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		for j := 0; j < reflect.TypeOf(test.Out).NumField(); j++ {
			perTestTrace(2, fmt.Sprintf("	in : %d", reflect.ValueOf(test.Out).Field(j).Int()))
		}
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeEnum(t *testing.T) {

	for i, test := range enumTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		perTestTrace(2, fmt.Sprintf("	in : %d", reflect.ValueOf(test.Out).Field(0).Uint()))
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodePtr(t *testing.T) {

	for i, test := range ptrTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		perTestTrace(2, fmt.Sprintf("	in : %v", reflect.ValueOf(test.Out).Field(0).Elem()))
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeSequenceOf(t *testing.T) {

	for i, test := range seqofTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		for j := 0; j < reflect.ValueOf(test.Out).Field(0).Len(); j++ {
			perTestTrace(2, fmt.Sprintf("	in : %v", reflect.ValueOf(test.Out).Field(0).Index(j)))
		}
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeStructPrintableString(t *testing.T) {

	for i, test := range printableStringStructTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		for j := 0; j < reflect.ValueOf(test.Out).NumField(); j++ {
			perTestTrace(2, fmt.Sprintf("	in : %s", reflect.ValueOf(test.Out).Field(j).String()))
		}
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeChoice(t *testing.T) {

	for i, test := range choiceTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		perTestTrace(2, fmt.Sprintf("	in : %v", reflect.ValueOf(test.Out).Field(0)))
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}

func TestEncodeOpenType(t *testing.T) {

	for i, test := range openTypeTestData {
		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		perTestTrace(2, fmt.Sprintf("	in : %v", reflect.ValueOf(test.Out)))
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}
func TestEncodeBoolean(t *testing.T) {
	for i, test := range boolTestData {

		perTestTrace(1, fmt.Sprintf("[TEST %d]\n", i+1))
		out, err := Marshal(test.Out)
		perTestTrace(2, fmt.Sprintf("	in : %t", reflect.ValueOf(test.Out).Field(0).Bool()))
		perTestTrace(2, fmt.Sprintf("	out : 0x%0x", out))
		perTestTrace(2, fmt.Sprintf("	exp : 0x%0x", test.in))
		if err != nil {
			perTestTrace(1, fmt.Sprintln("  [ERROR]: "+err.Error()))
		} else if reflect.DeepEqual(test.in, out) {
			perTestTrace(1, fmt.Sprintf("[PASS]\n"))
			continue
		}
		perTestTrace(1, fmt.Sprintf("[FAIL]\n"))
		t.Errorf("TEST %d is FAILED", i+1)

	}
}
