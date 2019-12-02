package router

import (
	"github.com/labstack/echo"
)

// GetRouter ルーティング
func GetRouter(e *echo.Echo) {

	// api := e.Group("/api")
	e.GET("/", "public/index.html")

	// //tfc db areaテーブル取得サンプル
	// g2 := g.Group("/area")
	// g2.GET("", controller.ShowAllAreas)
	// g2.GET("/:id", controller.ShowFindArea)
	// g2.GET("/param", controller.ShowFindArea2)
	// g2.POST("/all", controller.ShowAllAreas)
	// g2.POST("/id", controller.ShowFindAreaPost)
	// g2.POST("/param", controller.ShowFindArea2Post)

	// //DBのリレーションサンプル
	// g4 := g.Group("/members")
	// g4.GET("/pref", controller.ShowMemberPref)
	// g4.GET("/preforder", controller.ShowMemberPrefOrders)

	// //ファイルアップロードサンプル
	// e.POST("/upload", controller.Upload)

	// //ファイルダウンロードサンプル
	// g3 := e.Group("/download")
	// g3.GET("/url", controller.URLDownload)

	// //xml送受信サンプル(tfc会員情報取得)
	// e.GET("/xml", controller.GetAPIXml)
	// e.GET("/xml2", controller.GetAPIXml2)
	// e.GET("/xml3", controller.GetAPIXml3)

	// //APIトークンサンプル
	// e.GET("/token", controller.TokenAPI)
	// e.POST("/tokenheader", controller.HeaderToken)

	// //メール送信サンプル
	// e.GET("/mail", controller.AdminMailIso2022Jp)

	// //ポイント基盤 ポイント付与利用サンプル
	// g.POST("/point", controller.Point)

	// //バリデーションサンプル
	// e.POST("/validate", controller.Validate)
	// e.POST("/validate2", controller.Validate2)

	log.Fatal(e.Start(":8081"))

}
