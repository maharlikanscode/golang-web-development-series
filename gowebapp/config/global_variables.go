package config

// Settings is where the common settings.go constant variables.
type Settings struct {
	SiteFullName, SiteSlogan, SiteBaseURL, SiteTopMenuLogo, SiteProperDomainName,
	SiteShortName, SiteEmail, SitePhoneNumbers, SiteCompanyAddress string
	SiteYear int
}

// SiteSettings defines all constant variables from the settings.go
var SiteSettings = Settings{
	SiteFullName:         SiteFullName,
	SiteSlogan:           SiteSlogan,
	SiteBaseURL:          SiteBaseURL,
	SiteTopMenuLogo:      SiteTopMenuLogo,
	SiteProperDomainName: SiteProperDomainName,
	SiteShortName:        SiteShortName,
	SiteEmail:            SiteEmail,
	SitePhoneNumbers:     SitePhoneNumbers,
	SiteCompanyAddress:   SiteCompanyAddress,
	SiteYear:             SiteYear,
}
