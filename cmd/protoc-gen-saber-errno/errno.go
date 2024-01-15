package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/things-go/protogen-saber/internal/protoerrno"
	"github.com/things-go/protogen-saber/internal/protoutil"
)

const (
	fmtPackage = protogen.GoImportPath("fmt")
)

func runProtoGen(gen *protogen.Plugin) error {
	if args.ErrorsPackage == "" {
		log.Fatal("errors package import path must be give with '--saber-errno_out=epk=xxx'")
	}
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	usedTemplate, err := GetUsedTemplate()
	if err != nil {
		return err
	}
	for _, file := range gen.Files {
		if !file.Generate || len(file.Enums) == 0 {
			continue
		}
		enums := protoerrno.IntoEnums("", file.Enums)
		enums = append(enums, protoerrno.IntoEnumsFromMessage("", file.Messages)...)
		if len(enums) == 0 {
			continue
		}

		filename := file.GeneratedFilenamePrefix + ".errno.pb.go"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		g.P("// Code generated by protoc-gen-saber-errno. DO NOT EDIT.")
		g.P("// versions:")
		g.P("//   - protoc-gen-saber-errno ", version)
		g.P("//   - protoc                 ", protoutil.ProtocVersion(gen))
		if file.Proto.GetOptions().GetDeprecated() {
			g.P("// ", file.Desc.Path(), " is a deprecated file.")
		} else {
			g.P("// source: ", file.Desc.Path())
		}
		g.P()
		g.P("package ", file.GoPackageName)
		g.P()
		g.P("// Reference imports to suppress errors if they are not otherwise used.")
		g.P("var _ = ", fmtPackage.Ident("Errorf"))
		g.P("var _ = ", protogen.GoImportPath(args.ErrorsPackage).Ident("New"))
		g.P()
		g.P("// This is a compile-time assertion to ensure that this generated file")
		g.P("// is compatible with the errors package it is being compiled against.")
		g.P()

		mt := &errorWrapper{Errors: enums}
		err := mt.execute(usedTemplate, g)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr,
				"\u001B[31mWARN\u001B[m: execute template failed. %v\n", err)
		}
	}
	return nil
}
