package formats

import (
	"github.com/mcherdakov/serialization-test/internal/entity"
	"github.com/vmihailenco/msgpack/v5"
)

type msgPackFormat struct{}

func (f *msgPackFormat) Serialize(value entity.AnimalShelter) ([]byte, error) {
	return msgpack.Marshal(value)
}

func (f *msgPackFormat) Deserialize(value []byte, result *entity.AnimalShelter) error {
	return msgpack.Unmarshal(value, result)
}
