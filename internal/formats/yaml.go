package formats

import (
	"github.com/mcherdakov/serialization-test/internal/entity"
	yaml "gopkg.in/yaml.v3"
)

type yamlFormat struct{}

func (f *yamlFormat) Serialize(value entity.AnimalShelter) ([]byte, error) {
	return yaml.Marshal(value)
}

func (f *yamlFormat) Deserialize(value []byte, result *entity.AnimalShelter) error {
	return yaml.Unmarshal(value, result)
}
