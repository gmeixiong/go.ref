// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build java android
// +build cgo

package main

import (
	"unsafe"
)

// #include "jni_wrapper.h"
import "C"

var (
	jVM *C.JavaVM
)

// JNI_OnLoad is called when System.loadLibrary is called. We need to cache the
// *JavaVM because that's the only way to get hold of a JNIEnv that is needed
// for any JNI operation.
//
// Reference: https://developer.android.com/training/articles/perf-jni.html#native_libraries
//
//export JNI_OnLoad
func JNI_OnLoad(vm *C.JavaVM, reserved unsafe.Pointer) C.jint {
	var env *C.JNIEnv
	if C.GetEnv(vm, unsafe.Pointer(&env), C.JNI_VERSION_1_6) != C.JNI_OK {
		return C.JNI_ERR
	}
	jVM = vm
	v23_syncbase_Init()
	return C.JNI_VERSION_1_6
}

func Java_io_v_syncbase_internal_Service_GetPermissions(env *C.JNIEnv, cls C.jclass) C.jobject {
	return nil
}
func Java_io_v_syncbase_internal_Service_SetPermissions(env *C.JNIEnv, cls C.jclass, obj C.jobject) {}
func Java_io_v_syncbase_internal_Service_ListDatabases(env *C.JNIEnv, cls C.jclass) C.jobject {
	return nil
}

func Java_io_v_syncbase_internal_Database_GetPermissions(env *C.JNIEnv, cls C.jclass, name C.jstring) C.jobject {
	return nil
}
func Java_io_v_syncbase_internal_Database_SetPermissions(env *C.JNIEnv, cls C.jclass, name C.jstring, perms C.jobject) {
}
func Java_io_v_syncbase_internal_Database_Create(env *C.JNIEnv, cls C.jclass, name C.jstring, perms C.jobject) {
}
func Java_io_v_syncbase_internal_Database_Destroy(env *C.JNIEnv, cls C.jclass, name C.jstring) {}
func Java_io_v_syncbase_internal_Database_Exists(env *C.JNIEnv, cls C.jclass, name C.jstring) C.jboolean {
	return 0
}
func Java_io_v_syncbase_internal_Database_BeginBatch(env *C.JNIEnv, cls C.jclass, name C.jstring, opts C.jobject) C.jstring {
	return nil
}
func Java_io_v_syncbase_internal_Database_ListCollections(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) C.jobject {
	return nil
}
func Java_io_v_syncbase_internal_Database_Commit(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) {
}
func Java_io_v_syncbase_internal_Database_Abort(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) {
}
func Java_io_v_syncbase_internal_Database_GetResumeMarker(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) C.jbyteArray {
	return nil
}
func Java_io_v_syncbase_internal_Database_ListSyncgroups(env *C.JNIEnv, cls C.jclass, name C.jstring) C.jobject {
	return nil
}
func Java_io_v_syncbase_internal_Database_CreateSyncgroup(env *C.JNIEnv, cls C.jclass, name C.jstring, sgId C.jobject, spec C.jobject, info C.jobject) {
}
func Java_io_v_syncbase_internal_Database_JoinSyncgroup(env *C.JNIEnv, cls C.jclass, name C.jstring, sgId C.jobject, info C.jobject) C.jobject {
	return nil
}
func Java_io_v_syncbase_internal_Database_LeaveSyncgroup(env *C.JNIEnv, cls C.jclass, name C.jstring, sgId C.jobject) {
}
func Java_io_v_syncbase_internal_Database_DestroySyncgroup(env *C.JNIEnv, cls C.jclass, name C.jstring, sgId C.jobject) {
}
func Java_io_v_syncbase_internal_Database_EjectFromSyncgroup(env *C.JNIEnv, cls C.jclass, name C.jstring, sgId C.jobject, member C.jstring) {
}
func Java_io_v_syncbase_internal_Database_GetSyncgroupSpec(env *C.JNIEnv, cls C.jclass, name C.jstring, sgId C.jobject) C.jobject {
	return nil
}
func Java_io_v_syncbase_internal_Database_SetSyncgroupSpec(env *C.JNIEnv, cls C.jclass, name C.jstring, sgId C.jobject, spec C.jobject) {
}
func Java_io_v_syncbase_internal_Database_GetSyncgroupMembers(env *C.JNIEnv, cls C.jclass, name C.jstring, sgId C.jobject) C.jobject {
	return nil
}

func Java_io_v_syncbase_internal_Collection_GetPermissions(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) C.jobject {
	return nil
}
func Java_io_v_syncbase_internal_Collection_SetPermissions(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring, perms C.jobject) {
}
func Java_io_v_syncbase_internal_Collection_Create(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring, perms C.jobject) {
}
func Java_io_v_syncbase_internal_Collection_Destroy(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) {
}
func Java_io_v_syncbase_internal_Collection_Exists(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) C.jboolean {
	return 0
}
func Java_io_v_syncbase_internal_Collection_DeleteRange(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring, start C.jbyteArray, limit C.jbyteArray) {
}
func Java_io_v_syncbase_internal_Collection_Scan(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring, start C.jbyteArray, limit C.jbyteArray, callbacks C.jobject) {
}

func Java_io_v_syncbase_internal_Row_Exists(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) C.jboolean {
	return 0
}
func Java_io_v_syncbase_internal_Row_Get(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) C.jbyteArray {
	return nil
}
func Java_io_v_syncbase_internal_Row_Put(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring, value C.jbyteArray) {
}
func Java_io_v_syncbase_internal_Row_Delete(env *C.JNIEnv, cls C.jclass, name C.jstring, handle C.jstring) {
}