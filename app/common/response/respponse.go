package response

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type LoginResponse struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type Product struct {
	UserID      int64  `json:"user_id"`
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
}

type GetProductResponse struct {
	Trade []Product `json:"trade"`
}
