package se_user

type AuthorityEnum string

const (
	ROLE_USER   AuthorityEnum = "ROLE_USER"
	ROLE_ADMIN  AuthorityEnum = "ROLE_ADMIN"
	ROLE_MASTER AuthorityEnum = "ROLE_MASTER"
)

//func (authorityEnum *SeAuthorityEnum) GetValue() string {
//	if ROLE_USER == authorityEnum {
//		return "USER"
//	}
//}