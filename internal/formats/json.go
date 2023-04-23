package formats

import (
	"encoding/json"

	"github.com/mcherdakov/serialization-test/internal/entity"
)

type jsonFormat struct{}

func (f *jsonFormat) Serialize(value entity.AnimalShelter) ([]byte, error) {
	return json.Marshal(value)
}

func (f *jsonFormat) Deserialize(value []byte, result *entity.AnimalShelter) error {
	return json.Unmarshal(value, result)
}
