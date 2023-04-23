package formats

import (
	"github.com/mcherdakov/serialization-test/internal/entity"
	"github.com/mcherdakov/serialization-test/internal/gen/proto"
	gproto "google.golang.org/protobuf/proto"
)

type protoFormat struct{}

func (f *protoFormat) Serialize(value entity.AnimalShelter) ([]byte, error) {
	return gproto.Marshal(value.ToProto())
}

func (f *protoFormat) Deserialize(value []byte, result *entity.AnimalShelter) error {
	var resProto proto.AnimalShelter

	if err := gproto.Unmarshal(value, &resProto); err != nil {
		return err
	}

	*result = entity.FromProto(&resProto)
	return nil
}
