package main
import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)
//vaildator 请求校验

type RegisterReq struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username        string   `validate:"gt=0"`
	// 同上
	PasswordNew     string   `validate:"gt=0"`
	// eqfield 跨字段相等校验
	PasswordRepeat  string   `validate:"eqfield=PasswordNew"`
	// 合法 email 格式校验
	Email           string   `validate:"email"`
}

func main(){
	var req = new(RegisterReq)
	req.Username="Xargin"
	req.PasswordNew="ohno"
	req.PasswordRepeat ="ohn"
	req.Email  = "alex@abc.com"

	var v *validator.Validate
	v = validator.New()
	err:=v.Struct(req)
	fmt.Println("err:",err) // Key: 'RegisterReq.PasswordRepeat' Error:Field validation for 'PasswordRepeat' failed on the 'eqfield' tag

}

