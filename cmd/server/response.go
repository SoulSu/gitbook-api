package server

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s *SrvHandle) ResponseJSON() {

}
