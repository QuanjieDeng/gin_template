package test



import (
	"encoding/json"
	"github.com/goinggo/mapstructure"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"gvp/models"
	"gvp/pkg/logging"
	"gvp/pkg/setting"
	"gvp/routers"
	"gvp/pkg/app"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var  testBind =  models.GameService{
	Game:"测试绑定",
	ServiceId:1,
	GameId:1,
}

var  bind_db_id  int


func  init(){
	setting.Setup()
	logging.Setup()
	models.Setup()

	router = routers.InitRouter()
}



func TestBind( t *testing.T){
	Convey("测试游戏-服务商绑定关系",t, func() {

		Convey("获取token值", func() {
			var ret bool
			ret,TokenValue =  getToken()
			So(ret,ShouldBeTrue)
			So(TokenValue,ShouldNotBeNil)
		})


		Convey("测试添加绑定-错误字段", func() {
			w:= httptest.NewRecorder()

			var target = "/api/v1/bind"
			target =  target+"?token="+TokenValue

			var tmp_bind   struct {
				Game string
			}
			tmp_bind.Game = "测试绑定"
			tmp,_ :=  json.Marshal(tmp_bind)
			req:= httptest.NewRequest("POST",target,strings.NewReader(string(tmp[:])))
			req.Header.Add("Content-type","application/json")
			router.ServeHTTP(w,req)
			So(w.Code,ShouldNotEqual,200)
		})

		Convey("测试添加绑定-正确字段-错误的参数的范围", func() {
			w:= httptest.NewRecorder()

			var target = "/api/v1/bind"
			target =  target+"?token="+TokenValue

			stubs :=   gostub.Stub(&(testBind.GameId),0)
			defer  stubs.Reset()

			tmp,_ :=  json.Marshal(testBind)

			req:= httptest.NewRequest("POST",target,strings.NewReader(string(tmp[:])))
			req.Header.Add("Content-type","application/json")
			router.ServeHTTP(w,req)
			So(w.Code,ShouldNotEqual,200)
		})

		Convey("测试成功添加绑定关系", func() {
			w:= httptest.NewRecorder()

			var target = "/api/v1/bind"
			target =  target+"?token="+TokenValue

			tmp,_ :=  json.Marshal(testBind)

			req:= httptest.NewRequest("POST",target,strings.NewReader(string(tmp[:])))
			req.Header.Add("Content-type","application/json")
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
		})

		Convey("测试拉取绑定关系列表", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/bind"
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("GET",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
			if w.Code  == 200{
				var  res   app.Response
				err :=  json.Unmarshal(w.Body.Bytes(),&res)

				if err !=  nil{
					t.Error("拉取绑定关系列表进行json解析报错")
				}

				total  := res.Data.(map[string]interface{})["total"].(float64)
				if   total >= 1{
					lists := res.Data.(map[string]interface{})["lists"].([]interface {})

					for _,v:=range lists  {
						var  tmpbind  models.GameService
						mapstructure.Decode(v, &tmpbind)
						bind_db_id =  tmpbind.ID
						logging.Info(bind_db_id)
					}
				}
			}
		})

		Convey("测试拉取单个绑定关系", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/bind/"+strconv.Itoa(bind_db_id)
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("GET",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
		})

		Convey("删除不存在的绑定关系", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/bind/1000000"
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("DELETE",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldNotEqual,200)
		})

		Convey("测试删除绑定关系", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/bind/"+strconv.Itoa(bind_db_id)
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("DELETE",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
		})

	})
}