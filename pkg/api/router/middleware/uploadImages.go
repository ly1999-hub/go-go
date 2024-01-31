package middleware

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"google.golang.org/api/option"
)

func Upload(c echo.Context) *model.FileUploadInfo {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	BUCKET_NAME := os.Getenv("BUCKET_NAME")
	var fileString = "D:/Golang/GoAPI/serviceAccountKey.json"
	opt := option.WithCredentialsFile(fileString)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	client, err := app.Storage(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	bucketHandler, err := client.Bucket(BUCKET_NAME)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	image, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	src, err := image.Open()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer src.Close()
	id := uuid.New()
	id_string := id.String()
	nameFile := image.Filename
	name := strings.Split(nameFile, ".")
	newNameFile := name[0] + id_string + "." + name[1]
	objHandler := bucketHandler.Object(newNameFile)
	write := objHandler.NewWriter(context.Background())

	write.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	defer write.Close()
	if _, err := io.Copy(write, src); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	string_url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", BUCKET_NAME, image.Filename, id_string)
	fmt.Print(string_url)
	fmt.Print(id_string)

	fileUploadInfor := model.FileUploadInfo{
		Filename: image.Filename,
		Path:     string_url,
		Ext:      filepath.Ext(image.Filename),
	}
	return &fileUploadInfor
}
