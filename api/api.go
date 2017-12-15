package api

import (
	"fmt"
	goast "go/ast"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gokit/httpkit/static"
	"github.com/influx6/faux/fmtwriter"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// HTTPGen generates http crud for giving struct declaration.
func HTTPGen(toDir string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	updateAction := str
	createAction := str

	var (
		isSameCreate = true
		isSameUpdate = true
	)

	if len(str.Associations) != 0 {
		if newAction, ok := str.Associations["New"]; ok {
			if action, err := ast.FindStructType(pkgDeclr, newAction.TypeName); err == nil && action.Declr != nil {
				if action.Object != str.Object {
					isSameCreate = false
				}

				createAction = action
			}
		}

		if upAction, ok := str.Associations["Update"]; ok {
			if action, err := ast.FindStructType(pkgDeclr, upAction.TypeName); err == nil && action.Declr != nil {
				if action.Object != str.Object {
					isSameUpdate = false
				}

				updateAction = action
			}
		}
	}

	packageName := fmt.Sprintf("%sapi", strings.ToLower(str.Object.Name.Name))
	var hasPublicID bool

	// Validate we have a `PublicID` field.
	{
	fieldLoop:
		for _, field := range str.Struct.Fields.List {
			typeIdent, ok := field.Type.(*goast.Ident)

			// if we are not a ast.Ident then skip
			if !ok {
				continue
			}

			// If typeName is not a string, skip.
			if typeIdent.Name != "string" {
				continue
			}

			for _, indent := range field.Names {
				if indent.Name == "PublicID" {
					hasPublicID = true
					break fieldLoop
				}
			}
		}
	}

	if !hasPublicID {
		return nil, fmt.Errorf(`Struct has no 'PublicID' field with 'string' type
		 Add 'PublicID string json:"public_id"' to struct %q
		`, str.Object.Name.Name)
	}

	httpGen := gen.Block(
		gen.Commentary(
			gen.SourceText(`Package `+packageName+` provides a auto-generated package which contains a http restful CRUD API for the specific {{.Object.Name}} struct in package {{.Package}}.`, str),
		),
		gen.Package(
			gen.Name(packageName),
			gen.Imports(
				gen.Import("fmt", ""),
				gen.Import("net/http", ""),
				gen.Import("encoding/json", ""),
				gen.Import("github.com/dimfeld/httptreemux", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/httputil", "httputil"),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("http-api.tml", true)),
					template.FuncMap{
						"map":       ast.MapOutFields,
						"mapValues": ast.MapOutValues,
						"mapJSON":   ast.MapOutFieldsToJSON,
						"hasFunc":   ast.HasFunctionFor(pkgDeclr),
					},
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	httpReadmeGen := gen.Block(
		gen.Block(
			gen.SourceTextWith(
				string(static.MustReadFile("http-api-readme.tml", true)),
				template.FuncMap{
					"map":       ast.MapOutFields,
					"mapValues": ast.MapOutValues,
					"mapJSON":   ast.MapOutFieldsToJSON,
					"hasFunc":   ast.HasFunctionFor(pkgDeclr),
				},
				struct {
					Pkg          *ast.PackageDeclaration
					Struct       ast.StructDeclaration
					CreateAction ast.StructDeclaration
					UpdateAction ast.StructDeclaration
				}{
					Pkg:          &pkgDeclr,
					Struct:       str,
					CreateAction: createAction,
					UpdateAction: updateAction,
				},
			),
		),
	)

	httpJSONGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("encoding/json", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("http-api-json.tml", true)),
					template.FuncMap{
						"map":       ast.MapOutFields,
						"mapValues": ast.MapOutValues,
						"mapJSON":   ast.MapOutFieldsToJSON,
						"hasFunc":   ast.HasFunctionFor(pkgDeclr),
					},
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	httpTestGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("fmt", ""),
				gen.Import("bytes", ""),
				gen.Import("testing", ""),
				gen.Import("encoding/json", ""),
				gen.Import("net/http", ""),
				gen.Import("net/http/httptest", ""),
				gen.Import("github.com/dimfeld/httptreemux", ""),
				gen.Import("github.com/influx6/faux/tests", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import(filepath.Join(str.Path, toDir, packageName), "httpapi"),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("http-api-test.tml", true)),
					template.FuncMap{
						"map":       ast.MapOutFields,
						"mapValues": ast.MapOutValues,
						"hasFunc":   ast.HasFunctionFor(pkgDeclr),
						"randField": ast.RandomFieldAssign,
					},
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	httpMockGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("errors", ""),
				gen.Import("testing", ""),
				gen.Import("encoding/json", ""),
				gen.Import("golang.org/x/sync/syncmap", ""),
				gen.Import("github.com/influx6/faux/tests", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("http-api-mock.tml", true)),
					template.FuncMap{
						"map":       ast.MapOutFields,
						"mapValues": ast.MapOutValues,
						"hasFunc":   ast.HasFunctionFor(pkgDeclr),
					},
					struct {
						Pkg             *ast.PackageDeclaration
						Struct          ast.StructDeclaration
						CreateAction    ast.StructDeclaration
						UpdateAction    ast.StructDeclaration
						CreateIsSimilar bool
						UpdateIsSimilar bool
					}{
						Pkg:             &pkgDeclr,
						Struct:          str,
						CreateAction:    createAction,
						UpdateAction:    updateAction,
						CreateIsSimilar: isSameCreate,
						UpdateIsSimilar: isSameUpdate,
					},
				),
			),
		),
	)

	httpMockHelperGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("errors", ""),
				gen.Import("testing", ""),
				gen.Import("encoding/json", ""),
				gen.Import("github.com/influx6/faux/tests", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("http-api-mock-functions.tml", true)),
					template.FuncMap{
						"map":       ast.MapOutFields,
						"mapValues": ast.MapOutValues,
						"hasFunc":   ast.HasFunctionFor(pkgDeclr),
					},
					struct {
						Pkg             *ast.PackageDeclaration
						Struct          ast.StructDeclaration
						CreateAction    ast.StructDeclaration
						UpdateAction    ast.StructDeclaration
						CreateIsSimilar bool
						UpdateIsSimilar bool
					}{
						Pkg:             &pkgDeclr,
						Struct:          str,
						CreateAction:    createAction,
						UpdateAction:    updateAction,
						CreateIsSimilar: isSameCreate,
						UpdateIsSimilar: isSameUpdate,
					},
				),
			),
		),
	)

	writers := []gen.WriteDirective{
		{
			Writer:   httpReadmeGen,
			FileName: "readme.md",
			Dir:      packageName,
			// DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(httpMockGen, true, true),
			FileName: fmt.Sprintf("%s_mock_test.go", packageName),
			Dir:      packageName,
			// DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(httpTestGen, true, true),
			FileName: fmt.Sprintf("%s_test.go", packageName),
			Dir:      packageName,
			// DontOverride: true,
		},
		{
			Writer:       fmtwriter.New(httpJSONGen, true, true),
			FileName:     fmt.Sprintf("%s_fixtures_test.go", packageName),
			Dir:          packageName,
			DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(httpGen, true, true),
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
			// DontOverride: true,
		},
	}

	if !isSameCreate || !isSameUpdate {
		writers = append(writers, gen.WriteDirective{
			Writer:       fmtwriter.New(httpMockHelperGen, true, true),
			FileName:     fmt.Sprintf("%s_create_update_mock_test.go", packageName),
			Dir:          packageName,
			DontOverride: true,
		})
	}

	return writers, nil
}
