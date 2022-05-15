package se_photo

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


//Get video by idUser and idVideo video folder
func getImageSvgFromServerByName(context *gin.Context) {
	image := context.Param("image")
	pathVideo := fmt.Sprintf("repository/image/svg/%s.svg", image)
	context.File(pathVideo)
	context.Header("Content-Type", "image/svg+xml")
}


//Get video by idUser and idVideo video folder
func getImagePngFromServerByName(context *gin.Context) {
	image := context.Param("image")
	pathVideo := fmt.Sprintf("repository/image/png/%s.png", image)
	context.File(pathVideo)
	context.Header("Content-Type", "image/png")
}

//func updatePhotoInDatabase(context *gin.Context, r *http.Request) {
//	//r.ParseMultipartForm()
//	//idUser := context.Param("idUser")
//	file, fileHeader, err := context.Request.FormFile("photo")
//
//	if err != nil {
//		fmt.Println("Error Retrieving the File")
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
//	fmt.Printf("File Size: %+v\n", fileHeader.Size)
//	fmt.Printf("MIME Header: %+v\n", fileHeader.Header)
//
//	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
//	errorHelper.FmtError(err)
//	fileBytes, err := ioutil.ReadAll(file)
//	errorHelper.FmtError(err)
//	tempFile.Write(fileBytes)
//	context.JSON(http.StatusOK, "Successfully Uploaded File")
//	defer file.Close()
//}

//
//func updatePhotoInDatabase(context *gin.Context) (mongo.UpdateResult, error) {
//	file, _, _ := context.Request.FormFile("photo")
//	idUser := context.Param("idUser")
//	idUserObject, _ := primitive.ObjectIDFromHex(idUser)
//	savePhoto(context, "photo")
//	toByte, _ := ioutil.ReadAll(file)
//	encodeToString := base64.StdEncoding.EncodeToString(toByte)
//	return updatePhoto(idUserObject, encodeToString)
//}
//

////Get image from local sei_server
//func recoveryImageFromServer(context *gin.Context){
//	file := getLocalImageByName("profile")
//	toBytes, _ := ioutil.ReadAll(file)
//	context.Writer.Write(toBytes)
//}

////Save image in local path
//func savePhotoWithQuality(photo []byte, imageName string, quality int){
//	reader := bytes.NewReader(photo)
//	img, _, _ := image.Decode(reader)
//	out := createFileJpeg(imageName)
//	defer out.Close()
//	var opts jpeg.Options
//	opts.Quality = quality
//	err := jpeg.Encode(out, img, &opts)
//	errorHelper.CheckPanic(err)
//}

////Create a file in the sei_server
//func createFileJpeg(fileName string) *os.File {
//	file, err := os.Create(fmt.Sprintf("./%s.jpg", fileName))
//	errorHelper.CheckPanic(err)
//	return file
//}

//func getPhotoByName(fileName string) image.Image {
//	file, _ := os.Open(fmt.Sprintf("./%s.jpg", fileName))
//	img, _ := jpeg.Decode(file)
//	return img
//}
//
////Get a local jpeg with name
//func getLocalImageByName(fileName string) *os.File {
//	file1, _ := os.Open(fmt.Sprintf("%s.jpg", fileName))
//	return file1
//}
//
////CreateImageBase64 create a image to show in tag image html in format base64
//func CreateImageBase64(imageBytes []byte) string{
//	imageBase64 := base64.StdEncoding.EncodeToString(imageBytes)
//	return ImageSchemeBase64 + imageBase64
//}



//func updatePhotoInDatabase(context *gin.Context) mongo.UpdateResult {
//	file, _, _ := context.Request.FormFile("photo")
//	idUser := context.Param("idUser")
//	idUserObject, _ := primitive.ObjectIDFromHex(idUser)
//	savePhoto(context, "photo")
//	toByte, _ := ioutil.ReadAll(file)
//	encodeToString := base64.StdEncoding.EncodeToString(toByte)
//	return updatePhoto(idUserObject, encodeToString)
//}


////Get image from local sei_server
//func recoveryImageFromDatabase(context *gin.Context){
//	idUser := context.Param("idUser")
//	idUserObject, _ := primitive.ObjectIDFromHex(idUser)
//	sei-profile := findProfile(idUserObject)
//	context.Writer.Write(sei-profile.PhotoURL)
//}

////Get image from local sei_server
//func recoveryImageFromServer(context *gin.Context){
//	file := getLocalImageByName("profile")
//	toBytes, _ := ioutil.ReadAll(file)
//	context.Writer.Write(toBytes)
//}
//
//
////Save image in local path
//func savePhoto(context *gin.Context, fileName string){
//	fileHeader, err := context.FormFile(fileName)
//	errorHelper.CheckError(err)
//	destine := path.Join("./", fileHeader.Filename)
//	context.SaveUploadedFile(fileHeader, destine)
//}
//
////Save image in local path
//func savePhotoWithQuality(photo []byte, imageName string, quality int){
//	reader := bytes.NewReader(photo)
//	img, _, _ := image.Decode(reader)
//	out := createFileJpeg(imageName)
//	defer out.Close()
//	var opts jpeg.Options
//	opts.Quality = quality
//	err := jpeg.Encode(out, img, &opts)
//	errorHelper.CheckError(err)
//}
//
////Create a file in the sei_server
//func createFileJpeg(fileName string) *os.File {
//	file, err := os.Create(fmt.Sprintf("./%s.jpg", fileName))
//	errorHelper.CheckError(err)
//	return file
//}
//
//
//func getPhotoByName(fileName string) image.Image {
//	file, _ := os.Open(fmt.Sprintf("./%s.jpg", fileName))
//	img, _ := jpeg.Decode(file)
//	return img
//}
//
////Get a local jpeg with name
//func getLocalImageByName(fileName string) *os.File {
//	file1, _ := os.Open(fmt.Sprintf("%s.jpg", fileName))
//	return file1
//}
//
////CreateImageBase64 create a image to show in tag image html in format base64
//func CreateImageBase64(imageBytes []byte) string{
//	imageBase64 := base64.StdEncoding.EncodeToString(imageBytes)
//	return ImageSchemeBase64 + imageBase64
//}