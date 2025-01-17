package asa

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// addParamsToQuery 使用反射将结构体的字段添加到查询参数中
func addParamsToQuery(query url.Values, params interface{}) error {
	v := reflect.ValueOf(params)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("params must be a non-nil pointer to a struct")
	}
	return addParamsToQueryRecursive(query, v.Elem())
}

func addParamsToQueryRecursive(query url.Values, v reflect.Value) error {
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 处理匿名结构体
		if fieldType.Anonymous {
			if err := addParamsToQueryRecursive(query, field); err != nil {
				return err
			}
			continue
		}

		jsonTag := fieldType.Tag.Get("form")
		paramTag := fieldType.Tag.Get("param")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		jsonName := strings.Split(jsonTag, ",")[0]

		// 处理默认值
		if strings.HasPrefix(paramTag, "default=") && field.IsZero() {
			defaultValue := strings.TrimPrefix(paramTag, "default=")
			switch field.Kind() {
			case reflect.Int, reflect.Int64:
				if val, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
					field.SetInt(val)
				}
			case reflect.String:
				field.SetString(defaultValue)
			default:
				panic("unhandled default case")
			}
		}

		// 根据字段类型和标签处理参数
		switch {
		case field.Kind() == reflect.Slice, field.Kind() == reflect.Array:
			if !field.IsZero() {
				jsonValue, err := json.Marshal(field.Interface())
				if err != nil {
					return err
				}
				query.Set(jsonName, string(jsonValue))
			}
		case field.Kind() == reflect.Struct:
			if !field.IsZero() {
				jsonValue, err := json.Marshal(field.Interface())
				if err != nil {
					return err
				}
				query.Set(jsonName, string(jsonValue))
			}
		case field.Kind() == reflect.Int, field.Kind() == reflect.Int32, field.Kind() == reflect.Int64:
			if !field.IsZero() {
				query.Set(jsonName, strconv.FormatInt(field.Int(), 10))
			}
		case field.Kind() == reflect.String:
			if !field.IsZero() && field.String() != "" {
				query.Set(jsonName, field.String())
			}
		case field.Kind() == reflect.Bool:
			if !field.IsZero() {
				query.Set(jsonName, fmt.Sprintf("%v", field.Bool()))
			}
		default:
			panic("unhandled field type")
		}
	}
	return nil
}

func parsePrivateKey(blob []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(blob)
	if block == nil {
		return nil, ErrMissingPEM
	}

	parsedKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedKey, nil
}
