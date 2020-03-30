package test


import (
	"encoding/json"
	"github.com/goinggo/mapstructure"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"gvp/models"
	"gvp/pkg/app"
	"gvp/pkg/logging"
	"gvp/pkg/setting"
	"gvp/routers"
	"net/http/httptest"
	"strings"
	"testing"
)


//测试准备数据
var  testGame =  models.Game{
	Game:"测试游戏",
	GameType:1,
	TelNum:"15801727928",
	CreatedBy:"测试",
	ModifiedBy:"测试",
	State:1,
}



func  init(){
	setting.Setup()
	logging.Setup()
	models.Setup()

	router = routers.InitRouter()
}

var  game_db_id  int


type game_id struct {
	ID int `gorm:"primary_key" json:"id"`
}


func   TestGame(t *testing.T){
	Convey("游戏相关测试",t, func() {
		Convey("获取token值", func() {
			var ret bool
			ret,TokenValue =  getToken()
			So(ret,ShouldBeTrue)
			So(TokenValue,ShouldNotBeNil)
		})

		Convey("增加游戏,缺少必要参数", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game"
			target =  target+"?token="+TokenValue

			var tmp_vs   struct {
				Game string
			}
			tmp_vs.Game = "测试游戏"
			tmp,_ :=  json.Marshal(tmp_vs)
			req:= httptest.NewRequest("POST",target,strings.NewReader(string(tmp[:])))
			req.Header.Add("Content-type","application/json")
			router.ServeHTTP(w,req)
			So(w.Code,ShouldNotEqual,200)
		})

		Convey("增加游戏，错误的参数内容", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game"
			target =  target+"?token="+TokenValue

			stubs :=  gostub.Stub(&(testGame.State),12)
			defer stubs.Reset()

			tmp,_ :=  json.Marshal(testGame)
			req:= httptest.NewRequest("POST", target,strings.NewReader(string(tmp[:])))
			req.Header.Add("Content-type","application/json")
			router.ServeHTTP(w,req)
			So(w.Code,ShouldNotEqual,200)
		})

		Convey("增加游戏，成功添加", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game"
			target =  target+"?token="+TokenValue

			tmp,_ :=  json.Marshal(testGame)
			req:= httptest.NewRequest("POST", target,strings.NewReader(string(tmp[:])))
			req.Header.Add("Content-type","application/json")
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
		})

		Convey("获取所有的游戏列表", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game"
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("GET",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
			if w.Code  == 200{
				var  res   app.Response
				err :=  json.Unmarshal(w.Body.Bytes(),&res)

				if err !=  nil{
					t.Error("拉取游戏列表进行json解析报错")
				}

				//logging.Info(res)
				total  := res.Data.(map[string]interface{})["total"].(float64)
				if   total >= 1{
					lists := res.Data.(map[string]interface{})["lists"].([]interface {})

					for _,v:=range lists  {
						var  tmpgame game_id
						mapstructure.Decode(v, &tmpgame)
						game_db_id =  tmpgame.ID
						logging.Info(game_db_id)
					}
				}
			}

		})



		Convey("获取单个不存在的游戏", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game/1000000"
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("GET",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldNotEqual,200)
		})

		Convey("获取单个存在的游戏", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game/"+strconv.Itoa(game_db_id)
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("GET",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
		})

		Convey("更新游戏信息", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game/"+strconv.Itoa(game_db_id)
			logging.Info("更新游戏信息:",target)
			target =  target+"?token="+TokenValue
			testGame.ID =   game_db_id
			stubs :=   gostub.Stub(&(testGame.Game),"测试修改游戏名称")
			defer  stubs.Reset()

			tmp,_ :=  json.Marshal(testGame)

			req:= httptest.NewRequest("PUT", target,strings.NewReader(string(tmp[:])))
			req.Header.Add("Content-type","application/json")
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
		})



		Convey("删除单个不存在的游戏", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game/1000000"
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("DELETE",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldNotEqual,200)
		})

		Convey("删除存在的游戏", func() {
			w:= httptest.NewRecorder()
			var target = "/api/v1/game/"+strconv.Itoa(game_db_id)
			logging.Info("删除存在的游戏:",target)
			target =  target+"?token="+TokenValue
			req:= httptest.NewRequest("DELETE",target,nil)
			router.ServeHTTP(w,req)
			So(w.Code,ShouldEqual,200)
		})
	})
}
