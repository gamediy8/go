package services

import (
	"fmt"
	"gin-web/utils"
)

type Login struct {
	Account  string
	Password string
}

func (this *Login) LoginLogic() *utils.Data {
	var a = utils.Data{
		Result: fmt.Sprintf("%s %s", this.Account, this.Password),
	}
	return &a
}
