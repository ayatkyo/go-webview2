//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2_19Vtbl struct {
	IUnknownVtbl
	GetMemoryUsageTargetLevel ComProc
	PutMemoryUsageTargetLevel ComProc
}

type ICoreWebView2_19 struct {
	Vtbl *ICoreWebView2_19Vtbl
}

func (i *ICoreWebView2_19) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2_19() *ICoreWebView2_19 {
	var result *ICoreWebView2_19

	iidICoreWebView2_19 := NewGUID("{6921F954-79B0-437F-A997-C85811897C68}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_19)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2_19) GetMemoryUsageTargetLevel() (COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL, error) {

	var level COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL

	hr, _, err := i.Vtbl.GetMemoryUsageTargetLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&level)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL{}, syscall.Errno(hr)
	}
	return level, err
}

func (i *ICoreWebView2_19) PutMemoryUsageTargetLevel(level COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL) error {

	hr, _, err := i.Vtbl.PutMemoryUsageTargetLevel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(level),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
