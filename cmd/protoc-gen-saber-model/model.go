package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/things-go/ens"
	"github.com/things-go/ens/codegen"
	mysqlDriver "github.com/things-go/ens/driver/mysql"
	"github.com/things-go/protogen-saber/internal/protoseaql"
	"github.com/things-go/protogen-saber/internal/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func runProtoGen(gen *protogen.Plugin) error {
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		tables, err := protoseaql.IntoTable(f.Messages, args.DisableOrComment)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mERROR\u001B[m: %v\n", err)
		}
		if len(tables) == 0 {
			continue
		}

		dir := filepath.Dir(f.GeneratedFilenamePrefix)
		if args.TrimPrefix {
			dir = ""
		}
		for _, tb := range tables {
			filename := filepath.Join(dir, tb.Name) + ".go"
			g := gen.NewGeneratedFile(filename, f.GoImportPath)
			setGeneratedFileHeader(g, gen, []protoutil.Source{
				{Path: f.Desc.Path(), IsDeprecated: f.Proto.GetOptions().GetDeprecated()},
			})
			buf := &strings.Builder{}
			err = protoseaql.Execute(buf, &protoseaql.Schema{
				Tables: []protoseaql.Table{tb},
			})
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mERROR\u001B[m: %v\n", err)
				continue
			}
			d := mysqlDriver.SQL{
				CreateTableSQL: buf.String(),
			}
			schmaer, err := d.GetSchema()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mERROR\u001B[m: %v\n", err)
				continue
			}
			schema := schmaer.Build(&ens.Option{
				EnableInt:          false,                                       // no need
				EnableIntegerInt:   false,                                       // no need
				EnableBoolInt:      false,                                       // no need
				DisableNullToPoint: false,                                       // no need
				EnableForeignKey:   false,                                       // no need
				Tags:               map[string]string{"json": ens.TagSnakeCase}, // TODO: add more
				EnableGogo:         false,                                       // no need
				EnableSea:          false,                                       // no need
			})
			packageName := string(f.GoPackageName)
			if args.Package != "" {
				packageName = args.Package
			}
			gen := codegen.
				New(
					schema.Entities,
					codegen.WithPackageName(packageName),
					codegen.WithDisableDocComment(true),
				).
				GenModel()
			data, err := gen.FormatSource()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mERROR, format %s source, \u001B[m: %v\n", filename, err)
				data = gen.Bytes()
			}
			_, _ = g.Write(data)
		}
	}
	return nil
}

func setGeneratedFileHeader(g *protogen.GeneratedFile, gen *protogen.Plugin, sources []protoutil.Source) {
	g.P("// Code generated by protoc-gen-saber-model. DO NOT EDIT.")
	g.P("// versions:")
	g.P("//   - protoc-gen-saber-model ", version)
	g.P("//   - protoc                  ", protoutil.ProtocVersion(gen))
	for _, v := range sources {
		if v.IsDeprecated {
			g.P("// ", v.Path, " is a deprecated file.")
		} else {
			g.P("// source: ", v.Path)
		}
	}
}
