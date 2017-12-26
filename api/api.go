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

// Action defines a type that holds the name of a giving action type.
type Action struct {
	Object   string
	Package  string
	IsStruct bool
	Struct   ast.StructDeclaration
	Type     ast.TypeDeclaration
}

// HTTPGen generates http crud for giving struct declaration.
func HTTPGen(toPackage string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	var updateAction, createAction Action

	if newActionName := an.Param("New"); newActionName != "" {
		if action, ok := pkg.StructFor(pkgDeclr.Path, newActionName); ok && action.Declr != nil {
			createAction.Package = action.Package
			createAction.Object = action.Object.Name.Name
			createAction.Struct = action
			createAction.IsStruct = true
		} else if action, ok := pkg.TypeFor(pkgDeclr.Path, newActionName); ok && action.Declr != nil {
			createAction.Package = action.Package
			createAction.Object = action.Object.Name.Name
			createAction.Type = action
		} else {
			return nil, fmt.Errorf("New type %+q is not found in struct or type declarations", newActionName)
		}
	} else {
		createAction.Package = str.Package
		createAction.Object = str.Object.Name.Name
		createAction.Struct = str
		createAction.IsStruct = true
	}

	if updateActionName := an.Param("Update"); updateActionName != "" {
		if action, ok := pkg.StructFor(pkgDeclr.Path, updateActionName); ok && action.Declr != nil {
			updateAction.Package = action.Package
			updateAction.Object = action.Object.Name.Name
			updateAction.Struct = action
			updateAction.IsStruct = true
		} else if action, ok := pkg.TypeFor(pkgDeclr.Path, updateActionName); ok && action.Declr != nil {
			updateAction.Package = action.Package
			updateAction.Object = action.Object.Name.Name
			updateAction.Type = action
		} else {
			return nil, fmt.Errorf("Update type %+q is not found in struct or type declarations", updateActionName)
		}
	} else {
		updateAction.Package = str.Package
		updateAction.Object = str.Object.Name.Name
		updateAction.Struct = str
		updateAction.IsStruct = true
	}

	packageName := fmt.Sprintf("%sapi", strings.ToLower(str.Object.Name.Name))
	packageFinalPath := filepath.Join(toPackage, packageName)
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
				gen.Import("context", ""),
				gen.Import("net/http", ""),
				gen.Import("encoding/json", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/httputil", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("http-api.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"map":           ast.MapOutFields,
							"mapValues":     ast.MapOutValues,
							"mapJSON":       ast.MapOutFieldsToJSON,
							"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
							"hasFunc":       pkgDeclr.HasFunctionFor,
						},
					),
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction Action
						UpdateAction Action
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
				gen.ToTemplateFuncs(
					ast.ASTTemplatFuncs,
					template.FuncMap{
						"map":           ast.MapOutFields,
						"mapValues":     ast.MapOutValues,
						"mapJSON":       ast.MapOutFieldsToJSON,
						"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
						"hasFunc":       pkgDeclr.HasFunctionFor,
					},
				),
				struct {
					Pkg          *ast.PackageDeclaration
					Struct       ast.StructDeclaration
					CreateAction Action
					UpdateAction Action
					PackageName  string
					PackagePath  string
				}{
					PackagePath:  packageFinalPath,
					PackageName:  packageName,
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
			gen.Name("fixtures"),
			gen.Imports(
				gen.Import("encoding/json", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("http-api-json.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"map":           ast.MapOutFields,
							"mapValues":     ast.MapOutValues,
							"mapJSON":       ast.MapOutFieldsToJSON,
							"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
							"hasFunc":       pkgDeclr.HasFunctionFor,
						},
					),
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction Action
						UpdateAction Action
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

	return []gen.WriteDirective{
		{
			Writer:   httpReadmeGen,
			FileName: "readme.md",
			Dir:      packageName,
		},
		{
			Writer:       fmtwriter.New(httpJSONGen, true, true),
			FileName:     fmt.Sprintf("%s_fixtures.go", packageName),
			Dir:          filepath.Join(packageName, "fixtures"),
			DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(httpGen, true, true),
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
		},
	}, nil
}
