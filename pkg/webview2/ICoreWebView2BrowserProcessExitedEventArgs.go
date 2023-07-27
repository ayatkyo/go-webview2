//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2BrowserProcessExitedEventArgsVtbl struct {
	IUnknownVtbl
	GetBrowserProcessExitKind ComProc
	GetBrowserProcessId       ComProc
}

type ICoreWebView2BrowserProcessExitedEventArgs struct {
	Vtbl *ICoreWebView2BrowserProcessExitedEventArgsVtbl
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) GetBrowserProcessExitKind() (COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND, error) {

	var browserProcessExitKind COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND

	hr, _, err := i.Vtbl.GetBrowserProcessExitKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&browserProcessExitKind)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND{}, syscall.Errno(hr)
	}
	return browserProcessExitKind, err
}

func (i *ICoreWebView2BrowserProcessExitedEventArgs) GetBrowserProcessId() (uint32, error) {

	var value uint32

	hr, _, err := i.Vtbl.GetBrowserProcessId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return 0, syscall.Errno(hr)
	}
	return value, err
}
