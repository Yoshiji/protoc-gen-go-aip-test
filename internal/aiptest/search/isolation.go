package search

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/ident"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
	"github.com/einride/protoc-gen-go-aip-test/internal/util"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/protobuf/compiler/protogen"
)

var isolation = suite.Test{
	Name: "isolation",
	Doc: []string{
		"If parent is provided the method must only return resources",
		"under that parent.",
	},

	OnlyIf: func(scope suite.Scope) bool {
		_, hasSearch := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeSearch)
		return hasSearch && util.HasParent(scope.Resource)
	},
	Generate: func(f *protogen.GeneratedFile, scope suite.Scope) error {
		searchMethod, _ := util.StandardMethod(scope.Service, scope.Resource, aipreflect.MethodTypeSearch)
		responseResources := aipreflect.GrammaticalName(scope.Resource.GetPlural()).UpperCamelCase()
		util.MethodSearch{
			Resource: scope.Resource,
			Method:   searchMethod,
			Parent:   "parent",
			PageSize: "999",
		}.Generate(f, "response", "err", ":=")
		f.P(ident.AssertNilError, "(t, err)")
		f.P(ident.AssertDeepEqual, "(")
		f.P("t,")
		f.P("parentMsgs,")
		f.P("response.", responseResources, ",")
		f.P(ident.CmpoptsSortSlices, "(func(a,b *", scope.Message.GoIdent, ") bool {")
		f.P("return a.Name < b.Name")
		f.P("}),")
		f.P(ident.ProtocmpTransform, "(),")
		f.P(")")
		return nil
	},
}