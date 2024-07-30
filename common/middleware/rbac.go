package middleware

import (
	"log"
	"net/http"
	"zero-chat/common/result"
	"zero-chat/common/xerr"

	stdcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/auth/casbin"
)

var e *stdcasbin.Enforcer

func init() {
	a := xormadapter.NewAdapter("mysql", "root:PXDN93VRKUm8TeE7@tcp(127.0.0.1:33061)/")
	// load the casbin model and policy from files, database is also supported.
	m, err := model.NewModelFromFile("./rbac_model.conf")
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}
	e, err = stdcasbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	// define your router, and use the Casbin auth middleware.
	// the access that is denied by auth will return HTTP 403 error.
	// set username as the user unique identity field.
	authorizer := casbin.NewAuthorizer(e, casbin.WithUidField("username"))
	conf := rest.RestConf{}
	server := rest.MustNewServer(conf)
	server.Use(rest.ToMiddleware(authorizer))
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sub := "user1" // 用户角色
		obj := r.URL.Path
		act := r.Method

		if ok, _ := e.Enforce(sub, obj, act); !ok {
			result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.NO_ACCESS_TO_RESOURCE))
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
