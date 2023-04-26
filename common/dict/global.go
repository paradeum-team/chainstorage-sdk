package dict

import (
	"time"
)

const (
	SysTimeFmt       = "2006-01-02 15:04:05"
	ClientTypePcWeb  = "pcweb"
	ClientTypeMobile = "mobile"
	ClientTypeTablet = "tablet"
	ClientTypeOther  = "other"

	LoginUserInfoKey = "LoginUserInfo"
	DefaultLanguage  = "en"

	ArticlesType = "article"
	BountyType   = "bounty"
)

// 数据状态类型 1:enabled;2:deleted;
const (
	RecordEnabled = iota + 0
	RecordDeleted
)

// 数据状态类型映射关系
var RecordStatus = map[int]string{
	RecordEnabled: "使用中",
	RecordDeleted: "已删除",
}

// 强制覆盖设置
const (
	KeepObject = iota + 0
	ForecOverwrite
)

// 强制覆盖设置映射关系
var OverwriteSetting = map[int]string{
	KeepObject:     "保持原对象",
	ForecOverwrite: "强制覆盖",
}

// 是否星标 0:ObjectNotMarked;1:ObjectMarked;
const (
	ObjectNotMarked = iota + 0
	ObjectMarked
)

// 对象标记映射关系
var ObjectMarkedType = map[int]string{
	ObjectNotMarked: "未标记",
	ObjectMarked:    "已标记",
}

// 邮箱验证状态（1：已验证；0：未验证）
const (
	MailNotVerify = iota + 0
	MailVerified
)

// 邮箱验证状态映射关系
var MailVerifyStatus = map[int]string{
	MailNotVerify: "未验证",
	MailVerified:  "已验证",
}

// 文件扩展名
const (
	FileExtensionEmpty  = "empty"
	FileExtensionDir    = "dir"
	FileExtensionAac    = ".aac"
	FileExtensionAbw    = ".abw"
	FileExtensionArc    = ".arc"
	FileExtensionAvi    = ".avi"
	FileExtensionAzw    = ".azw"
	FileExtensionBin    = ".bin"
	FileExtensionBmp    = ".bmp"
	FileExtensionBz     = ".bz"
	FileExtensionBz2    = ".bz2"
	FileExtensionCsh    = ".csh"
	FileExtensionCss    = ".css"
	FileExtensionCsv    = ".csv"
	FileExtensionDoc    = ".doc"
	FileExtensionDocx   = ".docx"
	FileExtensionEot    = ".eot"
	FileExtensionEpub   = ".epub"
	FileExtensionGif    = ".gif"
	FileExtensionHtm    = ".htm"
	FileExtensionHtml   = ".html"
	FileExtensionIco    = ".ico"
	FileExtensionIcs    = ".ics"
	FileExtensionJar    = ".jar"
	FileExtensionJpeg   = ".jpeg"
	FileExtensionJpg    = ".jpg"
	FileExtensionJs     = ".js"
	FileExtensionJson   = ".json"
	FileExtensionJsonld = ".jsonld"
	FileExtensionMid    = ".mid"
	FileExtensionMidi   = ".midi"
	FileExtensionMjs    = ".mjs"
	FileExtensionMp3    = ".mp3"
	FileExtensionMpeg   = ".mpeg"
	FileExtensionMpkg   = ".mpkg"
	FileExtensionOdp    = ".odp"
	FileExtensionOds    = ".ods"
	FileExtensionOdt    = ".odt"
	FileExtensionOga    = ".oga"
	FileExtensionOgv    = ".ogv"
	FileExtensionOgx    = ".ogx"
	FileExtensionOtf    = ".otf"
	FileExtensionPng    = ".png"
	FileExtensionPdf    = ".pdf"
	FileExtensionPpt    = ".ppt"
	FileExtensionPptx   = ".pptx"
	FileExtensionRar    = ".rar"
	FileExtensionRtf    = ".rtf"
	FileExtensionSh     = ".sh"
	FileExtensionSvg    = ".svg"
	FileExtensionSwf    = ".swf"
	FileExtensionTar    = ".tar"
	FileExtensionTif    = ".tif"
	FileExtensionTiff   = ".tiff"
	FileExtensionTtf    = ".ttf"
	FileExtensionTxt    = ".txt"
	FileExtensionVsd    = ".vsd"
	FileExtensionWav    = ".wav"
	FileExtensionWeba   = ".weba"
	FileExtensionWebm   = ".webm"
	FileExtensionWebp   = ".webp"
	FileExtensionWoff   = ".woff"
	FileExtensionWoff2  = ".woff2"
	FileExtensionXhtml  = ".xhtml"
	FileExtensionXls    = ".xls"
	FileExtensionXlsx   = ".xlsx"
	FileExtensionXml    = ".xml"
	FileExtensionXul    = ".xul"
	FileExtensionZip    = ".zip"
	FileExtension3gp    = ".3gp"
	FileExtension3g2    = ".3g2"
	FileExtension7z     = ".7z"
	FileExtensionMp4    = ".mp4"
	FileExtensionApk    = ".apk"
	FileExtensionExe    = ".exe"
	FileExtensionDmg    = ".dmg"
)

// 文件类型
const (
	ObjectTypeCodeEmpty = 10000
	ObjectTypeCodeDir   = 20000
	ObjectTypeCodeAac   = iota + 29999
	ObjectTypeCodeAbw
	ObjectTypeCodeArc
	ObjectTypeCodeAvi
	ObjectTypeCodeAzw
	ObjectTypeCodeBin
	ObjectTypeCodeBmp
	ObjectTypeCodeBz
	ObjectTypeCodeBz2
	ObjectTypeCodeCsh
	ObjectTypeCodeCss
	ObjectTypeCodeCsv
	ObjectTypeCodeDoc
	ObjectTypeCodeDocx
	ObjectTypeCodeEot
	ObjectTypeCodeEpub
	ObjectTypeCodeGif
	ObjectTypeCodeHtm
	ObjectTypeCodeHtml
	ObjectTypeCodeIco
	ObjectTypeCodeIcs
	ObjectTypeCodeJar
	ObjectTypeCodeJpeg
	ObjectTypeCodeJpg
	ObjectTypeCodeJs
	ObjectTypeCodeJson
	ObjectTypeCodeJsonld
	ObjectTypeCodeMid
	ObjectTypeCodeMidi
	ObjectTypeCodeMjs
	ObjectTypeCodeMp3
	ObjectTypeCodeMpeg
	ObjectTypeCodeMpkg
	ObjectTypeCodeOdp
	ObjectTypeCodeOds
	ObjectTypeCodeOdt
	ObjectTypeCodeOga
	ObjectTypeCodeOgv
	ObjectTypeCodeOgx
	ObjectTypeCodeOtf
	ObjectTypeCodePng
	ObjectTypeCodePdf
	ObjectTypeCodePpt
	ObjectTypeCodePptx
	ObjectTypeCodeRar
	ObjectTypeCodeRtf
	ObjectTypeCodeSh
	ObjectTypeCodeSvg
	ObjectTypeCodeSwf
	ObjectTypeCodeTar
	ObjectTypeCodeTif
	ObjectTypeCodeTiff
	ObjectTypeCodeTtf
	ObjectTypeCodeTxt
	ObjectTypeCodeVsd
	ObjectTypeCodeWav
	ObjectTypeCodeWeba
	ObjectTypeCodeWebm
	ObjectTypeCodeWebp
	ObjectTypeCodeWoff
	ObjectTypeCodeWoff2
	ObjectTypeCodeXhtml
	ObjectTypeCodeXls
	ObjectTypeCodeXlsx
	ObjectTypeCodeXml
	ObjectTypeCodeXul
	ObjectTypeCodeZip
	ObjectTypeCode3gp
	ObjectTypeCode3g2
	ObjectTypeCode7z
	ObjectTypeCodeMp4
	ObjectTypeCodeApk
	ObjectTypeCodeExe
	ObjectTypeCodeDmg
)

