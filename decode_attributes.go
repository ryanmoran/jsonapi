package jsonapi

import (
	"encoding/json"
	"reflect"
	"strings"
)

type DecodeField struct {
	Value     reflect.Value
	OmitEmpty bool
}

type DecodeAttributes struct {
	d Decodable
}

func NewDecodeAttributes(d Decodable) DecodeAttributes {
	return DecodeAttributes{d}
}

func (da DecodeAttributes) UnmarshalJSON(data []byte) error {
	var attributes map[string]json.RawMessage

	err := json.Unmarshal(data, &attributes)
	if err != nil {
		return err
	}

	dValue := reflect.ValueOf(da.d).Elem()
	dType := reflect.TypeOf(da.d).Elem()

	fieldMap := map[string]DecodeField{}
	for i := 0; i < dValue.NumField(); i++ {
		fieldStruct := dType.Field(i)
		fieldValue := dValue.Field(i)
		tag, ok := fieldStruct.Tag.Lookup("jsonapi")
		if !ok {
			continue
		}

		parts := strings.Split(tag, ",")
		omitEmpty := false
		for _, part := range parts {
			if part == "omitempty" {
				omitEmpty = true
			}
		}

		fieldMap[parts[0]] = DecodeField{
			Value:     fieldValue,
			OmitEmpty: omitEmpty,
		}
	}

	for k, v := range attributes {
		field, ok := fieldMap[k]
		if !ok || !field.Value.CanAddr() {
			continue
		}

		addr := field.Value.Addr().Interface()
		if err := json.Unmarshal(v, addr); err != nil {
			return err
		}
	}

	return nil
}
