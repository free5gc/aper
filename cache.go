package aper

import (
	"fmt"
	"reflect"
	"sync"
)

var fieldParametersCache sync.Map
var structFieldCache structFieldCacheT

type structFieldCacheT struct {
	m sync.Map
}

func (f *structFieldCacheT) load(k reflect.Type) ([]structFieldElem, error) {
	v, ok := f.m.Load(k)
	if ok {
		return v.([]structFieldElem), nil
	}

	numField := k.NumField()
	result := make([]structFieldElem, 0, numField)
	for i := 0; i < numField; i++ {
		field := k.Field(i)
		if field.PkgPath != "" {
			return nil, fmt.Errorf("struct contains unexported fields : " + field.PkgPath)
		}

		result = append(result, structFieldElem{
			FieldName:       field.Name,
			FieldParameters: parseFieldParameters(field.Tag.Get("aper")),
		})
	}
	f.m.Store(k, result)

	return result, nil
}

type structFieldElem struct {
	FieldName       string
	FieldParameters fieldParameters
}
