package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := NewMyArray()
	fmt.Printf("Initial value of arry %v \n", arr)

	arr.push("Item One")
	arr.push("Item Two")

	fmt.Printf("After two item add %v \n", arr)

	itemTwo := arr.pop()

	fmt.Printf("After pop value of array %v and poped item %v \n", arr, itemTwo)
}

type MyArray struct {
	data   map[int]interface{}
	length int64
}

func NewMyArray() *MyArray {
	arr := new(MyArray)
	arr.data = make(map[int]interface{})
	arr.length = 0
	return arr
}

func (ar *MyArray) push(item interface{}) int {
	ar.data[int(ar.length)] = item
	ar.length++
	return int(ar.length)
}

func (ar *MyArray) pop() (popItem interface{}) {
	key := int(ar.length - 1)
	popItem = ar.data[key]
	delete(ar.data, key)
	ar.length--
	return
}

func (ar *MyArray) get(key int) (interface{}, error) {
	popItem, exist := ar.data[key]
	if !exist {
		return nil, errors.New(fmt.Sprintf("Provided key %d not exist", key))
	}
	return popItem, nil
}

func (ar *MyArray) shift(index int) {
	for i := index; i < int(ar.length)-1; i++ {
		ar.data[i] = ar.data[i+1]
	}
	delete(ar.data, int(ar.length)-1)
	ar.length--
}

func (ar *MyArray) unshift(index int, item interface{}) {
	for i := int(ar.length); i > index; i-- {
		ar.data[i] = ar.data[i-1]
	}

	ar.data[index] = item
	ar.length++
}

func (ar *MyArray) getIndex(item interface{}) int {
	for key, val := range ar.data {
		if val == item {
			return key
		}
	}
	return -1
}

func (ar *MyArray) deleteFromIndex(index int) (interface{}, error) {
	itm, exist := ar.data[index]
	if !exist {
		return nil, errors.New(fmt.Sprintf("Cant delete index %d cause it does not exist", index))
	}
	ar.shift(index)
	return itm, nil
}
