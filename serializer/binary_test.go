package serializer

import (
	"testing"

	"github.com/abhishekvaid/grpc/pcbook/pb"
	"github.com/abhishekvaid/grpc/pcbook/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestBinarySerialization(t *testing.T) {

	t.Parallel()

	laptop1 := sample.NewLaptop()
	laptop2 := &pb.Laptop{}

	fname := "../tmp/laptop.bin"

	t.Run("Write proto message to a file", func(t *testing.T) {
		err := WriteProtobufToBinaryFile(fname, laptop1)
		require.NoError(t, err)
	})

	t.Run("Reading proto message from a binary file", func(t *testing.T) {
		err := ReadBinaryFileToProtoBuf(fname, laptop2)
		require.NoError(t, err)
		require.True(t, proto.Equal(laptop1, laptop2))
	})

}

func TestJsonSerialization(t *testing.T) {

	t.Parallel()

	laptop1 := sample.NewLaptop()
	// laptop2 := &pb.Laptop{}

	fname := "../tmp/laptop.json"

	t.Run("Write proto message to a file", func(t *testing.T) {
		err := WriteProtobufToJsonFile(fname, laptop1)
		require.NoError(t, err)
	})

	// t.Run("Reading proto message from a binary file", func(t *testing.T) {
	// 	err := ReadBinaryFileToProtoBuf(fname, laptop2)
	// 	require.NoError(t, err)
	// 	require.True(t, proto.Equal(laptop1, laptop2))
	// })

}
