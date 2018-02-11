// Copyright 2018 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package mime

import "mime"

func init() {
	mime.AddExtensionType(".html", "text/html")
	mime.AddExtensionType(".htm", "text/html")
	mime.AddExtensionType(".shtml", "text/html")
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".xml", "text/xml")
	mime.AddExtensionType(".gif", "image/gif")
	mime.AddExtensionType(".jpeg", "image/jpeg")
	mime.AddExtensionType(".jpg", "image/jpeg")
	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".atom", "application/atom+xml")
	mime.AddExtensionType(".rss", "application/rss+xml")
	mime.AddExtensionType(".mml", "text/mathml")
	mime.AddExtensionType(".txt", "text/plain")
	mime.AddExtensionType(".jad", "text/vnd.sun.j2me.app-descriptor")
	mime.AddExtensionType(".wml", "text/vnd.wap.wml")
	mime.AddExtensionType(".htc", "text/x-component")
	mime.AddExtensionType(".png", "image/png")
	mime.AddExtensionType(".svg", "image/svg+xml")
	mime.AddExtensionType(".svgz", "image/svg+xml")
	mime.AddExtensionType(".tif", "image/tiff")
	mime.AddExtensionType(".tiff", "image/tiff")
	mime.AddExtensionType(".wbmp", "image/vnd.wap.wbmp")
	mime.AddExtensionType(".webp", "image/webp")
	mime.AddExtensionType(".ico", "image/x-icon")
	mime.AddExtensionType(".jng", "image/x-jng")
	mime.AddExtensionType(".bmp", "image/x-ms-bmp")
	mime.AddExtensionType(".woff", "application/font-woff")
	mime.AddExtensionType(".jar", "application/java-archive")
	mime.AddExtensionType(".war", "application/java-archive")
	mime.AddExtensionType(".ear", "application/java-archive")
	mime.AddExtensionType(".json", "application/json")
	mime.AddExtensionType(".hqx", "application/mac-binhex40")
	mime.AddExtensionType(".doc", "application/msword")
	mime.AddExtensionType(".pdf", "application/pdf")
	mime.AddExtensionType(".ps", "application/postscript")
	mime.AddExtensionType(".eps", "application/postscript")
	mime.AddExtensionType(".ai", "application/postscript")
	mime.AddExtensionType(".rtf", "application/rtf")
	mime.AddExtensionType(".m3u8", "application/vnd.apple.mpegurl")
	mime.AddExtensionType(".kml", "application/vnd.google-earth.kml+xml")
	mime.AddExtensionType(".kmz", "application/vnd.google-earth.kmz")
	mime.AddExtensionType(".xls", "application/vnd.ms-excel")
	mime.AddExtensionType(".eot", "application/vnd.ms-fontobject")
	mime.AddExtensionType(".ppt", "application/vnd.ms-powerpoint")
	mime.AddExtensionType(".odg", "application/vnd.oasis.opendocument.graphics")
	mime.AddExtensionType(".odp", "application/vnd.oasis.opendocument.presentation")
	mime.AddExtensionType(".ods", "application/vnd.oasis.opendocument.spreadsheet")
	mime.AddExtensionType(".odt", "application/vnd.oasis.opendocument.text")
	mime.AddExtensionType(".pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation")
	mime.AddExtensionType(".xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	mime.AddExtensionType(".docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	mime.AddExtensionType(".wmlc", "application/vnd.wap.wmlc")
	mime.AddExtensionType(".7z", "application/x-7z-compressed")
	mime.AddExtensionType(".cco", "application/x-cocoa")
	mime.AddExtensionType(".jardiff", "application/x-java-archive-diff")
	mime.AddExtensionType(".jnlp", "application/x-java-jnlp-file")
	mime.AddExtensionType(".run", "application/x-makeself")
	mime.AddExtensionType(".pl", "application/x-perl")
	mime.AddExtensionType(".pm", "application/x-perl")
	mime.AddExtensionType(".prc", "application/x-pilot")
	mime.AddExtensionType(".pdb", "application/x-pilot")
	mime.AddExtensionType(".rar", "application/x-rar-compressed")
	mime.AddExtensionType(".rpm", "application/x-redhat-package-manager")
	mime.AddExtensionType(".sea", "application/x-sea")
	mime.AddExtensionType(".swf", "application/x-shockwave-flash")
	mime.AddExtensionType(".sit", "application/x-stuffit")
	mime.AddExtensionType(".tcl", "application/x-tcl")
	mime.AddExtensionType(".tk", "application/x-tcl")
	mime.AddExtensionType(".der", "application/x-x509-ca-cert")
	mime.AddExtensionType(".pem", "application/x-x509-ca-cert")
	mime.AddExtensionType(".crt", "application/x-x509-ca-cert")
	mime.AddExtensionType(".xpi", "application/x-xpinstall")
	mime.AddExtensionType(".xhtml", "application/xhtml+xml")
	mime.AddExtensionType(".xspf", "application/xspf+xml")
	mime.AddExtensionType(".zip", "application/zip")
	mime.AddExtensionType(".bin", "application/octet-stream")
	mime.AddExtensionType(".exe", "application/octet-stream")
	mime.AddExtensionType(".dll", "application/octet-stream")
	mime.AddExtensionType(".deb", "application/octet-stream")
	mime.AddExtensionType(".dmg", "application/octet-stream")
	mime.AddExtensionType(".iso", "application/octet-stream")
	mime.AddExtensionType(".img", "application/octet-stream")
	mime.AddExtensionType(".msi", "application/octet-stream")
	mime.AddExtensionType(".msp", "application/octet-stream")
	mime.AddExtensionType(".msm", "application/octet-stream")
	mime.AddExtensionType(".mid", "audio/midi")
	mime.AddExtensionType(".midi", "audio/midi")
	mime.AddExtensionType(".kar", "audio/midi")
	mime.AddExtensionType(".mp3", "audio/mpeg")
	mime.AddExtensionType(".ogg", "audio/ogg")
	mime.AddExtensionType(".m4a", "audio/x-m4a")
	mime.AddExtensionType(".ra", "audio/x-realaudio")
	mime.AddExtensionType(".3gpp", "video/3gpp")
	mime.AddExtensionType(".3gp", "video/3gpp")
	mime.AddExtensionType(".ts", "video/mp2t")
	mime.AddExtensionType(".mp4", "video/mp4")
	mime.AddExtensionType(".mpeg", "video/mpeg")
	mime.AddExtensionType(".mpg", "video/mpeg")
	mime.AddExtensionType(".mov", "video/quicktime")
	mime.AddExtensionType(".webm", "video/webm")
	mime.AddExtensionType(".flv", "video/x-flv")
	mime.AddExtensionType(".m4v", "video/x-m4v")
	mime.AddExtensionType(".mng", "video/x-mng")
	mime.AddExtensionType(".asx", "video/x-ms-asf")
	mime.AddExtensionType(".asf", "video/x-ms-asf")
	mime.AddExtensionType(".wmv", "video/x-ms-wmv")
	mime.AddExtensionType(".avi", "video/x-msvideo")
}