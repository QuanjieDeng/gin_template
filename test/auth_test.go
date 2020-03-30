package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"gvp/models"
	"gvp/pkg/logging"
	"gvp/pkg/setting"
	"gvp/routers"
	"net/http/httptest"
	"testing"
)



func  init(){
	setting.Setup()
	logging.Setup()
	models.Setup()

	router = routers.InitRouter()
}





func TestGetAuth( t *testing.T){
	Convey("测试获取token",t, func() {
		Convey("测试错误参数", func() {
			w:= httptest.NewRecorder()
			req:= httptest.NewRequest("GET","/auth?username1=admin&password=admin",nil)
			router.ServeHTTP(w,req)

			So(200,ShouldNotEqual,w.Code)
		})
		Convey("测试正确参数", func() {
			w:= httptest.NewRecorder()
			req:= httptest.NewRequest("GET","/auth?username=admin&password=admin",nil)
			router.ServeHTTP(w,req)
			So(200,ShouldEqual,w.Code)

		})

	})
}

