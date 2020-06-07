package http

import (
	"fmt"
	"io/ioutil"

	"switch-onchain/internal/server/http/proto"

	"github.com/gin-gonic/gin"
)

// @Summary节点健康检查接口
// @Description 节点健康检查接口
// @Tags Health
// @Produce json
// @Success 200 object proto.Response
// @Router / [get]
func GET_Echo(c *gin.Context) {
	response := "server echo"

	reqBody, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		errs := fmt.Errorf("failed to read http body error: %s", err.Error())
		proto.ResponseAbortLogic(c, proto.ERROR, errs)
		return
	}
	if len(reqBody) > 0 {
		response = string(reqBody[:])
	}
	proto.ResponseSuccessLogic(c, proto.SUCCESS, response)
}
