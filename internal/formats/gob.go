package formats

import (
	"bytes"
	"encoding/gob"

	"github.com/mcherdakov/serialization-test/internal/entity"
)

type gobFormat struct{}

func (f *gobFormat) Serialize(value entity.AnimalShelter) ([]byte, error) {
	buf := bytes.Buffer{}

	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(value); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (f *gobFormat) Deserialize(value []byte, result *entity.AnimalShelter) error {
	decoder := gob.NewDecoder(bytes.NewBuffer(value))
	return decoder.Decode(result)
}
