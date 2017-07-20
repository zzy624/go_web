package user

import (
	"database/sql"
	"net/http"

	"mycode/go_web/app/common/response"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func PostLogin(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	if username == "zhangsan" && password == "123123" {
		return ctx.JSON(http.StatusOK, &response.Response{Code: "0", Msg: "", Data: &response.LoginResponse{UserID: 1, UserName: "张三"}})
	}
	return ctx.JSON(http.StatusOK, &response.Response{Code: "13", Msg: "用户名活密码错误！", Data: nil})
}

func PostAddProduct(ctx echo.Context) error {
	userID := ctx.FormValue("user_id")
	productName := ctx.FormValue("product_name")
	db, err := sql.Open("sqlite3", "mycode/go_web/db.sqlite3")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO product(user_id, product_name) values(?,?)")
	checkErr(err)

	res, err := stmt.Exec(userID, productName)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	log.Info(id)

	return ctx.JSON(http.StatusOK, &response.Response{Code: "0", Msg: "添加商品成功！", Data: nil})

}

func GetProduct(ctx echo.Context) error {
	userID := ctx.Param("product_id")
	db, err := sql.Open("sqlite3", "mycode/go_web/db.sqlite3")
	checkErr(err)

	rows, err := db.Query("SELECT id,user_id,product_name FROM product WHERE user_id = ?", userID)
	checkErr(err)
	productList := make([]response.Product, 0)
	for rows.Next() {
		var trade response.Product
		err = rows.Scan(&trade.ProductID, &trade.UserID, &trade.ProductName)
		checkErr(err)
		productList = append(productList, trade)

		return ctx.JSON(http.StatusOK, &response.Response{Code: "0", Msg: "", Data: &response.GetProductResponse{Trade: productList}})
	}
	return nil
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