// 链存对象类型
const (
	LinkedStorageTypeOther     = "Other"
	LinkedStorageTypeDirectory = "Directory"
	LinkedStorageTypeAudio     = "Audio"
	LinkedStorageTypeVideo     = "Video"
	LinkedStorageTypeImage     = "Image"
	LinkedStorageTypeDocument  = "Document"
	LinkedStorageTypeJson      = "Json"
)

// 文件扩展名与链存对象映射关系
var FileExtensionLinkedStorageObjectMapping = map[string]map[string]interface{}{
	FileExtensionEmpty:  {"Id": 10000, "FileExtension": "", "DocumentType": "", "MimeType": "", "LinkedStorageType": "", "LinkedStorageTypeName": "其他"},
	FileExtensionDir:    {"Id": 20000, "FileExtension": "", "DocumentType": "", "MimeType": "", "LinkedStorageType": "Directory", "LinkedStorageTypeName": "目录"},
	FileExtensionAac:    {"Id": 30001, "FileExtension": ".aac", "DocumentType": "AAC audio", "MimeType": "audio/aac", "LinkedStorageType": "Audio", "LinkedStorageTypeName": "音频"},
	FileExtensionAbw:    {"Id": 30002, "FileExtension": ".abw", "DocumentType": "AbiWord document", "MimeType": "application/x-abiword", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionArc:    {"Id": 30003, "FileExtension": ".arc", "DocumentType": "Archive document (multiple files embedded)", "MimeType": "application/x-freearc", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionAvi:    {"Id": 30004, "FileExtension": ".avi", "DocumentType": "AVI: Audio Video Interleave", "MimeType": "video/x-msvideo", "LinkedStorageType": "Video", "LinkedStorageTypeName": "视频"},
	FileExtensionAzw:    {"Id": 30005, "FileExtension": ".azw", "DocumentType": "Amazon Kindle eBook format", "MimeType": "application/vnd.amazon.ebook", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionBin:    {"Id": 30006, "FileExtension": ".bin", "DocumentType": "Any kind of binary data", "MimeType": "application/octet-stream", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionBmp:    {"Id": 30007, "FileExtension": ".bmp", "DocumentType": "Windows OS/2 Bitmap Graphics", "MimeType": "image/bmp", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionBz:     {"Id": 30008, "FileExtension": ".bz", "DocumentType": "BZip archive", "MimeType": "application/x-bzip", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionBz2:    {"Id": 30009, "FileExtension": ".bz2", "DocumentType": "BZip2 archive", "MimeType": "application/x-bzip2", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionCsh:    {"Id": 30010, "FileExtension": ".csh", "DocumentType": "C-Shell script", "MimeType": "application/x-csh", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionCss:    {"Id": 30011, "FileExtension": ".css", "DocumentType": "Cascading Style Sheets (CSS)", "MimeType": "text/css", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionCsv:    {"Id": 30012, "FileExtension": ".csv", "DocumentType": "Comma-separated values (CSV)", "MimeType": "text/csv", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionDoc:    {"Id": 30013, "FileExtension": ".doc", "DocumentType": "Microsoft Word", "MimeType": "application/msword", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionDocx:   {"Id": 30014, "FileExtension": ".docx", "DocumentType": "Microsoft Word (OpenXML)", "MimeType": "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionEot:    {"Id": 30015, "FileExtension": ".eot", "DocumentType": "MS Embedded OpenType fonts", "MimeType": "application/vnd.ms-fontobject", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionEpub:   {"Id": 30016, "FileExtension": ".epub", "DocumentType": "Electronic publication (EPUB)", "MimeType": "application/epub+zip", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionGif:    {"Id": 30017, "FileExtension": ".gif", "DocumentType": "Graphics Interchange Format (GIF)", "MimeType": "image/gif", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionHtm:    {"Id": 30018, "FileExtension": ".htm", "DocumentType": "HyperText Markup Language (HTML)", "MimeType": "text/html", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionHtml:   {"Id": 30019, "FileExtension": ".html", "DocumentType": "HyperText Markup Language (HTML)", "MimeType": "text/html", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionIco:    {"Id": 30020, "FileExtension": ".ico", "DocumentType": "Icon format", "MimeType": "image/vnd.microsoft.icon", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionIcs:    {"Id": 30021, "FileExtension": ".ics", "DocumentType": "iCalendar format", "MimeType": "text/calendar", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionJar:    {"Id": 30022, "FileExtension": ".jar", "DocumentType": "Java Archive (JAR)", "MimeType": "application/java-archive", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionJpeg:   {"Id": 30023, "FileExtension": ".jpeg", "DocumentType": "JPEG images", "MimeType": "image/jpeg", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionJpg:    {"Id": 30024, "FileExtension": ".jpg", "DocumentType": "JPEG images", "MimeType": "image/jpeg", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionJs:     {"Id": 30025, "FileExtension": ".js", "DocumentType": "JavaScript", "MimeType": "text/javascript", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionJson:   {"Id": 30026, "FileExtension": ".json", "DocumentType": "JSON format", "MimeType": "application/json", "LinkedStorageType": "JSON", "LinkedStorageTypeName": ""},
	FileExtensionJsonld: {"Id": 30027, "FileExtension": ".jsonld", "DocumentType": "JSON-LD format", "MimeType": "application/ld+json", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionMid:    {"Id": 30028, "FileExtension": ".mid", "DocumentType": "Musical Instrument Digital Interface (MIDI)", "MimeType": "audio/midi audio/x-midi", "LinkedStorageType": "Audio", "LinkedStorageTypeName": "音频"},
	FileExtensionMidi:   {"Id": 30029, "FileExtension": ".midi", "DocumentType": "Musical Instrument Digital Interface (MIDI)", "MimeType": "audio/midi audio/x-midi", "LinkedStorageType": "Audio", "LinkedStorageTypeName": "音频"},
	FileExtensionMjs:    {"Id": 30030, "FileExtension": ".mjs", "DocumentType": "JavaScript module", "MimeType": "text/javascript", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionMp3:    {"Id": 30031, "FileExtension": ".mp3", "DocumentType": "MP3 audio", "MimeType": "audio/mpeg", "LinkedStorageType": "Audio", "LinkedStorageTypeName": "音频"},
	FileExtensionMpeg:   {"Id": 30032, "FileExtension": ".mpeg", "DocumentType": "MPEG Video", "MimeType": "video/mpeg", "LinkedStorageType": "Video", "LinkedStorageTypeName": "视频"},
	FileExtensionMpkg:   {"Id": 30033, "FileExtension": ".mpkg", "DocumentType": "Apple Installer Package", "MimeType": "application/vnd.apple.installer+xml", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionOdp:    {"Id": 30034, "FileExtension": ".odp", "DocumentType": "OpenDocument presentation document", "MimeType": "application/vnd.oasis.opendocument.presentation", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionOds:    {"Id": 30035, "FileExtension": ".ods", "DocumentType": "OpenDocument spreadsheet document", "MimeType": "application/vnd.oasis.opendocument.spreadsheet", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionOdt:    {"Id": 30036, "FileExtension": ".odt", "DocumentType": "OpenDocument text document", "MimeType": "application/vnd.oasis.opendocument.text", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionOga:    {"Id": 30037, "FileExtension": ".oga", "DocumentType": "OGG audio", "MimeType": "audio/ogg", "LinkedStorageType": "Audio", "LinkedStorageTypeName": "音频"},
	FileExtensionOgv:    {"Id": 30038, "FileExtension": ".ogv", "DocumentType": "OGG video", "MimeType": "video/ogg", "LinkedStorageType": "Video", "LinkedStorageTypeName": "视频"},
	FileExtensionOgx:    {"Id": 30039, "FileExtension": ".ogx", "DocumentType": "OGG", "MimeType": "application/ogg", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionOtf:    {"Id": 30040, "FileExtension": ".otf", "DocumentType": "OpenType font", "MimeType": "font/otf", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionPng:    {"Id": 30041, "FileExtension": ".png", "DocumentType": "Portable Network Graphics", "MimeType": "image/png", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionPdf:    {"Id": 30042, "FileExtension": ".pdf", "DocumentType": "Adobe Portable Document Format (PDF)", "MimeType": "application/pdf", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionPpt:    {"Id": 30043, "FileExtension": ".ppt", "DocumentType": "Microsoft PowerPoint", "MimeType": "application/vnd.ms-powerpoint", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionPptx:   {"Id": 30044, "FileExtension": ".pptx", "DocumentType": "Microsoft PowerPoint (OpenXML)", "MimeType": "application/vnd.openxmlformats-officedocument.presentationml.presentation", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionRar:    {"Id": 30045, "FileExtension": ".rar", "DocumentType": "RAR archive", "MimeType": "application/x-rar-compressed", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionRtf:    {"Id": 30046, "FileExtension": ".rtf", "DocumentType": "Rich Text Format (RTF)", "MimeType": "application/rtf", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionSh:     {"Id": 30047, "FileExtension": ".sh", "DocumentType": "Bourne shell script", "MimeType": "application/x-sh", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionSvg:    {"Id": 30048, "FileExtension": ".svg", "DocumentType": "Scalable Vector Graphics (SVG)", "MimeType": "image/svg+xml", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionSwf:    {"Id": 30049, "FileExtension": ".swf", "DocumentType": "Small web format (SWF) or Adobe Flash document", "MimeType": "application/x-shockwave-flash", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionTar:    {"Id": 30050, "FileExtension": ".tar", "DocumentType": "Tape Archive (TAR)", "MimeType": "application/x-tar", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionTif:    {"Id": 30051, "FileExtension": ".tif", "DocumentType": "Tagged Image File Format (TIFF)", "MimeType": "image/tiff", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionTiff:   {"Id": 30052, "FileExtension": ".tiff", "DocumentType": "Tagged Image File Format (TIFF)", "MimeType": "image/tiff", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionTtf:    {"Id": 30053, "FileExtension": ".ttf", "DocumentType": "TrueType Font", "MimeType": "font/ttf", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionTxt:    {"Id": 30054, "FileExtension": ".txt", "DocumentType": "Text, (generally ASCII or ISO 8859-n)", "MimeType": "text/plain", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionVsd:    {"Id": 30055, "FileExtension": ".vsd", "DocumentType": "Microsoft Visio", "MimeType": "application/vnd.visio", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionWav:    {"Id": 30056, "FileExtension": ".wav", "DocumentType": "Waveform Audio Format", "MimeType": "audio/wav", "LinkedStorageType": "Audio", "LinkedStorageTypeName": "音频"},
	FileExtensionWeba:   {"Id": 30057, "FileExtension": ".weba", "DocumentType": "WEBM audio", "MimeType": "audio/webm", "LinkedStorageType": "Audio", "LinkedStorageTypeName": "音频"},
	FileExtensionWebm:   {"Id": 30058, "FileExtension": ".webm", "DocumentType": "WEBM video", "MimeType": "video/webm", "LinkedStorageType": "Video", "LinkedStorageTypeName": "视频"},
	FileExtensionWebp:   {"Id": 30059, "FileExtension": ".webp", "DocumentType": "WEBP image", "MimeType": "image/webp", "LinkedStorageType": "Image", "LinkedStorageTypeName": "图像"},
	FileExtensionWoff:   {"Id": 30060, "FileExtension": ".woff", "DocumentType": "Web Open Font Format (WOFF)", "MimeType": "font/woff", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionWoff2:  {"Id": 30061, "FileExtension": ".woff2", "DocumentType": "Web Open Font Format (WOFF)", "MimeType": "font/woff2", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionXhtml:  {"Id": 30062, "FileExtension": ".xhtml", "DocumentType": "XHTML", "MimeType": "application/xhtml+xml", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionXls:    {"Id": 30063, "FileExtension": ".xls", "DocumentType": "Microsoft Excel", "MimeType": "application/vnd.ms-excel", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionXlsx:   {"Id": 30064, "FileExtension": ".xlsx", "DocumentType": "Microsoft Excel (OpenXML)", "MimeType": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "LinkedStorageType": "Document", "LinkedStorageTypeName": "文档"},
	FileExtensionXml:    {"Id": 30065, "FileExtension": ".xml", "DocumentType": "XML", "MimeType": "application/xml", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionXul:    {"Id": 30066, "FileExtension": ".xul", "DocumentType": "XUL", "MimeType": "application/vnd.mozilla.xul+xml", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionZip:    {"Id": 30067, "FileExtension": ".zip", "DocumentType": "ZIP", "MimeType": "archive application/zip", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtension3gp:    {"Id": 30068, "FileExtension": ".3gp", "DocumentType": "3GPP audio/video container", "MimeType": "video/3gpp audio/3gpp", "LinkedStorageType": "Video", "LinkedStorageTypeName": "视频"},
	FileExtension3g2:    {"Id": 30069, "FileExtension": ".3g2", "DocumentType": "3GPP2 audio/video container", "MimeType": "video/3gpp2 audio/3gpp2", "LinkedStorageType": "Audio", "LinkedStorageTypeName": "音频"},
	FileExtension7z:     {"Id": 30070, "FileExtension": ".7z", "DocumentType": "7-zip", "MimeType": "archive application/x-7z-compressed", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionMp4:    {"Id": 30071, "FileExtension": ".mp4", "DocumentType": "MP4 Video", "MimeType": "video/mp4v-es", "LinkedStorageType": "Video", "LinkedStorageTypeName": "视频"},
	FileExtensionApk:    {"Id": 30072, "FileExtension": ".apk", "DocumentType": "Android Package Archive", "MimeType": "application/vnd.android.package-archive", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionExe:    {"Id": 30073, "FileExtension": ".exe", "DocumentType": "Microsoft Application", "MimeType": "application/x-msdownload", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
	FileExtensionDmg:    {"Id": 30074, "FileExtension": ".dmg", "DocumentType": "Apple Disk Image", "MimeType": "application/x-apple-diskimage", "LinkedStorageType": "Other", "LinkedStorageTypeName": "其他"},
}

// 对象类型代码与链存对象映射关系
var ObjectTypeCodeLinkedStorageObjectMapping = map[int]map[string]interface{}{
	ObjectTypeCodeEmpty:  {"id": 10000, "fileExtension": "", "documentType": "", "mimeType": "", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeDir:    {"id": 20000, "fileExtension": "", "documentType": "", "mimeType": "", "linkedStorageType": "Directory", "linkedStorageTypeName": "目录"},
	ObjectTypeCodeAac:    {"id": 30001, "fileExtension": ".aac", "documentType": "AAC audio", "mimeType": "audio/aac", "linkedStorageType": "Audio", "linkedStorageTypeName": "音频"},
	ObjectTypeCodeAbw:    {"id": 30002, "fileExtension": ".abw", "documentType": "AbiWord document", "mimeType": "application/x-abiword", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeArc:    {"id": 30003, "fileExtension": ".arc", "documentType": "Archive document (multiple files embedded)", "mimeType": "application/x-freearc", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeAvi:    {"id": 30004, "fileExtension": ".avi", "documentType": "AVI: Audio Video Interleave", "mimeType": "video/x-msvideo", "linkedStorageType": "Video", "linkedStorageTypeName": "视频"},
	ObjectTypeCodeAzw:    {"id": 30005, "fileExtension": ".azw", "documentType": "Amazon Kindle eBook format", "mimeType": "application/vnd.amazon.ebook", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeBin:    {"id": 30006, "fileExtension": ".bin", "documentType": "Any kind of binary data", "mimeType": "application/octet-stream", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeBmp:    {"id": 30007, "fileExtension": ".bmp", "documentType": "Windows OS/2 Bitmap Graphics", "mimeType": "image/bmp", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeBz:     {"id": 30008, "fileExtension": ".bz", "documentType": "BZip archive", "mimeType": "application/x-bzip", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeBz2:    {"id": 30009, "fileExtension": ".bz2", "documentType": "BZip2 archive", "mimeType": "application/x-bzip2", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeCsh:    {"id": 30010, "fileExtension": ".csh", "documentType": "C-Shell script", "mimeType": "application/x-csh", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeCss:    {"id": 30011, "fileExtension": ".css", "documentType": "Cascading Style Sheets (CSS)", "mimeType": "text/css", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeCsv:    {"id": 30012, "fileExtension": ".csv", "documentType": "Comma-separated values (CSV)", "mimeType": "text/csv", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeDoc:    {"id": 30013, "fileExtension": ".doc", "documentType": "Microsoft Word", "mimeType": "application/msword", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeDocx:   {"id": 30014, "fileExtension": ".docx", "documentType": "Microsoft Word (OpenXML)", "mimeType": "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeEot:    {"id": 30015, "fileExtension": ".eot", "documentType": "MS Embedded OpenType fonts", "mimeType": "application/vnd.ms-fontobject", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeEpub:   {"id": 30016, "fileExtension": ".epub", "documentType": "Electronic publication (EPUB)", "mimeType": "application/epub+zip", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeGif:    {"id": 30017, "fileExtension": ".gif", "documentType": "Graphics Interchange Format (GIF)", "mimeType": "image/gif", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeHtm:    {"id": 30018, "fileExtension": ".htm", "documentType": "HyperText Markup Language (HTML)", "mimeType": "text/html", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeHtml:   {"id": 30019, "fileExtension": ".html", "documentType": "HyperText Markup Language (HTML)", "mimeType": "text/html", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeIco:    {"id": 30020, "fileExtension": ".ico", "documentType": "Icon format", "mimeType": "image/vnd.microsoft.icon", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeIcs:    {"id": 30021, "fileExtension": ".ics", "documentType": "iCalendar format", "mimeType": "text/calendar", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeJar:    {"id": 30022, "fileExtension": ".jar", "documentType": "Java Archive (JAR)", "mimeType": "application/java-archive", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeJpeg:   {"id": 30023, "fileExtension": ".jpeg", "documentType": "JPEG images", "mimeType": "image/jpeg", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeJpg:    {"id": 30024, "fileExtension": ".jpg", "documentType": "JPEG images", "mimeType": "image/jpeg", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeJs:     {"id": 30025, "fileExtension": ".js", "documentType": "JavaScript", "mimeType": "text/javascript", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeJson:   {"id": 30026, "fileExtension": ".json", "documentType": "JSON format", "mimeType": "application/json", "linkedStorageType": "JSON", "linkedStorageTypeName": ""},
	ObjectTypeCodeJsonld: {"id": 30027, "fileExtension": ".jsonld", "documentType": "JSON-LD format", "mimeType": "application/ld+json", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeMid:    {"id": 30028, "fileExtension": ".mid", "documentType": "Musical Instrument Digital Interface (MIDI)", "mimeType": "audio/midi audio/x-midi", "linkedStorageType": "Audio", "linkedStorageTypeName": "音频"},
	ObjectTypeCodeMidi:   {"id": 30029, "fileExtension": ".midi", "documentType": "Musical Instrument Digital Interface (MIDI)", "mimeType": "audio/midi audio/x-midi", "linkedStorageType": "Audio", "linkedStorageTypeName": "音频"},
	ObjectTypeCodeMjs:    {"id": 30030, "fileExtension": ".mjs", "documentType": "JavaScript module", "mimeType": "text/javascript", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeMp3:    {"id": 30031, "fileExtension": ".mp3", "documentType": "MP3 audio", "mimeType": "audio/mpeg", "linkedStorageType": "Audio", "linkedStorageTypeName": "音频"},
	ObjectTypeCodeMpeg:   {"id": 30032, "fileExtension": ".mpeg", "documentType": "MPEG Video", "mimeType": "video/mpeg", "linkedStorageType": "Video", "linkedStorageTypeName": "视频"},
	ObjectTypeCodeMpkg:   {"id": 30033, "fileExtension": ".mpkg", "documentType": "Apple Installer Package", "mimeType": "application/vnd.apple.installer+xml", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeOdp:    {"id": 30034, "fileExtension": ".odp", "documentType": "OpenDocument presentation document", "mimeType": "application/vnd.oasis.opendocument.presentation", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeOds:    {"id": 30035, "fileExtension": ".ods", "documentType": "OpenDocument spreadsheet document", "mimeType": "application/vnd.oasis.opendocument.spreadsheet", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeOdt:    {"id": 30036, "fileExtension": ".odt", "documentType": "OpenDocument text document", "mimeType": "application/vnd.oasis.opendocument.text", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeOga:    {"id": 30037, "fileExtension": ".oga", "documentType": "OGG audio", "mimeType": "audio/ogg", "linkedStorageType": "Audio", "linkedStorageTypeName": "音频"},
	ObjectTypeCodeOgv:    {"id": 30038, "fileExtension": ".ogv", "documentType": "OGG video", "mimeType": "video/ogg", "linkedStorageType": "Video", "linkedStorageTypeName": "视频"},
	ObjectTypeCodeOgx:    {"id": 30039, "fileExtension": ".ogx", "documentType": "OGG", "mimeType": "application/ogg", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeOtf:    {"id": 30040, "fileExtension": ".otf", "documentType": "OpenType font", "mimeType": "font/otf", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodePng:    {"id": 30041, "fileExtension": ".png", "documentType": "Portable Network Graphics", "mimeType": "image/png", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodePdf:    {"id": 30042, "fileExtension": ".pdf", "documentType": "Adobe Portable Document Format (PDF)", "mimeType": "application/pdf", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodePpt:    {"id": 30043, "fileExtension": ".ppt", "documentType": "Microsoft PowerPoint", "mimeType": "application/vnd.ms-powerpoint", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodePptx:   {"id": 30044, "fileExtension": ".pptx", "documentType": "Microsoft PowerPoint (OpenXML)", "mimeType": "application/vnd.openxmlformats-officedocument.presentationml.presentation", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeRar:    {"id": 30045, "fileExtension": ".rar", "documentType": "RAR archive", "mimeType": "application/x-rar-compressed", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeRtf:    {"id": 30046, "fileExtension": ".rtf", "documentType": "Rich Text Format (RTF)", "mimeType": "application/rtf", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeSh:     {"id": 30047, "fileExtension": ".sh", "documentType": "Bourne shell script", "mimeType": "application/x-sh", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeSvg:    {"id": 30048, "fileExtension": ".svg", "documentType": "Scalable Vector Graphics (SVG)", "mimeType": "image/svg+xml", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeSwf:    {"id": 30049, "fileExtension": ".swf", "documentType": "Small web format (SWF) or Adobe Flash document", "mimeType": "application/x-shockwave-flash", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeTar:    {"id": 30050, "fileExtension": ".tar", "documentType": "Tape Archive (TAR)", "mimeType": "application/x-tar", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeTif:    {"id": 30051, "fileExtension": ".tif", "documentType": "Tagged Image File Format (TIFF)", "mimeType": "image/tiff", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeTiff:   {"id": 30052, "fileExtension": ".tiff", "documentType": "Tagged Image File Format (TIFF)", "mimeType": "image/tiff", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeTtf:    {"id": 30053, "fileExtension": ".ttf", "documentType": "TrueType Font", "mimeType": "font/ttf", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeTxt:    {"id": 30054, "fileExtension": ".txt", "documentType": "Text, (generally ASCII or ISO 8859-n)", "mimeType": "text/plain", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeVsd:    {"id": 30055, "fileExtension": ".vsd", "documentType": "Microsoft Visio", "mimeType": "application/vnd.visio", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeWav:    {"id": 30056, "fileExtension": ".wav", "documentType": "Waveform Audio Format", "mimeType": "audio/wav", "linkedStorageType": "Audio", "linkedStorageTypeName": "音频"},
	ObjectTypeCodeWeba:   {"id": 30057, "fileExtension": ".weba", "documentType": "WEBM audio", "mimeType": "audio/webm", "linkedStorageType": "Audio", "linkedStorageTypeName": "音频"},
	ObjectTypeCodeWebm:   {"id": 30058, "fileExtension": ".webm", "documentType": "WEBM video", "mimeType": "video/webm", "linkedStorageType": "Video", "linkedStorageTypeName": "视频"},
	ObjectTypeCodeWebp:   {"id": 30059, "fileExtension": ".webp", "documentType": "WEBP image", "mimeType": "image/webp", "linkedStorageType": "Image", "linkedStorageTypeName": "图像"},
	ObjectTypeCodeWoff:   {"id": 30060, "fileExtension": ".woff", "documentType": "Web Open Font Format (WOFF)", "mimeType": "font/woff", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeWoff2:  {"id": 30061, "fileExtension": ".woff2", "documentType": "Web Open Font Format (WOFF)", "mimeType": "font/woff2", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeXhtml:  {"id": 30062, "fileExtension": ".xhtml", "documentType": "XHTML", "mimeType": "application/xhtml+xml", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeXls:    {"id": 30063, "fileExtension": ".xls", "documentType": "Microsoft Excel", "mimeType": "application/vnd.ms-excel", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeXlsx:   {"id": 30064, "fileExtension": ".xlsx", "documentType": "Microsoft Excel (OpenXML)", "mimeType": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "linkedStorageType": "Document", "linkedStorageTypeName": "文档"},
	ObjectTypeCodeXml:    {"id": 30065, "fileExtension": ".xml", "documentType": "XML", "mimeType": "application/xml", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeXul:    {"id": 30066, "fileExtension": ".xul", "documentType": "XUL", "mimeType": "application/vnd.mozilla.xul+xml", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeZip:    {"id": 30067, "fileExtension": ".zip", "documentType": "ZIP", "mimeType": "archive application/zip", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCode3gp:    {"id": 30068, "fileExtension": ".3gp", "documentType": "3GPP audio/video container", "mimeType": "video/3gpp audio/3gpp", "linkedStorageType": "Video", "linkedStorageTypeName": "视频"},
	ObjectTypeCode3g2:    {"id": 30069, "fileExtension": ".3g2", "documentType": "3GPP2 audio/video container", "mimeType": "video/3gpp2 audio/3gpp2", "linkedStorageType": "Audio", "linkedStorageTypeName": "音频"},
	ObjectTypeCode7z:     {"id": 30070, "fileExtension": ".7z", "documentType": "7-zip", "mimeType": "archive application/x-7z-compressed", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeMp4:    {"id": 30071, "fileExtension": ".mp4", "documentType": "MP4 Video", "mimeType": "video/mp4", "linkedStorageType": "Video", "linkedStorageTypeName": "视频"},
	ObjectTypeCodeApk:    {"id": 30072, "fileExtension": ".apk", "documentType": "Android Package Archive", "mimeType": "application/vnd.android.package-archive", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeExe:    {"id": 30073, "fileExtension": ".exe", "documentType": "Microsoft Application", "mimeType": "application/x-msdownload", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
	ObjectTypeCodeDmg:    {"id": 30074, "fileExtension": ".dmg", "documentType": "Apple Disk Image", "mimeType": "application/x-apple-diskimage", "linkedStorageType": "Other", "linkedStorageTypeName": "其他"},
}

// 文件扩展名与链存对象类型代码映射关系
var FileExtensionObjectTypeCodeMapping = map[string]int{
	FileExtensionEmpty:  ObjectTypeCodeEmpty,
	FileExtensionDir:    ObjectTypeCodeDir,
	FileExtensionAac:    ObjectTypeCodeAac,
	FileExtensionAbw:    ObjectTypeCodeAbw,
	FileExtensionArc:    ObjectTypeCodeArc,
	FileExtensionAvi:    ObjectTypeCodeAvi,
	FileExtensionAzw:    ObjectTypeCodeAzw,
	FileExtensionBin:    ObjectTypeCodeBin,
	FileExtensionBmp:    ObjectTypeCodeBmp,
	FileExtensionBz:     ObjectTypeCodeBz,
	FileExtensionBz2:    ObjectTypeCodeBz2,
	FileExtensionCsh:    ObjectTypeCodeCsh,
	FileExtensionCss:    ObjectTypeCodeCss,
	FileExtensionCsv:    ObjectTypeCodeCsv,
	FileExtensionDoc:    ObjectTypeCodeDoc,
	FileExtensionDocx:   ObjectTypeCodeDocx,
	FileExtensionEot:    ObjectTypeCodeEot,
	FileExtensionEpub:   ObjectTypeCodeEpub,
	FileExtensionGif:    ObjectTypeCodeGif,
	FileExtensionHtm:    ObjectTypeCodeHtm,
	FileExtensionHtml:   ObjectTypeCodeHtml,
	FileExtensionIco:    ObjectTypeCodeIco,
	FileExtensionIcs:    ObjectTypeCodeIcs,
	FileExtensionJar:    ObjectTypeCodeJar,
	FileExtensionJpeg:   ObjectTypeCodeJpeg,
	FileExtensionJpg:    ObjectTypeCodeJpg,
	FileExtensionJs:     ObjectTypeCodeJs,
	FileExtensionJson:   ObjectTypeCodeJson,
	FileExtensionJsonld: ObjectTypeCodeJsonld,
	FileExtensionMid:    ObjectTypeCodeMid,
	FileExtensionMidi:   ObjectTypeCodeMidi,
	FileExtensionMjs:    ObjectTypeCodeMjs,
	FileExtensionMp3:    ObjectTypeCodeMp3,
	FileExtensionMpeg:   ObjectTypeCodeMpeg,
	FileExtensionMpkg:   ObjectTypeCodeMpkg,
	FileExtensionOdp:    ObjectTypeCodeOdp,
	FileExtensionOds:    ObjectTypeCodeOds,
	FileExtensionOdt:    ObjectTypeCodeOdt,
	FileExtensionOga:    ObjectTypeCodeOga,
	FileExtensionOgv:    ObjectTypeCodeOgv,
	FileExtensionOgx:    ObjectTypeCodeOgx,
	FileExtensionOtf:    ObjectTypeCodeOtf,
	FileExtensionPng:    ObjectTypeCodePng,
	FileExtensionPdf:    ObjectTypeCodePdf,
	FileExtensionPpt:    ObjectTypeCodePpt,
	FileExtensionPptx:   ObjectTypeCodePptx,
	FileExtensionRar:    ObjectTypeCodeRar,
	FileExtensionRtf:    ObjectTypeCodeRtf,
	FileExtensionSh:     ObjectTypeCodeSh,
	FileExtensionSvg:    ObjectTypeCodeSvg,
	FileExtensionSwf:    ObjectTypeCodeSwf,
	FileExtensionTar:    ObjectTypeCodeTar,
	FileExtensionTif:    ObjectTypeCodeTif,
	FileExtensionTiff:   ObjectTypeCodeTiff,
	FileExtensionTtf:    ObjectTypeCodeTtf,
	FileExtensionTxt:    ObjectTypeCodeTxt,
	FileExtensionVsd:    ObjectTypeCodeVsd,
	FileExtensionWav:    ObjectTypeCodeWav,
	FileExtensionWeba:   ObjectTypeCodeWeba,
	FileExtensionWebm:   ObjectTypeCodeWebm,
	FileExtensionWebp:   ObjectTypeCodeWebp,
	FileExtensionWoff:   ObjectTypeCodeWoff,
	FileExtensionWoff2:  ObjectTypeCodeWoff2,
	FileExtensionXhtml:  ObjectTypeCodeXhtml,
	FileExtensionXls:    ObjectTypeCodeXls,
	FileExtensionXlsx:   ObjectTypeCodeXlsx,
	FileExtensionXml:    ObjectTypeCodeXml,
	FileExtensionXul:    ObjectTypeCodeXul,
	FileExtensionZip:    ObjectTypeCodeZip,
	FileExtension3gp:    ObjectTypeCode3gp,
	FileExtension3g2:    ObjectTypeCode3g2,
	FileExtension7z:     ObjectTypeCode7z,
	FileExtensionMp4:    ObjectTypeCodeMp4,
	FileExtensionApk:    ObjectTypeCodeApk,
	FileExtensionExe:    ObjectTypeCodeExe,
	FileExtensionDmg:    ObjectTypeCodeDmg,
}

// 文件扩展名与链存对象类型映射关系
var FileExtensionLinkedStorageTypeMapping = map[string]string{
	FileExtensionEmpty:  LinkedStorageTypeOther,
	FileExtensionDir:    LinkedStorageTypeDirectory,
	FileExtensionAac:    LinkedStorageTypeAudio,
	FileExtensionAbw:    LinkedStorageTypeOther,
	FileExtensionArc:    LinkedStorageTypeOther,
	FileExtensionAvi:    LinkedStorageTypeVideo,
	FileExtensionAzw:    LinkedStorageTypeOther,
	FileExtensionBin:    LinkedStorageTypeOther,
	FileExtensionBmp:    LinkedStorageTypeImage,
	FileExtensionBz:     LinkedStorageTypeOther,
	FileExtensionBz2:    LinkedStorageTypeOther,
	FileExtensionCsh:    LinkedStorageTypeOther,
	FileExtensionCss:    LinkedStorageTypeDocument,
	FileExtensionCsv:    LinkedStorageTypeDocument,
	FileExtensionDoc:    LinkedStorageTypeDocument,
	FileExtensionDocx:   LinkedStorageTypeDocument,
	FileExtensionEot:    LinkedStorageTypeOther,
	FileExtensionEpub:   LinkedStorageTypeOther,
	FileExtensionGif:    LinkedStorageTypeImage,
	FileExtensionHtm:    LinkedStorageTypeDocument,
	FileExtensionHtml:   LinkedStorageTypeDocument,
	FileExtensionIco:    LinkedStorageTypeImage,
	FileExtensionIcs:    LinkedStorageTypeDocument,
	FileExtensionJar:    LinkedStorageTypeOther,
	FileExtensionJpeg:   LinkedStorageTypeImage,
	FileExtensionJpg:    LinkedStorageTypeImage,
	FileExtensionJs:     LinkedStorageTypeDocument,
	FileExtensionJson:   LinkedStorageTypeJson,
	FileExtensionJsonld: LinkedStorageTypeOther,
	FileExtensionMid:    LinkedStorageTypeAudio,
	FileExtensionMidi:   LinkedStorageTypeAudio,
	FileExtensionMjs:    LinkedStorageTypeDocument,
	FileExtensionMp3:    LinkedStorageTypeAudio,
	FileExtensionMpeg:   LinkedStorageTypeVideo,
	FileExtensionMpkg:   LinkedStorageTypeOther,
	FileExtensionOdp:    LinkedStorageTypeDocument,
	FileExtensionOds:    LinkedStorageTypeDocument,
	FileExtensionOdt:    LinkedStorageTypeDocument,
	FileExtensionOga:    LinkedStorageTypeAudio,
	FileExtensionOgv:    LinkedStorageTypeVideo,
	FileExtensionOgx:    LinkedStorageTypeOther,
	FileExtensionOtf:    LinkedStorageTypeOther,
	FileExtensionPng:    LinkedStorageTypeImage,
	FileExtensionPdf:    LinkedStorageTypeDocument,
	FileExtensionPpt:    LinkedStorageTypeDocument,
	FileExtensionPptx:   LinkedStorageTypeDocument,
	FileExtensionRar:    LinkedStorageTypeOther,
	FileExtensionRtf:    LinkedStorageTypeOther,
	FileExtensionSh:     LinkedStorageTypeOther,
	FileExtensionSvg:    LinkedStorageTypeImage,
	FileExtensionSwf:    LinkedStorageTypeOther,
	FileExtensionTar:    LinkedStorageTypeOther,
	FileExtensionTif:    LinkedStorageTypeImage,
	FileExtensionTiff:   LinkedStorageTypeImage,
	FileExtensionTtf:    LinkedStorageTypeOther,
	FileExtensionTxt:    LinkedStorageTypeDocument,
	FileExtensionVsd:    LinkedStorageTypeDocument,
	FileExtensionWav:    LinkedStorageTypeAudio,
	FileExtensionWeba:   LinkedStorageTypeAudio,
	FileExtensionWebm:   LinkedStorageTypeVideo,
	FileExtensionWebp:   LinkedStorageTypeImage,
	FileExtensionWoff:   LinkedStorageTypeOther,
	FileExtensionWoff2:  LinkedStorageTypeOther,
	FileExtensionXhtml:  LinkedStorageTypeOther,
	FileExtensionXls:    LinkedStorageTypeDocument,
	FileExtensionXlsx:   LinkedStorageTypeDocument,
	FileExtensionXml:    LinkedStorageTypeOther,
	FileExtensionXul:    LinkedStorageTypeOther,
	FileExtensionZip:    LinkedStorageTypeOther,
	FileExtension3gp:    LinkedStorageTypeVideo,
	FileExtension3g2:    LinkedStorageTypeAudio,
	FileExtension7z:     LinkedStorageTypeOther,
	FileExtensionMp4:    LinkedStorageTypeVideo,
	FileExtensionApk:    LinkedStorageTypeOther,
	FileExtensionExe:    LinkedStorageTypeOther,
	FileExtensionDmg:    LinkedStorageTypeOther,
}

// 上传默认值
const (
	DefaultChunkSize                  int64 = 45613056                   // 256*1024*174 默认块大小
	DefaultChunkTTL                         = 86400 * time.Second        // 缓存块cache 1天
	DefaultDagBuildLockTimeout              = 10 * time.Second           // dag build 锁
	DefaultSaveObjectLockTimeout            = 5 * time.Second            // 保存 object 锁
	DefaultDagCodec                         = "dag-pb"                   // 默认dag 数据格式
	DefaultCashHashPrefix                   = "cache-hash-"              // 默认cache-hash-{{sha256}} 映射 cid 缓存前缀
	CidCounterPrefix                        = "cid-amount-"              // 上传缓存cid 记录key 前缀
	UsedResourceFileBufferIncrTimeout       = 10 * time.Second           // 用户上传文件缓冲记数器过期时间
	UsedResourceFileBufferIncrPrefix        = "usedresourcefile-buffer-" // 用户上传文件缓冲记数器前缀
	UsedSpaceBufferIncrTimeout              = 10 * time.Second           // 用户已经使用空间大小缓冲记数器过期时间
	UsedSpaceBufferIncrPrefix               = "usedspace-buffer-"        // 用户已经使用空间大小缓冲记数器前缀

	IpfsCidCachesKey        = "cid-cache-exprie-set"      // 临时上传 ipfs cid 缓存到redis set key
	IpfsCleanPinsExprieLock = "cleanIpfsPinsByExprieLock" // 定时清理ipfs 缓存锁 key
)

// 上传状态
const (
	Uploading             = iota + 1 // 1上传中
	UploadDagBuilding                // 2 dag构建中
	UploadReplicaCreating            // 3 创建副本中
	UploadDone                       // 4 上传完成
)

// 存储网络编码
const (
	StorageNetworkCodeIpfs = iota + 10001 // IPFS
)

// 桶策略编码
const (
	BucketPrincipleCodePrivate = iota + 10000 // 私有
	BucketPrincipleCodePublic                 // 公开
)

// 存储网络编码与名称映射关系
var StorageNetworkCodeMapping = map[int]string{
	StorageNetworkCodeIpfs: "IPFS",
}

// 桶策略编码与名称映射关系
var BucketPrincipleCodeMapping = map[int]string{
	BucketPrincipleCodePrivate: "私有",
	BucketPrincipleCodePublic:  "公开",
}

//var FileExtensionLinkedStorageObjectMapping = map[string]models.ObjectDocumentMapping{
//	FileExtensionEmpty:  {Id: 10000, FileExtension: "", DocumentType: "", MimeType: "", LinkedStorageType: "", LinkedStorageTypeName: "其他"},
//	FileExtensionDir:    {Id: 20000, FileExtension: "", DocumentType: "", MimeType: "", LinkedStorageType: "Directory", LinkedStorageTypeName: "目录"},
//	FileExtensionAac:    {Id: 30001, FileExtension: ".aac", DocumentType: "AAC audio", MimeType: "audio/aac", LinkedStorageType: "Audio", LinkedStorageTypeName: "音频"},
//	FileExtensionAbw:    {Id: 30002, FileExtension: ".abw", DocumentType: "AbiWord document", MimeType: "application/x-abiword", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionArc:    {Id: 30003, FileExtension: ".arc", DocumentType: "Archive document (multiple files embedded)", MimeType: "application/x-freearc", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionAvi:    {Id: 30004, FileExtension: ".avi", DocumentType: "AVI: Audio Video Interleave", MimeType: "video/x-msvideo", LinkedStorageType: "Video", LinkedStorageTypeName: "视频"},
//	FileExtensionAzw:    {Id: 30005, FileExtension: ".azw", DocumentType: "Amazon Kindle eBook format", MimeType: "application/vnd.amazon.ebook", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionBin:    {Id: 30006, FileExtension: ".bin", DocumentType: "Any kind of binary data", MimeType: "application/octet-stream", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionBmp:    {Id: 30007, FileExtension: ".bmp", DocumentType: "Windows OS/2 Bitmap Graphics", MimeType: "image/bmp", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionBz:     {Id: 30008, FileExtension: ".bz", DocumentType: "BZip archive", MimeType: "application/x-bzip", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionBz2:    {Id: 30009, FileExtension: ".bz2", DocumentType: "BZip2 archive", MimeType: "application/x-bzip2", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionCsh:    {Id: 30010, FileExtension: ".csh", DocumentType: "C-Shell script", MimeType: "application/x-csh", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionCss:    {Id: 30011, FileExtension: ".css", DocumentType: "Cascading Style Sheets (CSS)", MimeType: "text/css", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionCsv:    {Id: 30012, FileExtension: ".csv", DocumentType: "Comma-separated values (CSV)", MimeType: "text/csv", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionDoc:    {Id: 30013, FileExtension: ".doc", DocumentType: "Microsoft Word", MimeType: "application/msword", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionDocx:   {Id: 30014, FileExtension: ".docx", DocumentType: "Microsoft Word (OpenXML)", MimeType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionEot:    {Id: 30015, FileExtension: ".eot", DocumentType: "MS Embedded OpenType fonts", MimeType: "application/vnd.ms-fontobject", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionEpub:   {Id: 30016, FileExtension: ".epub", DocumentType: "Electronic publication (EPUB)", MimeType: "application/epub+zip", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionGif:    {Id: 30017, FileExtension: ".gif", DocumentType: "Graphics Interchange Format (GIF)", MimeType: "image/gif", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionHtm:    {Id: 30018, FileExtension: ".htm", DocumentType: "HyperText Markup Language (HTML)", MimeType: "text/html", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionHtml:   {Id: 30019, FileExtension: ".html", DocumentType: "HyperText Markup Language (HTML)", MimeType: "text/html", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionIco:    {Id: 30020, FileExtension: ".ico", DocumentType: "Icon format", MimeType: "image/vnd.microsoft.icon", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionIcs:    {Id: 30021, FileExtension: ".ics", DocumentType: "iCalendar format", MimeType: "text/calendar", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionJar:    {Id: 30022, FileExtension: ".jar", DocumentType: "Java Archive (JAR)", MimeType: "application/java-archive", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionJpeg:   {Id: 30023, FileExtension: ".jpeg", DocumentType: "JPEG images", MimeType: "image/jpeg", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionJpg:    {Id: 30024, FileExtension: ".jpg", DocumentType: "JPEG images", MimeType: "image/jpeg", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionJs:     {Id: 30025, FileExtension: ".js", DocumentType: "JavaScript", MimeType: "text/javascript", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionJson:   {Id: 30026, FileExtension: ".json", DocumentType: "JSON format", MimeType: "application/json", LinkedStorageType: "JSON", LinkedStorageTypeName: ""},
//	FileExtensionJsonld: {Id: 30027, FileExtension: ".jsonld", DocumentType: "JSON-LD format", MimeType: "application/ld+json", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionMid:    {Id: 30028, FileExtension: ".mid", DocumentType: "Musical Instrument Digital Interface (MIDI)", MimeType: "audio/midi audio/x-midi", LinkedStorageType: "Audio", LinkedStorageTypeName: "音频"},
//	FileExtensionMidi:   {Id: 30029, FileExtension: ".midi", DocumentType: "Musical Instrument Digital Interface (MIDI)", MimeType: "audio/midi audio/x-midi", LinkedStorageType: "Audio", LinkedStorageTypeName: "音频"},
//	FileExtensionMjs:    {Id: 30030, FileExtension: ".mjs", DocumentType: "JavaScript module", MimeType: "text/javascript", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionMp3:    {Id: 30031, FileExtension: ".mp3", DocumentType: "MP3 audio", MimeType: "audio/mpeg", LinkedStorageType: "Audio", LinkedStorageTypeName: "音频"},
//	FileExtensionMpeg:   {Id: 30032, FileExtension: ".mpeg", DocumentType: "MPEG Video", MimeType: "video/mpeg", LinkedStorageType: "Video", LinkedStorageTypeName: "视频"},
//	FileExtensionMpkg:   {Id: 30033, FileExtension: ".mpkg", DocumentType: "Apple Installer Package", MimeType: "application/vnd.apple.installer+xml", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionOdp:    {Id: 30034, FileExtension: ".odp", DocumentType: "OpenDocument presentation document", MimeType: "application/vnd.oasis.opendocument.presentation", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionOds:    {Id: 30035, FileExtension: ".ods", DocumentType: "OpenDocument spreadsheet document", MimeType: "application/vnd.oasis.opendocument.spreadsheet", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionOdt:    {Id: 30036, FileExtension: ".odt", DocumentType: "OpenDocument text document", MimeType: "application/vnd.oasis.opendocument.text", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionOga:    {Id: 30037, FileExtension: ".oga", DocumentType: "OGG audio", MimeType: "audio/ogg", LinkedStorageType: "Audio", LinkedStorageTypeName: "音频"},
//	FileExtensionOgv:    {Id: 30038, FileExtension: ".ogv", DocumentType: "OGG video", MimeType: "video/ogg", LinkedStorageType: "Video", LinkedStorageTypeName: "视频"},
//	FileExtensionOgx:    {Id: 30039, FileExtension: ".ogx", DocumentType: "OGG", MimeType: "application/ogg", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionOtf:    {Id: 30040, FileExtension: ".otf", DocumentType: "OpenType font", MimeType: "font/otf", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionPng:    {Id: 30041, FileExtension: ".png", DocumentType: "Portable Network Graphics", MimeType: "image/png", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionPdf:    {Id: 30042, FileExtension: ".pdf", DocumentType: "Adobe Portable Document Format (PDF)", MimeType: "application/pdf", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionPpt:    {Id: 30043, FileExtension: ".ppt", DocumentType: "Microsoft PowerPoint", MimeType: "application/vnd.ms-powerpoint", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionPptx:   {Id: 30044, FileExtension: ".pptx", DocumentType: "Microsoft PowerPoint (OpenXML)", MimeType: "application/vnd.openxmlformats-officedocument.presentationml.presentation", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionRar:    {Id: 30045, FileExtension: ".rar", DocumentType: "RAR archive", MimeType: "application/x-rar-compressed", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionRtf:    {Id: 30046, FileExtension: ".rtf", DocumentType: "Rich Text Format (RTF)", MimeType: "application/rtf", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionSh:     {Id: 30047, FileExtension: ".sh", DocumentType: "Bourne shell script", MimeType: "application/x-sh", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionSvg:    {Id: 30048, FileExtension: ".svg", DocumentType: "Scalable Vector Graphics (SVG)", MimeType: "image/svg+xml", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionSwf:    {Id: 30049, FileExtension: ".swf", DocumentType: "Small web format (SWF) or Adobe Flash document", MimeType: "application/x-shockwave-flash", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionTar:    {Id: 30050, FileExtension: ".tar", DocumentType: "Tape Archive (TAR)", MimeType: "application/x-tar", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionTif:    {Id: 30051, FileExtension: ".tif", DocumentType: "Tagged Image File Format (TIFF)", MimeType: "image/tiff", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionTiff:   {Id: 30052, FileExtension: ".tiff", DocumentType: "Tagged Image File Format (TIFF)", MimeType: "image/tiff", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionTtf:    {Id: 30053, FileExtension: ".ttf", DocumentType: "TrueType Font", MimeType: "font/ttf", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionTxt:    {Id: 30054, FileExtension: ".txt", DocumentType: "Text, (generally ASCII or ISO 8859-n)", MimeType: "text/plain", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionVsd:    {Id: 30055, FileExtension: ".vsd", DocumentType: "Microsoft Visio", MimeType: "application/vnd.visio", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionWav:    {Id: 30056, FileExtension: ".wav", DocumentType: "Waveform Audio Format", MimeType: "audio/wav", LinkedStorageType: "Audio", LinkedStorageTypeName: "音频"},
//	FileExtensionWeba:   {Id: 30057, FileExtension: ".weba", DocumentType: "WEBM audio", MimeType: "audio/webm", LinkedStorageType: "Audio", LinkedStorageTypeName: "音频"},
//	FileExtensionWebm:   {Id: 30058, FileExtension: ".webm", DocumentType: "WEBM video", MimeType: "video/webm", LinkedStorageType: "Video", LinkedStorageTypeName: "视频"},
//	FileExtensionWebp:   {Id: 30059, FileExtension: ".webp", DocumentType: "WEBP image", MimeType: "image/webp", LinkedStorageType: "Image", LinkedStorageTypeName: "图像"},
//	FileExtensionWoff:   {Id: 30060, FileExtension: ".woff", DocumentType: "Web Open Font Format (WOFF)", MimeType: "font/woff", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionWoff2:  {Id: 30061, FileExtension: ".woff2", DocumentType: "Web Open Font Format (WOFF)", MimeType: "font/woff2", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionXhtml:  {Id: 30062, FileExtension: ".xhtml", DocumentType: "XHTML", MimeType: "application/xhtml+xml", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionXls:    {Id: 30063, FileExtension: ".xls", DocumentType: "Microsoft Excel", MimeType: "application/vnd.ms-excel", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionXlsx:   {Id: 30064, FileExtension: ".xlsx", DocumentType: "Microsoft Excel (OpenXML)", MimeType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", LinkedStorageType: "Document", LinkedStorageTypeName: "文档"},
//	FileExtensionXml:    {Id: 30065, FileExtension: ".xml", DocumentType: "XML", MimeType: "application/xml", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionXul:    {Id: 30066, FileExtension: ".xul", DocumentType: "XUL", MimeType: "application/vnd.mozilla.xul+xml", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtensionZip:    {Id: 30067, FileExtension: ".zip", DocumentType: "ZIP", MimeType: "archive application/zip", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//	FileExtension3gp:    {Id: 30068, FileExtension: ".3gp", DocumentType: "3GPP audio/video container", MimeType: "video/3gpp audio/3gpp", LinkedStorageType: "Video", LinkedStorageTypeName: "视频"},
//	FileExtension3g2:    {Id: 30069, FileExtension: ".3g2", DocumentType: "3GPP2 audio/video container", MimeType: "video/3gpp2 audio/3gpp2", LinkedStorageType: "Audio", LinkedStorageTypeName: "音频"},
//	FileExtension7z:     {Id: 30070, FileExtension: ".7z", DocumentType: "7-zip", MimeType: "archive application/x-7z-compressed", LinkedStorageType: "Other", LinkedStorageTypeName: "其他"},
//}

// Api Key生成
const (
	// Minimum length for MinIO access key.
	AccessKeyMinLen = 3

	// Maximum length for MinIO access key.
	// There is no max length enforcement for access keys
	AccessKeyMaxLen = 20

	// Minimum length for MinIO secret key for both server
	SecretKeyMinLen = 8

	// Maximum secret key length for MinIO, this
	// is used when autogenerating new credentials.
	// There is no max length enforcement for secret keys
	SecretKeyMaxLen = 40

	// Alphanumeric table used for generating access keys.
	AlphaNumericTable = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Total length of the alphanumeric table.
	AlphaNumericTableLen = byte(len(AlphaNumericTable))
)

// ApiKey权限类型编码
const (
	ApiKeyPermissionTypeCodeCustom = iota + 10000 // 自定义设置
	ApiKeyPermissionTypeCodeAdmin                 // 管理员设置
)

// ApiKey权限类型编码与名称映射关系
var ApiKeyPermissionTypeCodeMapping = map[int]string{
	ApiKeyPermissionTypeCodeCustom: "自定义设置",
	ApiKeyPermissionTypeCodeAdmin:  "管理员设置",
}

// ApiKey权限类型
const (
	ApiKeyPermissionCodeBucketList      = iota + 10001 // 桶列表查询
	ApiKeyPermissionCodeBucketCreate                   // 桶创建
	ApiKeyPermissionCodeBucketEmpty                    // 桶清空查询
	ApiKeyPermissionCodeBucketDelete                   // 桶删除
	ApiKeyPermissionCodeObjectList                     // 对象列表查询
	ApiKeyPermissionCodeObjectUpload                   // 上传对象
	ApiKeyPermissionCodeObjectMark                     // 添加星标
	ApiKeyPermissionCodeObjectRename                   // 对象重命名
	ApiKeyPermissionCodeObjectReference                // CID 添加对象
	ApiKeyPermissionCodeObjectDelete                   // 删除对象
)

// ApiKey权限类型与名称映射关系
var ApiKeyPermissionCodeMapping = map[int]string{
	ApiKeyPermissionCodeBucketList:      "桶列表查询",
	ApiKeyPermissionCodeBucketCreate:    "桶创建",
	ApiKeyPermissionCodeBucketEmpty:     "桶清空查询",
	ApiKeyPermissionCodeBucketDelete:    "桶删除",
	ApiKeyPermissionCodeObjectList:      "对象列表查询",
	ApiKeyPermissionCodeObjectUpload:    "上传对象",
	ApiKeyPermissionCodeObjectMark:      "添加星标",
	ApiKeyPermissionCodeObjectRename:    "对象重命名",
	ApiKeyPermissionCodeObjectReference: "CID 添加对象",
	ApiKeyPermissionCodeObjectDelete:    "删除对象",
}

// PinningServiceAPI权限类型
const (
	PinningServiceApiPermissionCodeAddPinObject     = iota + 10001 // addPinObject
	PinningServiceApiPermissionCodeRemovePinObject                 // removePinObject
	PinningServiceApiPermissionCodeGetPinObject                    // getPinObject
	PinningServiceApiPermissionCodeReplacePinObject                // replacePinObject
)

// PinningServiceAPI权限类型与名称映射关系
var PinningServiceApiPermissionCodeMapping = map[int]string{
	PinningServiceApiPermissionCodeAddPinObject:     "addPinObject",
	PinningServiceApiPermissionCodeRemovePinObject:  "removePinObject",
	PinningServiceApiPermissionCodeGetPinObject:     "getPinObject",
	PinningServiceApiPermissionCodeReplacePinObject: "replacePinObject",
}
