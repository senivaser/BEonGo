package apiserver

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/senivaser/BEonGo/internal/app/model"
	"github.com/senivaser/BEonGo/internal/app/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *APIServer) CreateAccessToken(guid string) (interface{}, error) {

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["guid"] = guid
	atClaims["created"] = time.Now().Unix()
	atClaims["expired"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *APIServer) CreateRefreshToken(guid string) (interface{}, error) {

	os.Setenv("REFRESH_SECRET", "gtredfgtredf")
	atClaims := jwt.MapClaims{}
	atClaims["guid"] = guid
	atClaims["created"] = time.Now().Unix()
	atClaims["expired"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	_, err = s.store.User.UpdateBy("guid", guid, "refreshToken", token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *APIServer) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *APIServer) ValidateRefreshToken(refreshToken string) (interface{}, error) {

	var err error
	var user model.User

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	guid := token.Claims.(jwt.MapClaims)["guid"]

	if guid == nil {
		err = fmt.Errorf("Wrong GUID")
	} else {
		user, err = s.store.User.GetBy("guid", utils.ToString(guid))
		if user.Guid == guid {
			fmt.Println(1)
			return guid, nil
		}
		fmt.Println(2)
		return nil, err
	}

	return nil, err
}
