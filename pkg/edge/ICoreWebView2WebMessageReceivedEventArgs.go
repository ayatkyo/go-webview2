//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2WebMessageReceivedEventArgsVtbl struct {
	_IUnknownVtbl
	GetSource                ComProc
	GetWebMessageAsJSON      ComProc
	TryGetWebMessageAsString ComProc
	GetAdditionalObjects     ComProc
}

type iCoreWebView2WebMessageReceivedEventArgs struct {
	vtbl *iCoreWebView2WebMessageReceivedEventArgsVtbl
}

func (i *iCoreWebView2WebMessageReceivedEventArgs) TryGetWebMessageAsString() (string, error) {
	var err error
	var _message *uint16

	_, _, err = i.vtbl.TryGetWebMessageAsString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_message)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	}

	message := windows.UTF16PtrToString(_message)
	windows.CoTaskMemFree(unsafe.Pointer(_message))

	return message, nil
}

func (i *iCoreWebView2WebMessageReceivedEventArgs) GetAdditionalObjects() (*ICoreWebView2ObjectCollectionView, error) {
	var err error
	var value *ICoreWebView2ObjectCollectionView
	_, _, err = i.vtbl.GetAdditionalObjects.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}
