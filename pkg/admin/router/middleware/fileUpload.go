package middleware

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
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
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Println(err.Error())
			return response.R400(c, nil, err.Error())
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

func ChangeAvatar(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		// Khởi tạo client của Firebase Storage
		client, err := storage.NewClient(ctx, option.WithCredentialsFile("path/to/your/firebase/credential.json"))
		if err != nil {
			fmt.Println("Lỗi khi khởi tạo client:", err)
			return response.R400(c, nil, "")
		}
		defer client.Close()

		// Đường dẫn tới ảnh cần xóa trong Firebase Storage

		fileName := c.Get("file_name_delete").(string)
		// Xóa ảnh từ Firebase Storage
		err = client.Bucket("your-firebase-storage-bucket").Object(PathUpload + fileName).Delete(ctx)
		if err != nil {
			fmt.Println("Lỗi khi xóa ảnh từ Firebase Storage:", err)
			return response.R400(c, nil, "")
		}
		return next(c)
	}
}
