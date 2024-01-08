package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	htmltemplate "html/template"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ly1999-hub/go-go/internal/constant"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"github.com/ly1999-hub/go-go/pkg/admin/dao"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Admin struct{}

func (a Admin) GetMe(ctx context.Context, doc model.Admin) *model.AdminResponse {
	res := model.AdminResponse{
		Name:      doc.Name,
		Email:     doc.Email,
		Phone:     doc.Phone,
		Birthday:  doc.Birthday,
		Avatar:    doc.Avatar,
		Address:   doc.Address,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
		Active:    doc.Active,
		Login:     doc.Login,
		Role:      doc.Role,
	}
	return &res
}

func (a Admin) GetDetail(ctx context.Context, admin model.Admin, payload model.AdminDetail) (*model.AdminResponse, error) {
	var (
		dao = dao.Admin{}
	)
	doc := dao.FindByID(ctx, payload.ID)
	if doc.ID.IsZero() {
		return nil, errors.New(response.CommonNotFound)
	}
	res := model.AdminResponse{
		Name:      doc.Name,
		Email:     doc.Email,
		Phone:     doc.Phone,
		Birthday:  doc.Birthday,
		Avatar:    doc.Avatar,
		Address:   doc.Address,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
		Active:    doc.Active,
		Login:     doc.Login,
		Role:      doc.Role,
	}
	return &res, nil
}

func (a Admin) Create(ctx context.Context, payload model.AdminCreate) (*model.ResponseCreate, error) {
	var (
		daoAdmin = dao.Admin{}
		daoRole  = dao.Role{}
	)
	roleAdmin := daoRole.FindOne(ctx, bson.M{"code": "ROOT"})
	if roleAdmin.ID.IsZero() {
		if err := daoRole.InsertOne(ctx, model.Role{
			ID:          primitive.NewObjectID(),
			RoleName:    "ROOT",
			Description: "ROOT",
			Permissions: []string{"ADMIN", "USER"},
			Code:        "ROOT",
			CreatedAt:   time.Now().GoString(),
			UpdatedAt:   time.Now().GoString(),
		}); err != nil {
			log.Error("Error-Insert ROOT_ROLE", log.LogData{"err": err.Error()})
			return nil, err
		}
	}
	// 	if err := daoRole.InsertOne(ctx, model.Role{
	// 		ID:          primitive.NewObjectID(),
	// 		RoleName:    "ADMIN",
	// 		Description: "ADMIN",
	// 		Permissions: []string{"ADMIN"},
	// 		Code:        "ADMIN",
	// 		CreatedAt:   time.Now().GoString(),
	// 		UpdatedAt:   time.Now().GoString(),
	// 	}); err != nil {
	// 		log.Error("Error-Insert ROOT_ROLE", log.LogData{"err": err.Error()})
	// 		return nil, err
	// 	}
	// }
	root := daoAdmin.FindOne(ctx, bson.M{"root": true})
	if root.ID.IsZero() {
		if err := daoAdmin.InsertOne(ctx, model.Admin{
			ID:             primitive.NewObjectID(),
			Name:           "root",
			Email:          "root@gmail.com",
			Phone:          "0967606851",
			Role:           "ROOT",
			HashedPassword: util.HashedPassword("root"),
			Active:         true,
			Root:           true,
			Login:          false,
			Birthday:       "25/01/1999",
			Avatar:         "",
			Address:        "da nang",
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		}); err != nil {
			return nil, err
		}
	}
	if !a.checkExistByEmail(ctx, payload.Email).ID.IsZero() {
		log.Error("serviceRole-Update", log.LogData{
			"error": errors.New(response.CommonExistEmail),
		})
		return nil, errors.New(response.CommonExistEmail)
	}

	role := daoRole.FindOne(ctx, bson.M{"code": payload.Role})
	if role.ID.IsZero() {
		log.Error("Role not found ", log.LogData{"data": payload.Role})
		return nil, errors.New(response.CommonNotFound)
	}

	doc := model.Admin{
		ID:             primitive.NewObjectID(),
		Name:           payload.Name,
		Email:          payload.Email,
		Phone:          payload.Phone,
		HashedPassword: util.HashedPassword(payload.Password),
		Birthday:       payload.Birthday,
		Avatar:         payload.Avatar,
		Address:        payload.Address,
		Role:           role.ID.Hex(),
		Root:           false,
		Active:         true,
		CreatedAt:      time.Now().GoString(),
		UpdatedAt:      time.Now().GoString(),
	}
	err := daoAdmin.InsertOne(ctx, doc)
	if err != nil {
		log.Error("serviceAdmin-Create", log.LogData{"data": err.Error()})
		return nil, errors.New(response.CommonExistEmail)
	}
	return &model.ResponseCreate{ID: doc.ID.Hex()}, nil
}

