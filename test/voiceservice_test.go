package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/goinggo/mapstructure"
	"gvp/models"
	"gvp/pkg/app"
	"gvp/pkg/logging"
	"gvp/pkg/setting"
	"gvp/routers"
	"net/http/httptest"
	"strconv"

	//"strconv"
	"strings"
	"testing"
)

/*
		//获取单个服务商信息--ok
		apiv1.GET("/svc/:id", v1.GetVoiceService)
		//获取服务商列表
		apiv1.GET("/svc", v1.GetVoiceServices)
		//添加服务商--ok
		apiv1.POST("/svc", v1.AddVoiceService)
		//更新指定的服务商
		apiv1.PUT("/svc/:id", v1.UpdateVoiceService)
		//删除指定的服务商
		apiv1.DELETE("/svc/:id", v1.DelVoiceService)
*/
//测试准备数据


var testVoiceSVC =  models.VoiceService{
	Name:"测试服务商",
	AppId:"123123123",
	AppKey:"12312312",
	UserId:"12123",
	URL:"www.baidu.com",
	Description:"这是一个测试服务商",
	CreatedBy:"邓全杰",
	ModifiedBy:"邓全杰",
	State:1,
}



func  init(){
	setting.Setup()
	logging.Setup()
	models.Setup()

	router = routers.InitRouter()
}

var  router  *gin.Engine
var  TokenValue string
var  Svc_db_id  int


type aa struct {
	ID int `gorm:"primary_key" json:"id"`
}

func getToken( )(bool,string){
	w:= httptest.NewRecorder()
	req:= httptest.NewRequest("GET","/auth?username=admin&password=admin",nil)
	router.ServeHTTP(w,req)

	var    token  string
	var  res   app.Response
	err :=  json.Unmarshal(w.Body.Bytes(),&res)
	if err !=  nil{
		return  false,token
	}
	if w.Code ==200{
		token = res.Data.(map[string]interface{})["token"].(string)
	}
	return  true,token
}


func   TestVoiceSever(t *testing.T){
		Convey("服务商相关测试",t, func() {
			Convey("获取token值", func() {
				var ret bool
				ret,TokenValue =  getToken()
				So(ret,ShouldBeTrue)
				So(TokenValue,ShouldNotBeNil)
			})

			Convey("增加服务商,缺少必要参数", func() {
				w:= httptest.NewRecorder()
				var target = "/api/v1/svc"
				target =  target+"?token="+TokenValue

				var tmp_vs   struct {
					Name string
				}
				tmp_vs.Name = "测试服务商"
				tmp,_ :=  json.Marshal(tmp_vs)
				req:= httptest.NewRequest("POST",target,strings.NewReader(string(tmp[:])))
				req.Header.Add("Content-type","application/json")
				router.ServeHTTP(w,req)
				So(w.Code,ShouldNotEqual,200)
			})

			Convey("增加服务商，错误的参数内容", func() {
				w:= httptest.NewRecorder()
				var target = "/api/v1/svc"
				target =  target+"?token="+TokenValue

				stubs :=  gostub.Stub(&(testVoiceSVC.State),12)
				defer stubs.Reset()

				tmp,_ :=  json.Marshal(testVoiceSVC)
				req:= httptest.NewRequest("POST", target,strings.NewReader(string(tmp[:])))
				req.Header.Add("Content-type","application/json")
				router.ServeHTTP(w,req)
				So(w.Code,ShouldNotEqual,200)
			})

			Convey("增加服务商，成功添加", func() {
					w:= httptest.NewRecorder()
					var target = "/api/v1/svc"
					target =  target+"?token="+TokenValue

					tmp,_ :=  json.Marshal(testVoiceSVC)
					req:= httptest.NewRequest("POST", target,strings.NewReader(string(tmp[:])))
					req.Header.Add("Content-type","application/json")
					router.ServeHTTP(w,req)
					So(w.Code,ShouldEqual,200)
			})

			Convey("获取所有的服务商", func() {
				w:= httptest.NewRecorder()
				var target = "/api/v1/svc"
				target =  target+"?token="+TokenValue
				req:= httptest.NewRequest("GET",target,nil)
				router.ServeHTTP(w,req)
				So(w.Code,ShouldEqual,200)
				if w.Code  == 200{
					var  res   app.Response
					err :=  json.Unmarshal(w.Body.Bytes(),&res)

					if err !=  nil{
						t.Error("拉取服务商列表进行json解析报错")
					}

					//logging.Info(res)
					total  := res.Data.(map[string]interface{})["total"].(float64)
					if   total >= 1{
						lists := res.Data.(map[string]interface{})["lists"].([]interface {})

						for _,v:=range lists  {
							var  tmpsvc aa
							mapstructure.Decode(v, &tmpsvc)
							Svc_db_id =  tmpsvc.ID
							logging.Info(Svc_db_id)
						}
					}
				}

			})



			Convey("获取单个不存在的服务商", func() {
				w:= httptest.NewRecorder()
				var target = "/api/v1/svc/1000000"
				target =  target+"?token="+TokenValue
				req:= httptest.NewRequest("GET",target,nil)
				router.ServeHTTP(w,req)
				So(w.Code,ShouldNotEqual,200)
			})

			Convey("获取单个存在的服务商", func() {
				w:= httptest.NewRecorder()
				var target = "/api/v1/svc/"+strconv.Itoa(Svc_db_id)
				target =  target+"?token="+TokenValue
				req:= httptest.NewRequest("GET",target,nil)
				router.ServeHTTP(w,req)
				So(w.Code,ShouldEqual,200)
			})

			Convey("更新服务商信息", func() {
				w:= httptest.NewRecorder()
				var target = "/api/v1/svc/"+strconv.Itoa(Svc_db_id)
				logging.Info("更新服务商信息:",target)
				target =  target+"?token="+TokenValue
				testVoiceSVC.ID = Svc_db_id
				stubs :=   gostub.Stub(&(testVoiceSVC.Name),"测试修改服务商名称")
				defer  stubs.Reset()

				tmp,_ :=  json.Marshal(testVoiceSVC)

				req:= httptest.NewRequest("PUT", target,strings.NewReader(string(tmp[:])))
				req.Header.Add("Content-type","application/json")
				router.ServeHTTP(w,req)
				So(w.Code,ShouldEqual,200)
			})



			Convey("删除单个不存在的服务商", func() {
				w:= httptest.NewRecorder()
				var target = "/api/v1/svc/1000000"
				target =  target+"?token="+TokenValue
				req:= httptest.NewRequest("DELETE",target,nil)
				router.ServeHTTP(w,req)
				So(w.Code,ShouldNotEqual,200)
			})

			Convey("删除存在的服务商", func() {
				w:= httptest.NewRecorder()
				var target = "/api/v1/svc/"+strconv.Itoa(Svc_db_id)
				logging.Info("删除存在的服务商:",target)
				target =  target+"?token="+TokenValue
				req:= httptest.NewRequest("DELETE",target,nil)
				router.ServeHTTP(w,req)
				So(w.Code,ShouldEqual,200)
			})
		})
}

