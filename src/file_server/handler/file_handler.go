package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/xi123/libgo/logs"
	"github.com/xi123/uploader/src/global"
	"github.com/xi123/uploader/src/global/httpsrv"
)

func handlerFileJsonReq(body []byte) (*global.FileInfoResp, bool) {
	if len(body) == 0 {
		return &global.FileInfoResp{ErrCode: 3, ErrMsg: "no body"}, false
	}
	logs.Warnf("%v", string(body))
	req := global.FileInfoReq{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		logs.Errorf(err.Error())
		return &global.FileInfoResp{ErrCode: 4, ErrMsg: "parse body error"}, false
	}
	if req.Md5 == "" && len(req.Md5) != 32 {
		return &global.FileInfoResp{Md5: req.Md5, ErrCode: 1, ErrMsg: "parse param error"}, false
	}
	return QueryCacheFile(req.Md5)
}

func handlerFileQuery(query url.Values) (*global.FileInfoResp, bool) {
	var md5 string
	if query.Has("md5") && len(query["md5"]) > 0 {
		md5 = query["md5"][0]
	}
	if md5 == "" && len(md5) != 32 {
		return &global.FileInfoResp{Md5: md5, ErrCode: 1, ErrMsg: "parse param error"}, false
	}
	return QueryCacheFile(md5)
}

func FileinfoReq(w http.ResponseWriter, r *http.Request) {
	logs.Infof("%v %v %#v", r.Method, r.URL.String(), r.Header)
	switch strings.ToUpper(r.Method) {
	case "POST":
		switch r.Header.Get("Content-Type") {
		case "application/json":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logs.Errorf(err.Error())
				resp := &global.FileInfoResp{ErrCode: 2, ErrMsg: "read body error"}
				httpsrv.WriteResponse(w, r, resp)
				return
			}
			resp, _ := handlerFileJsonReq(body)
			httpsrv.WriteResponse(w, r, resp)
		default:
			resp, _ := handlerFileQuery(r.URL.Query())
			httpsrv.WriteResponse(w, r, resp)
		}
	case http.MethodGet:
		switch r.Header.Get("Content-Type") {
		case "application/json":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logs.Errorf(err.Error())
				resp := &global.FileInfoResp{ErrCode: 2, ErrMsg: "read body error"}
				httpsrv.WriteResponse(w, r, resp)
				return
			}
			resp, _ := handlerFileJsonReq(body)
			httpsrv.WriteResponse(w, r, resp)
		default:
			resp, _ := handlerFileQuery(r.URL.Query())
			httpsrv.WriteResponse(w, r, resp)
		}
	case http.MethodOptions:
		switch r.Header.Get("Content-Type") {
		case "application/json":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logs.Errorf(err.Error())
				resp := &global.FileInfoResp{ErrCode: 2, ErrMsg: "read body error"}
				httpsrv.WriteResponse(w, r, resp)
				return
			}
			resp, _ := handlerFileJsonReq(body)
			httpsrv.WriteResponse(w, r, resp)
		default:
			resp, _ := handlerFileQuery(r.URL.Query())
			httpsrv.WriteResponse(w, r, resp)
		}
	}
}

func handlerFileDetailJsonReq(body []byte) (*global.FileDetailResp, bool) {
	if len(body) == 0 {
		return &global.FileDetailResp{ErrCode: 3, ErrMsg: "no body"}, false
	}
	logs.Warnf("%v", string(body))
	req := global.FileDetailReq{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		logs.Errorf(err.Error())
		return &global.FileDetailResp{ErrCode: 4, ErrMsg: "parse body error"}, false
	}
	if req.Md5 == "" && len(req.Md5) != 32 {
		return &global.FileDetailResp{ErrCode: 1, ErrMsg: "parse param error"}, false
	}
	return QueryCacheFileDetail(req.Md5)
}

func handlerFileDetailQuery(query url.Values) (*global.FileDetailResp, bool) {
	var md5 string
	if query.Has("md5") && len(query["md5"]) > 0 {
		md5 = query["md5"][0]
	}
	if md5 == "" && len(md5) != 32 {
		return &global.FileDetailResp{ErrCode: 1, ErrMsg: "parse param error"}, false
	}
	return QueryCacheFileDetail(md5)
}

func FileDetailReq(w http.ResponseWriter, r *http.Request) {
	logs.Infof("%v %v %#v", r.Method, r.URL.String(), r.Header)
	switch strings.ToUpper(r.Method) {
	case "POST":
		switch r.Header.Get("Content-Type") {
		case "application/json":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logs.Errorf(err.Error())
				resp := &global.FileDetailResp{ErrCode: 2, ErrMsg: "read body error"}
				httpsrv.WriteResponse(w, r, resp)
				return
			}
			resp, _ := handlerFileDetailJsonReq(body)
			httpsrv.WriteResponse(w, r, resp)
		default:
			resp, _ := handlerFileDetailQuery(r.URL.Query())
			httpsrv.WriteResponse(w, r, resp)
		}
	case http.MethodGet:
		switch r.Header.Get("Content-Type") {
		case "application/json":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logs.Errorf(err.Error())
				resp := &global.FileDetailResp{ErrCode: 2, ErrMsg: "read body error"}
				httpsrv.WriteResponse(w, r, resp)
				return
			}
			resp, _ := handlerFileDetailJsonReq(body)
			httpsrv.WriteResponse(w, r, resp)
		default:
			resp, _ := handlerFileDetailQuery(r.URL.Query())
			httpsrv.WriteResponse(w, r, resp)
		}
	case http.MethodOptions:
		switch r.Header.Get("Content-Type") {
		case "application/json":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logs.Errorf(err.Error())
				resp := &global.FileDetailResp{ErrCode: 2, ErrMsg: "read body error"}
				httpsrv.WriteResponse(w, r, resp)
				return
			}
			resp, _ := handlerFileDetailJsonReq(body)
			httpsrv.WriteResponse(w, r, resp)
		default:
			resp, _ := handlerFileDetailQuery(r.URL.Query())
			httpsrv.WriteResponse(w, r, resp)
		}
	}
}
