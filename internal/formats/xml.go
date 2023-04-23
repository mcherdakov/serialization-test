package formats

import (
	"encoding/xml"

	"github.com/mcherdakov/serialization-test/internal/entity"
)

type xmlFormat struct{}

func (f *xmlFormat) Serialize(value entity.AnimalShelter) ([]byte, error) {
	return xml.Marshal(value)
}

func (f *xmlFormat) Deserialize(value []byte, result *entity.AnimalShelter) error {
	return xml.Unmarshal(value, result)
}
