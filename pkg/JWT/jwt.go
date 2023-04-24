package JWT

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/spf13/viper"
)

var mySecret = []byte("miku") //定义加密
// MyClaims 自定义声明结构体并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (aToken string, err error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * viper.GetDuration("auth.jwt_expire"))},
			Issuer:    "holo", // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象	// 使用指定的secret签名并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	//refresh token不需要存任何自定义数据
	/*rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Second * 30)},
		Issuer:    "bluebell",
	}).SignedString(mySecret)*/
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var token *jwt.Token
	//way1将解析结果保存到claims变量中,若token字符串合法但过期claims也会有数据，err会提示token过期
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { //校验token
		return mc, nil
	}
	/*//way2从ParseWithClaims返回的Token结构体中取出Claims结构体
	token, err := jwt.ParseWithClaims(tokenString, *MyClaims, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	*/
	return nil, errors.New("invalid token")
}

type EmailClaims struct {
	UserName      int64  `json:"user_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.RegisteredClaims
}

func GenerateEmailToken(username int64, Operation uint, email string, password string) (string, error) {
	claims := EmailClaims{
		UserName:      username,
		Email:         email,
		Password:      password,
		OperationType: Operation,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * viper.GetDuration("auth.jwt_expire"))},
			Issuer:    "holo",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(mySecret)
	return token, err
}

// ParseEmailToken 验证邮箱验证token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
