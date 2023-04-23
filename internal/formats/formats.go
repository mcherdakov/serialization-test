package formats

import "github.com/mcherdakov/serialization-test/internal/entity"

type FormatName string

const (
	Gob         FormatName = "gob" // "native" format
	JSON        FormatName = "json"
	XML         FormatName = "xml"
	Protobuf    FormatName = "protobuf"
	YAML        FormatName = "yaml"
	MessagePack FormatName = "message_pack"
)

type Format interface {
	Serialize(value entity.AnimalShelter) ([]byte, error)
	Deserialize(value []byte, result *entity.AnimalShelter) error
}

var FormatMapping = map[FormatName]Format{
	Gob:         &gobFormat{},
	JSON:        &jsonFormat{},
	XML:         &xmlFormat{},
	YAML:        &yamlFormat{},
	Protobuf:    &protoFormat{},
	MessagePack: &msgPackFormat{},
}
