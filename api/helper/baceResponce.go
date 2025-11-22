package helper

type Resp struct {
	Rsult   string `json:"result"`
	Sussecc bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Sussecc(rsult, message string, code int, sussecc bool) Resp {
	return Resp{
		Rsult:   rsult,
		Sussecc: true,
		Code:    code,
		Message: message,
	}
}

func NotFound(rsult, message string, code int, sussecc bool) Resp {
	return Resp{
		Rsult:   rsult,
		Sussecc: false,
		Code:    404,
		Message: "Not found Error",
	}
}
