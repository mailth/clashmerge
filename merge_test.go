package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestMerge(t *testing.T) {
	for i:=1;i<=2;i++{
		yamlFile := fmt.Sprintf("testdata/test%d-in.yaml", i)
		res,err := os.ReadFile(yamlFile)
		if err != nil {
			t.Fatalf("read file %s failed: %v", yamlFile, err)
		}
		var items []*ConfigItem
		err = yaml.Unmarshal(res, &items)
		if err != nil {
			t.Fatalf("unmarshal file %s failed: %v", yamlFile, err)
		}

		yamlFile = fmt.Sprintf("testdata/test%d-out.yaml", i)
		res,err = os.ReadFile(yamlFile)
		if err != nil {
			t.Fatalf("read file %s failed: %v", yamlFile, err)
		}
		var expect map[string]any
		err = yaml.Unmarshal(res, &expect)
		if err != nil {
			t.Fatalf("unmarshal file %s failed: %v", yamlFile, err)
		}
		actual,err := doMerge(nil, items)
		if err != nil {
			t.Fatalf("merge failed: %v", err)
		}
		if !reflect.DeepEqual(actual, expect) {
			t.Fatalf("merge result not match, expect %v, actual %v", expect, actual)
		}
	}
}
