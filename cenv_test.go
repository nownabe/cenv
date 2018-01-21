package cenv

import (
	"os"
	"testing"
)

type testBaseI interface {
	isErr() bool
	withEnv(f func())
}

type testBase struct {
	key string
	val string
	in  []string
	err bool
}

func (t testBase) isErr() bool {
	return t.err
}

func (t testBase) withEnv(f func()) {
	os.Setenv(t.key, t.val)
	f()
	os.Unsetenv(t.key)
}

type testFunc func() (bool, error)

type boolTest struct {
	testBase
	out bool
}

var boolTests = []boolTest{
	// True cases
	{testBase{"ENV_KEY", "1", []string{"ENV_KEY"}, false}, true},
	{testBase{"ENV_KEY", "1", []string{"env.key"}, false}, true},
	{testBase{"ENV_KEY", "1", []string{"env", "key"}, false}, true},
	{testBase{"ENV_KEY", "t", []string{"ENV_KEY"}, false}, true},
	{testBase{"ENV_KEY", "T", []string{"ENV_KEY"}, false}, true},
	{testBase{"ENV_KEY", "TRUE", []string{"ENV_KEY"}, false}, true},
	{testBase{"ENV_KEY", "true", []string{"ENV_KEY"}, false}, true},
	{testBase{"ENV_KEY", "True", []string{"ENV_KEY"}, false}, true},

	// False cases
	{testBase{"ENV_KEY", "0", []string{"ENV_KEY"}, false}, false},
	{testBase{"ENV_KEY", "f", []string{"ENV_KEY"}, false}, false},
	{testBase{"ENV_KEY", "F", []string{"ENV_KEY"}, false}, false},
	{testBase{"ENV_KEY", "FALSE", []string{"ENV_KEY"}, false}, false},
	{testBase{"ENV_KEY", "false", []string{"ENV_KEY"}, false}, false},
	{testBase{"ENV_KEY", "False", []string{"ENV_KEY"}, false}, false},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, false},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, false},
	{testBase{"ENV_KEY", "-1", []string{"ENV_KEY"}, true}, false},
	{testBase{"ENV_KEY", "2", []string{"ENV_KEY"}, true}, false},
	{testBase{"ENV_KEY", "TRue", []string{"ENV_KEY"}, true}, false},
	{testBase{"ENV_KEY", "FAlse", []string{"ENV_KEY"}, true}, false},
}

func TestBool(t *testing.T) {
	for _, test := range boolTests {
		f := func() (bool, error) {
			out, err := Bool(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustBool(test.in...), nil
		}

		testF(t, test, "Bool", f, mf)
	}

	f := func() { MustBool("noenv") }
	if !isPanic(f) {
		t.Errorf("MustBool(noenv) should panic.")
	}
}

type float32Test struct {
	testBase
	out float32
}

var float32Tests = []float32Test{
	// Valid cases
	{testBase{"ENV_KEY", "3.14", []string{"ENV_KEY"}, false}, 3.14},
	{testBase{"ENV_KEY", "3.14", []string{"env.key"}, false}, 3.14},
	{testBase{"ENV_KEY", "3.14", []string{"ENV", "KEY"}, false}, 3.14},
	{testBase{"ENV_KEY", "+3.14e10", []string{"ENV_KEY"}, false}, 3.14e10},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "3.4028236e38", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "-3.4028236e38", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, 0},
}

func TestFloat32(t *testing.T) {
	for _, test := range float32Tests {
		f := func() (bool, error) {
			out, err := Float32(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustFloat32(test.in...), nil
		}

		testF(t, test, "Float32", f, mf)
	}

	f := func() { MustFloat32("noenv") }
	if !isPanic(f) {
		t.Errorf("MustFloat32(noenv) should panic.")
	}
}

type float64Test struct {
	testBase
	out float64
}

var float64Tests = []float64Test{
	// Valid cases
	{testBase{"ENV_KEY", "3.14", []string{"ENV_KEY"}, false}, 3.14},
	{testBase{"ENV_KEY", "3.14", []string{"env.key"}, false}, 3.14},
	{testBase{"ENV_KEY", "3.14", []string{"ENV", "KEY"}, false}, 3.14},
	{testBase{"ENV_KEY", "+3.14e10", []string{"ENV_KEY"}, false}, 3.14e10},
	{testBase{"ENV_KEY", "3.4028236e38", []string{"ENV_KEY"}, false}, 3.4028236e38},
	{testBase{"ENV_KEY", "-3.4028236e38", []string{"ENV_KEY"}, false}, -3.4028236e38},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "1.7976931348623159e308", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "-1.7976931348623159e308", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, 0},
}

