package jsonapi

import (
	"encoding/json"
	"reflect"
)

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
		if !ok || !field.CanAddr() {
			continue
		}

		addr := field.Addr().Interface()
		if err := json.Unmarshal(v, addr); err != nil {
			return err
		}
	}

	return nil
}
