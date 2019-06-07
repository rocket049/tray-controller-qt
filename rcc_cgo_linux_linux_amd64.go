package main

/*
#cgo CFLAGS: -pipe -O2 -Wall -W   -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -O2 -std=gnu++11 -Wall -W   -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../src -I. -I/opt/Qt/5.12.0/gcc_64/include -I/opt/Qt/5.12.0/gcc_64/include/QtWidgets -I/opt/Qt/5.12.0/gcc_64/include/QtGui -I/opt/Qt/5.12.0/gcc_64/include/QtCore -I. -isystem /usr/include/libdrm -I/opt/Qt/5.12.0/gcc_64/mkspecs/linux-g++
#cgo LDFLAGS: -O1 
#cgo LDFLAGS:  -L/opt/Qt/5.12.0/gcc_64/plugins/platforms -lqxcb -L/opt/Qt/5.12.0/gcc_64/lib -L/opt/Qt/5.12.0/gcc_64/plugins/xcbglintegrations -lqxcb-glx-integration -lQt5XcbQpa -lQt5ServiceSupport -lQt5ThemeSupport -lQt5FontDatabaseSupport /usr/lib/x86_64-linux-gnu/libfontconfig.so /usr/lib/x86_64-linux-gnu/libfreetype.so -lQt5GlxSupport -lQt5EdidSupport -lQt5DBus /usr/lib/x86_64-linux-gnu/libxcb-glx.so /usr/lib/x86_64-linux-gnu/libX11-xcb.so /usr/lib/x86_64-linux-gnu/libxcb.so /usr/lib/x86_64-linux-gnu/libXrender.so /usr/lib/x86_64-linux-gnu/libXext.so /usr/lib/x86_64-linux-gnu/libX11.so /usr/lib/x86_64-linux-gnu/libm.so /opt/Qt/5.12.0/gcc_64/lib/libxcb-static.a /usr/lib/x86_64-linux-gnu/libxkbcommon.so -L/opt/Qt/5.12.0/gcc_64/plugins/imageformats -lqgif -lqicns -lqico -lqjpeg -lqtga -lqtiff -lqwbmp -lqwebp -lQt5Widgets -lQt5Gui /opt/Qt/5.12.0/gcc_64/lib/libqtlibpng.a /opt/Qt/5.12.0/gcc_64/lib/libqtharfbuzz.a -lQt5Core -lm /opt/Qt/5.12.0/gcc_64/lib/libqtpcre2.a /usr/lib/x86_64-linux-gnu/libdl.so /usr/lib/x86_64-linux-gnu/libgthread-2.0.so /usr/lib/x86_64-linux-gnu/libglib-2.0.so /usr/lib/x86_64-linux-gnu/libGL.so -lpthread
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
