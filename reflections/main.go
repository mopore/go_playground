package main

import (
	"log"
	"reflect"
)

func main(){
    // if you don't have a string variable, you can use this
    var stringType = reflect.TypeOf((*string)(nil)).Elem()
    var stringSliceType = reflect.TypeOf([]string(nil))

    stringSliceValue := reflect.MakeSlice(stringSliceType, 0, 10)

    stringValue := reflect.New(stringType).Elem()
    stringValue.SetString("Hello World")

    stringSliceValue = reflect.Append(stringSliceValue, stringValue)
    strings := stringSliceValue.Interface().([]string)

    log.Printf("strings: %v", strings)
}
