package yaml

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	_ yaml.Unmarshaler = &OrderedMap{}
	_ yaml.Marshaler   = &OrderedMap{}
)

type OrderedMap struct {
	Keys   []string
	Values map[string]*yaml.Node
}

func NewOrderedMap(node *yaml.Node) (*OrderedMap, error) {
	if node.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("expected mapping node")
	}

	om := &OrderedMap{
		Keys:   make([]string, 0),
		Values: make(map[string]*yaml.Node),
	}

	for i := 0; i < len(node.Content); i += 2 {
		key := node.Content[i].Value
		value := node.Content[i+1]

		om.Keys = append(om.Keys, key)
		om.Values[key] = value
	}

	return om, nil
}

// 按顺序遍历
func (om *OrderedMap) Range(fn func(key string, value *yaml.Node) error) error {
	for _, key := range om.Keys {
		if err := fn(key, om.Values[key]); err != nil {
			return err
		}
	}
	return nil
}

func (om OrderedMap) MarshalYAML() (any, error) {
	// 创建一个新的 MappingNode 来保持顺序
	result := &yaml.Node{
		Kind:    yaml.MappingNode,
		Style:   0, // 使用默认样式
		Tag:     "!!map",
		Content: make([]*yaml.Node, 0, len(om.Keys)*2),
	}

	// 按照 Keys 的顺序添加键值对
	for _, key := range om.Keys {
		if value, exists := om.Values[key]; exists {
			// 创建 key node，保持原有的样式
			keyNode := &yaml.Node{
				Kind:  yaml.ScalarNode,
				Style: 0,
				Tag:   "!!str",
				Value: key,
			}

			// 确保 value node 有正确的标签
			if value.Tag == "" {
				switch value.Kind {
				case yaml.ScalarNode:
					// 自动推断标签
					if value.Tag == "" {
						value.Tag = "!!str"
					}
				case yaml.MappingNode:
					value.Tag = "!!map"
				case yaml.SequenceNode:
					value.Tag = "!!seq"
				}
			}

			result.Content = append(result.Content, keyNode, value)
		}
	}

	return result, nil
}

func (om *OrderedMap) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("expected mapping node, got %v", node.Kind)
	}

	// 初始化
	om.Keys = make([]string, 0)
	om.Values = make(map[string]*yaml.Node)

	// 解析键值对，保持顺序
	for i := 0; i < len(node.Content); i += 2 {
		if i+1 >= len(node.Content) {
			return fmt.Errorf("malformed mapping node: odd number of content items")
		}

		keyNode := node.Content[i]
		valueNode := node.Content[i+1]

		if keyNode.Kind != yaml.ScalarNode {
			return fmt.Errorf("mapping key must be scalar, got %v", keyNode.Kind)
		}

		key := keyNode.Value

		// 避免重复的键
		if _, exists := om.Values[key]; !exists {
			om.Keys = append(om.Keys, key)
		}
		om.Values[key] = valueNode
	}

	return nil
}

// 添加或更新键值对
func (om *OrderedMap) Set(key string, value *yaml.Node) {
	if _, exists := om.Values[key]; !exists {
		om.Keys = append(om.Keys, key)
	}
	om.Values[key] = value
}

// 获取值
func (om *OrderedMap) Get(key string) (*yaml.Node, bool) {
	value, exists := om.Values[key]
	return value, exists
}

// 删除键值对
func (om *OrderedMap) Delete(key string) bool {
	if _, exists := om.Values[key]; !exists {
		return false
	}

	delete(om.Values, key)

	// 从 Keys 中移除
	for i, k := range om.Keys {
		if k == key {
			om.Keys = append(om.Keys[:i], om.Keys[i+1:]...)
			break
		}
	}

	return true
}

// 获取键的数量
func (om *OrderedMap) Len() int {
	return len(om.Keys)
}

// 检查是否包含键
func (om *OrderedMap) Has(key string) bool {
	_, exists := om.Values[key]
	return exists
}

// 转换为普通的 map[string]interface{}
func (om *OrderedMap) ToMap() (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for _, key := range om.Keys {
		if node, exists := om.Values[key]; exists {
			var value interface{}
			if err := node.Decode(&value); err != nil {
				return nil, fmt.Errorf("failed to decode value for key %s: %v", key, err)
			}
			result[key] = value
		}
	}

	return result, nil
}
