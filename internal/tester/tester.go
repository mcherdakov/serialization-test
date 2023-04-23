package tester

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mcherdakov/serialization-test/internal/entity"
	"github.com/mcherdakov/serialization-test/internal/formats"
)

type TestResult struct {
	Size                int
	SerializationTime   time.Duration
	DeserializationTime time.Duration
}

func RunTest(formatName string, runCount int) (string, error) {
	format := formats.FormatName(formatName)

	result, err := testFormatAverage(format, runCount)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s - %d - %s - %s\n",
		formatName,
		result.Size,
		result.SerializationTime,
		result.DeserializationTime,
	), nil
}

func testFormatAverage(formatName formats.FormatName, count int) (*TestResult, error) {
	var sizeSum int // in case there are randomized compression algorithms
	var serializationSum, deserializationSum time.Duration

	for i := 0; i < count; i++ {
		result, err := testFormat(formatName)
		if err != nil {
			return nil, err
		}

		serializationSum += result.SerializationTime
		deserializationSum += result.DeserializationTime
		sizeSum += result.Size
	}

	return &TestResult{
		Size:                sizeSum / count,
		SerializationTime:   serializationSum / time.Duration(count),
		DeserializationTime: deserializationSum / time.Duration(count),
	}, nil
}

func testFormat(formatName formats.FormatName) (*TestResult, error) {
	formatType, ok := formats.FormatMapping[formatName]
	if !ok {
		return nil, fmt.Errorf("invalid format name")
	}

	initial := entity.NewAnimalShelter()

	start := time.Now()
	serialized, err := formatType.Serialize(initial)
	if err != nil {
		return nil, err
	}
	serializationTime := time.Since(start)

	var deserialized entity.AnimalShelter

	start = time.Now()
	if err := formatType.Deserialize(serialized, &deserialized); err != nil {
		return nil, err
	}
	deserializationTime := time.Since(start)

	if !reflect.DeepEqual(initial, deserialized) {
		return nil, fmt.Errorf("deserialized is not equal to initial value")
	}

	return &TestResult{
		Size:                len(serialized),
		SerializationTime:   serializationTime,
		DeserializationTime: deserializationTime,
	}, nil
}
