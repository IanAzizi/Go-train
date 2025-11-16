package helper

import "net/http"

type Responce struct {
	Result  string `json:"result"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}

func SuccessBaceResponce(w http.ResponseWriter , Result string , Success bool , Code int) *Responce {
	return &Responce{
		Result:  "success",
		Success: true,
		Code:    http.StatusOK,
	}
}


func UnAthoraize() *Responce{
	return &Responce{
		Result: "Unathorazie",
		Success: false,
		Code: http.StatusUnauthorized,
	}
}



