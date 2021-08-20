package get

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

var notFound = suite.Test{
	Name: "not found",
	Doc: []string{
		"Method should fail with NotFound if the resource does not exist.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		_, hasGet := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
		create, hasCreate := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeCreate)
		return hasGet && hasCreate && !util.ReturnsLRO(create.Desc)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		getMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
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

		util.MethodGet{
			Resource: scope.Resource,
			Method:   getMethod,
			// appending to the resource name ensures it is valid
			Name: "created.Name + \"notfound\"",
		}.Generate(f, "_", "err", "=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.NotFound), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}