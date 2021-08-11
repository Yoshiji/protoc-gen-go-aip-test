package plugin

import "google.golang.org/protobuf/compiler/protogen"

type testCase struct {
	enabled bool
	name    string
	fn      func(file *protogen.GeneratedFile)
}

func disabledTestCase() testCase {
	return testCase{}
}

func newTestCase(name string, fn func(f *protogen.GeneratedFile)) testCase {
	return testCase{
		enabled: true,
		name:    name,
		fn:      fn,
	}
}

func (t testCase) Name() string {
	return t.name
}

func (t testCase) FuncName() string {
	return "test" + t.name
}