package middleware

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"google.golang.org/api/option"
)

const (
	PathUpload = "./internals/public/avatar-image/"
)

func UploadSingleFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		BUCKET_NAME := os.Getenv("BUCKET_NAME")

		var fileString = "D:/Golang/GoAPI/serviceAccountKey.json"
		opt := option.WithCredentialsFile(fileString)
		fmt.Print(opt)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Println(err.Error())
			return response.R400(c, nil, "")
		}
		client, err := app.Storage(context.TODO())
		if err != nil {
			fmt.Println(err.Error())
			return response.R400(c, nil, "")
		}
		bucketHandler, err := client.Bucket(BUCKET_NAME)
		if err != nil {
			fmt.Println(err.Error())
			return response.R400(c, nil, "")
		}
		image, err := c.FormFile("image")
		if err != nil {
			fmt.Println(err.Error())
			return response.R400(c, nil, "")
		}
		src, err := image.Open()
		if err != nil {
			fmt.Println(err.Error())
			return response.R400(c, nil, "")
		}
		defer src.Close()

		objHandler := bucketHandler.Object(image.Filename)
		write := objHandler.NewWriter(context.Background())
		id := uuid.New()
		id_string := id.String()
		write.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
		defer write.Close()
		if _, err := io.Copy(write, src); err != nil {
			fmt.Println(err.Error())
			return response.R400(c, nil, "")
		}
		string_url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", BUCKET_NAME, image.Filename, id_string)
		fmt.Print(string_url)
		fmt.Print(id_string)

		c.Set("file_avatar", model.FileUploadInfo{
			Filename: image.Filename,
			Path:     string_url,
			Ext:      filepath.Ext(image.Filename),
		})

		return next(c)
	}
}
