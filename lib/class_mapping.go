// Copyright 2011 Julian Phillips.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package py

// #include "utils.h"
import "C"

import (
	"os"
	"unsafe"
)

//export goClassMapLen
func goClassMapLen(obj unsafe.Pointer) C.Py_ssize_t {
	// Get the class context
	ctxt := getClassContext(obj)

	// Turn the function into something we can call
	f := (*func(unsafe.Pointer) int64)(unsafe.Pointer(&ctxt.mp_len))

	return C.Py_ssize_t((*f)(obj))
}

//export goClassMapGet
func goClassMapGet(obj, arg unsafe.Pointer) unsafe.Pointer {
	// Get the class context
	ctxt := getClassContext(obj)

	// Turn the function into something we can call
	f := (*func(unsafe.Pointer, Object) (Object, os.Error))(unsafe.Pointer(&ctxt.mp_get))

	key := newBaseObject((*C.PyObject)(arg)).actual()

	ret, err := (*f)(obj, key)
	if err != nil {
		raise(err)
		return nil
	}

	return unsafe.Pointer(c(ret))
}

//export goClassMapSet
func goClassMapSet(obj, arg1, arg2 unsafe.Pointer) int {
	// Get the class context
	ctxt := getClassContext(obj)

	// Turn the function into something we can call
	f := (*func(unsafe.Pointer, Object, Object) os.Error)(unsafe.Pointer(&ctxt.mp_set))

	key := newBaseObject((*C.PyObject)(arg1)).actual()
	value := newBaseObject((*C.PyObject)(arg2)).actual()

	err := (*f)(obj, key, value)
	if err != nil {
		raise(err)
		return -1
	}

	return 0
}