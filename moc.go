package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QtBridge_ITF interface {
	std_core.QObject_ITF
	QtBridge_PTR() *QtBridge
}

func (ptr *QtBridge) QtBridge_PTR() *QtBridge {
	return ptr
}

func (ptr *QtBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QtBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQtBridge(ptr QtBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtBridge_PTR().Pointer()
	}
	return nil
}

func NewQtBridgeFromPointer(ptr unsafe.Pointer) (n *QtBridge) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QtBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QtBridge:
			n = deduced

		case *std_core.QObject:
			n = &QtBridge{QObject: *deduced}

		default:
			n = new(QtBridge)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQtBridgea35b5a_Constructor
func callbackQtBridgea35b5a_Constructor(ptr unsafe.Pointer) {
	this := NewQtBridgeFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQtBridgea35b5a_UpdateRun
func callbackQtBridgea35b5a_UpdateRun(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateRun"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QtBridge) ConnectUpdateRun(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateRun") {
			C.QtBridgea35b5a_ConnectUpdateRun(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateRun"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "updateRun", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateRun", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QtBridge) DisconnectUpdateRun() {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_DisconnectUpdateRun(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateRun")
	}
}

func (ptr *QtBridge) UpdateRun() {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_UpdateRun(ptr.Pointer())
	}
}

//export callbackQtBridgea35b5a_UpdateStop
func callbackQtBridgea35b5a_UpdateStop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateStop"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QtBridge) ConnectUpdateStop(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateStop") {
			C.QtBridgea35b5a_ConnectUpdateStop(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateStop"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "updateStop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateStop", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QtBridge) DisconnectUpdateStop() {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_DisconnectUpdateStop(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateStop")
	}
}

func (ptr *QtBridge) UpdateStop() {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_UpdateStop(ptr.Pointer())
	}
}

func QtBridge_QRegisterMetaType() int {
	return int(int32(C.QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaType()))
}

func (ptr *QtBridge) QRegisterMetaType() int {
	return int(int32(C.QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaType()))
}

func QtBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaType2(typeNameC)))
}

func (ptr *QtBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaType2(typeNameC)))
}

func QtBridge_QmlRegisterType() int {
	return int(int32(C.QtBridgea35b5a_QtBridgea35b5a_QmlRegisterType()))
}

func (ptr *QtBridge) QmlRegisterType() int {
	return int(int32(C.QtBridgea35b5a_QtBridgea35b5a_QmlRegisterType()))
}

func QtBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QtBridgea35b5a_QtBridgea35b5a_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QtBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QtBridgea35b5a_QtBridgea35b5a_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QtBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QtBridgea35b5a___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QtBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QtBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QtBridgea35b5a___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QtBridge) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QtBridgea35b5a___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QtBridge) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QtBridge) __findChildren_newList2() unsafe.Pointer {
	return C.QtBridgea35b5a___findChildren_newList2(ptr.Pointer())
}

func (ptr *QtBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QtBridgea35b5a___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QtBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QtBridge) __findChildren_newList3() unsafe.Pointer {
	return C.QtBridgea35b5a___findChildren_newList3(ptr.Pointer())
}

func (ptr *QtBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QtBridgea35b5a___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QtBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QtBridge) __findChildren_newList() unsafe.Pointer {
	return C.QtBridgea35b5a___findChildren_newList(ptr.Pointer())
}

func (ptr *QtBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QtBridgea35b5a___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QtBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QtBridge) __children_newList() unsafe.Pointer {
	return C.QtBridgea35b5a___children_newList(ptr.Pointer())
}

func NewQtBridge(parent std_core.QObject_ITF) *QtBridge {
	tmpValue := NewQtBridgeFromPointer(C.QtBridgea35b5a_NewQtBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQtBridgea35b5a_DestroyQtBridge
func callbackQtBridgea35b5a_DestroyQtBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QtBridge"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQtBridgeFromPointer(ptr).DestroyQtBridgeDefault()
	}
}

func (ptr *QtBridge) ConnectDestroyQtBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QtBridge"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QtBridge", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QtBridge", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QtBridge) DisconnectDestroyQtBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QtBridge")
	}
}

func (ptr *QtBridge) DestroyQtBridge() {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_DestroyQtBridge(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QtBridge) DestroyQtBridgeDefault() {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_DestroyQtBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQtBridgea35b5a_Event
func callbackQtBridgea35b5a_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQtBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QtBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QtBridgea35b5a_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQtBridgea35b5a_EventFilter
func callbackQtBridgea35b5a_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQtBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QtBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QtBridgea35b5a_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQtBridgea35b5a_ChildEvent
func callbackQtBridgea35b5a_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQtBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QtBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQtBridgea35b5a_ConnectNotify
func callbackQtBridgea35b5a_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQtBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QtBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQtBridgea35b5a_CustomEvent
func callbackQtBridgea35b5a_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQtBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QtBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQtBridgea35b5a_DeleteLater
func callbackQtBridgea35b5a_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQtBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QtBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQtBridgea35b5a_Destroyed
func callbackQtBridgea35b5a_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackQtBridgea35b5a_DisconnectNotify
func callbackQtBridgea35b5a_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQtBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QtBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQtBridgea35b5a_ObjectNameChanged
func callbackQtBridgea35b5a_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQtBridgea35b5a_TimerEvent
func callbackQtBridgea35b5a_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQtBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QtBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QtBridgea35b5a_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
