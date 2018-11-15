package jsonapi

import "reflect"

type EncodeAttributes map[string]interface{}

func NewEncodeAttributes(m interface{}) EncodeAttributes {
	mType := reflect.TypeOf(m)
	mValue := reflect.ValueOf(m)

	attributes := EncodeAttributes{}
	for i := 0; i < mType.NumField(); i++ {
		if name, ok := mType.Field(i).Tag.Lookup("jsonapi"); ok {
			attributes[name] = mValue.Field(i).Interface()
		}
	}

	return attributes
}
