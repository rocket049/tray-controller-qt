

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class QtBridgea35b5a: public QObject
{
Q_OBJECT
public:
	QtBridgea35b5a(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaType();QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaTypes();callbackQtBridgea35b5a_Constructor(this);};
	void Signal_UpdateRun() { callbackQtBridgea35b5a_UpdateRun(this); };
	void Signal_UpdateStop() { callbackQtBridgea35b5a_UpdateStop(this); };
	 ~QtBridgea35b5a() { callbackQtBridgea35b5a_DestroyQtBridge(this); };
	bool event(QEvent * e) { return callbackQtBridgea35b5a_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQtBridgea35b5a_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQtBridgea35b5a_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQtBridgea35b5a_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQtBridgea35b5a_CustomEvent(this, event); };
	void deleteLater() { callbackQtBridgea35b5a_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQtBridgea35b5a_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQtBridgea35b5a_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQtBridgea35b5a_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQtBridgea35b5a_TimerEvent(this, event); };
signals:
	void updateRun();
	void updateStop();
public slots:
private:
};

Q_DECLARE_METATYPE(QtBridgea35b5a*)


void QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaTypes() {
}

void QtBridgea35b5a_ConnectUpdateRun(void* ptr)
{
	QObject::connect(static_cast<QtBridgea35b5a*>(ptr), static_cast<void (QtBridgea35b5a::*)()>(&QtBridgea35b5a::updateRun), static_cast<QtBridgea35b5a*>(ptr), static_cast<void (QtBridgea35b5a::*)()>(&QtBridgea35b5a::Signal_UpdateRun));
}

void QtBridgea35b5a_DisconnectUpdateRun(void* ptr)
{
	QObject::disconnect(static_cast<QtBridgea35b5a*>(ptr), static_cast<void (QtBridgea35b5a::*)()>(&QtBridgea35b5a::updateRun), static_cast<QtBridgea35b5a*>(ptr), static_cast<void (QtBridgea35b5a::*)()>(&QtBridgea35b5a::Signal_UpdateRun));
}

void QtBridgea35b5a_UpdateRun(void* ptr)
{
	static_cast<QtBridgea35b5a*>(ptr)->updateRun();
}

void QtBridgea35b5a_ConnectUpdateStop(void* ptr)
{
	QObject::connect(static_cast<QtBridgea35b5a*>(ptr), static_cast<void (QtBridgea35b5a::*)()>(&QtBridgea35b5a::updateStop), static_cast<QtBridgea35b5a*>(ptr), static_cast<void (QtBridgea35b5a::*)()>(&QtBridgea35b5a::Signal_UpdateStop));
}

void QtBridgea35b5a_DisconnectUpdateStop(void* ptr)
{
	QObject::disconnect(static_cast<QtBridgea35b5a*>(ptr), static_cast<void (QtBridgea35b5a::*)()>(&QtBridgea35b5a::updateStop), static_cast<QtBridgea35b5a*>(ptr), static_cast<void (QtBridgea35b5a::*)()>(&QtBridgea35b5a::Signal_UpdateStop));
}

void QtBridgea35b5a_UpdateStop(void* ptr)
{
	static_cast<QtBridgea35b5a*>(ptr)->updateStop();
}

int QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaType()
{
	return qRegisterMetaType<QtBridgea35b5a*>();
}

int QtBridgea35b5a_QtBridgea35b5a_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QtBridgea35b5a*>(const_cast<const char*>(typeName));
}

int QtBridgea35b5a_QtBridgea35b5a_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QtBridgea35b5a>();
#else
	return 0;
#endif
}

int QtBridgea35b5a_QtBridgea35b5a_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QtBridgea35b5a>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QtBridgea35b5a___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QtBridgea35b5a___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QtBridgea35b5a___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QtBridgea35b5a___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QtBridgea35b5a___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QtBridgea35b5a___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QtBridgea35b5a___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QtBridgea35b5a___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QtBridgea35b5a___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QtBridgea35b5a___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QtBridgea35b5a___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QtBridgea35b5a___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QtBridgea35b5a___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QtBridgea35b5a___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QtBridgea35b5a___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QtBridgea35b5a_NewQtBridge(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QtBridgea35b5a(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QtBridgea35b5a(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QtBridgea35b5a(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QtBridgea35b5a(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QtBridgea35b5a(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QtBridgea35b5a(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QtBridgea35b5a(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QtBridgea35b5a(static_cast<QWindow*>(parent));
	} else {
		return new QtBridgea35b5a(static_cast<QObject*>(parent));
	}
}

void QtBridgea35b5a_DestroyQtBridge(void* ptr)
{
	static_cast<QtBridgea35b5a*>(ptr)->~QtBridgea35b5a();
}

void QtBridgea35b5a_DestroyQtBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QtBridgea35b5a_EventDefault(void* ptr, void* e)
{
	return static_cast<QtBridgea35b5a*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QtBridgea35b5a_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QtBridgea35b5a*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QtBridgea35b5a_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QtBridgea35b5a*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QtBridgea35b5a_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QtBridgea35b5a*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QtBridgea35b5a_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QtBridgea35b5a*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QtBridgea35b5a_DeleteLaterDefault(void* ptr)
{
	static_cast<QtBridgea35b5a*>(ptr)->QObject::deleteLater();
}

void QtBridgea35b5a_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QtBridgea35b5a*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QtBridgea35b5a_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QtBridgea35b5a*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
