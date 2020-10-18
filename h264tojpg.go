package h264tojpg

/*
#include "h264tojpg.h"
#cgo LDFLAGS: -lavcodec -lavformat -lswscale -lavutil -lavfilter -lm
*/
import "C"

func H264tojpg(content []byte, savePath string) {
	C.h264_to_jpg((*C.uint8_t)(C.CBytes(content)), C.size_t(len(content)), C.CString(savePath))
}
