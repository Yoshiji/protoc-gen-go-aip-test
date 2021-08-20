package get

import (
	"strconv"

	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/compiler/protogen"
)

var invalidName = suite.Test{
	Name: "invalid name",
	Doc: []string{
		"Method should fail with InvalidArgument is provided name is not valid.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		_, ok := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
		return ok
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		getMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeGet)
		util.MethodGet{
			Resource: scope.Resource,
			Method:   getMethod,
			Name:     strconv.Quote("invalid resource name"),
		}.Generate(f, "_", "err", ":=")
		f.P(ident.AssertEqual, "(t, ", ident.Codes(codes.InvalidArgument), ",", ident.StatusCode, "(err), err)")
		return nil
	},
}
