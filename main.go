package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func run() error {
	var req plugin.CodeGeneratorRequest
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	if err = proto.Unmarshal(buf, &req); err != nil {
		return err
	}

	files := make(map[string]*descriptor.FileDescriptorProto)
	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}
	var resp plugin.CodeGeneratorResponse
	for _, fname := range req.FileToGenerate {
		f := files[fname]
		out := fname + ".dump"
		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(out),
			Content: proto.String(proto.MarshalTextString(f)),
		})
	}

	buf, err = proto.Marshal(&resp)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(buf)
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
