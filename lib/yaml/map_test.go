package yaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestOrderedMapMarshal(t *testing.T) {
	// 创建测试数据
	yamlData := `
name: John
age: 30
city: New York
email: john@example.com
`

	// 解析到 OrderedMap
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlData), &node)
	if err != nil {
		t.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// 找到实际的 MappingNode
	var mappingNode *yaml.Node
	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		mappingNode = node.Content[0]
	} else {
		mappingNode = &node
	}

	om, err := NewOrderedMap(mappingNode)
	if err != nil {
		t.Fatalf("Failed to create OrderedMap: %v", err)
	}

	// 序列化回 YAML
	result, err := yaml.Marshal(om)
	if err != nil {
		t.Fatalf("Failed to marshal OrderedMap: %v", err)
	}

	t.Logf("Original:\n%s", yamlData)
	t.Logf("Result:\n%s", string(result))

	// 检查顺序
	t.Logf("Keys order: %v", om.Keys)

	// 验证包含所有键
	resultStr := string(result)
	expectedKeys := []string{"name", "age", "city", "email"}
	for _, key := range expectedKeys {
		if !contains(resultStr, key) {
			t.Errorf("Result doesn't contain key: %s", key)
		}
	}
}

func TestOrderedMapDirectUsage(t *testing.T) {
	// 直接创建 OrderedMap
	om := &OrderedMap{
		Keys:   []string{"name", "age", "city"},
		Values: make(map[string]*yaml.Node),
	}

	// 添加值
	om.Values["name"] = &yaml.Node{Kind: yaml.ScalarNode, Value: "Alice"}
	om.Values["age"] = &yaml.Node{Kind: yaml.ScalarNode, Value: "25"}
	om.Values["city"] = &yaml.Node{Kind: yaml.ScalarNode, Value: "Tokyo"}

	// 序列化
	result, err := yaml.Marshal(om)
	if err != nil {
		t.Fatalf("Failed to marshal OrderedMap: %v", err)
	}

	t.Logf("Direct usage result:\n%s", string(result))

	// 验证结果
	resultStr := string(result)
	if !contains(resultStr, "name") || !contains(resultStr, "Alice") {
		t.Error("Result doesn't contain name: Alice")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			len(s) > len(substr) &&
				(s[:len(substr)] == substr ||
					s[len(s)-len(substr):] == substr ||
					containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
