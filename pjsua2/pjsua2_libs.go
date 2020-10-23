// +build !shared

package pjsua2

/*
#cgo CPPFLAGS: -DPJ_AUTOCONF=1 -O2 -DPJ_IS_BIG_ENDIAN=0
#cgo CPPFLAGS: -DPJ_IS_LITTLE_ENDIAN=1
#cgo CPPFLAGS: -DPJMEDIA_USE_OLD_FFMPEG=1
#cgo CPPFLAGS: -I/usr/local/include -I/usr/include
#cgo CPPFLAGS: -I/usr/include/c++/5 -I/usr/include/x86_64-linux-gnu/c++/5
#cgo CPPFLAGS: -I/root/pjproject/pjlib/include
#cgo CPPFLAGS: -I/root/pjproject/pjlib-util/include
#cgo CPPFLAGS: -I/root/pjproject/pjnath/include
#cgo CPPFLAGS: -I/root/pjproject/pjmedia/include
#cgo CPPFLAGS: -I/root/pjproject/pjsip/include
#cgo LDFLAGS: -L/usr/local/lib
#cgo LDFLAGS: -L/root/pjproject/pjlib/lib
#cgo LDFLAGS: -L/root/pjproject/pjlib-util/lib
#cgo LDFLAGS: -L/root/pjproject/pjnath/lib
#cgo LDFLAGS: -L/root/pjproject/pjmedia/lib
#cgo LDFLAGS: -L/root/pjproject/pjsip/lib
#cgo LDFLAGS: -L/root/pjproject/third_party/lib
#cgo LDFLAGS: -lm
#cgo LDFLAGS: -lpjsua2-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjsua-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjsip-ua-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjsip-simple-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjsip-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjmedia-codec-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjmedia-videodev-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjmedia-audiodev-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjmedia-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjnath-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpjlib-util-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lsrtp-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lresample-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lgsmcodec-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lspeex-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lilbccodec-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lg7221codec-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lyuv-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lwebrtc-x86_64-unknown-linux-gnu
#cgo LDFLAGS: -lpj-x86_64-unknown-linux-gnu
*/
import "C"
