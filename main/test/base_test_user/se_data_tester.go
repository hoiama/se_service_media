package base_test_user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var IdGroup1, _ = primitive.ObjectIDFromHex("01fa940e1424f4731b93f3c3")
var ExpiresAt1, _ = time.Parse(time.RFC3339, "2022-08-01T23:00:00.815-03:00")

//User1
var IdUser1, _ = primitive.ObjectIDFromHex("10d15a9bc52f71dcfbbedc10")
var Nickname1 = "Nickname1"
var Name1 = "Name1"
var Cellphone1 = "Cellphone1"
var Username1 = "user1@gmail.com"
var Password1 = "user1@gmail.com"
var StartDate1 = primitive.NewDateTimeFromTime(time.Now())
var Email1 = "user1@gmail.com"
var FiscalNumber1 = "1d15a9bc52f71"

//User2
var IdUser2, _ = primitive.ObjectIDFromHex("20d15a9bc52f71dcfbbedc10")
var Email2 = "user2@gmail.com"
var Cellphone2 = "Cellphone2"
var FiscalNumber2 = "2d15a9bc52f71"

var IdNotification1 = "10g15a9bc52f71dcfbbedc10"
var SubjectNotificationTemplate = "sei-home subjectNotificationTemplate"
//var BodyNotificationTemplate = fmt.Sprintf(
//	`sei-home bodyNotificationTemplate %s %s`, notification.NAME, notification.EMAIL)
