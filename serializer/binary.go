package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func WriteProtobufToBinaryFile(fname string, message proto.Message) error {

	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Cannot Marshal Proto Message To Binary %w", err)
	}
	err = ioutil.WriteFile(fname, data, 0644)
	if err != nil {
		return fmt.Errorf("Cannot write proto message to a file %w", err)
	}
	return nil

}

func ReadBinaryFileToProtoBuf(fname string, message proto.Message) error {

	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("Cannot read from file into Protobuf %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("Cannot UnMarshal data into Protobuf %w", err)
	}

	return nil

}
