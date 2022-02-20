package general

import "time"

const (
	SourceFromDB    string = "db"
	SourceFromCache string = "cache"
)

//List database name
const (
	DatabaseRead  = "read"
	DatabaseWrite = "write"
)

const (
	ImageMaxSize  int64 = 1024000
	FileMaxSize   int64 = 1024000
	MultiPartSize int64 = 1024000

	MimeTypeImage string = "image"
	MimeTypeVideo string = "video"

	ImageTypeJPEG string = "image/jpeg"
	ImageTypePNG  string = "image/png"

	VideoTypeFLV         string = "video/x-flv"
	VideoTypeMP4         string = "video/mp4"
	VideoTypeMPEGURL     string = "application/x-mpegURL"
	VideoTypeMP2T        string = "video/MP2T"
	VideoType3gpp        string = "video/3gpp"
	VideoTypeQuicktime   string = "video/quicktime"
	VideoTypeMSVideo     string = "video/x-msvideo"
	VideoTypeWMV         string = "video/x-ms-wmv"
	VideoTypeOctetStream string = "application/octet-stream"
)

const (
	FullTimeFormat        string = "2006-01-02 15:04:05"
	DisplayDateTimeFormat string = "02 Jan 2006 15:04:05"
	DateFormat            string = "2006-01-02"
)

const (
	NumJan = iota + 1
	NumFeb
	NumMar
	NumApr
	NumMay
	NumJune
	NumJuly
	NumAug
	NumSep
	NumOct
	NumNov
	NumDec
)

const (
	RomanJan  string = "I"
	RomanFeb  string = "II"
	RomanMar  string = "III"
	RomanApr  string = "IV"
	RomanMay  string = "V"
	RomanJune string = "VI"
	RomanJuly string = "VII"
	RomanAug  string = "VIII"
	RomanSep  string = "IX"
	RomanOct  string = "X"
	RomanNov  string = "XI"
	RomanDec  string = "XII"
)

const (
	AuthCookies string = "auth"
)

const (
	ENVProduction string = "production"
)

const (
	UpdatedBySystem int = 0
)

const (
	Time1Min = 1 * time.Minute
	Time5Min = 5 * time.Minute
	Time1Day = 24 * time.Hour
)