func TestFloat64(t *testing.T) {
	for _, test := range float64Tests {
		f := func() (bool, error) {
			out, err := Float64(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustFloat64(test.in...), nil
		}

		testF(t, test, "Float64", f, mf)
	}

	f := func() { MustFloat64("noenv") }
	if !isPanic(f) {
		t.Errorf("MustFloat64(noenv) should panic.")
	}
}

type intTest struct {
	testBase
	out int
}

var intTests = []intTest{
	// Valid cases
	{testBase{"ENV_KEY", "1", []string{"ENV_KEY"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"env.key"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"ENV", "KEY"}, false}, 1},
	{testBase{"ENV_KEY", "0", []string{"ENV_KEY"}, false}, 0},
	{testBase{"ENV_KEY", "-1", []string{"ENV_KEY"}, false}, -1},
	{testBase{"ENV_KEY", "2147483647", []string{"ENV_KEY"}, false}, 2147483647},
	{testBase{"ENV_KEY", "-2147483648", []string{"ENV_KEY"}, false}, -2147483648},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "3.14", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "9223372036854775808", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "-9223372036854775809", []string{"ENV_KEY"}, true}, 0},
}

func TestInt(t *testing.T) {
	for _, test := range intTests {
		f := func() (bool, error) {
			out, err := Int(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustInt(test.in...), nil
		}

		testF(t, test, "Int", f, mf)
	}

	f := func() { MustInt("noenv") }
	if !isPanic(f) {
		t.Errorf("MustInt(noenv) should panic.")
	}
}

type int32Test struct {
	testBase
	out int32
}

var int32Tests = []int32Test{
	// Valid cases
	{testBase{"ENV_KEY", "1", []string{"ENV_KEY"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"env.key"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"ENV", "KEY"}, false}, 1},
	{testBase{"ENV_KEY", "0", []string{"ENV_KEY"}, false}, 0},
	{testBase{"ENV_KEY", "-1", []string{"ENV_KEY"}, false}, -1},
	{testBase{"ENV_KEY", "2147483647", []string{"ENV_KEY"}, false}, 2147483647},
	{testBase{"ENV_KEY", "-2147483648", []string{"ENV_KEY"}, false}, -2147483648},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "3.14", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "2147483648", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "-2147483649", []string{"ENV_KEY"}, true}, 0},
}

func TestInt32(t *testing.T) {
	for _, test := range int32Tests {
		f := func() (bool, error) {
			out, err := Int32(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustInt32(test.in...), nil
		}

		testF(t, test, "Int32", f, mf)
	}

	f := func() { MustInt32("noenv") }
	if !isPanic(f) {
		t.Errorf("MustInt32(noenv) should panic.")
	}
}

type int64Test struct {
	testBase
	out int64
}

var int64Tests = []int64Test{
	// Valid cases
	{testBase{"ENV_KEY", "1", []string{"ENV_KEY"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"env.key"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"ENV", "KEY"}, false}, 1},
	{testBase{"ENV_KEY", "0", []string{"ENV_KEY"}, false}, 0},
	{testBase{"ENV_KEY", "-1", []string{"ENV_KEY"}, false}, -1},
	{testBase{"ENV_KEY", "9223372036854775807", []string{"ENV_KEY"}, false}, 9223372036854775807},
	{testBase{"ENV_KEY", "-9223372036854775808", []string{"ENV_KEY"}, false}, -9223372036854775808},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "3.14", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "9223372036854775808", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "-9223372036854775809", []string{"ENV_KEY"}, true}, 0},
}

func TestInt64(t *testing.T) {
	for _, test := range int64Tests {
		f := func() (bool, error) {
			out, err := Int64(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustInt64(test.in...), nil
		}

		testF(t, test, "Int64", f, mf)
	}

	f := func() { MustInt64("noenv") }
	if !isPanic(f) {
		t.Errorf("MustInt64(noenv) should panic.")
	}
}

type stringTest struct {
	testBase
	out string
}

var stringTests = []stringTest{
	// Valid cases
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, false}, "foo"},
	{testBase{"ENV_KEY", "foo", []string{"env.key"}, false}, "foo"},
	{testBase{"ENV_KEY", "foo", []string{"ENV", "KEY"}, false}, "foo"},
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		f := func() (bool, error) {
			out := String(test.in...)
			return test.out == out, nil
		}

		mf := func() (bool, error) {
			return test.out == String(test.in...), nil
		}

		testF(t, test, "String", f, mf)
	}

	f := func() { MustString("noenv") }
	if !isPanic(f) {
		t.Errorf("MustString(noenv) should panic.")
	}
}

