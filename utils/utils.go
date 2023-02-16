package utils

import (
	"DouyinSimpleProject/config"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("SGVsbG9Xb3JsZA")

type CustomClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

func GenToken(uid uint) (string, error) {
	claims := &CustomClaims{
		UserID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "HelloWorld",
			Subject:   "douyin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func ValidToken(tokenString string) (*CustomClaims, error) {
	if tokenString == "" {
		return nil, errors.New("empty token")
	}

	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return nil, errors.New("expired token")
	}
	return claims, nil
}

func String2uint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(num), nil
}

func GetFileURL(filename string) string {
	return fmt.Sprintf(
		"http://%s:%s/static/%s",
		config.SERVER_HOST, config.SERVER_PORT, filename,
	)
}

// ExtractImageFromVideo extract the first frame from video,
// and return the cover path like `./public/2_3.jpg`
func ExtractImageFromVideo(videoName, suffix string) string {
	videoPath := filepath.Join(config.STATIC_ROOT_PATH, videoName+suffix)
	coverPath := filepath.Join(config.STATIC_ROOT_PATH, videoName+".jpg")
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vframes", "1", coverPath)
	if _, err := cmd.Output(); err != nil {
		return config.DEFAULT_COVER_FILENAME
	} else {
		if _, err := os.Stat(coverPath); errors.Is(err, os.ErrNotExist) {
			// ffmpeg execute successfully, but no such generated cover image
			// we use default cover image instead.
			return config.DEFAULT_COVER_FILENAME
		}
		return videoName + ".jpg"
	}
}

// UnmarshalJSONTimeStamp converts JS timestamp to `time.Time`
//
// Reference: https://gist.github.com/alextanhongpin/3b6b2ee47665ac9c1c32c805b86380a6
func UnmarshalJSTimeStamp(rawTime string) time.Time {
	timeStamp, _ := strconv.ParseInt(rawTime, 10, 64)
	t := time.Unix(timeStamp/1000, (timeStamp%1000)*1000*1000)
	return t
}

func MarshalJSTimeStamp(t time.Time) int64 {
	timeStamp := t.UnixNano() / 1e6
	return timeStamp
}

func GetCommentReponseDate(t time.Time) string {
	return t.Format("01-02")
}

// HashAndSalt --- Encrypt password
func HashAndSalt(pwdStr string) (pwdHash string, err error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	pwdHash = string(hash)
	return pwdHash, nil
}

// ComparePasswords Verify password
func ComparePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	return err == nil
}
