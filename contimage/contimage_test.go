package contimage

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestHiddenBytesPresence(t *testing.T) {
	for _, test := range []struct {
		name  string
		value *uint64
	}{
		{name: "unavailable"},
		{name: "zero", value: proto.Uint64(0)},
		{name: "nonzero", value: proto.Uint64(42)},
	} {
		t.Run(test.name, func(t *testing.T) {
			input := &ContainerImage_ContainerImageLayer{HiddenBytes: test.value}
			wire, err := proto.Marshal(input)
			if err != nil {
				t.Fatal(err)
			}
			output := new(ContainerImage_ContainerImageLayer)
			if err := proto.Unmarshal(wire, output); err != nil {
				t.Fatal(err)
			}
			if !proto.Equal(input, output) {
				t.Fatalf("hidden bytes presence or value changed: got %v, want %v", output, input)
			}
			field := output.ProtoReflect().Descriptor().Fields().ByName("hidden_bytes")
			if output.ProtoReflect().Has(field) != (test.value != nil) {
				t.Fatal("hidden bytes field presence did not survive serialization")
			}
		})
	}
}
