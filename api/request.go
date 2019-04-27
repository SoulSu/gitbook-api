package api

import (
	"sync"

	"github.com/parnurzeal/gorequest"
)

// the gitbook api address
const API_ROOT = "https://api.gitbook.com/"

type Requester interface {
	Uri() string
	Method() string
	Response() Responseer
}

type Responseer interface {
	SetError([]error)
	Error() []error

	SetHttpCode(int)
	HttpCode() int

	Success() bool
}

type CommonResponse struct {
	errs     []error
	httpCode int
}

func (r *CommonResponse) SetError(errs []error) {
	r.errs = errs
}

func (r *CommonResponse) Error() []error {
	return r.errs
}

func (r *CommonResponse) SetHttpCode(code int) {
	r.httpCode = code
}

func (r *CommonResponse) HttpCode() int {
	return r.httpCode
}

func (r *CommonResponse) Success() bool {
	return r.httpCode == 200 && len(r.errs) <= 0
}

type Request struct {
	pool sync.Pool
}

var once sync.Once
var request *Request

func NewRequest() *Request {
	once.Do(func() {
		request = new(Request)
		request.pool = sync.Pool{
			New: func() interface{} {
				return gorequest.New()
			},
		}
	})
	return request
}

func (r *Request) get() *gorequest.SuperAgent {
	return r.pool.Get().(*gorequest.SuperAgent)
}

func (r *Request) put(req *gorequest.SuperAgent) {
	req.ClearSuperAgent()
	r.pool.Put(req)
}

func (r *Request) Do(api Requester) Responseer {
	reqUrl := API_ROOT + api.Uri()
	apiResp := api.Response()
	agent := r.get()
	defer r.put(agent)

	agent = agent.CustomMethod(api.Method(), reqUrl)
	agent.Send(nil).EndStruct(apiResp, func(response gorequest.Response, v interface{}, body []byte, errs []error) {
		apiResp.SetError(errs)
		apiResp.SetHttpCode(response.StatusCode)
	})

	return apiResp
}
