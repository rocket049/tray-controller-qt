

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class QtBridgedd03f6: public QObject
{
Q_OBJECT
public:
	QtBridgedd03f6(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QtBridgedd03f6_QtBridgedd03f6_QRegisterMetaType();QtBridgedd03f6_QtBridgedd03f6_QRegisterMetaTypes();callbackQtBridgedd03f6_Constructor(this);};
	void Signal_UpdateRun() { callbackQtBridgedd03f6_UpdateRun(this); };
	void Signal_UpdateStop() { callbackQtBridgedd03f6_UpdateStop(this); };
	 ~QtBridgedd03f6() { callbackQtBridgedd03f6_DestroyQtBridge(this); };
	bool event(QEvent * e) { return callbackQtBridgedd03f6_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQtBridgedd03f6_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQtBridgedd03f6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQtBridgedd03f6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQtBridgedd03f6_CustomEvent(this, event); };
	void deleteLater() { callbackQtBridgedd03f6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQtBridgedd03f6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQtBridgedd03f6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQtBridgedd03f6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQtBridgedd03f6_TimerEvent(this, event); };
signals:
	void updateRun();
	void updateStop();
public slots:
private:
};

Q_DECLARE_METATYPE(QtBridgedd03f6*)


void QtBridgedd03f6_QtBridgedd03f6_QRegisterMetaTypes() {
}

void QtBridgedd03f6_ConnectUpdateRun(void* ptr)
{
	QObject::connect(static_cast<QtBridgedd03f6*>(ptr), static_cast<void (QtBridgedd03f6::*)()>(&QtBridgedd03f6::updateRun), static_cast<QtBridgedd03f6*>(ptr), static_cast<void (QtBridgedd03f6::*)()>(&QtBridgedd03f6::Signal_UpdateRun));
}

void QtBridgedd03f6_DisconnectUpdateRun(void* ptr)
{
	QObject::disconnect(static_cast<QtBridgedd03f6*>(ptr), static_cast<void (QtBridgedd03f6::*)()>(&QtBridgedd03f6::updateRun), static_cast<QtBridgedd03f6*>(ptr), static_cast<void (QtBridgedd03f6::*)()>(&QtBridgedd03f6::Signal_UpdateRun));
}

void QtBridgedd03f6_UpdateRun(void* ptr)
{
	static_cast<QtBridgedd03f6*>(ptr)->updateRun();
}

void QtBridgedd03f6_ConnectUpdateStop(void* ptr)
{
	QObject::connect(static_cast<QtBridgedd03f6*>(ptr), static_cast<void (QtBridgedd03f6::*)()>(&QtBridgedd03f6::updateStop), static_cast<QtBridgedd03f6*>(ptr), static_cast<void (QtBridgedd03f6::*)()>(&QtBridgedd03f6::Signal_UpdateStop));
}

void QtBridgedd03f6_DisconnectUpdateStop(void* ptr)
{
	QObject::disconnect(static_cast<QtBridgedd03f6*>(ptr), static_cast<void (QtBridgedd03f6::*)()>(&QtBridgedd03f6::updateStop), static_cast<QtBridgedd03f6*>(ptr), static_cast<void (QtBridgedd03f6::*)()>(&QtBridgedd03f6::Signal_UpdateStop));
}

void QtBridgedd03f6_UpdateStop(void* ptr)
{
	static_cast<QtBridgedd03f6*>(ptr)->updateStop();
}

int QtBridgedd03f6_QtBridgedd03f6_QRegisterMetaType()
{
	return qRegisterMetaType<QtBridgedd03f6*>();
}

int QtBridgedd03f6_QtBridgedd03f6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QtBridgedd03f6*>(const_cast<const char*>(typeName));
}

int QtBridgedd03f6_QtBridgedd03f6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QtBridgedd03f6>();
#else
	return 0;
#endif
}

int QtBridgedd03f6_QtBridgedd03f6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QtBridgedd03f6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QtBridgedd03f6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QtBridgedd03f6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QtBridgedd03f6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QtBridgedd03f6___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QtBridgedd03f6___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QtBridgedd03f6___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QtBridgedd03f6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QtBridgedd03f6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QtBridgedd03f6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QtBridgedd03f6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QtBridgedd03f6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QtBridgedd03f6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QtBridgedd03f6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QtBridgedd03f6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QtBridgedd03f6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QtBridgedd03f6_NewQtBridge(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QtBridgedd03f6(static_cast<QWindow*>(parent));
	} else {
		return new QtBridgedd03f6(static_cast<QObject*>(parent));
	}
}

void QtBridgedd03f6_DestroyQtBridge(void* ptr)
{
	static_cast<QtBridgedd03f6*>(ptr)->~QtBridgedd03f6();
}

void QtBridgedd03f6_DestroyQtBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QtBridgedd03f6_EventDefault(void* ptr, void* e)
{
	return static_cast<QtBridgedd03f6*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QtBridgedd03f6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QtBridgedd03f6*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QtBridgedd03f6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QtBridgedd03f6*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QtBridgedd03f6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QtBridgedd03f6*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QtBridgedd03f6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QtBridgedd03f6*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QtBridgedd03f6_DeleteLaterDefault(void* ptr)
{
	static_cast<QtBridgedd03f6*>(ptr)->QObject::deleteLater();
}

void QtBridgedd03f6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QtBridgedd03f6*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QtBridgedd03f6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QtBridgedd03f6*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
