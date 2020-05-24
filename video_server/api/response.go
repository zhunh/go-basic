package main

import (
	"encoding/json"
	"going/video_server/api/defs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errRsp defs.ErrorResponse)  {
	w.WriteHeader(errRsp.HttpSC)

	resStr, _:= json.Marshal(&errRsp.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int)  {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
