package batchget

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var ordered = suite.Test{
	Name: "ordered",
	Doc: []string{
		"The order of resources in the response must be the same as the names in the request.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		batchGetMethod, hasBatchGet := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		return hasBatchGet &&
			!util.IsAlternativeBatchGet(batchGetMethod.Desc)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		batchGetMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeBatchGet)
		responseResources := aipreflect.GrammaticalName(scope.Resource.GetPlural()).UpperCamelCase()
		f.P("for _, order := range [][]*", scope.Message.GoIdent, "{")
		f.P("{created00, created01, created02},")
		f.P("{created01, created00, created02},")
		f.P("{created02, created01, created00},")
		f.P("} {")
		util.MethodBatchGet{
			Resource: scope.Resource,
			Method:   batchGetMethod,
			Parent:   "parent",
			Names:    []string{"order[0].GetName()", "order[1].GetName()", "order[2].GetName()"},
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(t, order, response.", responseResources, ",", ident.ProtocmpTransform, "())")
		f.P("}")
		return nil
	},
}