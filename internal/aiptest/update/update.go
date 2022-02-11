package update

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/onlyif"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

// Suite for the Updaet method.
// nolint: gochecknoglobals
var Suite = suite.Suite{
	Name: "Update",
	Tests: []suite.Test{
		missingName,
		invalidName,
		updateTime,
		persisted,
		preserveCreateTime,
	},
	TestGroups: []suite.TestGroup{
		withResourceGroup,
	},
}

// nolint: gochecknoglobals
var withResourceGroup = suite.TestGroup{
	OnlyIf: suite.OnlyIfs(
		onlyif.HasMethod(aipreflect.MethodTypeCreate),
		onlyif.MethodNotLRO(aipreflect.MethodTypeCreate),
	),
	GenerateBefore: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		createMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		if util.HasParent(scope.Resource) {
			f.P("parent := ", ident.FixtureNextParent, "(t, false)")
		}
		util.MethodCreate{
			Resource: scope.Resource,
			Method:   createMethod,
			Parent:   "parent",
		}.Generate(f, "created", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		return nil
	},
	Tests: []suite.Test{
		notFound,
		invalidUpdateMask,
		requiredFields,
		// TODO: add test for supplying wildcard as name
		// TODO: add test for etags
	},
}
