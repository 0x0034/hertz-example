// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	user_example "github.com/0x0034/hertz-example/biz/router/user/example"
	user_exmple "github.com/0x0034/hertz-example/biz/router/user/exmple"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	user_example.Register(r)

	user_exmple.Register(r)

}
