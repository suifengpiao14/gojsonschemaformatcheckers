package gojsonschemavalidator

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

// Validate 验证
func Validate(input string, jsonLoader gojsonschema.JSONLoader) (err error) {
	if input == "" {
		jsonschema, err := jsonLoader.LoadJSON()
		if err != nil {
			return err
		}
		jsonMap, ok := jsonschema.(map[string]interface{})
		if !ok {
			err = errors.Errorf("can not convert jsonLoader.LoadJSON() to map[string]interface{}")
			return err
		}
		typ, ok := jsonMap["type"]
		if !ok {
			err = errors.Errorf("jsonschema missing property type :%v", jsonschema)
			return err
		}
		typStr, ok := typ.(string)
		if !ok {
			err = errors.Errorf("can not convert  jsonschema type to string :%v", typ)
			return err

		}
		switch strings.ToLower(typStr) {
		case "object":
			input = "{}"
		case "array":
			input = "[]"
		default:
			err = errors.Errorf("invalid jsonschema type:%v", typStr)
			return err
		}

	}
	documentLoader := gojsonschema.NewStringLoader(input)
	result, err := gojsonschema.Validate(jsonLoader, documentLoader)
	if err != nil {
		return err
	}
	if result.Valid() {
		return nil
	}

	msgArr := make([]string, 0)
	for _, resultError := range result.Errors() {
		msgArr = append(msgArr, resultError.String())
	}
	err = errors.Errorf("input args validate errors: %s", strings.Join(msgArr, ","))
	return err
}
