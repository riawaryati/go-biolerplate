package general

import (
	"github.com/furee/backend/constants/general"
)

type SectionService struct {
	App           AppAccount       `json:",omitempty"`
	Route         RouteAccount     `json:",omitempty"`
	Database      DatabaseAccount  `json:",omitempty"`
	Redis         RedisAccount     `json:",omitempty"`
	Authorization AuthAccount      `json:",omitempty"`
	Toggle        ToggleAccount    `json:",omitempty"`
	KeyData       KeyAccount       `json:",omitempty"`
	Minio         MinioSecret      `json:",omitempty"`
	NSQProducer   NSQProducer      `json:",omitempty"`
	NSQConsumer   NSQConsumer      `json:",omitempty"`
	PartnerSecret PartnerSecret    `json:",omitempty"`
	Logistic      LogisticSecret   `json:",omitempty"`
	Whitelist     WhitelistAccount `json:",omitempty"`
}

type AppAccount struct {
	Name         string `json:",omitempty"`
	Environtment string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	SecretKey    string `json:",omitempty"`
	Endpoint     string `json:",omitempty"`
}

type RouteAccount struct {
	Methods []string    `json:",omitempty"`
	Headers []string    `json:",omitempty"`
	Origins RouteOrigin `json:",omitempty"`
}

type RouteOrigin struct {
	InternalTools string `json:",omitempty"`
}

type DatabaseAccount struct {
	Read  DBDetailAccount `json:",omitempty"`
	Write DBDetailAccount `json:",omitempty"`
}
type DBDetailAccount struct {
	Username     string `json:",omitempty"`
	Password     string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	DBName       string `json:",omitempty"`
	MaxIdleConns int    `json:",omitempty"`
	MaxOpenConns int    `json:",omitempty"`
	MaxLifeTime  int    `json:",omitempty"`
	Timeout      string `json:",omitempty"`
	SSLMode      string `json:",omitempty"`
}

type RedisAccount struct {
	Username     string `json:",omitempty"`
	Password     string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         int    `json:",omitempty"`
	MinIdleConns int    `json:",omitempty"`
	Timeout      string `json:",omitempty"`
}

type AuthAccount struct {
	JWT    JWTCredential    `json:",omitempty"`
	Public PublicCredential `json:",omitempty"`
}

type JWTCredential struct {
	IsActive              bool   `json:",omitempty"`
	AccessTokenSecretKey  string `json:",omitempty"`
	AccessTokenDuration   int    `json:",omitempty"`
	RefreshTokenSecretKey string `json:",omitempty"`
	RefreshTokenDuration  int    `json:",omitempty"`
}

type PublicCredential struct {
	SecretKey string `json:",omitempty"`
}
type ToggleAccount struct {
	IsUseJWT bool `json:",omitempty"`
}

type S3CredentialKey struct {
	AccessKeyID     string            `json:",omitempty"`
	SecretAccessKey string            `json:",omitempty"`
	Region          string            `json:",omitempty"`
	BucketName      string            `json:",omitempty"`
	TempFolder      string            `json:",omitempty"`
	FileSource      map[string]string `json:",omitempty"`
	URLDuration     int               `json:",omitempty"`
	Endpoint        string            `josn:",omitempty"`
}

type MinioSecret struct {
	BucketName string            `json:", omitempty"`
	Endpoint   string            `json:", omitempty"`
	Key        string            `json:", omitempty"`
	Secret     string            `json:", omitempty"`
	Region     string            `json:", omitempty"`
	TempFolder string            `json:",omitempty"`
	BaseURL    string            `json:",omitempty"`
	FileSource map[string]string `json:",omitempty"`
}

type NSQProducer struct {
	NSQD string `json:",omitempty"`
}

type NSQConsumer struct {
	NSQLookupD string `json:",omitempty"`
	Detail     map[string]NSQConsumerDetail
}

type NSQConsumerDetail struct {
	MaxInFlight int
	MaxAttempts uint16
	Concurrency int
	Requeue     int
	Backoff     bool
}

type PartnerSecret struct {
	ActivePayment string                `json:",omitempty"`
	Xendit        XenditSecret          `json:",omitempty"`
	MessageBird   MessageBirdCredential `json:",omitempty"`
}

type XenditSecret struct {
	XenditCallbackToken   string `json:",omitempty"`
	XenditAPIMoneyInWrite string `json:",omitempty"`
	BaseURL               string `json:",omitempty"`
	CreateQRISURL         string `json:",omitempty"`
	Authorization         string `json:",omitempty"`
	QRISCallbackURL       string `json:",omitempty"`
}

type InternalCookieData struct {
	Auth string `json:"auth"`
	ID   int64  `json:"id"`
	Name int64  `json:"name"`
	Role int64  `json:"role"`
}

type LogisticSecret struct {
	JETExpress    JETExpressCredential `json:",omitempty"`
	Borzo         BorzoCredential      `json:",omitempty"`
	JNECredential JNECredential        `json:",omitempty"`
}

type JETExpressCredential struct {
	APIKey   string `json:",omitempty"`
	ClientID string `json:",omitempty"`
	BaseURL  string `json:",omitempty"`
	RateURL  string `json:",omitempty"`
}

type BorzoCredential struct {
	Token            string `json:",omitempty"`
	BaseURL          string `json:",omitempty"`
	RateURL          string `json:",omitempty"`
	RequestPickupURL string `json:",omitempty"`
}

type JNECredential struct {
	Username string `json:",omitempty"`
	APIKey   string `json:",omitempty"`
	BaseURL  string `json:",omitempty"`
	RateURL  string `json:",omitempty"`
}

type MessageBirdCredential struct {
	WhatsappOTPURL string `json:",omitempty"`
	Authorization  string `json:",omitempty"`
	NameSpace      string `json:",omitempty"`
	TemplateName   string `json:",omitempty"`
	FromID         string `json:",omitempty"`
	LanguageCode   string `json:",omitempty"`
	Type           string `json:",omitempty"`
}

type KeyAccount struct {
	User string `json:",omitempty"`
}

type WhitelistAccount struct {
	Region WhitelistRegionData `json:",omitempty"`
	User   WhitelistUserData   `json:",omitempty"`
}

type WhitelistRegionData struct {
	City map[int64]int64 `json:",omitempty"`
}

type WhitelistUserData struct {
	Phone map[string]string `json:",omitempty"`
	OTP   string            `json:",omitempty"`
}

func IsAllowImageType(imageType string) bool {
	switch imageType {
	case general.ImageTypeJPEG, general.ImageTypePNG:
		return true
	default:
		return false
	}
}

func IsAllowMediaType(imageType string) (string, bool) {
	switch imageType {
	case general.ImageTypeJPEG, general.ImageTypePNG:
		return "image", true
	case general.VideoType3gpp, general.VideoTypeFLV, general.VideoTypeMP2T, general.VideoTypeMP4, general.VideoTypeMPEGURL, general.VideoTypeMSVideo, general.VideoTypeQuicktime, general.VideoTypeWMV, general.VideoTypeOctetStream:
		return "video", true
	default:
		return "", false
	}
}

func IsAllowFileType(fileType string) bool {
	switch fileType {
	case general.ImageTypeJPEG, general.ImageTypePNG:
		return true
	default:
		return false
	}
}
