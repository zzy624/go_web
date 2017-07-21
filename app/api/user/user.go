package user

import (
	"net/http"
	"strconv"

	db "mycode/go_web/app/common/model"
	"mycode/go_web/app/common/response"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func PostLogin(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	var uid int

	cmd := `SElECT count(*) from userinfo where username = $1 AND password = $2`
	err := db.Db.QueryRow(cmd, username, password).Scan(&uid)

	if err != nil {
		log.Error(err)
	}
	if uid == 0 {
		return ctx.JSON(http.StatusOK, &response.Response{Code: "13", Msg: "用户名或密码错误！", Data: nil})
	}
	return ctx.JSON(http.StatusOK, &response.Response{Code: "0", Msg: "", Data: &response.LoginResponse{UserID: 1, UserName: "张三"}})
}

func PostAddProduct(ctx echo.Context) error {
	userID := ctx.FormValue("user_id")
	productName := ctx.FormValue("product_name")
	stmt, err := db.Db.Prepare("INSERT INTO trade(uid, tradename) values($1,$2)")

	if err != nil {
		log.Error("INSERT INTO trade failed", err)
		return ctx.JSON(http.StatusOK, &response.Response{Code: "13", Msg: "添加商品失败！", Data: nil})
	}

	_, err = stmt.Exec(userID, productName)
	if err != nil {
		log.Error("INSERT INTO trade failed", err, "userID", userID, "productName", productName)
		return ctx.JSON(http.StatusOK, &response.Response{Code: "13", Msg: "添加商品失败！", Data: nil})

	}
	return ctx.JSON(http.StatusOK, &response.Response{Code: "0", Msg: "添加商品成功！", Data: nil})

}

func GetProduct(ctx echo.Context) error {
	userID := ctx.FormValue("user_id")
	intUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		log.Error(err, userID)
		return ctx.JSON(http.StatusOK, &response.Response{Code: "13", Msg: "参数输入有误！", Data: nil})
	}

	rows, err := db.Db.Query("SELECT tid,uid,tradename FROM trade WHERE uid = $1", intUserID)
	defer rows.Close()

	if err != nil {
		log.Error("SELECT id,user_id,product_name FROM trade >>error", err, "userID", userID)
		return ctx.JSON(http.StatusOK, &response.Response{Code: "13", Msg: "没有查询到商品！", Data: nil})
	}
	productList := make([]response.Product, 0)
	for rows.Next() {
		var trade response.Product
		err = rows.Scan(&trade.ProductID, &trade.UserID, &trade.ProductName)
		if err != nil {
			log.Error(err)
			return ctx.JSON(http.StatusOK, &response.Response{Code: "13", Msg: "没有查询到商品！", Data: nil})
		}
		productList = append(productList, trade)
	}
	return ctx.JSON(http.StatusOK, &response.Response{Code: "0", Msg: "", Data: &response.GetProductResponse{Trade: productList}})

}
