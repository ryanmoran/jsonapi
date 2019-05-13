package jsonapi

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type DecodeAttributes struct {
	d Decodable
}

func NewDecodeAttributes(d Decodable) DecodeAttributes {
	return DecodeAttributes{d}
}

func (da DecodeAttributes) UnmarshalJSON(data []byte) error {
	var attributes map[string]interface{}

	err := json.Unmarshal(data, &attributes)
	if err != nil {
		return err
	}

	dValue := reflect.ValueOf(da.d).Elem()
	dType := reflect.TypeOf(da.d).Elem()

	fieldMap := map[string]reflect.Value{}
	for i := 0; i < dValue.NumField(); i++ {
		fieldStruct := dType.Field(i)
		fieldValue := dValue.Field(i)
		tag, ok := fieldStruct.Tag.Lookup("jsonapi")
		if !ok {
			continue
		}

		fieldMap[tag] = fieldValue
	}

	for k, v := range attributes {
		field, ok := fieldMap[k]
		if !ok || !field.CanSet() {
			continue
		}

		if field.Type().AssignableTo(reflect.TypeOf(json.RawMessage{})) {
			v, err = json.Marshal(v)
			if err != nil {
				return err
			}
		}

		value := reflect.ValueOf(v)
		if value.Kind() != field.Kind() {
			return NewDecodeError(da.d, fmt.Sprintf("field %q types %q and %q do not match", k, field.Kind(), value.Kind()))
		}

		field.Set(value)
	}

	return nil
}
