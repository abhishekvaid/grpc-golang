package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func protobufToJsonString(message proto.Message) (string, error) {
	marshler := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
	}
	return marshler.MarshalToString(message)
}

func WriteProtobufToJsonFile(fname string, message proto.Message) error {

	jsonString, err := protobufToJsonString(message)

	if err != nil {
		return fmt.Errorf("Cannot Marshal Proto Message To Json String %w", err)
	}
	err = ioutil.WriteFile(fname, []byte(jsonString), 0644)

	if err != nil {
		return fmt.Errorf("Cannot write proto message to a json file %w", err)
	}
	return nil
}
