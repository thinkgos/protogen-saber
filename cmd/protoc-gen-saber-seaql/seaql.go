package main

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/things-go/protogen-saber/internal/protoseaql"
	"github.com/things-go/protogen-saber/internal/protoutil"
)

func runProtoGen(gen *protogen.Plugin) error {
	var mergeTables []protoseaql.Table
	var sources []protoutil.Source

	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	hasMerge := args.Merge || args.ExtraMerge
	if hasMerge {
		mergeTables = make([]protoseaql.Table, 0, len(gen.Files)*4)
		sources = make([]protoutil.Source, 0, len(gen.Files))
	}

	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		tables, err := protoseaql.IntoTable(f.Messages)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "\u001B[31mERROR\u001B[m: %v\n", err)
		}
		if len(tables) == 0 {
			continue
		}
		if hasMerge {
			sources = append(sources, protoutil.Source{
				Path:         f.Desc.Path(),
				IsDeprecated: f.Proto.GetOptions().GetDeprecated(),
			})
			mergeTables = append(mergeTables, tables...)
			if !args.ExtraMerge {
				continue
			}
		}

		dir := filepath.Dir(f.GeneratedFilenamePrefix)
		if args.TrimPrefix {
			dir = ""
		}
		for _, tb := range tables {
			g := gen.NewGeneratedFile(filepath.Join(dir, tb.Name)+".sql", f.GoImportPath)
			setGeneratedFileHeader(g, gen, []protoutil.Source{
				{Path: f.Desc.Path(), IsDeprecated: f.Proto.GetOptions().GetDeprecated()},
			})
			_ = protoseaql.Execute(g, &protoseaql.Schema{
				Tables: []protoseaql.Table{tb},
			})
		}
	}
	if hasMerge && len(mergeTables) > 0 {
		g := gen.NewGeneratedFile(args.Filename+".sql", "")
		setGeneratedFileHeader(g, gen, sources)
		return protoseaql.Execute(g, &protoseaql.Schema{
			Tables: mergeTables,
		})
	}
	return nil
}

func setGeneratedFileHeader(g *protogen.GeneratedFile, gen *protogen.Plugin, sources []protoutil.Source) {
	g.P("-- Code generated by protoc-gen-saber-seaql. DO NOT EDIT.")
	g.P("-- versions:")
	g.P("--   - protoc-gen-saber-seaql ", version)
	g.P("--   - protoc                 ", protoutil.ProtocVersion(gen))
	for _, v := range sources {
		if v.IsDeprecated {
			g.P("-- ", v.Path, " is a deprecated file.")
		} else {
			g.P("-- source: ", v.Path)
		}
	}
	g.P()
}
