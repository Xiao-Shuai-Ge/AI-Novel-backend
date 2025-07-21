package entity

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
)

const (
	CAPTCHA_LENGTH = 6
)

type Captcha struct {
	Email string
	Code  string
}

func NewCaptcha(email string) Captcha {
	return Captcha{Email: email}
}

func (c *Captcha) CheckEmail() (err error) {
	// 用正则表达式检查邮箱格式
	if c.Email == "" {
		return errors.New("邮箱地址不能为空")
	}

	// 邮箱格式正则表达式
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 编译正则表达式
	re := regexp.MustCompile(emailRegex)

	// 进行匹配
	if !re.MatchString(c.Email) {
		return errors.New("邮箱格式无效，请输入正确的邮箱地址")
	}

	return nil
}

func (c *Captcha) SetRandCode() {
	code := ""
	for i := 0; i < CAPTCHA_LENGTH; i++ {
		// 正确方法：生成单个数字字符
		code += strconv.Itoa(rand.Intn(10))
	}
	c.Code = code
}

func (c *Captcha) GetCode() string {
	return c.Code
}

func (c *Captcha) GetEmail() string {
	return c.Email
}