func (a Admin) checkExistByEmail(ctx context.Context, email string) model.Admin {
	var (
		dao = dao.Admin{}
	)
	doc := dao.FindOne(ctx, bson.M{"email": email})
	return doc
}

func (a Admin) LoginByEmail(ctx context.Context, payload model.LoginByEmail) (*model.ResLoginAdmin, error) {
	admin := a.checkExistByEmail(ctx, payload.Email)

	if admin.ID.IsZero() {
		log.Error("ServiceAdmin-LoginByEmail ", log.LogData{"error": errors.New(response.CommonNotFound)})
		return nil, errors.New(response.CommonNotFound)
	}
	if !util.CheckPassword(payload.Password, admin.HashedPassword) {
		return nil, errors.New(response.CommonErrorPassword)
	}
	if !admin.Active {
		log.Error("ServiceAdmin-LoginByEmail ", log.LogData{"error": errors.New(response.CommonNotFound)})
		return nil, errors.New(response.CommonNotActive)
	}
	return generateToken(admin)
}

func generateToken(admin model.Admin) (*model.ResLoginAdmin, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": admin.ID,
		"exp": time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	fmt.Print(tokenString)
	if err != nil {
		log.Error("ServiceAdmin-generateToken ", log.LogData{"error": err.Error()})
		return nil, errors.New(response.CommonNotFound)
	}
	return &model.ResLoginAdmin{ID: admin.ID.Hex(), Token: tokenString}, nil
}

func (a Admin) GetAll(ctx context.Context, payload model.All) (res model.ResponseList) {
	var (
		wg  sync.WaitGroup
		dao = dao.Admin{}
	)

	if payload.Limit == 0 {
		res.Limit = 20
	}
	opts := options.Find()
	opts.SetLimit(payload.Limit).SetSkip(payload.Page * res.Limit)
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := dao.CountByCondition(ctx, bson.M{})
		res.Total = count
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		list := make([]model.Admin, 0)
		list = dao.FindByCondtion(ctx, bson.M{}, opts)
		res.List = list
	}()
	wg.Wait()
	return res
}

func (a Admin) ForGetPassword(ctx context.Context, payload model.AdminForGetPassword) (result *model.ResponseUpdate, err error) {
	var (
		dao = dao.Admin{}
	)
	doc := a.checkExistByEmail(ctx, payload.Email)
	fmt.Print(doc)
	if doc.ID.IsZero() {
		return nil, errors.New(response.CommonNotFound)
	}
	if !doc.Active {
		return nil, errors.New(response.CommonNotActive)
	}

	newPassword := a.generateRandomCode(6)
	fmt.Print(newPassword)
	if len(newPassword) != 6 {
		return nil, errors.New(response.CommonErrorService)
	}

	subject := "MẬT KHẨU MỚI CỦA BẠN"
	hashedPassword := util.HashedPassword(newPassword)
	if err = dao.UpdateByID(ctx, doc.ID, bson.M{"$set": bson.M{"hashed_password": hashedPassword}}); err != nil {
		err = errors.New(response.CommonErrorService)
		return
	}

	if SendEmailSendGrid(doc.Email, subject, newPassword) != nil {
		err = errors.New(response.CommonErrorSendgrid)
		return
	}

	return &model.ResponseUpdate{ID: doc.ID.Hex()}, nil
}

func SendEmailSendGrid(toEmail string, subject string, content string) error {
	var contentHTML bytes.Buffer

	html, err := htmltemplate.New("forgotPassword.html").Parse(constant.TemplateHTMLForgotPassword)
	if err != nil {

		return errors.New(response.CommonErrorService)
	}
	err = html.Execute(&contentHTML, struct {
		URI          string
		NEW_PASSWORD string
	}{URI: "http://localhost:3000/login", NEW_PASSWORD: content})
	if err != nil {

		return errors.New(response.CommonErrorSendgrid)
	}
	from := mail.NewEmail("From", "nhly123123456789@gmail.com")
	to := mail.NewEmail("To", toEmail)
	//htmlContent := mail.NewContent("text/html", contentHTML.String())
	message := mail.NewSingleEmail(from, subject, to, content, contentHTML.String())
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return nil
}

func (s Admin) generateRandomCode(length int) string {
	charset := os.Getenv("CHARSET")
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error generating random code:", err)
		return ""
	}

	for i := 0; i < length; i++ {
		b[i] = charset[b[i]%byte(len(charset))]
	}

	return string(b)
}

func (s Admin) UploadAvatar(ctx context.Context, doc model.Admin, file model.FileUploadInfo) (result *model.ResponseUpdate, err error) {
	var (
		dao = dao.Admin{}
	)
	if err = dao.UpdateByID(ctx, doc.ID, bson.M{
		"$set": bson.M{
			"avatar": file.Path,
		},
	}); err != nil {
		return nil, errors.New(response.CommonErrorService)
	}
	return &model.ResponseUpdate{ID: doc.ID.Hex()}, nil
}
