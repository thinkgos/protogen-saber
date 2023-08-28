package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.0.4"

var args = &struct {
	ShowVersion bool   // 显示当前版本
	TrimPrefix  bool   // 去掉文件前缀
	Merge       bool   // 合并到一个文件
	Filename    string // 合并文件名, 默认 create_table
}{
	ShowVersion: false,
	TrimPrefix:  false,
	Merge:       false,
	Filename:    "",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.BoolVar(&args.TrimPrefix, "trim_prefix", false, "trim filename prefix")
	flag.BoolVar(&args.Merge, "merge", false, "merge in a file")
	flag.StringVar(&args.Filename, "filename", "create_table", "filename when merge enabled")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-seaql %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
