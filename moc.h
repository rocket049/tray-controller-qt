

#pragma once

#ifndef GO_MOC_dd03f6_H
#define GO_MOC_dd03f6_H

#include <stdint.h>

#ifdef __cplusplus
class QtBridgedd03f6;
void QtBridgedd03f6_QtBridgedd03f6_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void QtBridgedd03f6_ConnectUpdateRun(void* ptr);
void QtBridgedd03f6_DisconnectUpdateRun(void* ptr);
void QtBridgedd03f6_UpdateRun(void* ptr);
void QtBridgedd03f6_ConnectUpdateStop(void* ptr);
void QtBridgedd03f6_DisconnectUpdateStop(void* ptr);
void QtBridgedd03f6_UpdateStop(void* ptr);
int QtBridgedd03f6_QtBridgedd03f6_QRegisterMetaType();
int QtBridgedd03f6_QtBridgedd03f6_QRegisterMetaType2(char* typeName);
int QtBridgedd03f6_QtBridgedd03f6_QmlRegisterType();
int QtBridgedd03f6_QtBridgedd03f6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QtBridgedd03f6___dynamicPropertyNames_atList(void* ptr, int i);
void QtBridgedd03f6___dynamicPropertyNames_setList(void* ptr, void* i);
void* QtBridgedd03f6___dynamicPropertyNames_newList(void* ptr);
void* QtBridgedd03f6___findChildren_atList2(void* ptr, int i);
void QtBridgedd03f6___findChildren_setList2(void* ptr, void* i);
void* QtBridgedd03f6___findChildren_newList2(void* ptr);
void* QtBridgedd03f6___findChildren_atList3(void* ptr, int i);
void QtBridgedd03f6___findChildren_setList3(void* ptr, void* i);
void* QtBridgedd03f6___findChildren_newList3(void* ptr);
void* QtBridgedd03f6___findChildren_atList(void* ptr, int i);
void QtBridgedd03f6___findChildren_setList(void* ptr, void* i);
void* QtBridgedd03f6___findChildren_newList(void* ptr);
void* QtBridgedd03f6___children_atList(void* ptr, int i);
void QtBridgedd03f6___children_setList(void* ptr, void* i);
void* QtBridgedd03f6___children_newList(void* ptr);
void* QtBridgedd03f6_NewQtBridge(void* parent);
void QtBridgedd03f6_DestroyQtBridge(void* ptr);
void QtBridgedd03f6_DestroyQtBridgeDefault(void* ptr);
char QtBridgedd03f6_EventDefault(void* ptr, void* e);
char QtBridgedd03f6_EventFilterDefault(void* ptr, void* watched, void* event);
void QtBridgedd03f6_ChildEventDefault(void* ptr, void* event);
void QtBridgedd03f6_ConnectNotifyDefault(void* ptr, void* sign);
void QtBridgedd03f6_CustomEventDefault(void* ptr, void* event);
void QtBridgedd03f6_DeleteLaterDefault(void* ptr);
void QtBridgedd03f6_DisconnectNotifyDefault(void* ptr, void* sign);
void QtBridgedd03f6_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif