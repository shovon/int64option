package int64option

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestReturnSomething(t *testing.T) {
	var expected int64 = 42
	s := Something(expected)
	v, err := s.Return()
	if err != nil {
		t.Fatal(err)
	}
	if v != expected {
		fmt.Printf("Expected %d but got %d\n", 42, v)
		t.Fail()
	}
}

func TestReturnNothing(t *testing.T) {
	n := Nothing()
	_, err := n.Return()
	if err == nil {
		fmt.Println("Expected to get an error but got nil")
		t.Fail()
	}
}

func TestGoStringSomething(t *testing.T) {
	var expected int64 = 42
	expectedStr := fmt.Sprintf("int64option.Something(%d)", expected)
	s := Something(expected)
	str := fmt.Sprintf("%#v", s)
	if str != expectedStr {
		fmt.Printf("Expected %q but got %q\n", expectedStr, str)
		t.Fail()
	}
}

func TestGoStringNothing(t *testing.T) {
	n := Nothing()
	expectedStr := "int64option.Nothing()"
	str := fmt.Sprintf("%#v", n)
	if str != expectedStr {
		fmt.Printf("Expected %q but got %q\n", expectedStr, str)
		t.Fail()
	}
}

func TestStringSomething(t *testing.T) {
	var expected int64 = 42
	expectedStr := fmt.Sprintf("int64option.Something(%d)", expected)
	s := Something(expected)
	str := strings.Trim(fmt.Sprintf(" %s", s), " ")
	if str != expectedStr {
		fmt.Printf("Expected %q but got %q\n", expectedStr, str)
		t.Fail()
	}
}

func TestStringNothing(t *testing.T) {
	expectedStr := "⧼nothing⧽"
	n := Nothing()
	str := strings.Trim(fmt.Sprintf(" %s", n), " ")
	if str != expectedStr {
		fmt.Printf("Expected %q but got %q\n", expectedStr, str)
		t.Fail()
	}
}

func TestMarshalSomething(t *testing.T) {
	expectedStr := `"something(42)"`
	s := Something(42)
	b, err := s.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	str := string(b)
	if str != expectedStr {
		fmt.Printf("Expected %q but got %q\n", expectedStr, str)
	}
}

func TestMarshalNothing(t *testing.T) {
	expectedStr := `"something(42)"`
	n := Nothing()
	b, err := n.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	str := string(b)
	if str != expectedStr {
		fmt.Printf("Expected %q but got %q\n", expectedStr, str)
	}
}

func TestUnmarshalSomething(t *testing.T) {
	expectedNum := 42
	j := []byte(fmt.Sprintf(`"something(%d)"`, expectedNum))
	var s Type
	err := json.Unmarshal(j, &s)
	if err != nil {
		t.Fatal(err)
	}
	v, err := s.Return()
	if err != nil {
		t.Fatal(err)
	}
	if v != int64(expectedNum) {
		fmt.Printf("Expected %d, but got %d\n", expectedNum, v)
		t.Fail()
	}
}

func TestUnmarshalNothing(t *testing.T) {
	j := []byte(`"nothing()"`)
	var s Type
	err := json.Unmarshal(j, &s)
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.Return()
	if err == nil {
		fmt.Println("Expected to get an error, but got nil")
		t.Fail()
	}
}

func TestValueSomething(t *testing.T) {
	var expected = "something(64)"
	v := Something(64)
	actual, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	actualParsed, ok := actual.(string)
	if !ok {
		fmt.Println("Expected a type string, but got something else")
		t.Fail()
	}
	if actualParsed != expected {
		fmt.Printf("Expected %q but got %q", expected, actualParsed)
		t.Fail()
	}
}

func TestValueNothing(t *testing.T) {
	var expected = "nothing()"
	v := Nothing()
	actual, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	actualParsed, ok := actual.(string)
	if !ok {
		fmt.Println("Expected a type string, but got something else")
		t.Fail()
	}
	if actualParsed != expected {
		fmt.Printf("Expected %q but got %q", expected, actualParsed)
		t.Fail()
	}
}

func TestScanStringSomething(t *testing.T) {
	var expected = int64(42)
	actual := Type{}
	actual.Scan("something(42)")
	r, err := actual.Return()
	if err != nil {
		t.Fatal(err)
	}
	if r != expected {
		fmt.Printf("Expected %d but got %d\n", expected, r)
	}
}

func TestScanByteSomething(t *testing.T) {
	var expected = int64(42)
	actual := Type{}
	err := actual.Scan([]byte("something(42)"))
	if err != nil {
		t.Fatal(err)
	}
	r, err := actual.Return()
	if err != nil {
		t.Fatal(err)
	}
	if r != expected {
		fmt.Printf("Expected %d but got %d\n", expected, r)
	}
}

func TestScanStringNothing(t *testing.T) {
	actual := Type{}
	actual.Scan("nothing()")
	_, err := actual.Return()
	if err == nil {
		fmt.Printf("Expected to fail\n")
		t.Fail()
	}
}

func TestScanByteNothing(t *testing.T) {
	actual := Type{}
	actual.Scan([]byte("nothing()"))
	_, err := actual.Return()
	if err == nil {
		fmt.Printf("Expected to fail\n")
		t.Fail()
	}
}
