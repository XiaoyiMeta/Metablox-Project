package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"go-metablox/pkg/base"
)

// BizContext request context structure
type BizContext struct {
	Session *ghttp.Session // Current Session management object
	Data    g.Map          // Custom KV variables, business modules are set according to needs, not fixed
}

type CurrentUser struct {
	Id                   int64       `json:"id"                   description:"ID"`
	Uid                  string      `json:"uid"                  description:"UID"`
	DisplayName          string      `json:"displayName"          description:"User Nick Name"`
	Username             string      `json:"username"             description:"Login account"`
	Password             string      `json:"-"                    description:"Password"`
	PhoneNumber          string      `json:"phoneNumber"          description:"Phone number"`
	WalletAccount        string      `json:"walletAccount"        description:"Wallet login account"`
	RealName             string      `json:"realName"             description:"Real name"`
	PhotoUrl             string      `json:"photoUrl"             description:"Avatar"`
	BirthDate            base.Date   `json:"birthDate"            description:"Birthday"`
	Gender               int         `json:"gender"               description:"Gender(0-Unknown,1-Male,2-Female)"`
	Email                string      `json:"email"                description:"Email address"`
	EmailVerified        bool        `json:"emailVerified"        description:"Whether the email has been verified"`
	IdentificationType   string      `json:"identificationType"   description:"Identification type"`
	IdentificationNumber string      `json:"identificationNumber" description:"Identification number"`
	Disabled             bool        `json:"disabled"             description:"User status(1-Normal,0-disabled)"`
	LastLoginAt          base.Time   `json:"lastLoginAt"          description:"Last login time on the account"`
	LastPasswordChangeAt base.Time   `json:"lastPasswordChangAt"  description:"Last password change time"`
	TenantId             string      `json:"tenantId"             description:"Multi-tenant identification"`
	ClientId             string      `json:"-"                    description:"Device ID"`
	CountryCode          string      `json:"countryCode"          description:"Country Code"`
	LanguageCode         string      `json:"languageCode"         description:"Language code"`
	DeletedAt            interface{} `json:"-"                    description:"Deletion time"`
	CreatedAt            interface{} `json:"createdAt"            description:"Creation time"`
	UpdatedAt            interface{} `json:"updatedAt"            description:"Last update time"`
}
