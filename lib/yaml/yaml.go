package yaml

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

func MarshalIndent(data any, indent int) ([]byte, error) {
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(indent) // 设置缩进空格数

	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	encoder.Close()
	return buf.Bytes(), nil
}

func Marshal(data any) ([]byte, error) {
	return yaml.Marshal(data)
}

func Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}
