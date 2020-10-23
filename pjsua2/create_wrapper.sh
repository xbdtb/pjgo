#!/bin/bash

swig -c++ -go -cgo -intgosize 64 \
	-outcurrentdir \
	-I/usr/local/include \
	-I/root/pjproject/pjsip/include \
	/root/pjproject/pjsip-apps/src/swig/pjsua2.i

# modify path above to match your installation.
