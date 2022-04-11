package base_test_user

import (
	"Rep_Pess_Sharever2/ses_service_user/main/src/entity/se_user"
)

var SeUserLogin1 = se_user.SeUserLogin{
	Username: Username1,
	Password: Nickname1,
	Roles:    []se_user.AuthorityEnum{se_user.ROLE_USER},
}

var SesUser1 = se_user.SesUser{
	Id:               IdUser1,
	Nickname:         Nickname1,
	Name:             Name1,
	FiscalNumber:     FiscalNumber1,
	Cellphone:        Cellphone1,
	StartDateUser:    StartDate1,
	SeUserStatusEnum: se_user.ACTIVE,
	AuthorityEnum:    se_user.ROLE_USER,
	Origin:           "Origin1",
	Score:            0,
	Birthday:         StartDate1,
	IndicatedBy:      "IndicatedBy1",
	SeUserLogin:      SeUserLogin1,
	SeCountryEnum:    se_user.BRAZIL,
}
