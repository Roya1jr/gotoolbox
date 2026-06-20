package gtbhttp

// Headers
const (
	// General headers
	HeaderCacheControl     = "Cache-Control"
	HeaderConnection       = "Connection"
	HeaderContentLength    = "Content-Length"
	HeaderContentType      = "Content-Type"
	HeaderDate             = "Date"
	HeaderPragma           = "Pragma"
	HeaderTrailer          = "Trailer"
	HeaderTransferEncoding = "Transfer-Encoding"
	HeaderUpgrade          = "Upgrade"
	HeaderVia              = "Via"
	HeaderWarning          = "Warning"

	// Request headers
	HeaderAccept             = "Accept"
	HeaderAcceptCharset      = "Accept-Charset"
	HeaderAcceptEncoding     = "Accept-Encoding"
	HeaderAcceptLanguage     = "Accept-Language"
	HeaderAuthorization      = "Authorization"
	HeaderCookie             = "Cookie"
	HeaderExpect             = "Expect"
	HeaderFrom               = "From"
	HeaderHost               = "Host"
	HeaderIfMatch            = "If-Match"
	HeaderIfModifiedSince    = "If-Modified-Since"
	HeaderIfNoneMatch        = "If-None-Match"
	HeaderIfRange            = "If-Range"
	HeaderIfUnmodifiedSince  = "If-Unmodified-Since"
	HeaderMaxForwards        = "Max-Forwards"
	HeaderProxyAuthorization = "Proxy-Authorization"
	HeaderRange              = "Range"
	HeaderReferer            = "Referer"
	HeaderUserAgent          = "User-Agent"

	// Response headers
	HeaderAcceptRanges      = "Accept-Ranges"
	HeaderAge               = "Age"
	HeaderETag              = "ETag"
	HeaderLocation          = "Location"
	HeaderProxyAuthenticate = "Proxy-Authenticate"
	HeaderRetryAfter        = "Retry-After"
	HeaderServer            = "Server"
	HeaderSetCookie         = "Set-Cookie"
	HeaderVary              = "Vary"
	HeaderWWWAuthenticate   = "WWW-Authenticate"

	// Entity headers
	HeaderAllow           = "Allow"
	HeaderContentEncoding = "Content-Encoding"
	HeaderContentLanguage = "Content-Language"
	HeaderContentLocation = "Content-Location"
	HeaderContentRange    = "Content-Range"
	HeaderExpires         = "Expires"
	HeaderLastModified    = "Last-Modified"

	// CORS headers
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"

	// WebSocket headers
	HeaderSecWebSocketKey      = "Sec-WebSocket-Key"
	HeaderSecWebSocketVersion  = "Sec-WebSocket-Version"
	HeaderSecWebSocketProtocol = "Sec-WebSocket-Protocol"
	HeaderSecWebSocketAccept   = "Sec-WebSocket-Accept"
)

// Mime types

const (
	// Application MIME types
	MIMEApplicationSOAPXML            = "application/soap+xml"
	MIMEApplicationJSON               = "application/json"
	MIMEApplicationXML                = "application/xml"
	MIMEApplicationXWWWFormURLEncoded = "application/x-www-form-urlencoded"
	MIMEApplicationOctetStream        = "application/octet-stream"
	MIMEApplicationPDF                = "application/pdf"
	MIMEApplicationZip                = "application/zip"

	// Text MIME types
	MIMETextPlain = "text/plain"
	MIMETextHTML  = "text/html"
	MIMETextCSS   = "text/css"
	MIMETextCSV   = "text/csv"

	// Image MIME types
	MIMEImageJPEG = "image/jpeg"
	MIMEImagePNG  = "image/png"
	MIMEImageGIF  = "image/gif"

	// Audio MIME types
	MIMEAudioMPEG = "audio/mpeg"
	MIMEAudioOGG  = "audio/ogg"

	// Video MIME types
	MIMEVideoMP4  = "video/mp4"
	MIMEVideoMPEG = "video/mpeg"
	MIMEVideoWebM = "video/webm"
)
