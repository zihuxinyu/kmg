package webTypeAdmin

import (
	"fmt"
	"html/template"
	"reflect"
)

//just for same with golang reflect system
//path -> pass to RefType
type ptrType struct {
	commonType
	elemType typeInterface
}

func (t *ptrType) init() {
	if t.elemType != nil {
		return
	}
	t.elemType = mustNewTypeFromReflect(t.getReflectType().Elem())
}
func (t *ptrType) Html(v reflect.Value) template.HTML {
	t.init()
	if v.IsNil() {
		return theTemplate.MustExecuteNameToHtml("NilPtr", nil)
	}
	return t.elemType.Html(v.Elem())
}
func (t *ptrType) getSubValueByString(v reflect.Value, k string) (reflect.Value, error) {
	t.init()
	if v.IsNil() {
		return reflect.Value{}, fmt.Errorf("[getSubValueByString] nil pointer while k:%s", k)
	}
	return t.elemType.getSubValueByString(v.Elem(), k)
}
func (t *ptrType) delete(v reflect.Value, k string) error {
	t.init()
	if v.IsNil() {
		return fmt.Errorf("[ptrType.delete] nil pointer while k:%s", k)
	}
	return t.elemType.delete(v.Elem(), k)
}
func (t *ptrType) create(v reflect.Value, k string) error {
	t.init()
	if v.IsNil() {
		v.Set(reflect.New(t.getReflectType().Elem()))
		return nil
	}
	return t.elemType.create(v.Elem(), k)
}
func (t *ptrType) save(v reflect.Value, value string) error {
	t.init()
	if v.IsNil() {
		return fmt.Errorf("[ptrType.save] nil pointer while value:%s", value)
	}
	return t.elemType.save(v.Elem(), value)
}
