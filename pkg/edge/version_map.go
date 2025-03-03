//go:build windows

package edge

type Version struct {
	SDKVersion         string
	ReleaseNotes           string
	RuntimeVersion string
	Notes          string
}

var versionMapping = map[string]Version{
	"1.0.1901.177": {
		SDKVersion:     "1.0.1901.177",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101901177",
		RuntimeVersion: "115.0.1901.177",
		Notes: 			`
<!-- ------------------------------ -->
#### General


<!-- ------------------------------ -->
###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
* 'NavigationKind' gets the navigation kind of each navigation, such as Back/Forward, Reload, or navigation to a new document. 

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2NavigationStartingEventArgs' Class:
   * [CoreWebView2NavigationStartingEventArgs.NavigationKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2navigationstartingeventargs.navigationkind?view=webview2-dotnet-1.0.1901.177&preserve-view=true)
* [CoreWebView2NavigationKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2navigationkind?view=webview2-dotnet-1.0.1901.177&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2NavigationStartingEventArgs' Class:
   * [CoreWebView2NavigationStartingEventArgs.NavigationKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2navigationstartingeventargs?view=webview2-winrt-1.0.1901.177&preserve-view=true#navigationkind)
* [CoreWebView2NavigationKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2navigationkind?view=webview2-winrt-1.0.1901.177&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2NavigationStartingEventArgs3 interface](/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs3?view=webview2-1.0.1901.177&preserve-view=true)
    * [ICoreWebView2NavigationStartingEventArgs3::get_NavigationKind property](/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs3?view=webview2-1.0.1901.177&preserve-view=true#get_navigationkind)<!--no put-->
* [COREWEBVIEW2_NAVIGATION_KIND enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1901.177&preserve-view=true#corewebview2_navigation_kind)

---


<!-- ------------------------------ -->
* The 'ServiceWorkers' enum value in the 'BrowsingDataKinds' enum specifies service workers that are registered for an origin.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2BrowsingDataKinds' Enum:
   * [CoreWebView2BrowsingDataKinds.ServiceWorkers Enum Value](/dotnet/api/microsoft.web.webview2.core.corewebview2browsingdatakinds?view=webview2-dotnet-1.0.1901.177&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2BrowsingDataKinds' Enum:
   * [CoreWebView2BrowsingDataKinds.ServiceWorkers Enum Value](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2browsingdatakinds?view=webview2-winrt-1.0.1901.177&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* 'COREWEBVIEW2_BROWSING_DATA_KINDS' enum:
   * [COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS enum value](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1901.177&preserve-view=true#corewebview2_browsing_data_kinds)

---


<!-- ------------------------------ -->
###### Bug fixes

* Fixed a bug where the entire toolbar is blank when hiding the Bookmarks, Search, and PageSelector buttons simultaneously (Runtime). [Issue 2866](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2866)

<!-- end of ## 1.0.1901.177 July -->


<!-- ====================================================================== -->
Release Date: July 24, 2023`,
	},
	"1.0.1988-prerelease": {
		SDKVersion:     "1.0.1988-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101988-prerelease",
		RuntimeVersion: "117.0.1988.0",
		Notes: 			`
#### General

<!-- ------------------------------ -->
###### Experimental features


<!-- ------------------------------ -->
* Supports desktop notifications through WebView2:

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.NotificationReceived Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.notificationreceived?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
* [CoreWebView2Notification Class](/dotnet/api/microsoft.web.webview2.core.corewebview2notification?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
* [CoreWebView2NotificationReceivedEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2notificationreceivedeventargs?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.NotificationReceived Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#notificationreceived)
* [CoreWebView2Notification Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2notification?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true)
* [CoreWebView2NotificationReceivedEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2notificationreceivedeventargs?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental22 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimental22?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2Experimental22::add_NotificationReceived event](/microsoft-edge/webview2/reference/win32/icorewebview2experimental22?view=webview2-1.0.1988-prerelease&preserve-view=true#add_notificationreceived)
   * [ICoreWebView2Experimental22::remove_NotificationReceived event](/microsoft-edge/webview2/reference/win32/icorewebview2experimental22?view=webview2-1.0.1988-prerelease&preserve-view=true#remove_notificationreceived)
* [ICoreWebView2ExperimentalNotification interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalnotification?view=webview2-1.0.1988-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalNotificationCloseRequestedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalnotificationcloserequestedeventhandler?view=webview2-1.0.1988-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalNotificationReceivedEventArgs interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalnotificationreceivedeventargs?view=webview2-1.0.1988-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalNotificationReceivedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalnotificationreceivedeventhandler?view=webview2-1.0.1988-prerelease&preserve-view=true)

---


<!-- ------------------------------ -->
* Supports monitoring iframe's runtime memory usage by getting process info details of iframes.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.FrameId Property](/dotnet/api/microsoft.web.webview2.core.corewebview2.frameid?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
* 'CoreWebView2Environment' Class:
   * [CoreWebView2Environment.GetProcessInfosWithDetailsAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.getprocessinfoswithdetailsasync?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
* 'CoreWebView2Frame' Class:
   * [CoreWebView2Frame.FrameId Property](/dotnet/api/microsoft.web.webview2.core.corewebview2frame.frameid?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
* 'CoreWebView2FrameInfo' Class:
   * [CoreWebView2FrameInfo.FrameId Property](/dotnet/api/microsoft.web.webview2.core.corewebview2frameinfo.frameid?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
   * [CoreWebView2FrameInfo.FrameKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2frameinfo.framekind?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
   * [CoreWebView2FrameInfo.ParentFrameInfo Property](/dotnet/api/microsoft.web.webview2.core.corewebview2frameinfo.parentframeinfo?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
* [CoreWebView2FrameKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2framekind?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
   * 'Iframe'
   * 'MainFrame'
   * 'Other'
* 'CoreWebView2ProcessInfo' Class:
   * [CoreWebView2ProcessInfo.AssociatedFrameInfos Property](/dotnet/api/microsoft.web.webview2.core.corewebview2processinfo.associatedframeinfos?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.FrameId Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#frameid)
* 'CoreWebView2Environment' Class:
   * [CoreWebView2Environment.GetProcessInfosWithDetailsAsync Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#getprocessinfoswithdetailsasync)
* 'CoreWebView2Frame' Class:
   * [CoreWebView2Frame.FrameId Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frame?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#frameid)
* 'CoreWebView2FrameInfo' Class:
   * [CoreWebView2FrameInfo.FrameId Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frameinfo?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#frameid)
   * [CoreWebView2FrameInfo.FrameKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frameinfo?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#framekind)
   * [CoreWebView2FrameInfo.ParentFrameInfo Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frameinfo?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#parentframeinfo)
* [CoreWebView2FrameKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2framekind?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true)
   * 'Other'
   * 'MainFrame'
   * 'Iframe'
* 'CoreWebView2ProcessInfo' Class:
   * [CoreWebView2ProcessInfo.AssociatedFrameInfos Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2processinfo?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#associatedframeinfos)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental23 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimental23?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2Experimental23::get_FrameId property](/microsoft-edge/webview2/reference/win32/icorewebview2experimental23?view=webview2-1.0.1988-prerelease&preserve-view=true#get_frameid)<!--no put-->
* [ICoreWebView2ExperimentalEnvironment11 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment11?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalEnvironment11::GetProcessInfosWithDetails method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment11?view=webview2-1.0.1988-prerelease&preserve-view=true#getprocessinfoswithdetails)
* [ICoreWebView2ExperimentalFrame5 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe5?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalFrame5::get_FrameId property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe5?view=webview2-1.0.1988-prerelease&preserve-view=true#get_frameid)<!--no put-->
* [ICoreWebView2ExperimentalFrameInfo interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframeinfo?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2FrameInfo::get_FrameId property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframeinfo?view=webview2-1.0.1988-prerelease&preserve-view=true#get_frameid)<!--no put-->
   * [ICoreWebView2FrameInfo::get_FrameKind property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframeinfo?view=webview2-1.0.1988-prerelease&preserve-view=true#get_framekind)<!--no put-->
   * [ICoreWebView2FrameInfo::get_ParentFrameInfo property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframeinfo?view=webview2-1.0.1988-prerelease&preserve-view=true#get_parentframeinfo)<!--no put-->
* [COREWEBVIEW2_FRAME_KIND enum](/microsoft-edge/webview2/reference/win32/webview2experimental-idl?view=webview2-1.0.1988-prerelease&preserve-view=true#corewebview2_frame_kind)
   * 'COREWEBVIEW2_FRAME_KIND_OTHER'
   * 'COREWEBVIEW2_FRAME_KIND_MAIN_FRAME'
   * 'COREWEBVIEW2_FRAME_KIND_IFRAME'
* [ICoreWebView2ExperimentalGetProcessInfosWithDetailsCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalgetprocessinfoswithdetailscompletedhandler?view=webview2-1.0.1988-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalProcessInfo interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprocessinfo?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalProcessInfo::get_AssociatedFrameInfos property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprocessinfo?view=webview2-1.0.1988-prerelease&preserve-view=true#get_associatedframeinfos)<!--no put-->

---


<!-- ------------------------------ -->
* Supports extensions in WebView2. 

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2BrowserExtension Class](/dotnet/api/microsoft.web.webview2.core.corewebview2browserextension?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.AreBrowserExtensionsEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.arebrowserextensionsenabled?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
* 'CoreWebView2Profile' Class:
   * [CoreWebView2Profile.AddBrowserExtensionAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.addbrowserextensionasync?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
   * [CoreWebView2Profile.GetBrowserExtensionsAsync method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.getbrowserextensionsasync?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2BrowserExtension Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2browserextension?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true)
* 'CoreWebView2EnvironmentOptions' Class:
   * [CoreWebView2EnvironmentOptions.AreBrowserExtensionsEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#arebrowserextensionsenabled)
* 'CoreWebView2Profile' Class:
   * [CoreWebView2Profile.AddBrowserExtensionAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true#addbrowserextensionasync)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalBrowserExtension interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalbrowserextension?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalBrowserExtension::get_IsEnabled property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalbrowserextension?view=webview2-1.0.1988-prerelease&preserve-view=true#get_isenabled)<!--no put-->
   * [ICoreWebView2ExperimentalBrowserExtension::get_Name property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalbrowserextension?view=webview2-1.0.1988-prerelease&preserve-view=true#get_name)<!--no put-->
* [ICoreWebView2ExperimentalBrowserExtensionEnableCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalbrowserextensionenablecompletedhandler?view=webview2-1.0.1988-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalBrowserExtensionList interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalbrowserextensionlist?view=webview2-1.0.1988-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalBrowserExtensionRemoveCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalbrowserextensionremovecompletedhandler?view=webview2-1.0.1988-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalEnvironmentOptions interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalEnvironmentOptions::get_AreBrowserExtensionsEnabled property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions?view=webview2-1.0.1988-prerelease&preserve-view=true#get_arebrowserextensionsenabled)
   * [ICoreWebView2ExperimentalEnvironmentOptions::put_AreBrowserExtensionsEnabled property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions?view=webview2-1.0.1988-prerelease&preserve-view=true#put_arebrowserextensionsenabled)
* [ICoreWebView2ExperimentalProfile12 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile12?view=webview2-1.0.1988-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalProfile12::AddBrowserExtension method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile12?view=webview2-1.0.1988-prerelease&preserve-view=true#addbrowserextension)
   * [ICoreWebView2ExperimentalProfile12::GetBrowserExtensions method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile12?view=webview2-1.0.1988-prerelease&preserve-view=true#getbrowserextensions)
* [ICoreWebView2ExperimentalProfileAddBrowserExtensionCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofileaddbrowserextensioncompletedhandler?view=webview2-1.0.1988-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalProfileGetBrowserExtensionsCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofilegetbrowserextensionscompletedhandler?view=webview2-1.0.1988-prerelease&preserve-view=true)

---


<!-- ------------------------------ -->
* The 'TextDirectionKind' enum specifies the text direction as left to right or right to left.

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2TextDirectionKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2textdirectionkind?view=webview2-dotnet-1.0.1988-prerelease&preserve-view=true)
   * 'Default'
   * 'LeftToRight'
   * 'RightToLeft'

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2TextDirectionKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2textdirectionkind?view=webview2-winrt-1.0.1988-prerelease&preserve-view=true)
   * 'Default'
   * 'LeftToRight'
   * 'RightToLeft'

##### [Win32/C++](#tab/win32cpp)

* [COREWEBVIEW2_TEXT_DIRECTION_KIND enum](/microsoft-edge/webview2/reference/win32/webview2experimental-idl?view=webview2-1.0.1988-prerelease&preserve-view=true#corewebview2_text_direction_kind)
   * 'COREWEBVIEW2_TEXT_DIRECTION_KIND_DEFAULT'
   * 'COREWEBVIEW2_TEXT_DIRECTION_KIND_LEFT_TO_RIGHT'
   * 'COREWEBVIEW2_TEXT_DIRECTION_KIND_RIGHT_TO_LEFT'

---


<!-- ------------------------------ -->
###### Bug fixes

* Fixed a 'CoreWebView2Frame.ExecuteScriptAsync' hang that occurred when a frame was destroyed during script execution. [Issue 3124](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3124)
* Fixed a 'COMException' when reading 'WebResourceResponse' content after a redirect. [Issue 3229](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3229)
* Fixed a regression where calling 'CoreWebView2.AddHostObjectToScript' twice for the same name hangs (Runtime). [Issue 3539](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3539)
* Fixed an issue where 'PrintAsync' fails when 'PrinterName' contains Chinese characters. [Issue 3379](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3379)
* Fixed an issue to disable the context menu in print pages when 'AreDefaultContextMenusEnabled' is set to 'false'. [Issue 3548](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3548)
* Removed visual search from the web capture context menu (Runtime). [Issue 3426](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3426)
* Fixed an issue that caused 'PrintAsync' and 'PrintToPdfStreamAsync' to fail when print settings are 'null'.
* Removed the **Launch game** button from the default **No Internet Connection** error page (Runtime).
* Fixed an issue to ensure that 'WebVivew2Loader' can be loaded from a UNC path. [Issue 3465](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3465)
* Fixed invalid 'CoreWebView2PdfToolbarItems.FullScreen' and 'CoreWebView2PdfToolbarItems.MoreSettings'.
* Added a lock for host object access from multithread (Runtime).
* Fixed configuration options that ('CoreWebView2PdfToolbarItems.MoreSettings', 'CoreWebView2PdfToolbarItems.FullScreen') are not valid in PDF preview mode. [Issue 3324](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3324)
* Removed the **Hide all annotations** option in PDF **Settings and more** (Runtime).
* Removed the **Show all saved passwords** context menu item (Runtime).

<!-- end of ## 1.0.1988-prerelease July -->


<!-- ====================================================================== -->
Release Date: June 5, 2023`,
	},
	"1.0.1823.32": {
		SDKVersion:     "1.0.1823.32",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10182332",
		RuntimeVersion: "114.0.1823.32",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ------------------------------ -->
###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
* The 'LaunchingExternalUriScheme' event is raised when there's an attempt to launch a URI scheme that is registered with the OS (an external URI scheme).

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.LaunchingExternalUriScheme Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.launchingexternalurischeme?view=webview2-dotnet-1.0.1823.32&preserve-view=true)
* [CoreWebView2LaunchingExternalUriSchemeEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs?view=webview2-dotnet-1.0.1823.32&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Cancel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.cancel?view=webview2-dotnet-1.0.1823.32&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.InitiatingOrigin Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.initiatingorigin?view=webview2-dotnet-1.0.1823.32&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.IsUserInitiated Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.isuserinitiated?view=webview2-dotnet-1.0.1823.32&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Uri Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.uri?view=webview2-dotnet-1.0.1823.32&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.GetDeferral Method](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.getdeferral?view=webview2-dotnet-1.0.1823.32&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.LaunchingExternalUriScheme Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1823.32&preserve-view=true#launchingexternalurischeme)
* [CoreWebView2LaunchingExternalUriSchemeEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1823.32&preserve-view=true)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Cancel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1823.32&preserve-view=true#cancel)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.InitiatingOrigin Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1823.32&preserve-view=true#initiatingorigin)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.IsUserInitiated Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1823.32&preserve-view=true#isuserinitiated)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Uri Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1823.32&preserve-view=true#uri)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.GetDeferral Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1823.32&preserve-view=true#getdeferral)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_18](/microsoft-edge/webview2/reference/win32/icorewebview2_18?view=webview2-1.0.1823.32&preserve-view=true)
    * [ICoreWebView2_18::add_LaunchingExternalUriScheme](/microsoft-edge/webview2/reference/win32/icorewebview2_18?view=webview2-1.0.1823.32&preserve-view=true#add_launchingexternalurischeme)
    * [ICoreWebView2_18::remove_LaunchingExternalUriScheme](/microsoft-edge/webview2/reference/win32/icorewebview2_18?view=webview2-1.0.1823.32&preserve-view=true#remove_launchingexternalurischeme)
* [ICoreWebView2LaunchingExternalUriSchemeEventHandler](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventhandler?view=webview2-1.0.1823.32&preserve-view=true)
* [ICoreWebView2LaunchingExternalUriSchemeEventArgs](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1823.32&preserve-view=true)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::get_Cancel](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1823.32&preserve-view=true#get_cancel)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::get_InitiatingOrigin](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1823.32&preserve-view=true#get_initiatingorigin)<!--no put-->
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::get_IsUserInitiated](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1823.32&preserve-view=true#get_isuserinitiated)<!--no put-->
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::get_Uri](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1823.32&preserve-view=true#get_uri)<!--no put-->
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::GetDeferral](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1823.32&preserve-view=true#getdeferral)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::put_Cancel](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1823.32&preserve-view=true#put_cancel)

---


<!-- ------------------------------ -->
* 'MemoryUsageTargetLevel' specifies memory consumption levels, such as 'low' or 'normal'.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.MemoryUsageTargetLevel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2.memoryusagetargetlevel?view=webview2-dotnet-1.0.1823.32&preserve-view=true)
* [CoreWebView2MemoryUsageTargetLevel Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2memoryusagetargetlevel?view=webview2-dotnet-1.0.1823.32&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.MemoryUsageTargetLevel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2#memoryusagetargetlevel?view=webview2-winrt-1.0.1823.32&preserve-view=true)
* [CoreWebView2MemoryUsageTargetLevel Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2memoryusagetargetlevel?view=webview2-winrt-1.0.1823.32&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_19 interface](/microsoft-edge/webview2/reference/win32/icorewebview2_19?view=webview2-1.0.1823.32&preserve-view=true)
    * [ICoreWebView2_19::get_MemoryUsageTargetLevel](/microsoft-edge/webview2/reference/win32/icorewebview2_19?view=webview2-1.0.1823.32&preserve-view=true#get_memoryusagetargetlevel)
    * [ICoreWebView2_19::put_MemoryUsageTargetLevel](/microsoft-edge/webview2/reference/win32/icorewebview2_19?view=webview2-1.0.1823.32&preserve-view=true#put_memoryusagetargetlevel)
* [COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1823.32&preserve-view=true#corewebview2_memory_usage_target_level)

---


<!-- ------------------------------ -->
* Enhanced support for multiple profiles, to allow configuring General Autofill and Password Autosave settings for different profiles.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2Profile' Class:
   * [CoreWebView2Profile.IsGeneralAutofillEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.isgeneralautofillenabled?view=webview2-dotnet-1.0.1823.32&preserve-view=true)
   * [CoreWebView2Profile.IsPasswordAutosaveEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.ispasswordautosaveenabled?view=webview2-dotnet-1.0.1823.32&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2Profile' Class:
   * [CoreWebView2Profile.IsGeneralAutofillEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1823.32&preserve-view=true#isgeneralautofillenabled)
   * [CoreWebView2Profile.IsPasswordAutosaveEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1823.32&preserve-view=true#ispasswordautosaveenabled)

##### [Win32/C++](#tab/win32cpp)

* 'ICoreWebView2Profile6' interface:
   * [ICoreWebView2Profile6::get_IsGeneralAutofillEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2profile6?view=webview2-1.0.1823.32&preserve-view=true#get_isgeneralautofillenabled)
   * [ICoreWebView2Profile6::get_IsPasswordAutosaveEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2profile6?view=webview2-1.0.1823.32&preserve-view=true#get_ispasswordautosaveenabled)
   * [ICoreWebView2Profile6::put_IsGeneralAutofillEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2profile6?view=webview2-1.0.1823.32&preserve-view=true#put_isgeneralautofillenabled)
   * [ICoreWebView2Profile6::put_IsPasswordAutosaveEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2profile6?view=webview2-1.0.1823.32&preserve-view=true#put_ispasswordautosaveenabled)

---


<!-- ====================================================================== -->
Release Date: June 12, 2023`,
	},
	"1.0.1905-prerelease": {
		SDKVersion:     "1.0.1905-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101905-prerelease",
		RuntimeVersion: "116.0.1905.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ------------------------------ -->
<!-- ###### Experimental features -->

<!-- no added experimental features this time  ------------------------------ -->


<!-- ------------------------------ -->
###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ------------------------------ -->
* 'NavigationKind' gets the navigation kind of each navigation, such as Back/Forward, Reload, or navigation to a new document. 

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2NavigationStartingEventArgs' Class:
   * [CoreWebView2NavigationStartingEventArgs.NavigationKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2navigationstartingeventargs.navigationkind?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)
* [CoreWebView2NavigationKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2navigationkind?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2NavigationStartingEventArgs' Class:
   * [CoreWebView2NavigationStartingEventArgs.NavigationKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2navigationstartingeventargs?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true#navigationkind)
* [CoreWebView2NavigationKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2navigationkind?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2NavigationStartingEventArgs3 interface](/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs3?view=webview2-1.0.1905-prerelease&preserve-view=true)
    * [ICoreWebView2NavigationStartingEventArgs3::get_NavigationKind property](/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs3?view=webview2-1.0.1905-prerelease&preserve-view=true#get_navigationkind)<!--no put-->
* [COREWEBVIEW2_NAVIGATION_KIND enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1905-prerelease&preserve-view=true#corewebview2_navigation_kind)

---


<!-- ------------------------------ -->
* The 'ServiceWorkers' enum value in the 'BrowsingDataKinds' enum specifies service workers that are registered for an origin.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2BrowsingDataKinds' Enum:
   * [CoreWebView2BrowsingDataKinds.ServiceWorkers Enum Value](/dotnet/api/microsoft.web.webview2.core.corewebview2browsingdatakinds?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2BrowsingDataKinds' Enum:
   * [CoreWebView2BrowsingDataKinds.ServiceWorkers Enum Value](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2browsingdatakinds?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* 'COREWEBVIEW2_BROWSING_DATA_KINDS' enum:
   * [COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS enum value](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1905-prerelease&preserve-view=true#corewebview2_browsing_data_kinds)

---


<!-- ------------------------------ -->
* The 'LaunchingExternalUriScheme' event is raised when there's an attempt to launch a URI scheme that is registered with the OS (an external URI scheme).

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.LaunchingExternalUriScheme Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.launchingexternalurischeme?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)
* [CoreWebView2LaunchingExternalUriSchemeEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Cancel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.cancel?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.InitiatingOrigin Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.initiatingorigin?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.IsUserInitiated Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.isuserinitiated?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Uri Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.uri?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)
    * [CoreWebView2LaunchingExternalUriSchemeEventArgs.GetDeferral Method](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.getdeferral?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.LaunchingExternalUriScheme Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true#launchingexternalurischeme)
* [CoreWebView2LaunchingExternalUriSchemeEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Cancel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true#cancel)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.InitiatingOrigin Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true#initiatingorigin)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.IsUserInitiated Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true#isuserinitiated)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Uri Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true#uri)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.GetDeferral Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true#getdeferral)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_18 interface](/microsoft-edge/webview2/reference/win32/icorewebview2_18?view=webview2-1.0.1905-prerelease&preserve-view=true)
    * [ICoreWebView2_18::add_LaunchingExternalUriScheme event](/microsoft-edge/webview2/reference/win32/icorewebview2_18?view=webview2-1.0.1905-prerelease&preserve-view=true#add_launchingexternalurischeme)
    * [ICoreWebView2_18::remove_LaunchingExternalUriScheme event](/microsoft-edge/webview2/reference/win32/icorewebview2_18?view=webview2-1.0.1905-prerelease&preserve-view=true#remove_launchingexternalurischeme)
* [ICoreWebView2LaunchingExternalUriSchemeEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventhandler?view=webview2-1.0.1905-prerelease&preserve-view=true)
* [ICoreWebView2LaunchingExternalUriSchemeEventArgs interface](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1905-prerelease&preserve-view=true)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::get_Cancel property](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1905-prerelease&preserve-view=true#get_cancel)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::get_InitiatingOrigin property](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1905-prerelease&preserve-view=true#get_initiatingorigin)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::get_IsUserInitiated property](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1905-prerelease&preserve-view=true#get_isuserinitiated)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::get_Uri property](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1905-prerelease&preserve-view=true#get_uri)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::GetDeferral method](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1905-prerelease&preserve-view=true#getdeferral)
    * [ICoreWebView2LaunchingExternalUriSchemeEventArgs::put_Cancel property](/microsoft-edge/webview2/reference/win32/icorewebview2launchingexternalurischemeeventargs?view=webview2-1.0.1905-prerelease&preserve-view=true#put_cancel)

---


<!-- ------------------------------ -->
* 'MemoryUsageTargetLevel' specifies memory consumption levels, such as 'low' or 'normal'.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.MemoryUsageTargetLevel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2.memoryusagetargetlevel?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)
* [CoreWebView2MemoryUsageTargetLevel Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2memoryusagetargetlevel?view=webview2-dotnet-1.0.1905-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.MemoryUsageTargetLevel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true#memoryusagetargetlevel)
* [CoreWebView2MemoryUsageTargetLevel Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2memoryusagetargetlevel?view=webview2-winrt-1.0.1905-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_19 interface](/microsoft-edge/webview2/reference/win32/icorewebview2_19?view=webview2-1.0.1905-prerelease&preserve-view=true)
    * [ICoreWebView2_19::get_MemoryUsageTargetLevel property](/microsoft-edge/webview2/reference/win32/icorewebview2_19?view=webview2-1.0.1905-prerelease&preserve-view=true#get_memoryusagetargetlevel)
    * [ICoreWebView2_19::put_MemoryUsageTargetLevel property](/microsoft-edge/webview2/reference/win32/icorewebview2_19?view=webview2-1.0.1905-prerelease&preserve-view=true#put_memoryusagetargetlevel)
* [COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1905-prerelease&preserve-view=true#corewebview2_memory_usage_target_level)

---


<!-- ------------------------------ -->
###### Bug fixes

* Using 'wv2winrt webhosthidden' entered an infinite loop when enumerating some 'webhosthidden' types.  (SDK-only)
* In code that's generated by the **wv2winrt** tool, when calling an async method, it would crash if it succeeded but returned 'null' instead of an 'IAsyncAction'.  (SDK-only)


<!-- ====================================================================== -->
Release Date: May 8, 2023`,
	},
	"1.0.1774.30": {
		SDKVersion:     "1.0.1774.30",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10177430",
		RuntimeVersion: "113.0.1774.30",
		Notes: 			`

<!-- ------------------------------ -->
#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
* The File API allows accessing a DOM 'File' object passed via 'WebMessage'.

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2File Class](/dotnet/api/microsoft.web.webview2.core.corewebview2file?view=webview2-dotnet-1.0.1774.30&preserve-view=true)
   * [CoreWebView2File.Path Property](/dotnet/api/microsoft.web.webview2.core.corewebview2file.path?view=webview2-dotnet-1.0.1774.30&preserve-view=true)
* 'CoreWebView2WebMessageReceivedEventArgs' Class:
   * [CoreWebView2WebMessageReceivedEventArgs.AdditionalObjects Property](/dotnet/api/microsoft.web.webview2.core.corewebview2webmessagereceivedeventargs.additionalobjects?view=webview2-dotnet-1.0.1774.30&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2File Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2file?view=webview2-winrt-1.0.1774.30&preserve-view=true)
   * [CoreWebView2File.Path Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2file?view=webview2-winrt-1.0.1774.30&preserve-view=true#path)
* 'CoreWebView2WebMessageReceivedEventArgs' Class:
   * [CoreWebView2WebMessageReceivedEventArgs.AdditionalObjects Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webmessagereceivedeventargs?view=webview2-winrt-1.0.1774.30&preserve-view=true#additionalobjects)
   
##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2File interface](/microsoft-edge/webview2/reference/win32/icorewebview2file?view=webview2-1.0.1774.30&preserve-view=true)
   * [ICoreWebView2File::get_Path](/microsoft-edge/webview2/reference/win32/icorewebview2file?view=webview2-1.0.1774.30&preserve-view=true#get_path)<!--no put-->
* [ICoreWebView2ObjectCollectionView interface](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview?view=webview2-1.0.1774.30&preserve-view=true)
   * [ICoreWebView2ObjectCollectionView::get_Count](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview?view=webview2-1.0.1774.30&preserve-view=true#get_count)<!--no put-->
   * [ICoreWebView2ObjectCollectionView::GetValueAtIndex](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview?view=webview2-1.0.1774.30&preserve-view=true#getvalueatindex)
* [ICoreWebView2WebMessageReceivedEventArgs2 interface](/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs2?view=webview2-1.0.1774.30&preserve-view=true)
   * [ICoreWebView2WebMessageReceivedEventArgs2::get_AdditionalObjects](/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs2?view=webview2-1.0.1774.30&preserve-view=true#get_additionalobjects)<!--no put-->

---


<!-- ------------------------------ -->
* The Profile Cookie Manager API supports profile management.  The 'CookieManager' property enables the host app to get the cookie manager for the profile.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2Profile' Class:
   * [CoreWebView2Profile.CookieManager Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.cookiemanager?view=webview2-dotnet-1.0.1774.30&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2Profile' Class:
   * [CoreWebView2Profile.CookieManager Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1774.30&preserve-view=true#cookiemanager)

##### [Win32/C++](#tab/win32cpp)

* 'ICoreWebView2Profile5' interface:
   * [ICoreWebView2Profile5::get_CookieManager](/microsoft-edge/webview2/reference/win32/icorewebview2profile5?view=webview2-1.0.1774.30&preserve-view=true#get_cookiemanager)<!--no put-->

---


<!-- ------------------------------ -->
###### Bug fixes

* Fixed an issue to allow an app to inject initial scripts by calling 'AddScriptToExecuteOnDocumentCreated' before a new window is created.  ([Issue #2491](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2491))

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.AddScriptToExecuteOnDocumentCreatedAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.addscripttoexecuteondocumentcreatedasync)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.AddScriptToExecuteOnDocumentCreatedAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2#addscripttoexecuteondocumentcreatedasync)

##### [Win32/C++](#tab/win32cpp)

* 'ICoreWebView2' interface:
   * [ICoreWebView2::AddScriptToExecuteOnDocumentCreated method](/microsoft-edge/webview2/reference/win32/icorewebview2#addscripttoexecuteondocumentcreated)

---

* (Runtime-only)  Fixed an issue that was causing the 'X-Edge-Shopping-Flag' header to be added to web requests that are coming from WebView2.  ([Issue #3365](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3365))


<!-- ====================================================================== -->
Release Date: May 8, 2023`,
	},
	"1.0.1829-prerelease": {
		SDKVersion:     "1.0.1829-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101829-prerelease",
		RuntimeVersion: "115.0.1829.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ------------------------------ -->
###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ------------------------------ -->
* Enhanced support for multiple profiles, to allow configuring General Autofill and Password Autosave settings for different profiles.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2Profile' Class:
   * [CoreWebView2Profile.IsGeneralAutofillEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.isgeneralautofillenabled?view=webview2-dotnet-1.0.1829-prerelease&preserve-view=true)
   * [CoreWebView2Profile.IsPasswordAutosaveEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.ispasswordautosaveenabled?view=webview2-dotnet-1.0.1829-prerelease&preserve-view=true)


##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2Profile' Class:
   * [CoreWebView2Profile.IsGeneralAutofillEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1829-prerelease&preserve-view=true#isgeneralautofillenabled)
   * [CoreWebView2Profile.IsPasswordAutosaveEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1829-prerelease&preserve-view=true#ispasswordautosaveenabled)


##### [Win32/C++](#tab/win32cpp)

* 'ICoreWebView2Profile6' interface:
   * [ICoreWebView2Profile6::get_IsPasswordAutosaveEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2profile6?view=webview2-1.0.1829-prerelease&preserve-view=true#get_ispasswordautosaveenabled)
   * [ICoreWebView2Profile6::put_IsPasswordAutosaveEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2profile6?view=webview2-1.0.1829-prerelease&preserve-view=true#put_ispasswordautosaveenabled)
   * [ICoreWebView2Profile6::get_IsGeneralAutofillEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2profile6?view=webview2-1.0.1829-prerelease&preserve-view=true#get_isgeneralautofillenabled)
   * [ICoreWebView2Profile6::put_IsGeneralAutofillEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2profile6?view=webview2-1.0.1829-prerelease&preserve-view=true#put_isgeneralautofillenabled)

---


<!-- ------------------------------ -->
###### Bug fixes

* Disabled the Chrome Web Store info banner that displays the option to allow extensions installation. ([Issue #3312](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3312))
* Fixed an issue where a custom menu item wasn't firing. ([Issue #3300](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3300))
* Fixed a crash during initialization when creating a WebView2 using WPF and SDK version 1.0.1722.32, which is now deprecated.  (See [SDK 1.0.1722.32 is deprecated](#sdk-10172232-is-deprecated) below.)  ([Issue #3375](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3375))
* (Runtime-only)  Fixed a bug in 'PostSharedBufferToScript' that stops after about 32000x1MB buffers are posted.  ([Issue #3360](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3360))

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.PostSharedBufferToScript Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.postsharedbuffertoscript)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class:
   * [CoreWebView2.PostSharedBufferToScript Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2#postsharedbuffertoscript)

##### [Win32/C++](#tab/win32cpp)

* 'ICoreWebView2_17' interface:
   * [ICoreWebView2_17::PostSharedBufferToScript method](/microsoft-edge/webview2/reference/win32/icorewebview2_17#postsharedbuffertoscript)

---

* (Runtime-only)  Fixed an issue where navigation will always take place within a 'ScriptDialogOpening' event callback.  ([Issue #3355](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3355))
* (Runtime-only)  Fixed an issue to support the 'BackForwardCache' flag.
* Fixed an issue with visual hosted owned windows, where clicking into the Find bar from outside the window didn't activate the Find bar.


<!-- ====================================================================== -->
Release Date: April 13, 2023`,
	},
	"1.0.1722.45": {
		SDKVersion:     "1.0.1722.45",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10172245",
		RuntimeVersion: "112.0.1722.45",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ------------------------------ -->
###### SDK 1.0.1722.32 is deprecated

WebView2 SDK 1.0.1722.32 is deprecated, and that package has been removed from the listing at NuGet.  Discontinue development with package 1.0.1722.32.  If your WebView2 app uses that package, we recommend that you move to a newer package, such as WebView2 SDK 1.0.1722.45 or later.


<!-- ------------------------------ -->
###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
* The Managing SmartScreen API controls whether SmartScreen is enabled.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2Settings'
   * [CoreWebView2Settings.IsReputationCheckingRequired Property](/dotnet/api/microsoft.web.webview2.core.corewebview2settings.isreputationcheckingrequired?view=webview2-dotnet-1.0.1722.45&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2Settings'
   * [CoreWebView2Settings.IsReputationCheckingRequired Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2settings?view=webview2-winrt-1.0.1722.45&preserve-view=true#isreputationcheckingrequired)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Settings8](/microsoft-edge/webview2/reference/win32/icorewebview2settings8?view=webview2-1.0.1722.45&preserve-view=true)
   * [ICoreWebView2Settings8::get_IsReputationCheckingRequired method](/microsoft-edge/webview2/reference/win32/icorewebview2settings8?view=webview2-1.0.1722.45&preserve-view=true#get_isreputationcheckingrequired)
   * [ICoreWebView2Settings8::put_IsReputationCheckingRequired method](/microsoft-edge/webview2/reference/win32/icorewebview2settings8?view=webview2-1.0.1722.45&preserve-view=true#put_isreputationcheckingrequired)

---


<!-- ------------------------------ -->
* The 'PermissionKind.WindowManagement' API indicates the kind of a permission request.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2PermissionKind' Enum
   * [CoreWebView2PermissionKind.WindowManagement Enum Value](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionkind?view=webview2-dotnet-1.0.1722.45&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2PermissionKind' Enum
   * [CoreWebView2PermissionKind.WindowManagement Enum Value](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionkind?view=webview2-winrt-1.0.1722.45&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* 'COREWEBVIEW2_PERMISSION_KIND' Enum
   * [COREWEBVIEW2_PERMISSION_KIND_WINDOW_MANAGEMENT enum value](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1722.45&preserve-view=true#corewebview2_permission_kind)

---


<!-- ====================================================================== -->
Release Date: April 10, 2023`,
	},
	"1.0.1777-prerelease": {
		SDKVersion:     "1.0.1777-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101777-prerelease",
		RuntimeVersion: "114.0.1777.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ------------------------------ -->
###### Experimental features

No experimental features are added in this Prerelease SDK.


<!-- ------------------------------ -->
###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ------------------------------ -->
* The File API allows accessing a DOM 'File' object passed via 'WebMessage'.

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2File](/dotnet/api/microsoft.web.webview2.core.corewebview2file?view=webview2-dotnet-1.0.1777-prerelease&preserve-view=true)
   * [CoreWebView2File.path property](/dotnet/api/microsoft.web.webview2.core.corewebview2file.path?view=webview2-dotnet-1.0.1777-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2file-path)

* 'CoreWebView2WebMessageReceivedEventArgs'
   * [CoreWebView2WebMessageReceivedEventArgs.AdditionalObjects property](/dotnet/api/microsoft.web.webview2.core.corewebview2webmessagereceivedeventargs.additionalobjects?view=webview2-dotnet-1.0.1777-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2webmessagereceivedeventargs-additionalobjects)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2File](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2file?view=webview2-winrt-1.0.1777-prerelease&preserve-view=true)
   * [CoreWebView2File.path property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2file?view=webview2-winrt-1.0.1777-prerelease&preserve-view=true#path)

* 'CoreWebView2WebMessageReceivedEventArgs'
   * [CoreWebView2WebMessageReceivedEventArgs.AdditionalObjects property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webmessagereceivedeventargs?view=webview2-winrt-1.0.1777-prerelease&preserve-view=true#additionalobjects)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2File](/microsoft-edge/webview2/reference/win32/icorewebview2file?view=webview2-1.0.1777-prerelease&preserve-view=true)
   * [ICoreWebView2File::get_path](/microsoft-edge/webview2/reference/win32/icorewebview2file?view=webview2-1.0.1777-prerelease&preserve-view=true#get_path)

* [ICoreWebView2WebMessageReceivedEventArgs2](/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs2?view=webview2-1.0.1777-prerelease&preserve-view=true)
   * [ICoreWebView2WebMessageReceivedEventArgs2::get_AdditionalObjects](/microsoft-edge/webview2/reference/win32/icorewebview2webmessagereceivedeventargs2?view=webview2-1.0.1777-prerelease&preserve-view=true#get_AdditionalObjects)

* [ICoreWebView2ObjectCollectionView](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview?view=webview2-1.0.1777-prerelease&preserve-view=true)
   * [ICoreWebView2ObjectCollectionView::get_Count](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview?view=webview2-1.0.1777-prerelease&preserve-view=true#get_Count)
   * [ICoreWebView2ObjectCollectionView::GetValueAtIndex](/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview?view=webview2-1.0.1777-prerelease&preserve-view=true#GetValueAtIndex)

---


<!-- ------------------------------ -->
* The Profile Cookie Manager API supports profile management.  The 'CookieManager' property enables the host app to get the cookie manager for the profile.

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2Profile'
   * [CoreWebView2Profile.CookieManager property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.cookiemanager?view=webview2-dotnet-1.0.1777-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2profile-cookiemanager)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2Profile'
   * [CoreWebView2Profile.CookieManager property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1777-prerelease&preserve-view=true#cookiemanager)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Profile5](/microsoft-edge/webview2/reference/win32/icorewebview2profile5?view=webview2-1.0.1777-prerelease&preserve-view=true)
   * [ICoreWebView2Profile5::get_CookieManager](/microsoft-edge/webview2/reference/win32/icorewebview2profile5?view=webview2-1.0.1777-prerelease&preserve-view=true#get_cookiemanager)

---


<!-- ------------------------------ -->
###### Bug fixes

* Fixed a crash when releasing the WebView from a different thread.  (Runtime-only)  ([Issue #3062](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3062))
* Fixed a bug where focus was trapped inside the WebView2 control when wrapped in a 'ContainerControl'.  ([Issue #2835](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2835))
* Fixed the issue by disabling the editable '.pdf' temporary cached data recovery function in WebView2.  ([Issue #3274](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3274))
* Disabled the Chrome Web Store info banner that displays the option to allow extensions installation.  ([Issue #3312](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3312))
* Fixed an issue with new download items not getting called out by screen readers.
* Fixed a bug where visual hosted owned windows didn't map mouse pointer input correctly.
* Fixed a bug where 'DownloadStarting' was getting raised for a canceled **Save As** dialog.  (Runtime-only)


<!-- ====================================================================== -->
Release Date: March 20, 2023`,
	},
	"1.0.1661.34": {
		SDKVersion:     "1.0.1661.34",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10166134",
		RuntimeVersion: "111.0.1661.34",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ---------- -->
###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
*  The SharedBuffer API:

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2' Class
   * [CoreWebView2.PostSharedBufferToScript Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.postsharedbuffertoscript?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

* 'CoreWebView2Environment' Class
   * [ICoreWebView2Environment.CreateSharedBuffer Method](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.createsharedbuffer?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

* 'CoreWebView2Frame' Class
   * [CoreWebView2Frame.PostSharedBufferToScript Method](/dotnet/api/microsoft.web.webview2.core.corewebview2frame.postsharedbuffertoscript?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

* [CoreWebView2SharedBuffer Class](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2SharedBuffer.Buffer Property](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.buffer?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2SharedBuffer.FileMappingHandle Property](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.filemappinghandle?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2SharedBuffer.Size Property](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.size?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2SharedBuffer.Close Method](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.close?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2SharedBuffer.Dispose Method](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.dispose?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2SharedBuffer.OpenStream Method](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.openstream?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

* [CoreWebView2SharedBufferAccess Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbufferaccess?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * 'ReadOnly'
   * 'ReadWrite'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2' Class
   * [CoreWebView2.PostSharedBufferToScript Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1661.34&preserve-view=true#postsharedbuffertoscript)

* 'CoreWebView2Environment' Class
   * [ICoreWebView2Environment.CreateSharedBuffer Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1661.34&preserve-view=true#createsharedbuffer)

* 'CoreWebView2Frame' Class
   * [CoreWebView2Frame.PostSharedBufferToScript Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frame?view=webview2-winrt-1.0.1661.34&preserve-view=true#postsharedbuffertoscript)

* [CoreWebView2SharedBuffer Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1661.34&preserve-view=true)
   * [CoreWebView2SharedBuffer.Buffer Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1661.34&preserve-view=true#buffer)
   * [CoreWebView2SharedBuffer.Size Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1661.34&preserve-view=true#size)
   * [CoreWebView2SharedBuffer.Close Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1661.34&preserve-view=true#close)
   * [CoreWebView2SharedBuffer.OpenStream Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1661.34&preserve-view=true#openstream)

* [CoreWebView2SharedBufferAccess Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbufferaccess?view=webview2-winrt-1.0.1661.34&preserve-view=true)
   * 'ReadOnly'
   * 'ReadWrite'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_17 interface](/microsoft-edge/webview2/reference/win32/icorewebview2_17?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2_17::PostSharedBufferToScript method](/microsoft-edge/webview2/reference/win32/icorewebview2_17?view=webview2-1.0.1661.34&preserve-view=true#postsharedbuffertoscript)

* [ICoreWebView2Environment12 interface](/microsoft-edge/webview2/reference/win32/icorewebview2environment12?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2Environment12::CreateSharedBuffer method](/microsoft-edge/webview2/reference/win32/icorewebview2environment12?view=webview2-1.0.1661.34&preserve-view=true#createsharedbuffer)

* [ICoreWebView2Frame4 interface](/microsoft-edge/webview2/reference/win32/icorewebview2frame4?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2Frame4::PostSharedBufferToScript method](/microsoft-edge/webview2/reference/win32/icorewebview2frame4?view=webview2-1.0.1661.34&preserve-view=true#postsharedbuffertoscript)

* [ICoreWebView2SharedBuffer interface](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2SharedBuffer::OpenStream method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1661.34&preserve-view=true#openstream)
   * [ICoreWebView2SharedBuffer::Close method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1661.34&preserve-view=true#close)
   * [ICoreWebView2SharedBuffer::get_Size method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1661.34&preserve-view=true#get_size)
   * [ICoreWebView2SharedBuffer::get_Buffer method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1661.34&preserve-view=true#get_buffer)
   * [ICoreWebView2SharedBuffer::get_FileMappingHandle method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1661.34&preserve-view=true#get_filemappinghandle)

* [COREWEBVIEW2_SHARED_BUFFER_ACCESS](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1661.34&preserve-view=true#corewebview2_shared_buffer_access)
   * 'COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY'
   * 'COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_WRITE'

---


<!-- ------------------------------ -->
*  APIs for managing permissions:

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2PermissionKind' Enum
   * [CoreWebView2PermissionKind.MidiSystemExclusiveMessages Enum Value](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionkind?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

* 'CoreWebView2PermissionRequestedEventArgs' Event
   * [CoreWebView2PermissionRequestedEventArgs.SavesInProfile Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionrequestedeventargs.savesinprofile?view=webview2-dotnet-1.0.1661.34&preserve-view=true)   

* [CoreWebView2PermissionSetting Class](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionkind?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionOrigin Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionorigin?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionState Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionstate?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

* 'CoreWebView2Profile' Class
   * [CoreWebView2Profile.GetNonDefaultPermissionSettingsAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.getnondefaultpermissionsettingsasync?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
   * [CoreWebView2Profile.SetPermissionStateAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.setpermissionstateasync?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2PermissionKind' Enum
   * [CoreWebView2PermissionKind.MidiSystemExclusiveMessages Enum Value](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionkind?view=webview2-winrt-1.0.1661.34&preserve-view=true)

* 'CoreWebView2PermissionRequestedEventArgs' Event
   * [CoreWebView2PermissionRequestedEventArgs.SavesInProfile Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionrequestedeventargs?view=webview2-winrt-1.0.1661.34&preserve-view=true#savesinprofile)

* [CoreWebView2PermissionSetting Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1661.34&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1661.34&preserve-view=true#permissionkind)
   * [CoreWebView2PermissionSetting.PermissionOrigin Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1661.34&preserve-view=true#permissionorigin)
   * [CoreWebView2PermissionSetting.PermissionState Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1661.34&preserve-view=true#permissionstate)

* 'CoreWebView2Profile' Class
   * [CoreWebView2Profile.GetNonDefaultPermissionSettingsAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1661.34&preserve-view=true#getnondefaultpermissionsettingsasync)
   * [CoreWebView2Profile.SetPermissionStateAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1661.34&preserve-view=true#setpermissionstateasync)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2getnondefaultpermissionsettingscompletedhandler?view=webview2-1.0.1661.34&preserve-view=true)

* [ICoreWebView2PermissionRequestedEventArgs3 interface](/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs3?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2PermissionRequestedEventArgs3::get_SavesInProfile](/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs3?view=webview2-1.0.1661.34&preserve-view=true#get_savesinprofile)
   * [ICoreWebView2PermissionRequestedEventArgs3::put_SavesInProfile](/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs3?view=webview2-1.0.1661.34&preserve-view=true#put_savesinprofile)

* [ICoreWebView2PermissionSetting interface](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2PermissionSetting::get_PermissionKind method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting?view=webview2-1.0.1661.34&preserve-view=true#get_permissionkind)
   * [ICoreWebView2PermissionSetting::get_PermissionOrigin method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting?view=webview2-1.0.1661.34&preserve-view=true#get_permissionorigin)
   * [ICoreWebView2PermissionSetting::get_PermissionState method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting?view=webview2-1.0.1661.34&preserve-view=true#get_permissionstate)

* [ICoreWebView2PermissionSettingCollectionView interface](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2PermissionSettingCollectionView::GetValueAtIndex method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview?view=webview2-1.0.1661.34&preserve-view=true#getvalueatindex)
   * [ICoreWebView2PermissionSettingCollectionView::get_Count method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview?view=webview2-1.0.1661.34&preserve-view=true#get_count)

* [ICoreWebView2Profile4 interface](/microsoft-edge/webview2/reference/win32/icorewebview2profile4?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2Profile4::GetNonDefaultPermissionSettings method](/microsoft-edge/webview2/reference/win32/icorewebview2profile4?view=webview2-1.0.1661.34&preserve-view=true#getnondefaultpermissionsettings)
   * [ICoreWebView2Profile4::SetPermissionState method](/microsoft-edge/webview2/reference/win32/icorewebview2profile4?view=webview2-1.0.1661.34&preserve-view=true#setpermissionstate)

* [ICoreWebView2SetPermissionStateCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2setpermissionstatecompletedhandler?view=webview2-1.0.1661.34&preserve-view=true)

* 'COREWEBVIEW2_PERMISSION_KIND' Enum
   * [COREWEBVIEW2_PERMISSION_KIND_MIDI_SYSTEM_EXCLUSIVE_MESSAGES enum value](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1661.34&preserve-view=true#corewebview2_permission_kind)

---


<!-- ------------------------------ -->
APIs for managing tracking prevention:

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2EnvironmentOptions' Class
   * [CoreWebView2EnvironmentOptions.EnableTrackingPrevention Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.enabletrackingprevention?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

* 'CoreWebView2Profile' Class
   * [CoreWebView2Profile.PreferredTrackingPreventionLevel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.preferredtrackingpreventionlevel?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

* [CoreWebView2TrackingPreventionLevel Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2trackingpreventionlevel?view=webview2-dotnet-1.0.1661.34&preserve-view=true)
    * 'None'
    * 'Basic'
    * 'Balanced'
    * 'Strict'

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2EnvironmentOptions' Class
   * [CoreWebView2EnvironmentOptions.EnableTrackingPrevention Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1661.34&preserve-view=true#enabletrackingprevention)

* 'CoreWebView2Profile' Class
   * [CoreWebView2Profile.PreferredTrackingPreventionLevel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1661.34&preserve-view=true#preferredtrackingpreventionlevel)

* [CoreWebView2TrackingPreventionLevel Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2trackingpreventionlevel?view=webview2-winrt-1.0.1661.34&preserve-view=true)
    * 'None'
    * 'Basic'
    * 'Balanced'
    * 'Strict'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions5 interface](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2EnvironmentOptions5::EnableTrackingPrevention property (get](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5?view=webview2-1.0.1661.34&preserve-view=true#get_enabletrackingprevention), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5?view=webview2-1.0.1661.34&preserve-view=true#put_enabletrackingprevention)

* [ICoreWebView2Profile3 interface](/microsoft-edge/webview2/reference/win32/icorewebview2profile3?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2Profile3::PreferredTrackingPreventionLevel property (get](/microsoft-edge/webview2/reference/win32/icorewebview2profile3?view=webview2-1.0.1661.34&preserve-view=true#get_preferredtrackingpreventionlevel), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2profile3?view=webview2-1.0.1661.34&preserve-view=true#put_preferredtrackingpreventionlevel)

* [COREWEBVIEW2_TRACKING_PREVENTION_LEVEL enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1661.34&preserve-view=true#corewebview2_tracking_prevention_level)
  * 'COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_NONE'
  * 'COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BASIC'
  * 'COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED'
  * 'COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_STRICT'

---


<!-- ------------------------------ -->
*  APIs to manage the value of the controller's script locale:

##### [.NET/C#](#tab/dotnetcsharp)

* 'CoreWebView2ControllerOptions' Class
   * [CoreWebView2ControllerOptions.ScriptLocale Property](/dotnet/api/microsoft.web.webview2.core.corewebview2controlleroptions.scriptlocale?view=webview2-dotnet-1.0.1661.34&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* 'CoreWebView2ControllerOptions' Class
   * [CoreWebView2ControllerOptions.ScriptLocale Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2controlleroptions?view=webview2-winrt-1.0.1661.34&preserve-view=true#scriptlocale)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ControllerOptions2 interface](/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2?view=webview2-1.0.1661.34&preserve-view=true)
   * [ICoreWebView2ControllerOptions2::get_ScriptLocale method](/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2?view=webview2-1.0.1661.34&preserve-view=true#get_scriptlocale)
   * [ICoreWebView2ControllerOptions2::put_ScriptLocale method](/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2?view=webview2-1.0.1661.34&preserve-view=true#put_scriptlocale)

---


<!-- ====================================================================== -->
Release Date: March 20, 2023`,
	},
	"1.0.1724-prerelease": {
		SDKVersion:     "1.0.1724-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101724-prerelease",
		RuntimeVersion: "113.0.1724.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General

<!-- ------------------------------ -->
###### Experimental Features


<!-- ------------------------------ -->
*  Added AdditionalObjects for WebMessage received:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2WebMessageReceivedEventArgs.AdditionalObjects Property](/dotnet/api/microsoft.web.webview2.core.corewebview2webmessagereceivedeventargs.additionalobjects?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2WebMessageReceivedEventArgs.AdditionalObjects Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webmessagereceivedeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#additionalobjects)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalWebMessageReceivedEventArgs::get_AdditionalObjects method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebmessagereceivedeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#get_additionalobjects)

---


<!-- ------------------------------ -->
*  Added Window Management permission type:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2PermissionKind.WindowManagement Enum Value](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionkind?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2PermissionKind.WindowManagement Enum Value](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionkind?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [COREWEBVIEW2_PERMISSION_KIND_WINDOW_MANAGEMENT Enum Value](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1724-prerelease&preserve-view=true#corewebview2_permission_kind)

---


<!-- ------------------------------ -->
*  Added support for launching external URIs:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.LaunchingExternalUriScheme Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.launchingexternalurischeme?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

* [CoreWebView2LaunchingExternalUriSchemeEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Cancel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.cancel?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.GetDeferral Method](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.getdeferral?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.InitiatingOrigin Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.initiatingorigin?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.IsUserInitiated Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.isuserinitiated?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Uri Property](/dotnet/api/microsoft.web.webview2.core.corewebview2launchingexternalurischemeeventargs.uri?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.LaunchingExternalUriScheme Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#launchingexternalurischeme)

* [CoreWebView2LaunchingExternalUriSchemeEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Cancel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#cancel)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.GetDeferral Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#getdeferral)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.InitiatingOrigin Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#initiatingorigin)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.IsUserInitiated Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#isuserinitiated)
   * [CoreWebView2LaunchingExternalUriSchemeEventArgs.Uri Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2launchingexternalurischemeeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#uri)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental21::add_LaunchingExternalUriScheme](/microsoft-edge/webview2/reference/win32/icorewebview2experimental21?view=webview2-1.0.1724-prerelease&preserve-view=true#add_launchingexternalurischeme)

* [ICoreWebView2Experimental21::remove_LaunchingExternalUriScheme](/microsoft-edge/webview2/reference/win32/icorewebview2experimental21?view=webview2-1.0.1724-prerelease&preserve-view=true#remove_launchingexternalurischeme)

* [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventHandler](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)
  * [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventHandler::Invoke](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true#invoke)

* [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventArgs](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true)
  * [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventArgs::get_Uri](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#get_uri)
  * [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventArgs::get_InitiatingOrigin](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#get_initiatingorigin)
  * [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventArgs::get_IsUserInitiated](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#get_isuserinitiated)
  * [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventArgs::get_Cancel](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#get_cancel)
  * [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventArgs::put_Cancel](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#put_cancel)
  * [ICoreWebView2ExperimentalLaunchingExternalUriSchemeEventArgs::GetDeferral](/microsoft-edge/webview2/reference/win32/icorewebview2experimentallaunchingexternalurischemeeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#getdeferral)

---


<!-- ------------------------------ -->
*  Added support for texture streaming:

##### [.NET/C#](#tab/dotnetcsharp)

The 'Environment' interface that returns the 'TextureStream' interface:
* [CoreWebView2Environment.CreateTextureStream Method](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.createtexturestream?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
* [CoreWebView2Environment.RenderAdapterLUIDChanged Event](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.renderadapterluidchanged?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
* [CoreWebView2Environment.RenderAdapterLUID Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.renderadapterluid?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

The 'TextureStream' interface:
* [CoreWebView2TextureStream Class](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.AddAllowedOrigin Method](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.addallowedorigin?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.CloseTexture Method](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.closetexture?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.CreateTexture Method](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.createtexture?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.ErrorReceived Event](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.errorreceived?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.GetAvailableTexture Method](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.getavailabletexture?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.Id Property](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.id?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.PresentTexture Method](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.presenttexture?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.RemoveAllowedOrigin Method](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.removeallowedorigin?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.SetD3DDevice Method](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.setd3ddevice?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.StartRequested Event](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.startrequested?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.Stop Method](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.stop?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.Stopped Event](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.stopped?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.WebTextureReceived Event](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.webtexturereceived?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.WebTextureStreamStopped Event](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestream.webtexturestreamstopped?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

ErrorReceivedEventArgs:
* [CoreWebView2TextureStreamErrorReceivedEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamerrorreceivedeventargs?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorReceivedEventArgs.Kind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamerrorreceivedeventargs.kind?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorReceivedEventArgs.texture Property](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamerrorreceivedeventargs.texture?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

WebTextureReceivedEventArgs:
* [CoreWebView2TextureStreamWebTextureReceivedEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamwebtexturereceivedeventargs?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
* [CoreWebView2TextureStreamWebTextureReceivedEventArgs.WebTexture Property](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamwebtexturereceivedeventargs.webtexture?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

TextureStream error kind enum:
* [CoreWebView2TextureStreamErrorKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamerrorkind?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorKind.CoreWebView2TextureStreamErrorNoVideoTrackStarted Enum Value](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamerrorkind?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorKind.CoreWebView2TextureStreamErrorTextureError Enum Value](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamerrorkind?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorKind.CoreWebView2TextureStreamErrorTextureInUse Enum Value](/dotnet/api/microsoft.web.webview2.core.corewebview2texturestreamerrorkind?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

The 'Texture' interface that the host writes to so that the Renderer will render on it:
* [CoreWebView2Texture Class](/dotnet/api/microsoft.web.webview2.core.corewebview2texture?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2Texture.Handle Property](/dotnet/api/microsoft.web.webview2.core.corewebview2texture.handle?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2Texture.Resource Property](/dotnet/api/microsoft.web.webview2.core.corewebview2texture.resource?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2Texture.Timestamp Property](/dotnet/api/microsoft.web.webview2.core.corewebview2texture.timestamp?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

The received 'WebTexture' interface that the Renderer writes to so that the host will read on it:
* [CoreWebView2WebTexture Class](/dotnet/api/microsoft.web.webview2.core.corewebview2webtexture?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2WebTexture.Handle Property](/dotnet/api/microsoft.web.webview2.core.corewebview2webtexture.handle?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2WebTexture.Resource Property](/dotnet/api/microsoft.web.webview2.core.corewebview2webtexture.resource?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2WebTexture.Timestamp Property](/dotnet/api/microsoft.web.webview2.core.corewebview2webtexture.timestamp?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

The 'Environment' interface that returns the 'TextureStream' interface:
* [CoreWebView2Environment.CreateTextureStream Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#createtexturestream)
* [CoreWebView2Environment.RenderAdapterLUIDChanged Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#renderadapterluidchanged)
* [CoreWebView2Environment.RenderAdapterLUID Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#renderadapterluid)

The 'TextureStream' interface:
* [CoreWebView2TextureStream Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
    * [CoreWebView2TextureStream.AddAllowedOrigin Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#addallowedorigin)
    * [CoreWebView2TextureStream.CloseTexture Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#closetexture)
    * [CoreWebView2TextureStream.CreateTexture Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#createtexture)
    * [CoreWebView2TextureStream.ErrorReceived Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#errorreceived)
    * [CoreWebView2TextureStream.GetAvailableTexture Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#getavailabletexture)
    * [CoreWebView2TextureStream.Id Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#id)
    * [CoreWebView2TextureStream.PresentTexture Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#presenttexture)
    * [CoreWebView2TextureStream.RemoveAllowedOrigin Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#removeallowedorigin)
    * [CoreWebView2TextureStream.SetD3DDevice Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#setd3ddevice)
    * [CoreWebView2TextureStream.StartRequested Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#startrequested)
    * [CoreWebView2TextureStream.Stop Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#stop)
    * [CoreWebView2TextureStream.Stopped Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#stopped)
    * [CoreWebView2TextureStream.WebTextureReceived Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#webtexturereceived)
    * [CoreWebView2TextureStream.WebTextureStreamStopped Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestream?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#webtexturestreamstopped)

ErrorReceivedEventArgs:
* [CoreWebView2TextureStreamErrorReceivedEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestreamerrorreceivedeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorReceivedEventArgs.Kind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestreamerrorreceivedeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#kind)

WebTextureReceivedEventArgs:
* [CoreWebView2TextureStreamWebTextureReceivedEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestreamwebtexturereceivedeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamWebTextureReceivedEventArgs.WebTexture Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestreamwebtexturereceivedeventargs?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#webtexture)

TextureStream error kind enum:
* [CoreWebView2TextureStreamErrorKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestreamerrorkind?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorKind.CoreWebView2TextureStreamErrorNoVideoTrackStarted Enum Value](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestreamerrorkind?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorKind.CoreWebView2TextureStreamErrorTextureError Enum Value](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestreamerrorkind?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2TextureStreamErrorKind.CoreWebView2TextureStreamErrorTextureInUse Enum Value](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texturestreamerrorkind?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)

The 'Texture' interface that the host writes to so that the Renderer will render on it:
* [CoreWebView2Texture Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texture?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2Texture.Resource Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texture?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#resource)
   * [CoreWebView2Texture.Timestamp Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2texture?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#timestamp)

The received 'WebTexture' interface that the Renderer writes to so that the host will read on it:
* [CoreWebView2WebTexture Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webtexture?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2WebTexture.Resource Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webtexture?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#resource)
   * [CoreWebView2WebTexture.Timestamp Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webtexture?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#timestamp)

##### [Win32/C++](#tab/win32cpp)

The 'Environment' interface that returns the 'TextureStream' interface:
* [ICoreWebView2ExperimentalEnvironment12 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment12?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalEnvironment12::CreateTextureStream](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment12?view=webview2-1.0.1724-prerelease&preserve-view=true#createtexturestream)
   * [ICoreWebView2ExperimentalEnvironment12::RenderAdapterLUID (get)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment12?view=webview2-1.0.1724-prerelease&preserve-view=true#get_renderadapterluid)
   * [ICoreWebView2ExperimentalEnvironment12::RenderAdapterLUIDChanged (add, remove)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment12?view=webview2-1.0.1724-prerelease&preserve-view=true#add_renderadapterluidchanged)

The 'TextureStream' interface:
* [ICoreWebView2ExperimentalTextureStream interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalTextureStream::add_ErrorReceived](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#add_errorreceived)
   * [ICoreWebView2ExperimentalTextureStream::add_StartRequested](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#add_startrequested)
   * [ICoreWebView2ExperimentalTextureStream::add_Stopped](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#add_stopped)
   * [ICoreWebView2ExperimentalTextureStream::add_WebTextureReceived](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#add_webtexturereceived)
   * [ICoreWebView2ExperimentalTextureStream::add_WebTextureStreamStopped](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#add_webtexturestreamstopped)
   * [ICoreWebView2ExperimentalTextureStream::AddAllowedOrigin](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#addallowedorigin)
   * [ICoreWebView2ExperimentalTextureStream::CloseTexture](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#closetexture)
   * [ICoreWebView2ExperimentalTextureStream::CreateTexture](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#createtexture)
   * [ICoreWebView2ExperimentalTextureStream::get_Id](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#get_id)
   * [ICoreWebView2ExperimentalTextureStream::GetAvailableTexture](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#getavailabletexture)
   * [ICoreWebView2ExperimentalTextureStream::PresentTexture](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#presenttexture)
   * [ICoreWebView2ExperimentalTextureStream::remove_ErrorReceived](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#remove_errorreceived)
   * [ICoreWebView2ExperimentalTextureStream::remove_StartRequested](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#remove_startrequested)
   * [ICoreWebView2ExperimentalTextureStream::remove_Stopped](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#remove_stopped)
   * [ICoreWebView2ExperimentalTextureStream::remove_WebTextureReceived](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#remove_webtexturereceived)
   * [ICoreWebView2ExperimentalTextureStream::remove_WebTextureStreamStopped](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#remove_webtexturestreamstopped)
   * [ICoreWebView2ExperimentalTextureStream::RemoveAllowedOrigin](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#removeallowedorigin)
   * [ICoreWebView2ExperimentalTextureStream::SetD3DDevice](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#setd3ddevice)
   * [ICoreWebView2ExperimentalTextureStream::Stop](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestream?view=webview2-1.0.1724-prerelease&preserve-view=true#stop)

Supplemental 'TextureStream*' interfaces:
* [ICoreWebView2ExperimentalTextureStreamStartRequestedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamstartrequestedeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalTextureStreamStoppedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamstoppedeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalTextureStreamErrorReceivedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamerrorreceivedeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalTextureStreamErrorReceivedEventArgs interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamerrorreceivedeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalTextureStreamErrorReceivedEventArgs::get_Kind](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamerrorreceivedeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#get_kind)
   * [ICoreWebView2ExperimentalTextureStreamErrorReceivedEventArgs::get_Texture](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamerrorreceivedeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#get_texture)
* [ICoreWebView2ExperimentalTextureStreamWebTextureReceivedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamwebtexturereceivedeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalTextureStreamWebTextureReceivedEventArgs interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamwebtexturereceivedeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalTextureStreamWebTextureReceivedEventArgs::get_WebTexture](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamwebtexturereceivedeventargs?view=webview2-1.0.1724-prerelease&preserve-view=true#get_webtexture)
* [ICoreWebView2ExperimentalTextureStreamWebTextureStreamStoppedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexturestreamwebtexturestreamstoppedeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)

TextureStream error kind enum:
* [COREWEBVIEW2_TEXTURE_STREAM_ERROR_KIND Enum](/microsoft-edge/webview2/reference/win32/webview2experimental-idl?view=webview2-1.0.1724-prerelease&preserve-view=true#corewebview2_texture_stream_error_kind)

Other interfaces ('RenderAdapter'):
* [ICoreWebView2ExperimentalRenderAdapterLUIDChangedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalrenderadapterluidchangedeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)

The 'Texture' interface that the host writes to so that the Renderer will render on it:
* [ICoreWebView2ExperimentalTexture interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexture?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalTexture::get_Handle](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexture?view=webview2-1.0.1724-prerelease&preserve-view=true#get_handle)
   * [ICoreWebView2ExperimentalTexture::get_Resource](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexture?view=webview2-1.0.1724-prerelease&preserve-view=true#get_resource)
   * [ICoreWebView2ExperimentalTexture::get_Timestamp](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexture?view=webview2-1.0.1724-prerelease&preserve-view=true#get_timestamp)
   * [ICoreWebView2ExperimentalTexture::put_Timestamp](/microsoft-edge/webview2/reference/win32/icorewebview2experimentaltexture?view=webview2-1.0.1724-prerelease&preserve-view=true#put_timestamp)

The received 'WebTexture' interface that the Renderer writes to so that the host will read on it:
* [ICoreWebView2ExperimentalWebTexture interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebtexture?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalWebTexture::get_Handle](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebtexture?view=webview2-1.0.1724-prerelease&preserve-view=true#get_handle)
   * [ICoreWebView2ExperimentalWebTexture::get_Resource](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebtexture?view=webview2-1.0.1724-prerelease&preserve-view=true#get_resource)
   * [ICoreWebView2ExperimentalWebTexture::get_Timestamp](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebtexture?view=webview2-1.0.1724-prerelease&preserve-view=true#get_timestamp)

---


<!-- ------------------------------ -->
*  Added support for profile management: custom data partition, cookie manager and profile deletion:

##### [.NET/C#](#tab/dotnetcsharp)

Added support for custom data partition:
* [CoreWebView2.CustomDataPartitionId Property](/dotnet/api/microsoft.web.webview2.core.corewebview2.customdatapartitionid?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
* [CoreWebView2Profile.ClearCustomDataPartitionAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.clearcustomdatapartitionasync?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

Added support for cookie manager:
* [CoreWebView2Profile.CookieManager Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.cookiemanager?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

Add support for managing profile deletion:
* [CoreWebView2Profile.Delete Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.delete?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
* [CoreWebView2Profile.Deleted Event](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.deleted?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

Added support for custom data partition:
* [CoreWebView2.CustomDataPartitionId Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#customdatapartitionid)
* [CoreWebView2Profile.ClearCustomDataPartitionAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#clearcustomdatapartitionasync)

Added support for cookie manager:
* [CoreWebView2Profile.CookieManager Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#cookiemanager)

Add support for managing profile deletion:
* [CoreWebView2Profile.Delete Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#delete)
* [CoreWebView2Profile.Deleted Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#deleted)

##### [Win32/C++](#tab/win32cpp)

Added support for custom data partition:
* [ICoreWebView2Experimental20 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimental20?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2Experimental20::get_CustomDataPartitionId](/microsoft-edge/webview2/reference/win32/icorewebview2experimental20?view=webview2-1.0.1724-prerelease&preserve-view=true#get_customdatapartitionid)
   * [ICoreWebView2Experimental20::put_CustomDataPartitionId](/microsoft-edge/webview2/reference/win32/icorewebview2experimental20?view=webview2-1.0.1724-prerelease&preserve-view=true#put_customdatapartitionid)
* [ICoreWebView2ExperimentalProfile7 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile7?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalProfile7::ClearCustomDataPartition](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile7?view=webview2-1.0.1724-prerelease&preserve-view=true#clearcustomdatapartition)
* [ICoreWebView2ExperimentalClearCustomDataPartitionCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalclearcustomdatapartitioncompletedhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)

Added support for cookie manager:
* [ICoreWebView2ExperimentalProfile8 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile8?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalProfile8::get_CookieManager](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile8?view=webview2-1.0.1724-prerelease&preserve-view=true#get_cookiemanager)

Add support for managing profile deletion:
* [ICoreWebView2ExperimentalProfile10 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile10?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalProfile10::Delete](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile10?view=webview2-1.0.1724-prerelease&preserve-view=true#delete)
   * [ICoreWebView2ExperimentalProfile10::add_Deleted](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile10?view=webview2-1.0.1724-prerelease&preserve-view=true#add_deleted)
   * [ICoreWebView2ExperimentalProfile10::remove_Deleted](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile10?view=webview2-1.0.1724-prerelease&preserve-view=true#remove_deleted)
* [ICoreWebView2ExperimentalProfileDeletedEventHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofiledeletedeventhandler?view=webview2-1.0.1724-prerelease&preserve-view=true)

---


<!-- ------------------------------ -->
###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ------------------------------ -->
*  Managing smartscreen API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2Settings Class](/dotnet/api/microsoft.web.webview2.core.corewebview2settings?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2Settings.IsReputationCheckingRequired Property](/dotnet/api/microsoft.web.webview2.core.corewebview2settings.isreputationcheckingrequired?view=webview2-dotnet-1.0.1724-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2Settings Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2settings?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true)
   * [CoreWebView2Settings.IsReputationCheckingRequired Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2settings?view=webview2-winrt-1.0.1724-prerelease&preserve-view=true#isreputationcheckingrequired)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Settings8 interface](/microsoft-edge/webview2/reference/win32/icorewebview2settings8?view=webview2-1.0.1724-prerelease&preserve-view=true)
   * [ICoreWebView2Settings8::get_IsReputationCheckingRequired](/microsoft-edge/webview2/reference/win32/icorewebview2settings8?view=webview2-1.0.1724-prerelease&preserve-view=true#get_isreputationcheckingrequired)
   * [ICoreWebView2Settings8::put_IsReputationCheckingRequired](/microsoft-edge/webview2/reference/win32/icorewebview2settings8?view=webview2-1.0.1724-prerelease&preserve-view=true#put_isreputationcheckingrequired)

---


<!-- ------------------------------ -->
###### Bug fixes

*  Fixed a bug in 'PrintAsync' and 'PrintToPdfStreamAsync' that throws an exception when print settings are null.
*  Improved handling of apps running elevated.  (Runtime-only)
*  Added support for window management permission kind.  (SDK and Runtime)
*  Reliability improvement.  (Runtime-only)


<!-- ====================================================================== -->
Release Date: February 15, 2023`,
	},
	"1.0.1587.40": {
		SDKVersion:     "1.0.1587.40",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10158740",
		RuntimeVersion: "110.0.1587.40",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ---------- -->
###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
*  Additional options used to create a WebView2 Environment to manage custom scheme registration:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2CustomSchemeRegistration Class](/dotnet/api/microsoft.web.webview2.core.corewebview2customschemeregistration?view=webview2-dotnet-1.0.1587.40&preserve-view=true)
* [CoreWebView2EnvironmentOptions Class](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions?view=webview2-dotnet-1.0.1587.40&preserve-view=true)
   * [CoreWebView2EnvironmentOptions.CustomSchemeRegistrations Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.customschemeregistrations?view=webview2-dotnet-1.0.1587.40&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2CustomSchemeRegistration Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2customschemeregistration?view=webview2-winrt-1.0.1587.40&preserve-view=true)
* [CoreWebView2EnvironmentOptions Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1587.40&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2CustomSchemeRegistration interface](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1587.40&preserve-view=true)
   * [ICoreWebView2CustomSchemeRegistration::GetAllowedOrigins method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1587.40&preserve-view=true#getallowedorigins)
   * [ICoreWebView2CustomSchemeRegistration::SetAllowedOrigins method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1587.40&preserve-view=true#setallowedorigins)
   * [ICoreWebView2CustomSchemeRegistration::get_HasAuthorityComponent method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1587.40&preserve-view=true#get_hasauthoritycomponent)
   * [ICoreWebView2CustomSchemeRegistration::put_HasAuthorityComponent method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1587.40&preserve-view=true#put_hasauthoritycomponent)
   * [ICoreWebView2CustomSchemeRegistration::get_SchemeName method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1587.40&preserve-view=true#get_schemename)<!--no put-->
   * [ICoreWebView2CustomSchemeRegistration::get_TreatAsSecure method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1587.40&preserve-view=true#get_treatassecure)
   * [ICoreWebView2CustomSchemeRegistration::put_TreatAsSecure method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1587.40&preserve-view=true#put_treatassecure)
* [ICoreWebView2EnvironmentOptions4 interface](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4?view=webview2-1.0.1587.40&preserve-view=true)
   * [ICoreWebView2EnvironmentOptions4::GetCustomSchemeRegistrations method](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4?view=webview2-1.0.1587.40&preserve-view=true#getcustomschemeregistrations)
   * [ICoreWebView2EnvironmentOptions4::SetCustomSchemeRegistrations method](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4?view=webview2-1.0.1587.40&preserve-view=true#setcustomschemeregistrations)

---


<!-- ====================================================================== -->
Release Date: February 15, 2023`,
	},
	"1.0.1671-prerelease": {
		SDKVersion:     "1.0.1671-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101671-prerelease",
		RuntimeVersion: "112.0.1671.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ---------- -->
###### Experimental features

*  Added support for the Experimental File API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2File Class](/dotnet/api/microsoft.web.webview2.core.corewebview2file?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2File.Path Property](/dotnet/api/microsoft.web.webview2.core.corewebview2file.path?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
* [CoreWebView2WebMessageReceivedEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2webmessagereceivedeventargs?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2File Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2file?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2File.Path Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2file?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#path)
* [CoreWebView2WebMessageReceivedEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webmessagereceivedeventargs?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalFile interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalfile?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalFile::get_Path method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalfile?view=webview2-1.0.1671-prerelease&preserve-view=true#get_path)
* [ICoreWebView2ExperimentalWebMessageReceivedEventArgs interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebmessagereceivedeventargs?view=webview2-1.0.1671-prerelease&preserve-view=true)

Also added support for Experimental Object Collection View API:

* [ICoreWebView2ExperimentalObjectCollectionView interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalobjectcollectionview?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalObjectCollectionView::get_Count method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalobjectcollectionview?view=webview2-1.0.1671-prerelease&preserve-view=true#get_count)
   * [ICoreWebView2ExperimentalObjectCollectionView::GetValueAtIndex method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalobjectcollectionview?view=webview2-1.0.1671-prerelease&preserve-view=true#getvalueatindex)

The above interface is currently being used for:

* [ICoreWebView2ExperimentalWebMessageReceivedEventArgs::get_AdditionalObjects method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebmessagereceivedeventargs?view=webview2-1.0.1671-prerelease&preserve-view=true#get_additionalobjects)

---


<!-- ---------- -->
###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ------------------------------ -->
*  The SharedBuffer API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2 Class](/dotnet/api/microsoft.web.webview2.core.corewebview2?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2.PostSharedBufferToScript Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.postsharedbuffertoscript?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
* [CoreWebView2Environment Class](/dotnet/api/microsoft.web.webview2.core.corewebview2environment?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2Environment.CreateSharedBuffer Method](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.createsharedbuffer?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
* [CoreWebView2Frame Class](/dotnet/api/microsoft.web.webview2.core.corewebview2frame?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2Frame.PostSharedBufferToScript Method](/dotnet/api/microsoft.web.webview2.core.corewebview2frame.postsharedbuffertoscript?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
* [CoreWebView2SharedBuffer Class](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2SharedBuffer.Buffer Property](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.buffer?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2SharedBuffer.FileMappingHandle Property](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.filemappinghandle?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2SharedBuffer.Size Property](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.size?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2SharedBuffer.Close Method](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.close?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2SharedBuffer.OpenStream Method](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer.openstream?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2 Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2.PostSharedBufferToScript Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#postsharedbuffertoscript)
* [CoreWebView2Environment Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2Environment.CreateSharedBuffer Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#createsharedbuffer)
* [CoreWebView2Frame Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frame?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2Frame.PostSharedBufferToScript Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frame?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#postsharedbuffertoscript)
* [CoreWebView2SharedBuffer Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2SharedBuffer.Buffer Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#buffer)
   * [CoreWebView2SharedBuffer.Size Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#size)
   * [CoreWebView2SharedBuffer.Close Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#close)
   * [CoreWebView2SharedBuffer.OpenStream Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#openstream)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_17 interface](/microsoft-edge/webview2/reference/win32/icorewebview2_17?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2_17::PostSharedBufferToScript method](/microsoft-edge/webview2/reference/win32/icorewebview2_17?view=webview2-1.0.1671-prerelease&preserve-view=true#postsharedbuffertoscript)
* [ICoreWebView2Environment12 interface](/microsoft-edge/webview2/reference/win32/icorewebview2environment12?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2Environment12::CreateSharedBuffer method](/microsoft-edge/webview2/reference/win32/icorewebview2environment12?view=webview2-1.0.1671-prerelease&preserve-view=true#createsharedbuffer)
* [ICoreWebView2Frame4 interface](/microsoft-edge/webview2/reference/win32/icorewebview2frame4?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2Frame4::PostSharedBufferToScript method](/microsoft-edge/webview2/reference/win32/icorewebview2frame4?view=webview2-1.0.1671-prerelease&preserve-view=true#postsharedbuffertoscript)
* [ICoreWebView2SharedBuffer interface](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2SharedBuffer::OpenStream method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1671-prerelease&preserve-view=true#openstream)
   * [ICoreWebView2SharedBuffer::Close method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1671-prerelease&preserve-view=true#close)
   * [ICoreWebView2SharedBuffer::get_Size method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1671-prerelease&preserve-view=true#get_size)
   * [ICoreWebView2SharedBuffer::get_Buffer method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1671-prerelease&preserve-view=true#get_buffer)
   * [ICoreWebView2SharedBuffer::get_FileMappingHandle method](/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer?view=webview2-1.0.1671-prerelease&preserve-view=true#get_filemappinghandle)

---


<!-- ------------------------------ -->
*  The Permission API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2PermissionSetting Class](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionkind?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionOrigin Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionorigin?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionState Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionstate?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
* [CoreWebView2Profile Class](/dotnet/api/microsoft.web.webview2.core.corewebview2profile?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2Profile.GetNonDefaultPermissionSettingsAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.getnondefaultpermissionsettingsasync?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2Profile.SetPermissionStateAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.setpermissionstateasync?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2PermissionSetting Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#permissionkind)
   * [CoreWebView2PermissionSetting.PermissionOrigin Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#permissionorigin)
   * [CoreWebView2PermissionSetting.PermissionState Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#permissionstate)
* [CoreWebView2Profile Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2Profile.GetNonDefaultPermissionSettingsAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#getnondefaultpermissionsettingsasync)
   * [CoreWebView2Profile.SetPermissionStateAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#setpermissionstateasync)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2getnondefaultpermissionsettingscompletedhandler?view=webview2-1.0.1671-prerelease&preserve-view=true)
* [ICoreWebView2PermissionSetting interface](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2PermissionSetting::get_PermissionKind method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting?view=webview2-1.0.1671-prerelease&preserve-view=true#get_permissionkind)
   * [ICoreWebView2PermissionSetting::get_PermissionOrigin method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting?view=webview2-1.0.1671-prerelease&preserve-view=true#get_permissionorigin)
   * [ICoreWebView2PermissionSetting::get_PermissionState method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting?view=webview2-1.0.1671-prerelease&preserve-view=true#get_permissionstate)
* [ICoreWebView2PermissionSettingCollectionView interface](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2PermissionSettingCollectionView::GetValueAtIndex method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview?view=webview2-1.0.1671-prerelease&preserve-view=true#getvalueatindex)
   * [ICoreWebView2PermissionSettingCollectionView::get_Count method](/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview?view=webview2-1.0.1671-prerelease&preserve-view=true#get_count)
* [ICoreWebView2Profile4 interface](/microsoft-edge/webview2/reference/win32/icorewebview2profile4?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2Profile4::SetPermissionState method](/microsoft-edge/webview2/reference/win32/icorewebview2profile4?view=webview2-1.0.1671-prerelease&preserve-view=true#setpermissionstate)
   * [ICoreWebView2Profile4::GetNonDefaultPermissionSettings method](/microsoft-edge/webview2/reference/win32/icorewebview2profile4?view=webview2-1.0.1671-prerelease&preserve-view=true#getnondefaultpermissionsettings)
* [ICoreWebView2SetPermissionStateCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2setpermissionstatecompletedhandler?view=webview2-1.0.1671-prerelease&preserve-view=true)

---


<!-- ------------------------------ -->
*  The ScriptLocale API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2ControllerOptions Class](/dotnet/api/microsoft.web.webview2.core.corewebview2controlleroptions?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2ControllerOptions.ScriptLocale Property](/dotnet/api/microsoft.web.webview2.core.corewebview2controlleroptions.scriptlocale?view=webview2-dotnet-1.0.1671-prerelease&preserve-view=true)

Previous name in 1619-prerelease:
* [CoreWebView2ControllerOptions.LocaleRegion Property](/dotnet/api/microsoft.web.webview2.core.corewebview2controlleroptions.localeregion?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)<!--keep 1619-->

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2ControllerOptions Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2controlleroptions?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true)
   * [CoreWebView2ControllerOptions.ScriptLocale Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2controlleroptions?view=webview2-winrt-1.0.1671-prerelease&preserve-view=true#scriptlocale)

Previous name in 1619-prerelease:
* [CoreWebView2ControllerOptions.LocaleRegion Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2controlleroptions?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#localeregion)<!--keep 1619-->

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ControllerOptions2 interface](/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2?view=webview2-1.0.1671-prerelease&preserve-view=true)
   * [ICoreWebView2ControllerOptions2::get_ScriptLocale method](/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2?view=webview2-1.0.1671-prerelease&preserve-view=true#get_scriptlocale)
   * [ICoreWebView2ControllerOptions2::put_ScriptLocale method](/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2?view=webview2-1.0.1671-prerelease&preserve-view=true#put_scriptlocale)

Previous name in 1619-prerelease:
* [ICoreWebView2ExperimentalControllerOptions::get_LocaleRegion method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontrolleroptions?view=webview2-1.0.1619-prerelease&preserve-view=true#get_localeregion)<!--keep 1619-->
* [ICoreWebView2ExperimentalControllerOptions::put_LocaleRegion method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontrolleroptions?view=webview2-1.0.1619-prerelease&preserve-view=true#put_localeregion)<!--keep 1619-->

---


<!-- ---------- -->
###### Bug fixes

*  Fixed a bug where WebView2 was not closing properly when a 'BeforeUnload' event was received.  (Runtime-only)  ([Issue #2677](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2677))
*  In the 'DownloadStarting' event, the 'ResultFilePath' previously wasn't showing the correct download location for UWP applications when the 'DownloadStarting' event handler was attached.  This has been fixed; the correct 'ResultFilePath' is now shown.
*  Fixed a bug where 'System.ArgumentException' was thrown when a call to the 'HostObject' method returns a non-generic task.  ([Issue #2787](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2787))
*  Fixed an issue in the 'SharedBuffer' API where the stream object didn't work well with 'StreamWriter'.  (Runtime-only)  ([Issue #3108](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3108))
*  DOM speech-synthesis APIs, such as 'SpeechSynthesis.getVoices()', will now work in UWP apps.  (Runtime-only)
*  Fixed a crash that occurred on frame destruction.  (Runtime-only)  ([Issue #3062](https://github.com/MicrosoftEdge/WebView2Feedback/issues/3062))
*  Fixed a bug where the app crashes when trying to call 'CreateWebResourceResponse' with a 'null' 'reason' phrase.  (Runtime-only)
*  The 'CoreWebView2.AddHostObjectToScript' option 'chrome.webview.hostObjects.options.ignoreMemberNotFoundError' now works in non-English locales.  (Runtime-only)
*  Fully enabled **Open file** dialog support for elevated apps on Windows 7.
*  Fixed a bug where owned windows were not appearing for UWP.


<!-- ====================================================================== -->
Release Date: January 17, 2023`,
	},
	"1.0.1518.46": {
		SDKVersion:     "1.0.1518.46",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10151846",
		RuntimeVersion: "109.0.1518.46",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ---------- -->
###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
*  The Print API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.PrintAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.printasync?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
* [CoreWebView2.PrintToPdfStreamAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.printtopdfstreamasync?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
* [CoreWebView2.ShowPrintUI Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.showprintui?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
* [CoreWebView2PrintSettings Class](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
   * [CoreWebView2PrintSettings.Collation Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.collation?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
   * [CoreWebView2PrintSettings.ColorMode Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.colormode?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
   * [CoreWebView2PrintSettings.Copies Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.copies?view=webview2-dotnet-1.0.1518.46&preserve-view=true#microsoft-web-webview2-core-corewebview2printsettings-copies)
   * [CoreWebView2PrintSettings.Duplex Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.duplex?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
   * [CoreWebView2PrintSettings.MediaSize Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.mediasize?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
   * [CoreWebView2PrintSettings.PageRanges Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.pageranges?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
   * [CoreWebView2PrintSettings.PagesPerSide Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.pagesperside?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
   * [CoreWebView2PrintSettings.PrinterName Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.printername?view=webview2-dotnet-1.0.1518.46&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.PrintAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1518.46&preserve-view=true#printasync)
* [CoreWebView2.PrintToPdfStreamAsync](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1518.46&preserve-view=true#printtopdfstreamasync)
* [CoreWebView2.ShowPrintUI Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1518.46&preserve-view=true#showprintui)
* [CoreWebView2PrintSettings Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true)
   * [CoreWebView2PrintSettings.Collation Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true#collation)
   * [CoreWebView2PrintSettings.ColorMode Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true#colormode)
   * [CoreWebView2PrintSettings.Copies Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true#copies)
   * [CoreWebView2PrintSettings.Duplex Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true#duplex)
   * [CoreWebView2PrintSettings.MediaSize Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true#mediasize)
   * [CoreWebView2PrintSettings.PageRanges Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true#pageranges)
   * [CoreWebView2PrintSettings.PagesPerSide Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true#pagesperside)
   * [CoreWebView2PrintSettings.PrinterName Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1518.46&preserve-view=true#printername)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_16 interface](/microsoft-edge/webview2/reference/win32/icorewebview2_16?view=webview2-1.0.1518.46&preserve-view=true)
   * [ICoreWebView2_16::Print method](/microsoft-edge/webview2/reference/win32/icorewebview2_16?view=webview2-1.0.1518.46&preserve-view=true#print)
   * [ICoreWebView2_16::PrintToPdfStream method](/microsoft-edge/webview2/reference/win32/icorewebview2_16?view=webview2-1.0.1518.46&preserve-view=true#printtopdfstream)
   * [ICoreWebView2_16::ShowPrintUI method](/microsoft-edge/webview2/reference/win32/icorewebview2_16?view=webview2-1.0.1518.46&preserve-view=true#showprintui)
* [ICoreWebView2PrintCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2printcompletedhandler?view=webview2-1.0.1518.46&preserve-view=true)
* [ICoreWebView2PrintToPdfStreamCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfstreamcompletedhandler?view=webview2-1.0.1518.46&preserve-view=true)
* [ICoreWebView2PrintSettings2 interface](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true)
   * [ICoreWebView2PrintSettings2::Collation property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#get_collation), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#put_collation)
   * [ICoreWebView2PrintSettings2::ColorMode property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#get_colormode), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#put_colormode)
   * [ICoreWebView2PrintSettings2::Copies property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#get_copies), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#put_copies)
   * [ICoreWebView2PrintSettings2::Duplex property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#get_duplex), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#put_duplex)
   * [ICoreWebView2PrintSettings2::MediaSize property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#get_mediasize), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#put_mediasize)
   * [ICoreWebView2PrintSettings2::PageRanges property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#get_pageranges), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#put_pageranges)
   * [ICoreWebView2PrintSettings2::PagesPerSide property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#get_pagesperside), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#put_pagesperside)
   * [ICoreWebView2PrintSettings2::PrinterName property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#get_printername), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1518.46&preserve-view=true#put_printername)

---


<!-- ------------------------------ -->
*  The Custom Crash Reporting API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2EnvironmentOptions.IsCustomCrashReportingEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.iscustomcrashreportingenabled?view=webview2-dotnet-1.0.1518.46&preserve-view=true)
* [CoreWebView2Environment.FailureReportFolderPath Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.failurereportfolderpath?view=webview2-dotnet-1.0.1518.46&preserve-view=true#microsoft-web-webview2-core-corewebview2environment-failurereportfolderpath)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2EnvironmentOptions.IsCustomCrashReportingEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1518.46&preserve-view=true#iscustomcrashreportingenabled)
* [CoreWebView2Environment.FailureReportFolderPath Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1518.46&preserve-view=true#failurereportfolderpath)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions3::IsCustomCrashReportingEnabled property (get](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3?view=webview2-1.0.1518.46&preserve-view=true#get_iscustomcrashreportingenabled), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3?view=webview2-1.0.1518.46&preserve-view=true#put_iscustomcrashreportingenabled)
* [ICoreWebView2Environment11::FailureReportFolderPath property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2environment11?view=webview2-1.0.1518.46&preserve-view=true#get_failurereportfolderpath)<!--no put-->

---


<!-- ====================================================================== -->
Release Date: January 19, 2023`,
	},
	"1.0.1619-prerelease": {
		SDKVersion:     "1.0.1619-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101619-prerelease",
		RuntimeVersion: "111.0.1619.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General

<!-- ---------- -->
###### Experimental features


<!-- ------------------------------ -->
*  Added support for the Permission management API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2PermissionRequestedEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionrequestedeventargs?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true&preserve-view=true)
   * [CoreWebView2PermissionRequestedEventArgs.SavesInProfile Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionrequestedeventargs.savesinprofile?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
* [CoreWebView2Profile Class](/dotnet/api/microsoft.web.webview2.core.corewebview2profile?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2Profile.GetNonDefaultPermissionSettingsAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.getnondefaultpermissionsettingsasync?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2Profile.SetPermissionStateAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.setpermissionstateasync?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
* [CoreWebView2PermissionSetting Class](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionKind Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionkind?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2PermissionKind Enum](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionkind?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
      * 'MultipleAutomaticDownloads'
      * 'FileReadWrite'
      * 'Autoplay'
      * 'LocalFonts'
      * 'MidiSystemExclusiveMessageAccess'
   * [CoreWebView2PermissionSetting.PermissionOrigin Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionorigin?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionState Property](/dotnet/api/microsoft.web.webview2.core.corewebview2permissionsetting.permissionstate?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2PermissionRequestedEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionrequestedeventargs?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2PermissionRequestedEventArgs.SavesInProfile Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionrequestedeventargs?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#savesinprofile)
* [CoreWebView2Profile Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2Profile.GetNonDefaultPermissionSettingsAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#getnondefaultpermissionsettingsasync)
   * [CoreWebView2Profile.SetPermissionStateAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#setpermissionstateasync)
* [CoreWebView2PermissionSetting Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2PermissionSetting.PermissionKind Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#permissionkind)
   * [CoreWebView2PermissionKind Enum](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionkind?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)
      * 'MultipleAutomaticDownloads'
      * 'FileReadWrite'
      * 'Autoplay'
      * 'LocalFonts'
      * 'MidiSystemExclusiveMessageAccess'
   * [CoreWebView2PermissionSetting.PermissionOrigin Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#permissionorigin)
   * [CoreWebView2PermissionSetting.PermissionState Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2permissionsetting?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#permissionstate)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalPermissionRequestedEventArgs3 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionrequestedeventargs3?view=webview2-1.0.1619-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalPermissionRequestedEventArgs3::SavesInProfile property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionrequestedeventargs3?view=webview2-1.0.1619-prerelease&preserve-view=true#get_savesinprofile), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionrequestedeventargs3?view=webview2-1.0.1619-prerelease&preserve-view=true#put_savesinprofile)
* [ICoreWebView2ExperimentalSetPermissionStateCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsetpermissionstatecompletedhandler?view=webview2-1.0.1619-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalGetNonDefaultPermissionSettingsCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalgetnondefaultpermissionsettingscompletedhandler?view=webview2-1.0.1619-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalProfile6 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile6?view=webview2-1.0.1619-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalProfile6::GetNonDefaultPermissionSettings method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile6?view=webview2-1.0.1619-prerelease&preserve-view=true#getnondefaultpermissionsettings)
   * [ICoreWebView2ExperimentalProfile6::SetPermissionState method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile6?view=webview2-1.0.1619-prerelease&preserve-view=true#setpermissionstate)
* [ICoreWebView2ExperimentalPermissionSettingCollectionView interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionsettingcollectionview?view=webview2-1.0.1619-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalPermissionSettingCollectionView::Count property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionsettingcollectionview?view=webview2-1.0.1619-prerelease&preserve-view=true#get_count)<!--no put-->
   * [ICoreWebView2ExperimentalPermissionSettingCollectionView::GetValueAtIndex method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionsettingcollectionview?view=webview2-1.0.1619-prerelease&preserve-view=true#getvalueatindex)
* [ICoreWebView2ExperimentalPermissionSetting interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionsetting?view=webview2-1.0.1619-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalPermissionSetting::PermissionKind property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionsetting?view=webview2-1.0.1619-prerelease&preserve-view=true#get_permissionkind)<!--no put-->
   * [COREWEBVIEW2_PERMISSION_KIND Enum](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-1.0.1619-prerelease&preserve-view=true#corewebview2_permission_kind)
      * 'COREWEBVIEW2_PERMISSION_KIND_MULTIPLE_AUTOMATIC_DOWNLOADS'
      * 'COREWEBVIEW2_PERMISSION_KIND_FILE_READ_WRITE'
      * 'COREWEBVIEW2_PERMISSION_KIND_AUTOPLAY'
      * 'COREWEBVIEW2_PERMISSION_KIND_LOCAL_FONTS'
      * 'COREWEBVIEW2_PERMISSION_KIND_MIDI_SYSTEM_EXCLUSIVE_MESSAGE_ACCESS'
   * [ICoreWebView2ExperimentalPermissionSetting::PermissionOrigin property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionsetting?view=webview2-1.0.1619-prerelease&preserve-view=true#get_permissionorigin)<!--no put-->
   * [ICoreWebView2ExperimentalPermissionSetting::PermissionState property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalpermissionsetting?view=webview2-1.0.1619-prerelease&preserve-view=true#get_permissionstate)<!--no put-->

---


<!-- ------------------------------ -->
*  Added support for API to disable back and forward navigation:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2NavigationStartingEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2navigationstartingeventargs?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2NavigationStartingEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2navigationstartingeventargs?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalNavigationStartingEventArgs2 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalnavigationstartingeventargs2?view=webview2-1.0.1619-prerelease&preserve-view=true)

---


<!-- ---------- -->
###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ------------------------------ -->
*  The Custom Scheme Registration API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2EnvironmentOptions Class](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2EnvironmentOptions.CustomSchemeRegistrations Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.customschemeregistrations?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
* [CoreWebView2CustomSchemeRegistration Class](/dotnet/api/microsoft.web.webview2.core.corewebview2customschemeregistration?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2EnvironmentOptions Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)
* [CoreWebView2CustomSchemeRegistration Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2customschemeregistration?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions4 interface](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4?view=webview2-1.0.1619-prerelease&preserve-view=true)
   * [ICoreWebView2EnvironmentOptions4::GetCustomSchemeRegistrations method](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4?view=webview2-1.0.1619-prerelease&preserve-view=true#getcustomschemeregistrations)
   * [ICoreWebView2EnvironmentOptions4::SetCustomSchemeRegistrations method](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4?view=webview2-1.0.1619-prerelease&preserve-view=true#setcustomschemeregistrations)
* [ICoreWebView2CustomSchemeRegistration interface](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1619-prerelease&preserve-view=true)
   * [ICoreWebView2CustomSchemeRegistration::GetAllowedOrigins method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1619-prerelease&preserve-view=true#getallowedorigins)
   * [ICoreWebView2CustomSchemeRegistration::SetAllowedOrigins method](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1619-prerelease&preserve-view=true#setallowedorigins)
   * [ICoreWebView2CustomSchemeRegistration::HasAuthorityComponent property (get](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1619-prerelease&preserve-view=true#get_hasauthoritycomponent), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1619-prerelease&preserve-view=true#put_hasauthoritycomponent)
   * [ICoreWebView2CustomSchemeRegistration::SchemeName property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1619-prerelease&preserve-view=true#get_schemename)<!--no put-->
   * [ICoreWebView2CustomSchemeRegistration::TreatAsSecure property (get](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1619-prerelease&preserve-view=true#get_treatassecure), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration?view=webview2-1.0.1619-prerelease&preserve-view=true#put_treatassecure)

---


<!-- ------------------------------ -->
*  The Tracking Prevention API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2EnvironmentOptions Class](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2EnvironmentOptions.EnableTrackingPrevention Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.enabletrackingprevention?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
* [CoreWebView2Profile Class](/dotnet/api/microsoft.web.webview2.core.corewebview2profile?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2Profile.PreferredTrackingPreventionLevel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.preferredtrackingpreventionlevel?view=webview2-dotnet-1.0.1619-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2EnvironmentOptions Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2EnvironmentOptions.EnableTrackingPrevention Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#enabletrackingprevention)
* [CoreWebView2Profile Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true)
   * [CoreWebView2Profile.PreferredTrackingPreventionLevel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1619-prerelease&preserve-view=true#preferredtrackingpreventionlevel)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions5 interface](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5?view=webview2-1.0.1619-prerelease&preserve-view=true)
   * [ICoreWebView2EnvironmentOptions5::EnableTrackingPrevention property (get](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5?view=webview2-1.0.1619-prerelease&preserve-view=true#get_enabletrackingprevention), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5?view=webview2-1.0.1619-prerelease&preserve-view=true#put_enabletrackingprevention)
* [ICoreWebView2Profile3 interface](/microsoft-edge/webview2/reference/win32/icorewebview2profile3?view=webview2-1.0.1619-prerelease&preserve-view=true)
   * [ICoreWebView2Profile3::PreferredTrackingPreventionLevel property (get](/microsoft-edge/webview2/reference/win32/icorewebview2profile3?view=webview2-1.0.1619-prerelease&preserve-view=true#get_preferredtrackingpreventionlevel), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2profile3?view=webview2-1.0.1619-prerelease&preserve-view=true#put_preferredtrackingpreventionlevel)

---


<!-- ---------- -->
###### Bug fixes

*  Disabled **Open link as Profile** in the WebView2 context menu.
*  Fixed post data missing in form submit with Ctrl-click. ([Issue #2652](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2652))
*  Fixed a bug where the user is not able to get the custom context menu on PDF Viewer. ([Issue #2607](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2607))
*  Fixed a bug where the entire toolbar is blank when simultaneously hiding the **Bookmarks**, **Search**, and **PageSelector** buttons. ([Issue #2866](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2866))
*  Fixed a bug where the app crashes when trying to move focus to WebView2 when it is disabled.
*  Fixed drag and drop within the WebView2 for composition-hosted WebViews.
*  Removed read-aloud icon in address bar in a WebView2 popup window.
*  Fixed unexpected items in the context menu of popup windows in WebView2.


<!-- ====================================================================== -->
Release Date: December 12, 2022`,
	},
	"1.0.1462.37": {
		SDKVersion:     "1.0.1462.37",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10146237",
		RuntimeVersion: "108.0.1462.37",
		Notes: 			`

<!-- ------------------------------ -->
#### General

This WebView2 SDK release has the same bug fixes as [Bug fixes for 1.0.1466-prerelease](#bug-fixes-for-101466-prerelease).


<!-- ====================================================================== -->
Release Date: December 12, 2022`,
	},
	"1.0.1549-prerelease": {
		SDKVersion:     "1.0.1549-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101549-prerelease",
		RuntimeVersion: "110.0.1549.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General

<!-- ---------- -->
###### Experimental features

*  Added support for the Locale Region API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2ControllerOptions.LocaleRegion Property](/dotnet/api/microsoft.web.webview2.core.corewebview2controlleroptions.localeregion?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2ControllerOptions.LocaleRegion Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2controlleroptions?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#localeregion)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalControllerOptions interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontrolleroptions?view=webview2-1.0.1549-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalControllerOptions::LocaleRegion property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontrolleroptions?view=webview2-1.0.1549-prerelease&preserve-view=true#get_localeregion), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontrolleroptions?view=webview2-1.0.1549-prerelease&preserve-view=true#put_localeregion)

---


*  Added support for the tracking prevention API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2EnvironmentOptions.EnableTrackingPrevention Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.enabletrackingprevention?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
* [CoreWebView2Profile.PreferredTrackingPreventionLevel Property](/dotnet/api/microsoft.web.webview2.core.corewebview2profile.preferredtrackingpreventionlevel?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2EnvironmentOptions.EnableTrackingPrevention Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#enabletrackingprevention)
* [CoreWebView2Profile.PreferredTrackingPreventionLevel Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2profile?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#preferredtrackingpreventionlevel)

##### [Win32/C++](#tab/win32cpp)

*  [ICoreWebView2ExperimentalEnvironmentOptions2 interface](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions2?view=webview2-1.0.1549-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalEnvironmentOptions2::EnableTrackingPrevention property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_enabletrackingprevention), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_enabletrackingprevention)
* [ICoreWebView2ExperimentalProfile5 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile5?view=webview2-1.0.1549-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalProfile5::PreferredTrackingPreventionLevel property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile5?view=webview2-1.0.1549-prerelease&preserve-view=true#get_preferredtrackingpreventionlevel), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile5?view=webview2-1.0.1549-prerelease&preserve-view=true#put_preferredtrackingpreventionlevel)

---


<!-- ---------- -->
###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


<!-- ------------------------------ -->
*  Added support for the Print API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.PrintAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.printasync?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
* [CoreWebView2.PrintToPdfStreamAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.printtopdfstreamasync?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
* [CoreWebView2.ShowPrintUI Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.showprintui?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
* [CoreWebView2PrintSettings Class](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.Collation Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.collation?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.ColorMode Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.colormode?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.Copies Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.copies?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2printsettings-copies)
   * [CoreWebView2PrintSettings.Duplex Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.duplex?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.MediaSize Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.mediasize?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.PageRanges Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.pageranges?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.PagesPerSide Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.pagesperside?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.PrinterName Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.printername?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.PrintAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#printasync)
* [CoreWebView2.PrintToPdfStreamAsync](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#printtopdfstreamasync)
* [CoreWebView2.ShowPrintUI Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#showprintui)
* [CoreWebView2PrintSettings Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.Collation Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#collation)
   * [CoreWebView2PrintSettings.ColorMode Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#colormode)
   * [CoreWebView2PrintSettings.Copies Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#copies)
   * [CoreWebView2PrintSettings.Duplex Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#duplex)
   * [CoreWebView2PrintSettings.MediaSize Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#mediasize)
   * [CoreWebView2PrintSettings.PageRanges Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#pageranges)
   * [CoreWebView2PrintSettings.PagesPerSide Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#pagesperside)
   * [CoreWebView2PrintSettings.PrinterName Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#printername)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_16 interface](/microsoft-edge/webview2/reference/win32/icorewebview2_16?view=webview2-1.0.1549-prerelease&preserve-view=true)
   * [ICoreWebView2_16::Print method](/microsoft-edge/webview2/reference/win32/icorewebview2_16?view=webview2-1.0.1549-prerelease&preserve-view=true#print)
   * [ICoreWebView2_16::PrintToPdfStream method](/microsoft-edge/webview2/reference/win32/icorewebview2_16?view=webview2-1.0.1549-prerelease&preserve-view=true#printtopdfstream)
   * [ICoreWebView2_16::ShowPrintUI method](/microsoft-edge/webview2/reference/win32/icorewebview2_16?view=webview2-1.0.1549-prerelease&preserve-view=true#showprintui)
* [ICoreWebView2PrintCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2printcompletedhandler?view=webview2-1.0.1549-prerelease&preserve-view=true)
* [ICoreWebView2PrintToPdfStreamCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfstreamcompletedhandler?view=webview2-1.0.1549-prerelease&preserve-view=true)
* [ICoreWebView2PrintSettings2 interface](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true)
   * [ICoreWebView2PrintSettings2::Collation property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_collation), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_collation)
   * [ICoreWebView2PrintSettings2::ColorMode property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_colormode), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_colormode)
   * [ICoreWebView2PrintSettings2::Copies property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_copies), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_copies)
   * [ICoreWebView2PrintSettings2::Duplex property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_duplex), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_duplex)
   * [ICoreWebView2PrintSettings2::MediaSize property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_mediasize), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_mediasize)
   * [ICoreWebView2PrintSettings2::PageRanges property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_pageranges), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_pageranges)
   * [ICoreWebView2PrintSettings2::PagesPerSide property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_pagesperside), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_pagesperside)
   * [ICoreWebView2PrintSettings2::PrinterName property (get](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#get_printername), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2?view=webview2-1.0.1549-prerelease&preserve-view=true#put_printername)

---

*  Added support for Custom Crash Reporting API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2EnvironmentOptions.IsCustomCrashReportingEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.iscustomcrashreportingenabled?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true)
* [CoreWebView2Environment.FailureReportFolderPath Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.failurereportfolderpath?view=webview2-dotnet-1.0.1549-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2environment-failurereportfolderpath)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2EnvironmentOptions.IsCustomCrashReportingEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#iscustomcrashreportingenabled)
* [CoreWebView2Environment.FailureReportFolderPath Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1549-prerelease&preserve-view=true#failurereportfolderpath)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2EnvironmentOptions3::IsCustomCrashReportingEnabled property (get](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3?view=webview2-1.0.1549-prerelease&preserve-view=true#get_iscustomcrashreportingenabled), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3?view=webview2-1.0.1549-prerelease&preserve-view=true#put_iscustomcrashreportingenabled)
* [ICoreWebView2Environment11::FailureReportFolderPath property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2environment11?view=webview2-1.0.1549-prerelease&preserve-view=true#get_failurereportfolderpath)<!--no put-->

---


<!-- ---------- -->
###### Bug fixes

*  Fixed some nullptr issues where now some public APIs which take nullptr as input parameters do not crash the WebView2.
*  Disabled "Open link as Profile" in the WebView2 context menu.
*  Fixed bug where the whole tool bar will be blank when hiding Bookmarks, Search, and PageSelector buttons simultaneously. ([Issue #2866](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2866))
*  Fix post data missing in form submit with control click. ([Issue #2652](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2652))
*  Fixed a bug where the user is not able to get the custom context menu on PDF Viewer. ([Issue #2607](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2607))
*  Fix drag/drop within the WebView2 for composition hosted WebViews.
*  Fixed a bug where the app crashes when trying to move focus to WebView2 when it is disabled.
*  Remove read aloud icon in address bar in a WebView2 popup window.
*  Fixed an issue where context menu shows unexpected items in WebView2 popup window.


<!-- ====================================================================== -->
Release Date: October 31, 2022`,
	},
	"1.0.1418.22": {
		SDKVersion:     "1.0.1418.22",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10141822",
		RuntimeVersion: "107.0.1418.22",
		Notes: 			`

<!-- ------------------------------ -->
#### General

This WebView2 SDK release has the same bug fixes as [Bug fixes for 1.0.1414-prerelease](#bug-fixes-for-101414-prerelease).


<!-- ====================================================================== -->
Release Date: October 31, 2022`,
	},
	"1.0.1466-prerelease": {
		SDKVersion:     "1.0.1466-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101466-prerelease",
		RuntimeVersion: "109.0.1466.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ---------- -->
###### Experimental features

* Added support for creating a shared memory based buffer with a specified size:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2SharedBuffer Class](/dotnet/api/microsoft.web.webview2.core.corewebview2sharedbuffer?view=webview2-dotnet-1.0.1466-prerelease&preserve-view=true)
    * 'Buffer'
    * 'FileMappingHandle'
    * 'Size'
    * 'Close'
    * 'Dispose'
    * 'OpenStream'

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2SharedBuffer Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2sharedbuffer?view=webview2-winrt-1.0.1466-prerelease&preserve-view=true)
    * 'Buffer'
    * 'Size'
    * 'Close'
    * 'OpenStream'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalSharedBuffer interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsharedbuffer?view=webview2-1.0.1466-prerelease&preserve-view=true)
    * 'Close'
    * 'get_Buffer'
    * 'get_FileMappingHandle'
    * 'get_Size'
    * 'OpenStream'

---

*  Added support for accessing a shared buffer object from the script of the main frame or 'iframe':

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.PostSharedBufferToScript Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.postsharedbuffertoscript?view=webview2-dotnet-1.0.1466-prerelease&preserve-view=true)
* [CoreWebView2Frame.PostSharedBufferToScript Method](/dotnet/api/microsoft.web.webview2.core.corewebview2frame.postsharedbuffertoscript?view=webview2-dotnet-1.0.1466-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.PostSharedBufferToScript Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1466-prerelease&preserve-view=true#postsharedbuffertoscript)
* [CoreWebView2Frame.PostSharedBufferToScript Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2frame?view=webview2-winrt-1.0.1466-prerelease&preserve-view=true#postsharedbuffertoscript)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental18::PostSharedBufferToScript method](/microsoft-edge/webview2/reference/win32/icorewebview2experimental18?view=webview2-1.0.1466-prerelease&preserve-view=true#postsharedbuffertoscript)
* [ICoreWebView2ExperimentalFrame4::PostSharedBufferToScript method](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe4?view=webview2-1.0.1466-prerelease&preserve-view=true#postsharedbuffertoscript)

---

*  Added support for running JavaScript code from the 'JavaScript' parameter in the current top-level document:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2ScriptException Class](/dotnet/api/microsoft.web.webview2.core.corewebview2scriptexception?view=webview2-dotnet-1.0.1466-prerelease&preserve-view=true)
   * 'ColumnNumber'
   * 'LineNumber'
   * 'Message'
   * 'Name'
   * 'ToJson'

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2ScriptException Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2scriptexception?view=webview2-winrt-1.0.1466-prerelease&preserve-view=true)
   * 'ColumnNumber'
   * 'LineNumber'
   * 'Message'
   * 'Name'
   * 'ToJson'

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalScriptException interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalscriptexception?view=webview2-1.0.1466-prerelease&preserve-view=true)
   * 'get_ColumnNumber'
   * 'get_LineNumber'
   * 'get_Message'
   * 'get_Name'
   * 'get_ToJson'

---


<!-- another section links to here -->
###### Bug fixes for 1.0.1466-prerelease

*   Fixed a bug in which the custom header title in print settings could be wrong. ([Issue #2093](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2093))
*   Display 'AllowedCertificateAuthorities' in 'add_ClientCertificateRequested' event as a 'Base64' string.  (Runtime-only)  ([Issue #2346](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2346))
*   Fixed a bug in which the default footer URI in print settings is missing. ([Issue #2851](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2851))
*   Fixed a bug that produces a null pointer exception that's related to print settings.  (Runtime-only)  ([Issue #2858](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2858))
*   Fixed a bug that reports navigation failure when redirecting to a server that has been configured with Client Certificate Authentication and when the 'WebResourceRequested' event is subscribed to.  (Runtime-only)
*   Fixed an 'AddHostObjectToScript' bug in which, when JavaScript calls an async method and then a synchronous method, the async method call might fail.


<!-- ====================================================================== -->
Release Date: October 11, 2022`,
	},
	"1.0.1370.28": {
		SDKVersion:     "1.0.1370.28",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10137028",
		RuntimeVersion: "106.0.1370.28",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ---------- -->
###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


<!-- ------------------------------ -->
*  The drag and drop API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2CompositionController.DragLeave Method](/dotnet/api/microsoft.web.webview2.core.corewebview2compositioncontroller.dragleave?view=webview2-dotnet-1.0.1370.28&preserve-view=true#microsoft-web-webview2-core-corewebview2compositioncontroller-dragleave)

##### [WinRT/C#](#tab/winrtcsharp)

* [ICoreWebView2CompositionControllerInterop2.DragEnter Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2compositioncontrollerinterop2?view=webview2-winrt-1.0.1370.28&preserve-view=true#dragenter)
* [ICoreWebView2CompositionControllerInterop2.DragLeave Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2compositioncontrollerinterop2?view=webview2-winrt-1.0.1370.28&preserve-view=true#dragleave)
* [ICoreWebView2CompositionControllerInterop2.DragOver Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2compositioncontrollerinterop2?view=webview2-winrt-1.0.1370.28&preserve-view=true#dragover)
* [ICoreWebView2CompositionControllerInterop2.Drop Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2compositioncontrollerinterop2?view=webview2-winrt-1.0.1370.28&preserve-view=true#drop)
* [CoreWebView2CompositionController.DragLeave Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2compositioncontroller?view=webview2-winrt-1.0.1370.28&preserve-view=true#dragleave)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2CompositionController3::DragEnter method](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3?view=webview2-1.0.1370.28&preserve-view=true#dragenter)
* [ICoreWebView2CompositionController3::DragLeave method](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3?view=webview2-1.0.1370.28&preserve-view=true#dragleave)
* [ICoreWebView2CompositionController3::DragOver method](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3?view=webview2-1.0.1370.28&preserve-view=true#dragover)
* [ICoreWebView2CompositionController3::Drop method](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3?view=webview2-1.0.1370.28&preserve-view=true#drop)

---


<!-- ====================================================================== -->
Release Date: October 11, 2022`,
	},
	"1.0.1414-prerelease": {
		SDKVersion:     "1.0.1414-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101414-prerelease",
		RuntimeVersion: "107.0.1414.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ---------- -->
###### Experimental features

*  Added support for the Print API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.PrintAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.printasync?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
* [CoreWebView2.PrintToPdfStreamAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.printtopdfstreamasync?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
* [CoreWebView2.ShowPrintUI Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.showprintui?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
* [CoreWebView2PrintSettings Class](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.Collation Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.collation?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.ColorMode Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.colormode?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.Copies Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.copies?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2printsettings-copies)
   * [CoreWebView2PrintSettings.Duplex Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.duplex?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.MediaSize Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.mediasize?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.PageRanges Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.pageranges?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.PagesPerSide Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.pagesperside?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.PrinterName Property](/dotnet/api/microsoft.web.webview2.core.corewebview2printsettings.printername?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.PrintAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#printasync)
* [CoreWebView2.PrintToPdfStreamAsync](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#printtopdfstreamasync)
* [CoreWebView2.ShowPrintUI Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#showprintui)
* [CoreWebView2PrintSettings Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true)
   * [CoreWebView2PrintSettings.Collation Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#collation)
   * [CoreWebView2PrintSettings.ColorMode Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#colormode)
   * [CoreWebView2PrintSettings.Copies Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#copies)
   * [CoreWebView2PrintSettings.Duplex Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#duplex)
   * [CoreWebView2PrintSettings.MediaSize Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#mediasize)
   * [CoreWebView2PrintSettings.PageRanges Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#pageranges)
   * [CoreWebView2PrintSettings.PagesPerSide Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#pagesperside)
   * [CoreWebView2PrintSettings.PrinterName Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2printsettings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#printername)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental17::Print method](/microsoft-edge/webview2/reference/win32/icorewebview2experimental17?view=webview2-1.0.1414-prerelease&preserve-view=true#print)
* [ICoreWebView2Experimental17::PrintToPdfStream method](/microsoft-edge/webview2/reference/win32/icorewebview2experimental17?view=webview2-1.0.1414-prerelease&preserve-view=true#printtopdfstream)
* [ICoreWebView2Experimental17::ShowPrintUI method](/microsoft-edge/webview2/reference/win32/icorewebview2experimental17?view=webview2-1.0.1414-prerelease&preserve-view=true#showprintui)
* [ICoreWebView2ExperimentalPrintCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintcompletedhandler?view=webview2-1.0.1414-prerelease&preserve-view=true)
* [ICorewebView2ExperimentalPrintToPdfStreamCompletedHandler interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprinttopdfstreamcompletedhandler?view=webview2-1.0.1414-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalPrintSettings2 interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true)
   * [ICoreWebView2ExperimentalPrintSettings2::Collation property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_collation), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_collation)
   * [ICoreWebView2ExperimentalPrintSettings2::ColorMode property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_colormode), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_colormode)
   * [ICoreWebView2ExperimentalPrintSettings2::Copies property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_copies), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_copies)
   * [ICoreWebView2ExperimentalPrintSettings2::Duplex property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_duplex), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_duplex)
   * [ICoreWebView2ExperimentalPrintSettings2::MediaSize property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_mediasize), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_mediasize)
   * [ICoreWebView2ExperimentalPrintSettings2::PageRanges property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_pageranges), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_pageranges)
   * [ICoreWebView2ExperimentalPrintSettings2::PagesPerSide property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_pagesperside), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_pagesperside)
   * [ICoreWebView2ExperimentalPrintSettings2::PrinterName property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_printername), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprintsettings2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_printername)

---

*  Added support for SmartScreen API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2Settings.IsReputationCheckingRequired Property](/dotnet/api/microsoft.web.webview2.core.corewebview2settings.isreputationcheckingrequired?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2settings-isreputationcheckingrequired)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2Settings.IsReputationCheckingRequired Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2settings?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#isreputationcheckingrequired)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalSettings7::IsReputationCheckingRequired property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsettings7?view=webview2-1.0.1414-prerelease&preserve-view=true#get_isreputationcheckingrequired), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsettings7?view=webview2-1.0.1414-prerelease&preserve-view=true#put_isreputationcheckingrequired)

---

*  Added support for Custom Crash Reporting API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2EnvironmentOptions.IsCustomCrashReportingEnabled Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions.iscustomcrashreportingenabled?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2environmentoptions-iscustomcrashreportingenabled)

* [CoreWebView2Environment.FailureReportFolderPath Property](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.failurereportfolderpath?view=webview2-dotnet-1.0.1414-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2environment-failurereportfolderpath)


##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2EnvironmentOptions.IsCustomCrashReportingEnabled Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#iscustomcrashreportingenabled)
* [CoreWebView2Environment.FailureReportFolderPath Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environment?view=webview2-winrt-1.0.1414-prerelease&preserve-view=true#failurereportfolderpath)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalEnvironmentOptions2::IsCustomCrashReportingEnabled property (get](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions2?view=webview2-1.0.1414-prerelease&preserve-view=true#get_iscustomcrashreportingenabled), [put)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions2?view=webview2-1.0.1414-prerelease&preserve-view=true#put_iscustomcrashreportingenabled)
* [ICoreWebView2ExperimentalEnvironment::FailureReportFolderPath property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment?view=webview2-1.0.1414-prerelease&preserve-view=true#get_failurereportfolderpath)<!--no put-->

---


<!-- another section links to here -->
###### Bug fixes for 1.0.1414-prerelease

*   Removed three-dot menu with a broken link from the downloads page.  (Runtime-only)  ([Issue #2753](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2753))
*   Fixed a bug in the WebView2 WinRT JS Projection tool (wv2winrt) where C++20 projects failed to compile.  ([Issue #2768](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2768))
*   Fixed a crash which could occur with the WebView2 WinRT API while closing down WebView2 if you subscribed to any events, especially the 'CoreWebView2.GetDevToolsEventReceiver' event.  (SDK-only)
*   Fixed a bug where it wasn't possible to dismiss the download popup after minimizing the window.  (Runtime-only)


<!-- ====================================================================== -->
Release Date: September 6, 2022`,
	},
	"1.0.1343.22": {
		SDKVersion:     "1.0.1343.22",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10134322",
		RuntimeVersion: "105.0.1343.22",
		Notes: 			`

<!-- ------------------------------ -->
#### General

This WebView2 SDK release has the same bug fixes as [Bug fixes for 1.0.1369-prerelease](#bug-fixes-for-101369-prerelease).


<!-- ====================================================================== -->
Release Date: September 6, 2022`,
	},
	"1.0.1369-prerelease": {
		SDKVersion:     "1.0.1369-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101369-prerelease",
		RuntimeVersion: "106.0.1369.0",
		Notes: 			`

<!-- ------------------------------ -->
#### General


<!-- ---------- -->
###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


*  The drag and drop API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2CompositionController.DragLeave Method](/dotnet/api/microsoft.web.webview2.core.corewebview2compositioncontroller.dragleave?view=webview2-dotnet-1.0.1369-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2compositioncontroller-dragleave)

##### [WinRT/C#](#tab/winrtcsharp)

* [ICoreWebView2CompositionControllerInterop2.DragEnter Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2compositioncontrollerinterop2?view=webview2-winrt-1.0.1369-prerelease&preserve-view=true#dragenter)
* [ICoreWebView2CompositionControllerInterop2.DragLeave Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2compositioncontrollerinterop2?view=webview2-winrt-1.0.1369-prerelease&preserve-view=true#dragleave)
* [ICoreWebView2CompositionControllerInterop2.DragOver Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2compositioncontrollerinterop2?view=webview2-winrt-1.0.1369-prerelease&preserve-view=true#dragover)
* [ICoreWebView2CompositionControllerInterop2.Drop Method](/microsoft-edge/webview2/reference/winrt/interop/icorewebview2compositioncontrollerinterop2?view=webview2-winrt-1.0.1369-prerelease&preserve-view=true#drop)
* [CoreWebView2CompositionController.DragLeave Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2compositioncontroller?view=webview2-winrt-1.0.1369-prerelease&preserve-view=true#dragleave)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2CompositionController3::DragEnter method](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3?view=webview2-1.0.1369-prerelease&preserve-view=true#dragenter)
* [ICoreWebView2CompositionController3::DragLeave method](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3?view=webview2-1.0.1369-prerelease&preserve-view=true#dragleave)
* [ICoreWebView2CompositionController3::DragOver method](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3?view=webview2-1.0.1369-prerelease&preserve-view=true#dragover)
* [ICoreWebView2CompositionController3::Drop method](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3?view=webview2-1.0.1369-prerelease&preserve-view=true#drop)

---


<!-- another section links to here -->
###### Bug fixes for 1.0.1369-prerelease

*  Fixed a bug where WPF apps would crash when windows with WebView2 were closed. ([Issue #640](https://github.com/MicrosoftEdge/WebView2Feedback/issues/640))

*  Fixed a bug that produced simultaneous WebView creation failure.  (Runtime-only)  [Issue #2703](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2703)

*  Fixed print settings paper size to support dimensions as small as 0.01 inches.  (Runtime-only)

*  Fixed a bug where the WebView2 print dialog reset the **Scale** setting to **Fit to printable area** every time. ([Issue #2523](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2523))

*  Fixed a bug in the **wv2winrt** tool where a WinMD file wasn't referenced in some projects.


<!-- ====================================================================== -->
Release Date: August 8, 2022`,
	},
	"1.0.1293.44": {
		SDKVersion:     "1.0.1293.44",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10129344",
		RuntimeVersion: "104.0.1293.44",
		Notes: 			`
#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


* The Favicon API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.FaviconChanged Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.faviconchanged?view=webview2-dotnet-1.0.1293.44&preserve-view=true)
* [CoreWebView2.FaviconUri Property](/dotnet/api/microsoft.web.webview2.core.corewebview2.faviconuri?view=webview2-dotnet-1.0.1293.44&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.FaviconChanged Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1293.44&preserve-view=true#faviconchanged)
* [CoreWebView2.FaviconUri Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1293.44&preserve-view=true#faviconuri)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_15::FaviconChanged event (add](/microsoft-edge/webview2/reference/win32/icorewebview2_15?view=webview2-1.0.1293.44&preserve-view=true#add_faviconchanged), [remove)](/microsoft-edge/webview2/reference/win32/icorewebview2_15?view=webview2-1.0.1293.44&preserve-view=true#remove_faviconchanged)
* [ICoreWebView2_15::FaviconUri property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2_15?view=webview2-1.0.1293.44&preserve-view=true#get_faviconuri)<!--no put-->

---


<!-- ====================================================================== -->
Release Date: August 8, 2022`,
	},
	"1.0.1340-prerelease": {
		SDKVersion:     "1.0.1340-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101340-prerelease",
		RuntimeVersion: "105.0.1340.0",
		Notes: 			`
#### General

###### Experimental features

*  Added support for 'WebResourceRequested' for workers which allows setting filters in order to receive 'WebResourceRequested' events for service workers, shared workers, and different origin iframes.

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.AddWebResourceRequestedFilter(RequestSourceKinds) Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.addwebresourcerequestedfilter?view=webview2-dotnet-1.0.1340-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2-addwebresourcerequestedfilter(system-string-microsoft-web-webview2-core-corewebview2webresourcecontext-microsoft-web-webview2-core-corewebview2webresourcerequestsourcekinds))
* [CoreWebView2.RemoveWebResourceRequestedFilter(RequestSourceKinds) Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.removewebresourcerequestedfilter?view=webview2-dotnet-1.0.1340-prerelease&preserve-view=true#microsoft-web-webview2-core-corewebview2-removewebresourcerequestedfilter(system-string-microsoft-web-webview2-core-corewebview2webresourcecontext-microsoft-web-webview2-core-corewebview2webresourcerequestsourcekinds))
* [CoreWebView2WebResourceRequestedEventArgs Class](/dotnet/api/microsoft.web.webview2.core.corewebview2webresourcerequestedeventargs?view=webview2-dotnet-1.0.1340-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.AddWebResourceRequestedFilter(requestSourceKinds) Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1340-prerelease&preserve-view=true#addwebresourcerequestedfilter)
* [CoreWebView2.RemoveWebResourceRequestedFilter(requestSourceKinds) Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1340-prerelease&preserve-view=true#removewebresourcerequestedfilter)
* [CoreWebView2WebResourceRequestedEventArgs Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2webresourcerequestedeventargs?view=webview2-winrt-1.0.1340-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2Experimental16.AddWebResourceRequestedFilterWithRequestSourceKinds method](/microsoft-edge/webview2/reference/win32/icorewebview2experimental16?view=webview2-1.0.1340-prerelease&preserve-view=true#addwebresourcerequestedfilterwithrequestsourcekinds)
* [ICoreWebView2Experimental16.RemoveWebResourceRequestedFilterWithRequestSourceKinds method](/microsoft-edge/webview2/reference/win32/icorewebview2experimental16?view=webview2-1.0.1340-prerelease&preserve-view=true#removewebresourcerequestedfilterwithrequestsourcekinds)
* [ICoreWebView2ExperimentalWebResourceRequestedEventArgs interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebresourcerequestedeventargs?view=webview2-1.0.1340-prerelease&preserve-view=true)

---

*  Added support for custom scheme registration which allows WebView2 apps to be able to handle 'WebResourceRequested' event for requests with the specified scheme and be able to navigate the WebView2 control to the custom scheme.

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2EnvironmentOptions Class](/dotnet/api/microsoft.web.webview2.core.corewebview2environmentoptions?view=webview2-dotnet-1.0.1340-prerelease&preserve-view=true)
* [CoreWebView2CustomSchemeRegistration Class](/dotnet/api/microsoft.web.webview2.core.corewebview2customschemeregistration?view=webview2-dotnet-1.0.1340-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2EnvironmentOptions Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2environmentoptions?view=webview2-winrt-1.0.1340-prerelease&preserve-view=true)
* [CoreWebView2CustomSchemeRegistration Class](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2customschemeregistration?view=webview2-winrt-1.0.1340-prerelease&preserve-view=true)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2ExperimentalEnvironmentOptions interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions?view=webview2-1.0.1340-prerelease&preserve-view=true)
* [ICoreWebView2ExperimentalCustomSchemeRegistration interface](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcustomschemeregistration?view=webview2-1.0.1340-prerelease&preserve-view=true)

---

###### Bug fixes

*   Added the ability for developers to explicitly specify the path from which to load the WebView2Loader.dll. ([Issue #767](https://github.com/MicrosoftEdge/WebView2Feedback/issues/767))

*   Added useful error messages when using 'CallDevToolsProtocolMethod'. ([Issue #1609](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1609))

*   Fixed a bug in finding and loading the 'WebView2Loader.dll' in some .NET apps. ([Issue #2372](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2372))

*   Fixed a bug where 'DownloadStarting' event was not fired when retrying a download. ([Issue #2489](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2489))

*   Fixed an issue in service worker caching if the path was too long. ([Issue #1900](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1900))

*   Improved performance for **wv2winrt** 'IMap' and 'IMapView' projections into JavaScript.

*   Adding support for HWND_MESSAGE to be used as WebView2 parent window to support headless scenarios.  ([Issue #202](https://github.com/MicrosoftEdge/WebView2Feedback/issues/202))

*   Improved handling of running as admin user apps.

*   Fixed online/offline status and notifications when using WebView2 in UWP apps.

*   GDI scaling can now be enabled for WebView2.  WebView2 will respect the GDI scaling setting of the hosting application without additional work needed from the app.  ([Issue #1700](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1700))

*   Fixed a bug where focus is not returned to the application after closing the find bar for windowed mode. ([Issue #1225](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1225))


<!-- ====================================================================== -->
Release Date: July 4, 2022`,
	},
	"1.0.1264.42": {
		SDKVersion:     "1.0.1264.42",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10126442",
		RuntimeVersion: "103.0.1264.42",
		Notes: 			`
#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


*  Added 'ContextMenuRequested'API to enable host app to create or modify their own context menu.

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.ContextMenuRequested Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.contextmenurequested?view=webview2-dotnet-1.0.1264.42&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.ContextMenuRequested Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1264.42&preserve-view=true#contextmenurequested)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_11::add_ContextMenuRequested event (add](/microsoft-edge/webview2/reference/win32/icorewebview2_11?view=webview2-1.0.1264.42&preserve-view=true#add_contextmenurequested), [remove)](/microsoft-edge/webview2/reference/win32/icorewebview2_11?view=webview2-1.0.1264.42&preserve-view=true#remove_contextmenurequested)

---


<!-- ====================================================================== -->
Release Date: July 4, 2022`,
	},
	"1.0.1305-prerelease": {
		SDKVersion:     "1.0.1305-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101305-prerelease",
		RuntimeVersion: "105.0.1305.0",
		Notes: 			`
#### General

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.


* The Favicon API:

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.FaviconChanged Event](/dotnet/api/microsoft.web.webview2.core.corewebview2.faviconchanged?view=webview2-dotnet-1.0.1305-prerelease&preserve-view=true)
* [CoreWebView2.FaviconUri Property](/dotnet/api/microsoft.web.webview2.core.corewebview2.faviconuri?view=webview2-dotnet-1.0.1305-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.FaviconChanged Event](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1305-prerelease&preserve-view=true#faviconchanged)
* [CoreWebView2.FaviconUri Property](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1305-prerelease&preserve-view=true#faviconuri)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_15::FaviconChanged event (add](/microsoft-edge/webview2/reference/win32/icorewebview2_15?view=webview2-1.0.1305-prerelease&preserve-view=true#add_faviconchanged), [remove)](/microsoft-edge/webview2/reference/win32/icorewebview2_15?view=webview2-1.0.1305-prerelease&preserve-view=true#remove_faviconchanged)
* [ICoreWebView2_15::FaviconUri property (get)](/microsoft-edge/webview2/reference/win32/icorewebview2_15?view=webview2-1.0.1305-prerelease&preserve-view=true#get_faviconuri)<!--no put-->

---

###### Bug fixes

*  Fixed an issue where 'PrintToPdfAsync' may hang for long time. ([Issue #1974](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1974))

##### [.NET/C#](#tab/dotnetcsharp)

* [CoreWebView2.PrintToPdfAsync Method](/dotnet/api/microsoft.web.webview2.core.corewebview2.printtopdfasync?view=webview2-dotnet-1.0.1305-prerelease&preserve-view=true)

##### [WinRT/C#](#tab/winrtcsharp)

* [CoreWebView2.PrintToPdfAsync Method](/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2?view=webview2-winrt-1.0.1305-prerelease&preserve-view=true#printtopdfasync)

##### [Win32/C++](#tab/win32cpp)

* [ICoreWebView2_7::PrintToPdf method](/microsoft-edge/webview2/reference/win32/icorewebview2_7?view=webview2-1.0.1305-prerelease&preserve-view=true#printtopdf)

---

* Fixed regression where WebView2 would steal focus from the app when the WebView2 was made visible. ([Issue #862](https://github.com/MicrosoftEdge/WebView2Feedback/issues/862))


<!-- ====================================================================== -->
Release Date: June 14, 2022`,
	},
	"1.0.1245.22": {
		SDKVersion:     "1.0.1245.22",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10124522",
		RuntimeVersion: "102.0.1245.22",
		Notes: 			`
There is no corresponding prerelease package.

#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.


* The [Server Certificate API](/microsoft-edge/webview2/reference/win32/icorewebview2_14?view=webview2-1.0.1245.22&preserve-view=true) which provides an option to trust the server's TLS certificate at the application level. It renders the page without prompting the user about TLS or providing the ability to cancel the web request.

*  The [ClearBrowsingData API](/microsoft-edge/webview2/reference/win32/icorewebview2profile2?view=webview2-1.0.1245.22&preserve-view=true) which allows developers to programmatically clear specific data types for a duration:
   * 'ClearBrowsingData'
   * 'ClearBrowsingDataAll'
   * 'ClearBrowsingDataInTimeRange'

*  The [HttpStatusCode API](/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs2?view=webview2-1.0.1245.22&preserve-view=true) which provides the HTTP status code for navigation requests in 'NavigationCompleted' events.

###### Bug fixes

*  Fixed an issue with the on-screen keyboard in which the keyboard doesn't reappear after it's closed by clicking the **X** button. Also fixed an issue in which the keyboard gets dismissed when users switch from one edit control to another within WebView2. ([Issue #460](https://github.com/MicrosoftEdge/WebView2Feedback/issues/460))
*  Fixed an issue when using a proxy from 'AddHostObjectToScript' in script.  If you call 'setHostProperty' and it failed, you could have received an internal error message structure rather than a JavaScript Error object.
*  Fixed regression where WebView2 would steal focus from the app when the WebView2 was made visible.  ([Issue #862](https://github.com/MicrosoftEdge/WebView2Feedback/issues/862))
*  Fixed a bug that caused increased memory usage with 'WebResourceRequested' events using large data.  ([Issue #2171](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2171))
*  Fixed 'StatusBarTextChanged' regression. The [StatusBarText API](/microsoft-edge/webview2/reference/win32/icorewebview2_12?view=webview2-1.0.1245.22&preserve-view=true) was made compatible with previous versions again. ([Issue #2414](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2414))
*  Better support for apps running as admin. ([Issue #2356](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2356))


<!-- ====================================================================== -->
Release Date: May 9, 2022`,
	},
	"1.0.1210.39": {
		SDKVersion:     "1.0.1210.39",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10121039",
		RuntimeVersion: "101.0.1210.39",
		Notes: 			`
#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

* Support for [multiple user profiles](/microsoft-edge/webview2/reference/win32/icorewebview2environment10?view=webview2-1.0.1210.39&preserve-view=true) in WebView2.

* [Theming API](/microsoft-edge/webview2/reference/win32/icorewebview2profile?view=webview2-1.0.1210.39&preserve-view=true) which provides a way to customize the WebView2 color theme as 'light', 'dark', or 'system'.

* [Default Download API](/microsoft-edge/webview2/reference/win32/icorewebview2profile?view=webview2-1.0.1210.39&preserve-view=true) which provides a way to customize the default download location.

<!-- ====================================================================== -->
Release Date: May 9, 2022`,
	},
	"1.0.1248-prerelease": {
		SDKVersion:     "1.0.1248-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101248-prerelease",
		RuntimeVersion: "102.0.1248.0",
		Notes: 			`
#### General

* Added support for WinRT Object projection into JavaScript by adding WinRT JS Projection tool (**wv2winrt**) in NuGet package. For instructions about using the WinRT JS Projection tool see [Call native-side WinRT code from web-side code](/microsoft-edge/webview2/how-to/winrt-from-js).

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

* The [Server Certificate API](/microsoft-edge/webview2/reference/win32/icorewebview2_14?view=webview2-1.0.1248-prerelease&preserve-view=true) which provides an option to trust the server's TLS certificate at the application level and render the page without prompting the user about TLS or providing the ability to cancel the web request.

* The [ClearBrowsingData API](/microsoft-edge/webview2/reference/win32/icorewebview2profile2?view=webview2-1.0.1248-prerelease&preserve-view=true) which allows developers to programmatically clear specific data types for a duration:
   * 'clearBrowsingDataInTimeRange'
   * 'clearBrowsingDataAll'

###### Bug fixes

* Fixed an unavoidable crash that occurred in the WPF control's 'OnWindowPositionChanged' event. ([Issue #1531](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1531))

* Fixed the issue with 'CoreWebView2EnvironmentOptions.ExclusiveUserDataFolderAccess' not working properly in .NET SDK. ([Issue #2363](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2363))

* Fixed a runtime regression that caused some Office Add-ins which use host objects to crash during operations that previously worked. ([Issue #2337](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2337))

* Fixed an issue where WebView2 content can become blurry when moving between monitors with different scaling.

* Fixed a regression to make sure that WebView2 creation fails quickly with 'HRESULT_FROM_WIN32(ERROR_INVALID_STATE)' instead of time out.

* Fixed a bug where changes from Chromium broke WebView2 background color.


<!-- ====================================================================== -->
Release Date: April 12, 2022`,
	},
	"1.0.1185.39": {
		SDKVersion:     "1.0.1185.39",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10118539",
		RuntimeVersion: "100.0.1185.39",
		Notes: 			`
#### General

* Renamed 'ICoreWebView2Certificate' to 'ICoreWebView2ClientCertificate'.

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

* The [CallDevToolsProtocolMethodForSession API](/microsoft-edge/webview2/reference/win32/icorewebview2_11?view=webview2-1.0.1185.39&preserve-view=true#calldevtoolsprotocolmethodforsession) that supports 'sessionId' for CDP method calls.

* The [StatusBarText API](/microsoft-edge/webview2/reference/win32/icorewebview2_12?view=webview2-1.0.1185.39&preserve-view=true):
    *  'add_StatusBarTextChanged'
    *  'get_StatusBarText'
    *  'remove_StatusBarTextChanged'

* The [AllowExternalDrop API](/microsoft-edge/webview2/reference/win32/icorewebview2controller4?view=webview2-1.0.1185.39&preserve-view=true) that supports enable/disable for external drop operations.

* The [HiddenPdfToolbarItems API](/microsoft-edge/webview2/reference/win32/icorewebview2settings7?view=webview2-1.0.1185.39&preserve-view=true) is available to customize PDF toolbar items.

* The [ExclusiveUserDataFolderAccess API](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions2?view=webview2-1.0.1185.39&preserve-view=true) allows control of whether or not other processes can create WebView2 from 'WebView2Environment' created with the same user data folder and therefore sharing the same WebView browser process instance.

* The [permission requested support for iframes](/microsoft-edge/webview2/reference/win32/icorewebview2frame3?view=webview2-1.0.1185.39&preserve-view=true):
    * 'add_PermissionRequested'
    * 'remove_PermissionRequested'


<!-- ====================================================================== -->
Release Date: April 12, 2022`,
	},
	"1.0.1222-prerelease": {
		SDKVersion:     "1.0.1222-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101222-prerelease",
		RuntimeVersion: "102.0.1222.0",
		Notes: 			`
#### General

###### Experimental features

* Added the [Server Certificate API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental15?view=webview2-1.0.1222-prerelease&preserve-view=true) which provides an option to trust the server's TLS certificate at the application level and render the page without prompting the user about TLS or providing the ability to cancel the web request.

* Added the [Favicon API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental12?view=webview2-1.0.1222-prerelease&preserve-view=true) which provides a way to get the favicon when it changes or is set at a website.

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

* Support for [multiple user profiles](/microsoft-edge/webview2/reference/win32/icorewebview2environment10?view=webview2-1.0.1222-prerelease&preserve-view=true) in WebView2.

* [Theming API](/microsoft-edge/webview2/reference/win32/icorewebview2profile?view=webview2-1.0.1222-prerelease&viewFallbackFrom=webview2-1.0.1185.39&preserve-view=true) which provides a way to customize the WebView2 color theme as 'light', 'dark', or 'system'.

* [Default Download API](/microsoft-edge/webview2/reference/win32/icorewebview2profile?view=webview2-1.0.1222-prerelease&viewFallbackFrom=webview2-1.0.1185.39&preserve-view=true) which provides a way to customize the default download location.

###### Bug fixes

* Fixed 'ZoomFactor' issue that incorrectly sets 'ZoomFactor' value to the maximum value when it is out of bounds.

* Fixed an issue in which WebView2 content can become blurry when moving between monitors with different scaling.

* Fixed a bug where 'MouseEvent.movementX' and 'MouseEvent.movementY' will always be **0** in visual hosting mode. ([Issue #2220](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2220))

* Fixed log in issue caused by a password regression in WebView2. ([Issue #2291](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2291))

* Fixed a failure caused when a user opens a new app window and the web page does not have a navigation entry assigned.

* Made a runtime change to fix a bug in WinUI 2 (UWP) in which owned windows were not showing up.

* Fixed 'ICoreWebView2Frame::PostWebMessage' functionality after source update. ([Issue #2267](https://github.com/MicrosoftEdge/WebView2Feedback/issues/2267))


<!-- ====================================================================== -->
Release Date: March 10, 2022`,
	},
	"1.0.1150.38": {
		SDKVersion:     "1.0.1150.38",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10115038",
		RuntimeVersion: "99.0.1150.38",
		Notes: 			`
#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

*   The [BasicAuthentication API](/microsoft-edge/webview2/reference/win32/icorewebview2_10?view=webview2-1.0.1150.38&preserve-view=true) that enables developers to handle Basic HTTP Authentication request and response.


<!-- ====================================================================== -->
Release Date: March 10, 2022`,
	},
	"1.0.1189-prerelease": {
		SDKVersion:     "1.0.1189-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101189-prerelease",
		RuntimeVersion: "100.0.1189.0",
		Notes: 			`
#### General

###### Experimental features

*   Added [ContextMenuRequested API](/microsoft-edge/webview2/reference/win32/icorewebview2_11?view=webview2-1.0.1189-prerelease&preserve-view=true) to enable host app to create or modify their own context menu.

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*    The [CallDevToolsProtocolMethodForSession API](/microsoft-edge/webview2/reference/win32/icorewebview2_11?view=webview2-1.0.1189-prerelease&preserve-view=true#calldevtoolsprotocolmethodforsession) that supports sessionId for CDP method calls.
*   The [StatusBarText API](/microsoft-edge/webview2/reference/win32/icorewebview2_12?view=webview2-1.0.1189-prerelease&preserve-view=true):
    *  'add_StatusBarTextChanged'
    *  'get_StatusBarText'
    *  'remove_StatusBarTextChanged'
*   The [AllowExternalDrop API](/microsoft-edge/webview2/reference/win32/icorewebview2controller4?view=webview2-1.0.1189-prerelease&preserve-view=true) that supports enable/disable external drop.
*    The [HiddenPdfToolbarItems API](/microsoft-edge/webview2/reference/win32/icorewebview2settings7?view=webview2-1.0.1189-prerelease&preserve-view=true) is available to customize the PDF toolbar items.
*  The [ExclusiveUserDataFolderAccess API](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions2?view=webview2-1.0.1189-prerelease&preserve-view=true) allows control of whether or not other processes can create WebView2 using the same user data folder.

###### Bug fixes

*   Fixed a bug where WebView2 app gets stuck occasionally with UWP.
*   Fixed a bug where focus is not returned to the application after closing the **Find** bar for windowed mode.
*   Fixed bug in which the 'DocumentTitleChanged' event was not being raised for backward/forward navigation in single-page apps.
*   Fixed bug in which the 'HistoryChanged' event was not being raised for Iframe navigation.


<!-- ====================================================================== -->
Release Date: February 6, 2022`,
	},
	"1.0.1108.44": {
		SDKVersion:     "1.0.1108.44",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10110844",
		RuntimeVersion: "98.0.1108.44",
		Notes: 			`
#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

*  The [AdditionalAllowedFrameAncestors API](/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventargs2?view=webview2-1.0.1108.44&preserve-view=true) that enable developers to provide additional allowed frame ancestors.

*  The [ProcessInfo APIs](/microsoft-edge/webview2/reference/win32/icorewebview2processinfo?view=webview2-1.0.1108.44&preserve-view=true) provide more information about WebView2 processes and process collections.

*  New [APIs for iframes](/microsoft-edge/webview2/reference/win32/icorewebview2frame2?view=webview2-1.0.1108.44&preserve-view=true&preserve-view=true):
   *  'add_NavigationStarting'
   *  'remove_NavigationStarting'
   *  'add_ContentLoading'
   *  'remove_ContentLoading'
   *  'add_NavigationCompleted'
   *  'remove_NavigationCompleted'
   *  'add_DOMContentLoaded'
   *  'remove_DOMContentLoaded'
   *  'ExecuteScript'
   *  'PostWebMessageAsJson'
   *  'PostWebMessageAsString'
   *  'add_WebMessageReceived'
   *  'remove_WebMessageReceived'


<!-- ====================================================================== -->
Release Date: February 6, 2022`,
	},
	"1.0.1158-prerelease": {
		SDKVersion:     "1.0.1158-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101158-prerelease",
		RuntimeVersion: "100.0.1158.0",
		Notes: 			`
#### General

###### Experimental features

*  Added [Status bar API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental13?view=webview2-1.0.1158-prerelease&preserve-view=true) to provide info when webiew is showing status message, URL, or empty string.
*  Added [CDP API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental14?view=webview2-1.0.1158-prerelease&preserve-view=true) to provide possibility for developers have multiple 'DevToolsProtocol' targets in WebView2.

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  Rename ICoreWebView2ClientCertificate to [ICoreWebView2Certificate](/microsoft-edge/webview2/reference/win32/icorewebview2certificate?view=webview2-1.0.1158-prerelease&preserve-view=true).
*  New [APIs for iframes](/microsoft-edge/webview2/reference/win32/icorewebview2frame3?view=webview2-1.0.1158-prerelease&preserve-view=true):
   *  'add_PermissionRequested'
   *  'remove_PermissionRequested'

###### Bug fixes

*  Fixed an issue causing erroneous warnings in the Visual Studio Error List window.  ([Issue #1722](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1722))
*  Fixed a bug where NewWindowRequested was not getting raised when opening PDF downloads.
*  Resolved a bug in WinUI 3 where select dropdowns would not show up.  ([Issue #1693](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1693))
*  Added the ability to toggle WebView2 mute state, even when there is no audio playing.


<!-- ====================================================================== -->
Release Date: January 13, 2022`,
	},
	"1.0.1072.54": {
		SDKVersion:     "1.0.1072.54",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10107254",
		RuntimeVersion: "97.0.1072.54",
		Notes: 			`
#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

*  The [Media API](/microsoft-edge/webview2/reference/win32/icorewebview2_8?view=webview2-1.0.1072.54&preserve-view=true#summary) that enables developers to mute/unmute media within WebView2.

*  The [Download Positioning and Anchoring API](/microsoft-edge/webview2/reference/win32/icorewebview2_9?view=webview2-1.0.1072.54&preserve-view=true) enables:
   *  Changing the position of the download dialog, relative to the WebView2 bounds.  You can anchor the download dialog to the **Download** button, instead of the default position, which is the top-right corner.
   *  Programmatically open and close the default download dialog.
   *  Making changes in response to the dialog opening and closing.


<!-- ====================================================================== -->
Release Date: January 13, 2022`,
	},
	"1.0.1133-prerelease": {
		SDKVersion:     "1.0.1133-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101133-prerelease",
		RuntimeVersion: "99.0.1133.0",
		Notes: 			`
#### General

###### Experimental features

*  Added support for [theming](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile2?view=webview2-1.0.1133-prerelease&preserve-view=true) (overall color scheme - light, dark, system) of WebView2.
*  Added a way to set [default download path](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile3?view=webview2-1.0.1133-prerelease&preserve-view=true).
*  Added support for [clearing browser data](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprofile4?view=webview2-1.0.1133-prerelease&preserve-view=true).
*  Added [permission requested](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe3?view=webview2-1.0.1133-prerelease&preserve-view=true) support for iframes.

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  New [APIs for iframes](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe2?view=webview2-1.0.1133-prerelease&preserve-view=true):
   *  'PostWebMessageAsJson'
   *  'PostWebMessageAsString'
   *  'add_WebMessageReceived'
   *  'remove_WebMessageReceived'
*  The ProcessInfo APIs provides more information about WebView2 [processes](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprocessinfo?view=webview2-1.0.1133-prerelease&preserve-view=true) and [process collections](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprocessinfocollection?view=webview2-1.0.1133-prerelease&preserve-view=true).
*  The [HTTP Authentication API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental10?view=webview2-1.0.1133-prerelease&preserve-view=true).

###### Bug fixes

*  Fixed a bug that prevented 'Set-Cookies' header from showing up in the 'WebResourceResponseReceived' event.
*  Resolved a bug where pop-ups and owned windows would jump to a different position before closing instead of closing
along with the app window. This bug was only active for a very short window of time.
*  Fixed focus issue after closing file picker dialog.
*  Fixed bug where Find on Page UI visibility did not change with WebView2 visibility.
*  Fixed bug where 'GetAvailableBrowserVersionString()' fails to locate/load 'WebView2Loader.dll'. ([Issue #1236](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1236))
*  Fixed size and position of the new window created with 'window.open' when 'NewWindowRequested' event was not
handled. ([Issue #1343](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1343))
*  Fixed bug where mini menu was still displaying on selected text when context menus were disabled. This change is Runtime-specific. ([Issue #1345](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1345))
*  Fixed bug where focus returns to wrong location after switching apps in WinForms.


<!-- ====================================================================== -->
Release Date: November 29, 2021`,
	},
	"1.0.1083-prerelease": {
		SDKVersion:     "1.0.1083-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101083-prerelease",
		RuntimeVersion: "97.0.1083.0",
		Notes: 			`
#### Experimental features

* Added the following [APIs for iframes](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe2?view=webview2-1.0.1083-prerelease&preserve-view=true) in WebView2:
   *  'PostWebMessageAsJson'
   *  'PostWebMessageAsString'
   *  'add_WebMessageReceived'
   *  'remove_WebMessageReceived'

* Added ProcessInfo APIs to provide more information about WebView2 [processes](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprocessinfo?view=webview2-1.0.1083-prerelease&preserve-view=true) and [process collections](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprocessinfocollection?view=webview2-1.0.1083-prerelease&preserve-view=true).

#### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  The [Media API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental9?view=webview2-1.0.1083-prerelease&preserve-view=true#summary) that enables developers to mute/unmute media within WebView2.
*  The [Download Positioning and Anchoring API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental11?view=webview2-1.0.1083-prerelease&preserve-view=true).  This API enables:
   *  Changing the position of the download dialog, relative to the WebView2 bounds.  You can anchor the download dialog to the **Download** button, instead of the default position, which is the top-right corner.
   *  Programmatically opening and closing the default download dialog.
   *  Making changes in response to the dialog opening and closing.

#### Bug fixes

*  Fixed a focus issue after closing the file picker dialog.
*  Fixed a bug where WebView2 doesn't receive spatial input on initial launch.
*  Fixed an issue that prevented single sign-on in WebView2.
*  Resolved a bug where the download dialog was not moving with the window on WPF and WinForms.
*  Updated compatible command line check to prevent needing a version check for optional switches.
*  Fixed an error that was causing "Microsoft Edge" branding to appear in the accessibility tree.


<!-- ====================================================================== -->
Release Date: November 29, 2021`,
	},
	"1.0.1054.31": {
		SDKVersion:     "1.0.1054.31",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10105431",
		RuntimeVersion: "96.0.1054.31",
		Notes: 			`
#### General

*  General reliability fixes.

#### Bug fixes

*  Turned off the Control-flow Enforcement Technology (CET) Shadow Stack feature for v96 WebView2 Runtime.
*  Fixed an issue that was causing slow startup times when launching in a .NET single-file application. ([Issue #1909](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1909))
*  Fixed a crash caused by Microsoft Edge browser policies getting incorrectly applied to WebView2 as well. ([Issue #1860](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1860))
*  Fixed a crash that occurred when a pop-up window with a download dialog was closed. ([Issue #1765](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1765)) & ([Issue #1723](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1723))


<!-- ====================================================================== -->
Release Date: October 29, 2021`,
	},
	"1.0.1056-prerelease": {
		SDKVersion:     "1.0.1056-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101056-prerelease",
		RuntimeVersion: "97.0.1056.0",
		Notes: 			`
#### General

*  General reliability improvements.

###### Experimental features

*  The [Download Positioning and Anchoring API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental11?view=webview2-1.0.1056-prerelease&preserve-view=true).  This API enables:
   *  Changing the position of the download dialog, relative to the WebView2 bounds.  You can anchor the download dialog to the **Download** button, instead of the default position, which is the top-right corner.
   *  Programmatically opening and closing the default download dialog.
   *  Making changes in response to the dialog opening and closing.
*  The [HTTP Authentication API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental10?view=webview2-1.0.1056-prerelease&preserve-view=true).

###### Bug fixes

*  The real process exit code is now provided as 'ExitCode' in 'ICoreWebView2ProcessFailedEventArgs2' for 'COREWEBVIEW2_PROCESS_FAILED_KIND_BROWSER_PROCESS_EXITED' process failure.
*  The '--js-flags' switch is now honored in the 'AdditionalBrowserArguments' that are provided in 'CoreWebView2EnvironmentOptions'.
*  Fixed access to the 'name' property for host objects in JavaScript. ([Issue #641](https://github.com/MicrosoftEdge/WebView2Feedback/issues/641))
*  Fixed an 'InvalidCastException' in the WPF control when it's implicitly initialized prior to the event loop starting. ([Issue #1577](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1577))


<!-- ====================================================================== -->
Release Date: October 25, 2021`,
	},
	"1.0.1020.30": {
		SDKVersion:     "1.0.1020.30",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10102030",
		RuntimeVersion: "95.0.1020.30",
		Notes: 			`
#### General

###### Bug fixes

*  Updated 'EnsureCoreWebView2Async' to not throw exceptions when the WPF source property is set. ([Issue #1781](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1781))
*  Fixed a bug where WebView2 crashes after interacting with multiple windows that show a download UI. ([Issue #1723](https://github.com/MicrosoftEdge/WebView2Feedback/issues/1723))

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

*  [PrintToPdf API](/microsoft-edge/webview2/reference/win32/icorewebview2_7?view=webview2-1.0.1020.30&preserve-view=true#printtopdf).


<!-- ====================================================================== -->
Release Date: September 27, 2021`,
	},
	"1.0.992.28": {
		SDKVersion:     "1.0.992.28",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1099228",
		RuntimeVersion: "94.0.992.31",
		Notes: 			`
#### General

###### Bug fixes

*  Fixed missing WebView2 DLLs (which led to initialization failure) when 'PlatformTarget' isn't set in the user's .NET project. ([Issue #1061](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1061))

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

*  [OpenTaskManagerWindow API](/microsoft-edge/webview2/reference/win32/icorewebview2_6?view=webview2-1.0.992.28&preserve-view=true#summary).
*  [isSwipeNavigationEnabled property](/microsoft-edge/webview2/reference/win32/icorewebview2settings6?view=webview2-1.0.992.28&preserve-view=true#summary).
*  [BrowserProcessExited API](/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs?view=webview2-1.0.992.28&preserve-view=true#summary).
*  [get_Name property](/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs2?view=webview2-1.0.992.28#get_name&preserve-view=true#summary) on 'ICoreWebView2NewWindowRequestedEventArgs2' interface.


<!-- ====================================================================== -->
Release Date: September 20, 2021`,
	},
	"1.0.1018-prerelease": {
		SDKVersion:     "1.0.1018-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101018-prerelease",
		RuntimeVersion: "95.0.1018.0",
		Notes: 			`
#### General

###### Experimental features

*  Added a [media API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental9?view=webview2-1.0.1018-prerelease&preserve-view=true#summary) that enables developers to mute/unmute media within WebView2.
*  Added support for [multiple user profiles](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment8?view=webview2-1.0.1018-prerelease&preserve-view=true) with WebView2.

###### Bug fixes

*  Fixed a bug where WebView2 stops rendering when the app is spanning monitors and the monitor scale changes.
*  Fixed a bug where closing the download UI crashes WebView2 when multiple download windows are open. ([Issue #1723](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1723))
*  Fixed a build/initialization error when PlatformTarget isn't set in the user's .NET project. ([Issue #730](https://github.com/MicrosoftEdge/WebViewFeedback/issues/730) and [Issue #1548](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1548))


<!-- ====================================================================== -->
Release Date: September 14, 2021`,
	},
	"1.0.1010-prerelease": {
		SDKVersion:     "1.0.1010-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#101010-prerelease",
		RuntimeVersion: "95.0.1010.0",
		Notes: 			`
#### General
*  WebView2 performance improvements.
*  Reliability fixes. ([Issue #1605](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1605) and [Issue #1678](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1678))
*  Added performance improvements during startup and when the host app is in the foreground.

###### Experimental features

*  Removed silent failures by using 'EnsureCoreWebView2Async', which throws an 'ArgumentException' when called multiple times with incompatible parameters.
*  Changed default handling of the [UserDataFolder](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment5?view=webview2-1.0.1010-prerelease&preserve-view=true#get_userdatafolder) property in the environment object.

   > [!CAUTION]
   > **Breaking Change**:  The default handling for the user data folder, if the developer doesn't specify where to put it, will be changing.  See [Announcement: User directory folder default handling updates](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1410).

*  Added [navigation & script APIs](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe?view=webview2-1.0.1010-prerelease&preserve-view=true) for iframes.
*  Added [MemoryUsageTargetLevel](/microsoft-edge/webview2/reference/win32/icorewebview2experimental5?view=webview2-1.0.1010-prerelease&preserve-view=true) which allows developers to specify memory consumption levels, such as low, or normal.
*  Added [ExclusiveUserDataFolderAccess](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions?view=webview2-1.0.1010-prerelease&preserve-view=true) to environment options.
*  Added [HiddenPdfToolbarItems](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsettings6?view=webview2-1.0.1010-prerelease&preserve-view=true) to customize PDF toolbar items.
*  Added [PrintToPdf](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalprinttopdfcompletedhandler?view=webview2-1.0.1010-prerelease&preserve-view=true), which allows printing the current page to PDF. Also, you can use optional custom settings with this new API.
*  Added [AllowExternalDrop](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcompositioncontroller3?view=webview2-1.0.1010-prerelease&preserve-view=true) property to allow the dragging and dropping of objects from outside a WebView2 control into it.
*  Added [ContextMenu APIs](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontextmenuitem?view=webview2-1.0.1010-prerelease&preserve-view=true) which allow customization of the WebView2 context menu.

###### Bug fixes

*  Improved how host objects exceptions are caught in your JavaScript code.
*  Replaced WebView2 icon with a generic icon in DevTools windows.
*  Turn on the Tab screen sharing option when 'MediaDevices.getDisplayMedia()' is used. ([Issue #1566](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1566))
*  Fixed a bug in the Client Certificate API, when the correct certificate was not selected. This is a Runtime change. ([Issue #1666](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1666))
*  Fixed bug where 'window.chrome.webview' was unavailable in new windows in the same parent domain. This change is Runtime-specific. ([Issue #1144](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1144))
*  Fixed a bug where dropdown menus or lists were displayed behind the window that has focus. ([Issue #411](https://github.com/MicrosoftEdge/WebViewFeedback/issues/411))
*  Fixed focus issues when using 'put_IsVisible(false)'. ([Issue #238](https://github.com/MicrosoftEdge/WebViewFeedback/issues/238))
*  Fixed a bug to apply 'SetVirtualHostNameToFolderMapping' to pop-up windows.
*  Fixed bugs where an 'IDispatch' objects were returned as 'IUnknown'.

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  'IsSwipeNavigationEnabled'
*  'BrowserProcessExited'
*  'OpenBrowserTaskManager'


<!-- ====================================================================== -->
Release Date: September 8, 2021`,
	},
	"1.0.961.33": {
		SDKVersion:     "1.0.961.33",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1096133",
		RuntimeVersion: "93.0.961.44",
		Notes: 			`
#### General

###### Bug fixes
*  Fixed a bug that caused 'ERR_SSL_CLIENT_AUTH_CERT_NEEDED' errors. This is a Runtime change.
*  Fixed a bug that special browser keys like **Refresh**, **Home**, **Back**, and so on can't be turned off using 'AreBrowserAcceleratorKeysEnabled'. This change is Runtime-specific.
*  Fixed a bug where the transparent background color isn't rendered.
*  Fixed a bug that caused a white flicker when loading WebView2.
*  Fixed a bug in WebView2 .NET controls where WebView2 windows were blank when created in the background. ([Issue #1077](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1077))
*  Fixed a bug where settings were not updated when the user navigated to or a new window displayed 'about:blank' pages. This is a Runtime change.

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

*  [Client Certificate API](/microsoft-edge/webview2/reference/win32/icorewebview2_5?view=webview2-1.0.961.33&preserve-view=true#add_clientcertificaterequested).


<!-- ====================================================================== -->
Release Date: July 26, 2021`,
	},
	"1.0.955-prerelease": {
		SDKVersion:     "1.0.955-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10955-prerelease",
		RuntimeVersion: "93.0.967.0",
		Notes: 			`
#### General
*  WebView2 performance improvements.
*  Added partial Event Tracing for Windows (ETW) support.
*  Removed Microsoft branding from 'edge://history'.
*  New default Download UI.

###### Experimental features

*  Added [OpenTaskManagerWindow](/microsoft-edge/webview2/reference/win32/icorewebview2experimental4?view=webview2-1.0.955-prerelease&preserve-view=true#opentaskmanagerwindow) to launch a WebView2 browser task manager.
*  Added [NewWindowRequestedEventArgs](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalnewwindowrequestedeventargs?view=webview2-1.0.955-prerelease&preserve-view=true#get_name).
*  Added support for virtual host name mapping to work with Service Workers.
*  Added [HiddenPdfToolbarItems](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsettings6?view=webview2-1.0.955-prerelease&preserve-view=true#get_hiddenpdftoolbaritems) to customize the PDF toolbar items.

###### Bug fixes

*  Fixed bug that broke the 'edge://downloads' and 'edge://history' pages. This change is Runtime-specific.
*  Fixed bugs to improve reliability in the WebView2Loader.dll.
*  Fixed bug in which 'NewWindowRequested' event handler launched two windows when handling links that use 'target=_blank'.
*  Fixed a bug in WebView2 visual hosting that flickered before startup.
*  Fixed bug when 'add_WebResourceRequested' didn't work on WebView2 controls created using 'add_NewWindowRequested'. ([Issue #616](https://github.com/MicrosoftEdge/WebViewFeedback/issues/616))
*  Allow the host app to set foreground on a different application in response to events including 'NavigationStarting', 'AddHostObjectToScript' methods, 'WebMessageReceived', and 'NewWindowRequested'. ([Issue #1092](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1092))
*  Fix bug to trigger the 'PermissionRequested' event for the microphone. This change is Runtime-specific.([Issue #1462](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1462))
*  Fixed bug when 'ExecuteScriptAsync' blocked after several successful runs. This change is Runtime-specific. ([Issue #1348](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1348))
*  Fixed bug preventing non-ASCII filenames from being used in 'ResultFilePath' in 'DownloadStartingEventArgs'. ([Issue #1428](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1428))
*  Fixed bug where the title bar on the default pop-up wasn't displayed completely. This change is Runtime-specific. ([Issue #1016](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1016))

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  [add_ClientCertificateRequested](/microsoft-edge/webview2/reference/win32/icorewebview2_5?view=webview2-1.0.955-prerelease&preserve-view=true#add_clientcertificaterequested)

#### .NET

###### Bug fixes
*  Fixed an issue in WebView2 .NET API reference documentation that caused only the first exception to be displayed.
*  .NET core libraries are now built in release mode. To debug, ensure you clear the **Just my code** checkbox.
*  Fixed a bug that crashed WebView2 on forms with child forms. The child form, with the find in page bar open, caused WebView2 to crash when the child form was closed. ([Issue #1097](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1097))


<!-- ====================================================================== -->
Release Date: July 26, 2021`,
	},
	"1.0.902.49": {
		SDKVersion:     "1.0.902.49",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1090249",
		RuntimeVersion: "92.0.902.49",
		Notes: 			`
#### General

###### Bug fixes
*  Fix bug that broke the 'IsBuiltInErrorPageEnabled' property, which turned off the error page that's displayed when there's a navigation failure or render process failure.  This change is Runtime-specific. ([Issue #634](https://github.com/MicrosoftEdge/WebViewFeedback/issues/634))
*  Fixed an issue where WebView2 controls took focus away from the user's focus.
*  Fixed bug when 'AddScriptToExecuteOnDocumentCreated' didn't work on child windows. ([Issue #935](https://github.com/MicrosoftEdge/WebViewFeedback/issues/935))
*  Fixed a bug that caused inactive tabs to be automatically discarded. ([Issue #637](https://github.com/MicrosoftEdge/WebViewFeedback/issues/637))
*  Fixed a bug when a navigation event was interrupted by another navigation event resulting in the Navigation ID of 'NavigationCompleted' events to be incorrect. ([Issue #1142](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1142))

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

*  [add_FrameCreated](/microsoft-edge/webview2/reference/win32/icorewebview2_4?view=webview2-1.0.902.49&preserve-view=true#add_framecreated).
*  [get_IsGeneralAutofillEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2settings4?view=webview2-1.0.902.49&preserve-view=true#get_isgeneralautofillenabled).
*  [get_IsPinchZoomEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2settings5?view=webview2-1.0.902.49&preserve-view=true#get_ispinchzoomenabled).
*  [The Download APIs](/microsoft-edge/webview2/reference/win32/icorewebview2_4?view=webview2-1.0.902-prerelease&preserve-view=true#add_downloadstarting).
*  [AddHostObjectToScriptWithOrigins](/microsoft-edge/webview2/reference/win32/icorewebview2frame?view=webview2-1.0.902-prerelease&preserve-view=true#addhostobjecttoscriptwithorigins) API with iframe element support.


<!-- ====================================================================== -->
Release Date: June 1, 2021`,
	},
	"1.0.902-prerelease": {
		SDKVersion:     "1.0.902-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10902-prerelease",
		RuntimeVersion: "92.0.902.0",
		Notes: 			`
#### General

*  Improved WebView2 startup performance and disk footprint.

###### Experimental features

*  Added [IsSwipeNavigationEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsettings5?view=webview2-1.0.902-prerelease&preserve-view=true#get_isswipenavigationenabled) property to enable or disable the ability of the end user to use swiping gesture on touch input-enabled devices to navigate in WebView2.
*  Added [BrowserProcessExited](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment4?view=webview2-1.0.902-prerelease&preserve-view=true#add_browserprocessexited) event.
*  Added [add_ClientCertificateRequested API](/microsoft-edge/webview2/reference/win32/icorewebview2experimental3?view=webview2-1.0.902-prerelease&preserve-view=true#add_clientcertificaterequested). It allows showing a client certificate dialog prompt if desired and enables access to required metadata to replace default client certificate dialog prompt.

###### Bug fixes

*  Fix a bug where mouse left click doesn't dismiss context menu. This change is Runtime-specific.
*  Fixed a bug where WebView2 creation fails when exe files for apps sharing the same user data folder have inconsistent version info.
*  Fixed a bug where special browser keys such as 'Refresh', 'Home', and 'Back' can't be disabled by 'AreBrowserAcceleratorKeysEnabled'. This change is Runtime-specific.
*  Fixed a bug in WebView2 .NET controls, where WebView2 windows are blank when created in the background. ([Issue #1077](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1077)).
*  Dismissing a file picker dialog by pressing **Enter** or **Esc** no longer crashes WPF applications using WebView2 control. ([Issue #1099](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1099)).
*  Fixed a bug that [AllowSingleSignOnUsingOSPrimaryAccount](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions#get_allowsinglesignonusingosprimaryaccount) doesn't work properly with WebView2 when a 'WebResourceRequested' event handler is attached. This change is Runtime-specific. ([Issue #1183](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1183)).
*  Downloading a file no longer breaks WebView2 'DefaultBackgroundColor' transparency. This change is Runtime-specific. ([Issue #1108](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1108)).
*  Removed screen sharing media picker message that contains Microsoft branding. ([Issue #940](https://github.com/MicrosoftEdge/WebViewFeedback/issues/940)).
*  Fixed a bug in WebView2 WinForm control where hiding the parent form doesn't hide the WebView2 control ([Issue #828](https://github.com/MicrosoftEdge/WebViewFeedback/issues/828) and [Issue #1079](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1079)).
*  Added static WS_CLIPCHILDREN style to WebView2's WPF windows. ([Issue #1013](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1013)).
*  Fixed a bug where right-clicking a link crashes the WebView2 host app. This change is Runtime-specific.
*  Fixed a reliability bug that could crash the host app process when moving to a newer Edge WebView2 Runtime version.
*  **DEPRECATION**: Officially deprecated the 'DefaultBackgroundColor' API for Windows 7.

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  [Download API](/microsoft-edge/webview2/reference/win32/icorewebview2_4?view=webview2-1.0.902-prerelease&preserve-view=true#add_downloadstarting).
*  [PinchZoom API](/microsoft-edge/webview2/reference/win32/icorewebview2settings5?view=webview2-1.0.902-prerelease&preserve-view=true#get_ispinchzoomenabled).
*  [AddFrameCreated](/microsoft-edge/webview2/reference/win32/icorewebview2_4?view=webview2-1.0.902-prerelease&preserve-view=true#add_framecreated).
*  [AddHostObjectToScriptWithOrigins](/microsoft-edge/webview2/reference/win32/icorewebview2frame?view=webview2-1.0.902-prerelease&preserve-view=true#addhostobjecttoscriptwithorigins) API promoted to Stable with iframe element support.
*  [Autofill API](/microsoft-edge/webview2/reference/win32/icorewebview2settings4?view=webview2-1.0.902-prerelease&preserve-view=true#get_isgeneralautofillenabled).
   > [!NOTE]
   > There is no current API to delete the locally stored general autofill and password autosave information.  Please provide a control to delete the data, which will involve deleting the entire user data folder.

#### .NET

###### Bug fixes

*  Fixed a bug in WebView2 WinForm control where WebView2 window visibility isn't updated properly after parent window is disposed. ([Issue #1282](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1282) and [Issue #828](https://github.com/MicrosoftEdge/WebViewFeedback/issues/828)).
*  Fixed a bug in WebView2 WPF control that Source property binding in WPF OneWay binding mode isn't working properly. ([Issue #619](https://github.com/MicrosoftEdge/WebViewFeedback/issues/619) and [Issue #608](https://github.com/MicrosoftEdge/WebViewFeedback/issues/608)).


<!-- ====================================================================== -->
Release Date: May 31, 2021`,
	},
	"1.0.864.35": {
		SDKVersion:     "1.0.864.35",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1086435",
		RuntimeVersion: "91.0.864.35",
		Notes: 			`
#### General

###### Bug fixes

*  Fixed a reliability bug that could crash the host app process when moving to a newer Edge WebView2 Runtime version.
*  Fixed a bug that prevented memory purge in some situations. This change is Runtime-specific.
*  Fixed error in 818 SDK release package where project couldn't find the 'WebView2.h' file. ([Issue #1209](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1209)).
*  Fixed a bug which caused WebResourceRequested event to be dropped for some requests with binary bodies.
*  Improve 'NewWindowRequested' documentation. ([Issue #448](https://github.com/MicrosoftEdge/WebViewFeedback/issues/448)).

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

*  [UserAgent API](/microsoft-edge/webview2/reference/win32/icorewebview2settings2?view=webview2-1.0.864.35&preserve-view=true#get_useragent)
*  [AreBrowserkeysenabled](/microsoft-edge/webview2/reference/win32/icorewebview2settings3?view=webview2-1.0.864.35&preserve-view=true#get_arebrowseracceleratorkeysenabled)

#### .NET

###### Bug fixes

*  Fixed a bug in WebView2 .NET controls that first header is missing when iterating 'CoreWebView2WebResourceRequest' headers collection. ([Issue #1123](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1123)).


<!-- ====================================================================== -->
Release Date: April 26, 2021`,
	},
	"1.0.865-prerelease": {
		SDKVersion:     "1.0.865-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10865-prerelease",
		RuntimeVersion: "91.0.865.0",
		Notes: 			`
#### General

###### Experimental features

*  Added [IsPinchZoomEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsettings4?view=webview2-1.0.865-prerelease&preserve-view=true#ispinchzoomenabled) setting. It allows you to turn on or off page scale zoom control in a setting.
*  Added Custom [add_DownloadStarting](/microsoft-edge/webview2/reference/win32/icorewebview2experimental2?view=webview2-1.0.865-prerelease&preserve-view=true#add_downloadstarting) API.  It allows you to block downloads, save to a different path, and access the required metadata to build custom download UI.
*  Added 'iframe' element support from [AddHostObjectToScriptWithOrigins](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalframe?view=webview2-1.0.865-prerelease&preserve-view=true#addhostobjecttoscriptwithorigins).
*  Added sample code for [WPF sample app](https://github.com/MicrosoftEdge/WebView2Samples/tree/main/SampleApps/WebView2WpfBrowser) to use the API to turn off browser function keys.
*  Added the [UpdateRuntime](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironment3?view=webview2-1.0.865-prerelease&preserve-view=true#updateruntime) API, to easily update the WebView2 Runtime.

###### Bug fixes

*  Fixed handler for a 'Chromium DevTools Protocol' message with 'POST' binary data in WebView2.
*  Turned off the 'OpenSaveAsAwareness' download UI, because it included links to 'edge://settings'.  ([Issue #1120](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1120)).
*  Removed branding from screen share dialog.  ([Issue #940](https://github.com/MicrosoftEdge/WebViewFeedback/issues/940)).
*  Fixed bug where the [SetWindowDisplayAffinity](/windows/win32/api/winuser/nf-winuser-setwindowdisplayaffinity) function broke WebView2 when it stopped screen capture in an WebView2 app.  ([Issue #841](https://github.com/MicrosoftEdge/WebViewFeedback/issues/841)).
*  Fixed bug for composition hosting where mouse input stopped working if any pen input was sent to WebView2.
*  Fixed bug that broke mouse input after any pen input.  This change is Runtime-specific.

#### .NET

###### Experimental features

*  Added WebView2 designer tool to WPF Toolbox.  ([Issue #210](https://github.com/MicrosoftEdge/WebViewFeedback/issues/210)).
*  Added WebView2 UI element in .NET Designer Mode.

###### Bug fixes

*  Improved COM Exception descriptions by wrapping each in a more detailed .NET exception.  ([Issue #338](https://github.com/MicrosoftEdge/WebViewFeedback/issues/338)).  This change is Runtime-specific.
*  Fixed bug caused when you select **Tab** to shift focus caused WebView2 control to crash in Microsoft Visual Studio Tools for Office.  ([Issue #589](https://github.com/MicrosoftEdge/WebViewFeedback/issues/589) and [Issue #933](https://github.com/MicrosoftEdge/WebViewFeedback/issues/933)).  This change is Runtime-specific.
*  Improved .NET framework loader down level to be more robust.  ([Issue #946](https://github.com/MicrosoftEdge/WebViewFeedback/issues/946)).
*  Fixed bug that caused crash when you try to refresh before first navigation completed.  ([Issue #1011](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1011)).
*  Fixed initialization so navigation occurs during 'CoreWebView2InitializationCompleted'.  ([Issue #1050](https://github.com/MicrosoftEdge/WebViewFeedback/issues/1050)).
*  Improved .NET browser process crash error handling.  You can now recreate controls after you handle a 'ProcessFailed' event, without a crash.  ([Issue #996](https://github.com/MicrosoftEdge/WebViewFeedback/issues/996)).


<!-- ====================================================================== -->
Release Date: April 21, 2021`,
	},
	"1.0.818.41": {
		SDKVersion:     "1.0.818.41",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1081841",
		RuntimeVersion: "90.0.818.41",
		Notes: 			`
#### General

###### Features

*  Extended the 'ProcessFailed' event.  It now raises for non-renderer child processes and frame renderers.
*  Added 'iframe' element support for 'AddScriptToExecuteOnDocumentCreated'.
*  Improved WebView2 code to be more resilient to '.exe' application files with malformatted version information.  ([Issue #850](https://github.com/MicrosoftEdge/WebViewFeedback/issues/850)).
*  Removed '--winhttp-proxy-resolver' from WebView2 browser process command-line, turned on other proxy command-line options for WebView2.


<!-- ====================================================================== -->
Release Date: March 8, 2021`,
	},
	"1.0.824-prerelease": {
		SDKVersion:     "1.0.824-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10824-prerelease",
		RuntimeVersion: "91.0.824.0",
		Notes: 			`
#### General

###### Features

*  Extended the 'ProcessFailed' event.  It now raises for non-renderer child processes and frame renderers.
*  Added experimental [AreBrowserAcceleratorKeysEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsettings2?view=webview2-1.0.824&preserve-view=true#get_arebrowseracceleratorkeysenabled) setting.  You can prevent the browser from responding to keyboard shortcuts related to navigation, printing, saving, and other browser-specific functions.
*  Added 'iframe' element support for 'AddScriptToExecuteOnDocumentCreated'.

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  [UserAgent](/microsoft-edge/webview2/reference/win32/icorewebview2_2?view=webview2-1.0.721-prerelease&preserve-view=true#add_webresourceresponsereceived).

*  Rasterization Scale APIs:
   *  [RasterizationScale property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontroller?view=webview2-1.0.721-prerelease&preserve-view=true#get_rasterizationscale)
   *  [RasterizationScaleChanged event](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontroller?view=webview2-1.0.721-prerelease&preserve-view=true#add_rasterizationscalechanged)
   *  [BoundsMode property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontroller?view=webview2-1.0.721-prerelease&preserve-view=true#get_boundsmode)
   *  [ShouldDetectMonitorScaleChanges property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontroller?view=webview2-1.0.721-prerelease&preserve-view=true#get_shoulddetectmonitorscalechanges)

###### Bug fixes

*  Expanded supported C++ and .NET project types such as MFC and ATL.  ([Issue #506](https://github.com/MicrosoftEdge/WebViewFeedback/issues/506), [Issue #669](https://github.com/MicrosoftEdge/WebViewFeedback/issues/669), and [Issue #851](https://github.com/MicrosoftEdge/WebViewFeedback/issues/851)).
*  Fixed a bug that Evergreen WebView2 Runtime leaks Inbound firewall entry.
*  Fixed setting Response during 'WebResourceRequested' event.  ([Issue #568](https://github.com/MicrosoftEdge/WebViewFeedback/issues/568)).
*  Fixed a bug that navigating to 'edge://' causes browser process to exit.  ([Issue #604](https://github.com/MicrosoftEdge/WebViewFeedback/issues/604)).
*  Fixed a bug that limited WebView2 bounds to size of screen in Visual Hosting mode.


<!-- ====================================================================== -->
Release Date: March 8, 2021`,
	},
	"1.0.774.44": {
		SDKVersion:     "1.0.774.44",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1077444",
		RuntimeVersion: "89.0.774.44",
		Notes: 			`
#### General

###### Features

*  Turned off various Microsoft Edge browser services in WebView2.
*  Visual Hosting APIs are now Generally Available.

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

   *  [DPI support](/microsoft-edge/webview2/reference/win32/icorewebview2_2?view=webview2-1.0.774.44&preserve-view=true#add_webresourceresponsereceived) related APIs
   *  Visual hosting APIs
   *  [SetVirtualHostNameToFolderMapping](/microsoft-edge/webview2/reference/win32/icorewebview2_3?view=webview2-1.0.774.44&preserve-view=true#setvirtualhostnametofoldermapping)
   *  [TrySuspend and Resume](/microsoft-edge/webview2/reference/win32/icorewebview2_3?view=webview2-1.0.774.44&preserve-view=true#trysuspend)
   *  [DefaultBackgroundColor](/microsoft-edge/webview2/reference/win32/icorewebview2controller2?view=webview2-1.0.774.44&preserve-view=true#get_defaultbackgroundcolor)

###### Bug fixes

*  Fixed a bug that limited WebView2 bounds to size of screen in Visual Hosting mode.


<!-- ====================================================================== -->
Release Date: February 10, 2021`,
	},
	"1.0.790-prerelease": {
		SDKVersion:     "1.0.790-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10790-prerelease",
		RuntimeVersion: "86.0.616.0",
		Notes: 			`
#### General

> [!IMPORTANT]
> **Breaking Change**:  WebView2 prerelease package 1.0.781 is deprecated.  Discontinue development with package 1.0.781.

> [!IMPORTANT]
> WebView2 prerelease package 0.9.430 is deprecated, and is removed with the next release.  If your WebView2 app uses the package, the WebView2 team recommends that you move to a newer package.

###### Features

*  Added [TrySuspend and Resume](/microsoft-edge/webview2/reference/win32/icorewebview2_3?view=webview2-1.0.790-prerelease&preserve-view=true#trysuspend) method to suspend and resume WebViews.
*  Added [SetVirtualHostNameToFolderMapping](/microsoft-edge/webview2/reference/win32/icorewebview2_3?view=webview2-1.0.790-prerelease&preserve-view=true#setvirtualhostnametofoldermapping) method that maps a virtual host name to a directory path.  ([Issue #37](https://github.com/MicrosoftEdge/WebViewFeedback/issues/37), [Issue #161](https://github.com/MicrosoftEdge/WebViewFeedback/issues/161), and [Issue #212](https://github.com/MicrosoftEdge/WebViewFeedback/issues/212)).
*  Added the [DefaultBackgroundColor](/microsoft-edge/webview2/reference/win32/icorewebview2controller2?view=webview2-1.0.790-prerelease&preserve-view=true#get_defaultbackgroundcolor) property to set the color and alpha-channel of the background.  ([Issue #414](https://github.com/MicrosoftEdge/WebViewFeedback/issues/414)).
*  Added the [UserAgent](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalsettings?view=webview2-1.0.790-prerelease&preserve-view=true#get_useragent) property to get or set the User Agent.  ([Issue #122](https://github.com/MicrosoftEdge/WebViewFeedback/issues/122)).
*  Replaced the 'CreateCookieWithCookie' method with the 'CopyCookie' method.
*  Added visual hosting support using the [ICoreWebView2CompositionController](/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller?view=webview2-1.0.790-prerelease&preserve-view=true) interface, which is created using the new 'CreateCoreWebView2CompositionController' method from 'ICoreWebView2Environment3'.

###### Bug fixes

*  Turned off the Microsoft Edge Shopping feature in WebView2.
*  Turned off the context menu in the PDF viewer when 'AreDefaultContextMenusEnabled' is 'false'.  ([Issue #605](https://github.com/MicrosoftEdge/WebViewFeedback/issues/605)).
*  Fixed a bug that returned 'E_NOINTERFACE' when querying 'ICoreWebView2' for 'ICoreWebView2Experimental'.  ([Issue #691](https://github.com/MicrosoftEdge/WebViewFeedback/issues/691)).
*  Fixed a bug that allowed navigation with malformed URIs when 'CoreWebView2NavigationStartingEventArgs.Cancel' is set to 'false'.  ([Issue #400](https://github.com/MicrosoftEdge/WebViewFeedback/issues/400)).
*  Fixed a bug that blocked 'window.print()' on pop-up windows with event handlers attached to 'NewWindowRequested' events.  ([Issue #409](https://github.com/MicrosoftEdge/WebViewFeedback/issues/409)).
*  Fixed Dynamic DPI issue when moving apps between different monitors.  ([Issue #58](https://github.com/MicrosoftEdge/WebViewFeedback/issues/58))
*  Improved the 'HRESULT' instances passed by [ICoreWebView2WebResourceResponseViewGetContentCompletedHandler::Invoke](/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseviewgetcontentcompletedhandler?view=webview2-1.0.790-prerelease&preserve-view=true#invoke).
*  Turned off autofill manage button.  ([Issue #585](https://github.com/MicrosoftEdge/WebViewFeedback/issues/585)).
*  Fixed Visual Studio crashes while you run 'WebView2.Dispose' when hosted in multiple windows.  ([Issue #816](https://github.com/MicrosoftEdge/WebViewFeedback/issues/816)) and [Issue #442](https://github.com/MicrosoftEdge/WebViewFeedback/issues/442)).
*  Fixed bug to display WebView2 control in Visual Studio Toolbox.  ([Issue #210](https://github.com/MicrosoftEdge/WebViewFeedback/issues/210)).
*  Reduced high CPU usage issues.  ([Issue #878](https://github.com/MicrosoftEdge/WebViewFeedback/issues/878)).
*  Fixed issues with deprecated 1.0.781-prerelease package.  ([Issue #875](https://github.com/MicrosoftEdge/WebViewFeedback/issues/875) and [Issue #878](https://github.com/MicrosoftEdge/WebViewFeedback/issues/878)).

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  Visual Hosting APIs
*  [SetVirtualHostNameToFolderMapping](/microsoft-edge/webview2/reference/win32/icorewebview2_3?view=webview2-1.0.790-prerelease&preserve-view=true#setvirtualhostnametofoldermapping)

#### .NET

###### Bug fixes

*  Fixed bug that crashed WebView2 apps that use the WPF SDK.  The crash occurred when pressing **F4** to close a window.  ([Issue #399](https://github.com/MicrosoftEdge/WebViewFeedback/issues/399)).
*  The WebView2 initialization screen is now transparent, instead of gray.  ([Issue #196](https://github.com/MicrosoftEdge/WebViewFeedback/issues/196)).


<!-- ====================================================================== -->
Release Date: January 25, 2021`,
	},
	"1.0.705.50": {
		SDKVersion:     "1.0.705.50",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1070550",
		RuntimeVersion: "86.0.616.0",
		Notes: 			`
#### General

###### Promotions

The following APIs have been promoted to Stable and are now included in this Release SDK.

   *  [WebResourceResponseReceived API](/microsoft-edge/webview2/reference/win32/icorewebview2_2?view=webview2-1.0.721-prerelease&preserve-view=true#add_webresourceresponsereceived)
   *  [NavigateWithWebResourceRequest API](/microsoft-edge/webview2/reference/win32/icorewebview2environment2?view=webview2-1.0.721-prerelease&preserve-view=true#createwebresourcerequest)
   *  [Cookie management API](/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager?view=webview2-1.0.721-prerelease&preserve-view=true)
   *  [DOMContentLoaded API](/microsoft-edge/webview2/reference/win32/icorewebview2_2?view=webview2-1.0.721-prerelease&preserve-view=true#add_domcontentloaded)
   *  [Environment property](/microsoft-edge/webview2/reference/win32/icorewebview2_2?view=webview2-1.0.721-prerelease&preserve-view=true#get_environment)


<!-- ====================================================================== -->
Release Date: December 8, 2020`,
	},
	"1.0.721-prerelease": {
		SDKVersion:     "1.0.721-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10721-prerelease",
		RuntimeVersion: "86.0.616.0",
		Notes: 			`
#### General

> [!IMPORTANT]
> **Breaking Change**:  WebView2 prerelease package 1.0.707 and package 0.9.628 are deprecated.  Discontinue development with package 1.0.707 and  package0.9.628.

###### Features

*  Added [WebView2 Group Policies](/deployedge/microsoft-edge-webview-policies).  For best practices, see [group policies for WebView2](concepts/enterprise.md#group-policies-for-webview2).
*  > [!IMPORTANT]
   > **Breaking Change**: Deprecated the old registry location.
   >
   > '''text
   > {Root}\Software\Policies\Microsoft\EmbeddedBrowserWebView\LoaderOverride\{AppId}
   > '''

*  Added support for [Drag and Drop](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcompositioncontroller3?view=webview2-1.0.721-prerelease&preserve-view=true) in WebView2.
*  Added APIs to handle DPI support.
   *  Added [RasterizationScale](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontroller?view=webview2-1.0.721-prerelease&preserve-view=true#get_rasterizationscale) property to change the DPI scale for WebView2 content and UI pop-ups, and associated [RasterizationScaleChanged](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontroller?view=webview2-1.0.721-prerelease&preserve-view=true#add_rasterizationscalechanged) event.
   *  Added [ShouldDetectMonitorScaleChanges](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontroller?view=webview2-1.0.721-prerelease&preserve-view=true#get_shoulddetectmonitorscalechanges) property to automatically update 'RasterizationScale' property if needed.
   *  Added [BoundsMode property](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcontroller?view=webview2-1.0.721-prerelease&preserve-view=true#get_boundsmode) to specify that the bounds are logic pixels and allow WebView2 to use 'RasterizationScale' for WebView2 pixel display, and WebView2 use the 'RasterizationScale' with the 'Bounds' to get the physical size.
*  Updated 'NewWindowRequested' event to handle **Ctrl+click** and **Shift+click**.  ([Issue #168](https://github.com/MicrosoftEdge/WebViewFeedback/issues/168) and [Issue #371](https://github.com/MicrosoftEdge/WebViewFeedback/issues/371)).

###### Promotions

The following APIs have been promoted from Experimental to Stable in this Prerelease SDK.

*  [WebResourceResponseReceived API](/microsoft-edge/webview2/reference/win32/icorewebview2_2?view=webview2-1.0.721-prerelease&preserve-view=true#add_webresourceresponsereceived)
*  [NavigateWithWebResourceRequest API](/microsoft-edge/webview2/reference/win32/icorewebview2environment2?view=webview2-1.0.721-prerelease&preserve-view=true#createwebresourcerequest)
*  [Cookie management API](/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager?view=webview2-1.0.721-prerelease&preserve-view=true)
*  [DOMContentLoaded API](/microsoft-edge/webview2/reference/win32/icorewebview2_2?view=webview2-1.0.721-prerelease&preserve-view=true#add_domcontentloaded)
*  [Environment property](/microsoft-edge/webview2/reference/win32/icorewebview2_2?view=webview2-1.0.721-prerelease&preserve-view=true#get_environment)

#### .NET

###### Features

*  Turned on WinForms designer in .NET Core 3.1+ and .NET 5.
*  Improved .NET cookie management.  ([Issue #611](https://github.com/MicrosoftEdge/WebViewFeedback/issues/611)).
*  Replaced 'CoreWebView2Ready' with [CoreWebView2InitializationCompleted](/dotnet/api/microsoft.web.webview2.core.corewebview2initializationcompletedeventargs).

###### Bug fixes

*  Added [AcceleratorKeyPressed](/dotnet/api/microsoft.web.webview2.wpf.webview2.acceleratorkeypressed) event to support 'AcceleratorKey' select in WebView2.  ([Issue #288](https://github.com/MicrosoftEdge/WebViewFeedback/issues/288)).
*  Removed unnecessary files from being output to WebView2 folders.  ([Issue #461](https://github.com/MicrosoftEdge/WebViewFeedback/issues/461)).
*  Improved host object API.  ([Issue #335](https://github.com/MicrosoftEdge/WebViewFeedback/issues/335) and [Issue #525](https://github.com/MicrosoftEdge/WebViewFeedback/issues/525)).


<!-- ====================================================================== -->
Release Date: November 20, 2020`,
	},
	"1.0.664.37": {
		SDKVersion:     "1.0.664.37",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1066437",
		RuntimeVersion: "86.0.616.0",
		Notes: 			`
#### General

> [!IMPORTANT]
> **Announcement**: .NET WPF/WinForms WebView2 SDKs are now Generally Available (GA).  Starting with this release, Release SDKs are forward-compatible.  For more details, see [GA announcement blog post](https://devblogs.microsoft.com/dotnet/announcing-general-availability-for-microsoft-edge-webview2-for-net-and-fixed-distribution-method).

###### Features

*  .NET WPF/WinForms WebView2 is now Generally Available (GA).
*  Fixed Distribution (Bring-your-own) mode reached GA.

#### .NET

###### Bug fixes

*  'CoreWebView2NewWindowRequestedEventArgs.Handled' prevents new window from being opened.  ([Issue #549](https://github.com/MicrosoftEdge/WebViewFeedback/issues/549) and [Issue #560](https://github.com/MicrosoftEdge/WebViewFeedback/issues/560)).


<!-- ====================================================================== -->
Release Date: October 19, 2020`,
	},
	"1.0.674-prerelease": {
		SDKVersion:     "1.0.674-prerelease",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#10674-prerelease",
		RuntimeVersion: "86.0.616.0",
		Notes: 			`
#### General

*  Added [NavigateWithWebResourceRequest](/microsoft-edge/webview2/reference/win32/icorewebview2experimental?view=webview2-1.0.674-prerelease&preserve-view=true#navigatewithwebresourcerequest) method to provide post data or other request headers during navigation.
*  Added [DOMContentLoaded](/microsoft-edge/webview2/reference/win32/icorewebview2experimental?view=webview2-1.0.674-prerelease&preserve-view=true#add_domcontentloaded) event that runs when the initial HTML document is loaded and parsed.
*  Added the [Environment](/microsoft-edge/webview2/reference/win32/icorewebview2experimental?view=webview2-1.0.674-prerelease&preserve-view=true#get_environment) property on WebView2.  This property exposes the WebView2 environment where an instance of WebView2 was created.
*  Added [cookie management](/microsoft-edge/webview2/reference/win32/icorewebview2experimental?view=webview2-1.0.674-prerelease&preserve-view=true#get_cookiemanager) APIs that allow developers to authenticate the WebView2 session, or retrieve cookies from WebView2 to authenticate other tools.  The WebView2 team plans to make language or framework-specific improvements.  See [API Review: Cookie Management](https://github.com/MicrosoftEdge/WebView2Announcement/issues/2).
*  Updated the [WebResourceResponseReceived](/microsoft-edge/webview2/reference/win32/icorewebview2experimental?view=webview2-1.0.674-prerelease&preserve-view=true#add_webresourceresponsereceived) event, and added immutable [WebResourceResponseView](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebresourceresponseview?view=webview2-1.0.674-prerelease&preserve-view=true) and [WebResourceResponseReceivedEventArgs::PopulateResponseContent](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebresourceresponsereceivedeventargs?view=webview2-0.9.628-prerelease&preserve-view=true#populateresponsecontent) to [WebResourceResponseView::GetContent](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwebresourceresponseview?view=webview2-1.0.674-prerelease&preserve-view=true#getcontent).
*  Turned off [Microsoft Defender Application Guard (WDAG)](/windows/security/threat-protection/microsoft-defender-application-guard/md-app-guard-overview) in WebView2.
*  Added [SystemCursorId](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalcompositioncontroller2?view=webview2-1.0.674-prerelease&preserve-view=true#get_systemcursorid) for Visual Hosting.
*  Added bug fixed for Input Method in Visual Hosting.
*  Removed include requirement for 'version.lib' when using WebView2 static library.

#### .NET

*  Updated [CoreWebView2](/dotnet/api/microsoft.web.webview2.core.corewebview2) class to expose the 'CoreWebView2Environment' variable.
*  Changed implementations of custom EventArgs classes in 'Microsoft.Web.WebView2.Core' namespace to subclasses of [System.EventArgs](/dotnet/api/system.eventargs) or [System.ComponentModel.CancelEventArgs](/dotnet/api/system.componentmodel.canceleventargs).  ([Issue #250](https://github.com/MicrosoftEdge/WebViewFeedback/issues/250))
*  Added support for [CoreWebView2CreationProperties](/dotnet/api/microsoft.web.webview2.winforms) in WinForms.  ([Issue #204](https://github.com/MicrosoftEdge/WebViewFeedback/issues/204)).
*  Added [WebResourceRequested](/dotnet/api/microsoft.web.webview2.core.corewebview2.webresourcerequested) .NET APIs.  ([Issue #219](https://github.com/MicrosoftEdge/WebViewFeedback/issues/219)).
*  Updated WinForms Designer [Source](/dotnet/api/microsoft.web.webview2.winforms.webview2.source) property to default or reset to null.  ([Issue #177](https://github.com/MicrosoftEdge/WebViewFeedback/issues/177)).
*  Updated WebView2 bounds in WebView2.Init() to support DPI modes that are less than 100%.  ([Issue #432](https://github.com/MicrosoftEdge/WebViewFeedback/issues/432)).
*  Updated [BuildWindowCore](/dotnet/api/microsoft.web.webview2.wpf.webview2.buildwindowcore) and [DestroyWindowCore](/dotnet/api/microsoft.web.webview2.wpf.webview2.destroywindowcore) to increase robustness.  ([Issue #382](https://github.com/MicrosoftEdge/WebViewFeedback/issues/382)).
*  Updated .NET Loader base to load on process bit instead of operating system architecture.  ([Issue #431](https://github.com/MicrosoftEdge/WebViewFeedback/issues/431)).
*  Renamed 'EdgeNotFoundException' to [WebView2RuntimeNotFoundException](/dotnet/api/microsoft.web.webview2.core.webview2runtimenotfoundexception).


<!-- ====================================================================== -->
Release Date: October 19, 2020`,
	},
	"1.0.622.22": {
		SDKVersion:     "1.0.622.22",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#1062222",
		RuntimeVersion: "86.0.616.0",
		Notes: 			`
#### General

> [!IMPORTANT]
> **Announcement**:  Win32 C/C++ WebView2 is now Generally Available (GA).  Starting this release, Release SDKs are forward-compatible.  See [GA announcement blog post](https://blogs.windows.com/msedgedev/edge-webview2-general-availability).

*  The Evergreen WebView2 Runtime and installer are GA.  The bootstrapper, the downlink link for the Bootstrapper, and the Standalone Installer for the Evergreen WebView2 Runtime are available on [Microsoft Edge WebView2](https://developer.microsoft.com/microsoft-edge/webview2/).  Sample code for the installation workflow is also available in the [WebView2Samples repo](https://github.com/MicrosoftEdge/WebView2Samples).

For more information about the Runtime, Evergreen distribution, and Fixed Version distribution, see [Distribute your app and the WebView2 Runtime](concepts/distribution.md).


<!-- ====================================================================== -->
Release Date: September 10, 2020`,
	},
	"0.9.622.11": {
		SDKVersion:     "0.9.622.11",
		ReleaseNotes:   "https://learn.microsoft.com/en-us/microsoft-edge/webview2/release-notes?tabs=win32cpp#0962211",
		RuntimeVersion: "85.0.538.0",
		Notes: 			`
#### General

*  > [!IMPORTANT]
   > **Announcement**: This SDK is the Release Candidate for WebView2 Win32 C/C++ GA.  The GA version is expected to use the same API interface and functionality.

*  Disconnected [browser policies](/deployedge/microsoft-edge-policies).
*  Added [AllowSingleSignOnUsingOSPrimaryAccount](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions?view=webview2-0.9.622&preserve-view=true#get_allowsinglesignonusingosprimaryaccount) property on WebView2 environment options to turn on conditional access for WebView2.
*  Updated 'ICoreWebView2NewWindowRequestedEventArgs' to include [WindowFeatures](/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs?view=webview2-0.9.622&preserve-view=true#get_windowfeatures) property, and the associated [ICoreWebView2WindowFeatures](/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures?view=webview2-0.9.622&preserve-view=true).  ([Issue #293](https://github.com/MicrosoftEdge/WebViewFeedback/issues/293)).
*  Updated 'System.Windows.Rect'  to use 'System.Drawing.Rectangle' instead of 'System.Windows.Rect' ([Issue #235](https://github.com/MicrosoftEdge/WebViewFeedback/issues/235)).
*  Updated NewWindowRequested event to handle 'window.open()' request without parameters.  ([Issue #293](https://github.com/MicrosoftEdge/WebViewFeedback/issues/293)).
*  [AdditionalBrowserArguments](/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions?view=webview2-0.9.622&preserve-view=true#put_additionalbrowserarguments) specified with 'ICoreWebView2EnvironmentOptions' aren't overridden with environment variables or registry values.  See [CreateCoreWebView2EnvironmentWithOptions](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-0.9.622&preserve-view=true#createcorewebview2environmentwithoptions).


<!-- ====================================================================== -->
Release Date: July 20, 2020

#### General

*  > [!IMPORTANT]
   > **Announcement**: Evergreen WebView2 Runtime and installer is released for preview.  See [Distribute your app and the WebView2 Runtime](concepts/distribution.md).

*  > [!IMPORTANT]
   > **Announcement**:  The following WebView2 SDK Versions are no longer supported after the next SDK release:
   >
   > *  [0.8.190](#08190)
   > *  [0.8.230](#08230)
   > *  [0.8.270](#08270)
   > *  [0.8.314](#08314)
   > *  [0.8.355](#08355)
   >
   > The WebView2 SDK Versions are also marked deprecated on nuget.org.  WebView2 recommends staying up to date with the latest version of WebView2.

*  Added WebView2 worker thread improvements.  ([Issue #318](https://github.com/MicrosoftEdge/WebViewFeedback/issues/318)).
*  Turned off the pop-up blocker in WebView2.  See the [IsUserInitiated](/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs?view=webview2-0.9.538&preserve-view=true#get_isuserinitiated) property in the 'NewWindowRequested' event.
*  Ensured WebView2 navigation starting event is run for 'about:blank'.  Now, 'NavigationStarting' events are run for all navigation, but cancellations for 'about:blank' or 'srcdoc' of the 'iframe' element aren't supported and ignored.
*  Blocked some 'edge://' URI schemes in WebView2.
*  Added experimental [IsSingleSignOnUsingOSPrimaryAccountEnabled](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalenvironmentoptions?view=webview2-0.9.538-prerelease&preserve-view=true#get_issinglesignonusingosprimaryaccountenabled) property on WebView2 environment options to turn on conditional access for WebView2.
*  Added experimental [WebResourceResponseReceived](/microsoft-edge/webview2/reference/win32/icorewebview2experimental?view=webview2-0.9.538-prerelease&preserve-view=true#add_webresourceresponsereceived) event that runs after the WebView2 receives and processes the response from a WebResource request.  Authentication headers, if any, are included in the response object.

#### .NET

*  Improved WPF focus handling.  ([Issue #185](https://github.com/MicrosoftEdge/WebViewFeedback/issues/185)).
*  Added 'ZoomFactor' property on WPF Webview2 Controller.


<!-- ====================================================================== -->

#### General

*  Dropping support for WebView2 SDK Version [0.8.149](#08149).  WebView2 recommends staying up to date with the latest version of WebView2.
*  Updated group policy to account for when the profile path of the Microsoft Edge browser is modified  ([#179](https://github.com/MicrosoftEdge/WebViewFeedback/issues/179)).

#### Win32 C/C++

*  Added [ICoreWebView2ExperimentalNewWindowRequestedEventArgs::get_WindowFeatures](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalnewwindowrequestedeventargs?view=webview2-0.9.538-prerelease&preserve-view=true#get_windowfeatures), which fires when 'window.open()' is run and associated with [ICoreWebView2ExperimentalWindowFeatures](/microsoft-edge/webview2/reference/win32/icorewebview2experimentalwindowfeatures?view=webview2-0.9.538-prerelease&preserve-view=true) ([#70](https://github.com/MicrosoftEdge/WebViewFeedback/issues/70)).
*  > [!IMPORTANT]
   > **Breaking Change**:  Deprecated [CreateCoreWebView2EnvironmentWithDetails](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-0.9.488&preserve-view=true#createcorewebview2environmentwithdetails) and replaced with [CreateCoreWebView2EnvironmentWithOptions](/microsoft-edge/webview2/reference/win32/webview2-idl?view=webview2-0.9.538&preserve-view=true#createcorewebview2environmentwithoptions).

*  > [!IMPORTANT]
   > **Breaking Change**:  In order to ensure the WebView2 API aligns with the Windows API naming conventions, the WebView2 team updated the names of the following.
   >
   > *  [AreRemoteObjectsAllowed](/microsoft-edge/webview2/reference/win32/icorewebview2settings?view=webview2-0.9.488&preserve-view=true#get_areremoteobjectsallowed) is now [AreHostObjectsAllowed](/microsoft-edge/webview2/reference/win32/icorewebview2settings?view=webview2-0.9.538&preserve-view=true#get_arehostobjectsallowed).

*  Updated [AddHostObjectToScript](/microsoft-edge/webview2/reference/win32/icorewebview2?view=webview2-0.9.538&preserve-view=true#addhostobjecttoscript).  The original host object serializer markers are now set to the proxy objects.  Then host object serializer markers are serialized back as a host object when passed as a parameter in the JavaScript callback ([#148](https://github.com/MicrosoftEdge/WebViewFeedback/issues/148)).

#### .NET (0.9.538 prerelease)

*  Released WinForms and WPF WebView2API Samples, which are comprehensive guides of the WebView2 SDK.  See [Samples Repo](https://github.com/MicrosoftEdge/WebView2Samples).
*  Added support for visual hosting and window features [experimental APIs](concepts/versioning.md#experimental-apis).
*  > [!IMPORTANT]
   > **Breaking Change**:  The following deferrals now implement IDisposable:  [ScriptDialogOpening](/dotnet/api/microsoft.web.webview2.core.corewebview2.scriptdialogopening), [NewWindowRequested](/dotnet/api/microsoft.web.webview2.core.corewebview2.newwindowrequested), [WebResourceRequested](/dotnet/api/microsoft.web.webview2.core.corewebview2.webresourcerequested), and [PermissionRequested](/dotnet/api/microsoft.web.webview2.core.corewebview2.permissionrequested).

*  Added [GetAvailableBrowserVersionString](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.getavailablebrowserversionstring) and [CompareBrowserVersions](/dotnet/api/microsoft.web.webview2.core.corewebview2environment.comparebrowserversions) as [CoreWebView2Environment](/dotnet/api/microsoft.web.webview2.core.corewebview2environment) statics.


<!-- ====================================================================== -->`,
	},
}