type uintTest struct {
	testBase
	out uint
}

var uintTests = []uintTest{
	// Valid cases
	{testBase{"ENV_KEY", "1", []string{"ENV_KEY"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"env.key"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"ENV", "KEY"}, false}, 1},
	{testBase{"ENV_KEY", "0", []string{"ENV_KEY"}, false}, 0},
	{testBase{"ENV_KEY", "4294967295", []string{"ENV_KEY"}, false}, 4294967295},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "-1", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "3.14", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "18446744073709551616", []string{"ENV_KEY"}, true}, 0},
}

func TestUint(t *testing.T) {
	for _, test := range uintTests {
		f := func() (bool, error) {
			out, err := Uint(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustUint(test.in...), nil
		}

		testF(t, test, "Uint", f, mf)
	}

	f := func() { MustUint("noenv") }
	if !isPanic(f) {
		t.Errorf("MustUint(noenv) should panic.")
	}
}

type uint32Test struct {
	testBase
	out uint32
}

var uint32Tests = []uint32Test{
	// Valid cases
	{testBase{"ENV_KEY", "1", []string{"ENV_KEY"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"env.key"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"ENV", "KEY"}, false}, 1},
	{testBase{"ENV_KEY", "0", []string{"ENV_KEY"}, false}, 0},
	{testBase{"ENV_KEY", "4294967295", []string{"ENV_KEY"}, false}, 4294967295},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "-1", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "3.14", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "4294967296", []string{"ENV_KEY"}, true}, 0},
}

func TestUint32(t *testing.T) {
	for _, test := range uint32Tests {
		f := func() (bool, error) {
			out, err := Uint32(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustUint32(test.in...), nil
		}

		testF(t, test, "Uint32", f, mf)
	}

	f := func() { MustUint32("noenv") }
	if !isPanic(f) {
		t.Errorf("MustUint32(noenv) should panic.")
	}
}

type uint64Test struct {
	testBase
	out uint64
}

var uint64Tests = []uint64Test{
	// Valid cases
	{testBase{"ENV_KEY", "1", []string{"ENV_KEY"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"env.key"}, false}, 1},
	{testBase{"ENV_KEY", "1", []string{"ENV", "KEY"}, false}, 1},
	{testBase{"ENV_KEY", "0", []string{"ENV_KEY"}, false}, 0},
	{testBase{"ENV_KEY", "18446744073709551615", []string{"ENV_KEY"}, false}, 18446744073709551615},

	// Invalid cases
	{testBase{"ENV_KEY", "", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "-1", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "foo", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "3.14", []string{"ENV_KEY"}, true}, 0},
	{testBase{"ENV_KEY", "18446744073709551616", []string{"ENV_KEY"}, true}, 0},
}

func TestUint64(t *testing.T) {
	for _, test := range uint64Tests {
		f := func() (bool, error) {
			out, err := Uint64(test.in...)
			return test.out == out, err
		}

		mf := func() (bool, error) {
			return test.out == MustUint64(test.in...), nil
		}

		testF(t, test, "Uint64", f, mf)
	}

	f := func() { MustUint64("noenv") }
	if !isPanic(f) {
		t.Errorf("MustUint64(noenv) should panic.")
	}
}

func testF(t *testing.T, test testBaseI, name string, f testFunc, mf testFunc) {
	test.withEnv(func() {
		e, err := f()
		if test.isErr() {
			if err == nil {
				t.Errorf("%s() with %v should return error.", name, test)
			}
		} else {
			if err != nil {
				t.Errorf("%s() with %v should not return error.", name, test)
			}
			if !e {
				t.Errorf("%s() with %v returns wrong result.", name, test)
			}
		}

		ip := isPanic(func() { e, _ = mf() })
		if test.isErr() {
			if !ip {
				t.Errorf("Must%s() with %v should panic.", name, test)
			}
		} else {
			if ip {
				t.Errorf("Must%s() with %v should not panic.", name, test)
			}
			if !e {
				t.Errorf("Must%s() with %v returns wrong result.", name, test)
			}
		}
	})
}

func isPanic(f func()) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			b = true
		}
	}()
	f()
	return
}
