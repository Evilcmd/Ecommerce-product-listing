package main

import (
	"fmt"
	"io"
	"net/http"
)

func (lbCfg *lbConfig) loadBalance(res http.ResponseWriter, req *http.Request) {
	if lbCfg.lastserver >= len(lbCfg.serversList) {
		lbCfg.lastserver = 0
	} else {
		lbCfg.lastserver = (lbCfg.lastserver + 1) % len(lbCfg.serversList)
	}
	port := lbCfg.serversList[lbCfg.lastserver]

	targetUrl := fmt.Sprintf("http://localhost:%v%v", port, req.URL.Path)
	new_req, _ := http.NewRequest(req.Method, targetUrl, req.Body)
	new_req.Header = req.Header.Clone()
	new_res, err := http.DefaultClient.Do(new_req)
	if err != nil {
		respondWithError(res, 406, fmt.Sprintf("error from loadbalancer: %v", err.Error()))
		return
	}
	defer new_res.Body.Close()

	bodyBytes, _ := io.ReadAll(new_res.Body)

	res.Write(bodyBytes)

	fmt.Println(targetUrl)

}
