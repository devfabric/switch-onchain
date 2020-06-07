package v1

import (
	"switch-onchain/internal/server/http/proto"

	"github.com/gin-gonic/gin"
)

// @Summary资源归属者删除权限
// @Description 资源归属者删除权限
// @Tags Auth
// @Accept  json
// @Produce json
// @Param RAC_TOKEN header string false "token令牌"
// @Param deleteResAuth body model.DeleteAsset true "请求信息"
// @Success 200 object proto.Response
// @Router /api/v1/res/deleteResAuth [post]
func (v1 *APIV1) API(c *gin.Context) {
	// var (
	// 	fabOrg     string
	// 	fabUser    string
	// 	deleteAuth model.DeleteAsset
	// )

	// err := c.BindJSON(&deleteAuth)
	// if err != nil {
	// 	log.Logger.Debugf("DeleteResAuth BindJSON Err:%v", err)
	// 	proto.ResponseAbortLogic(c, proto.INVALID_PARAMS, err)
	// 	return
	// }

	// if deleteAuth.Domain == "" ||
	// 	len(deleteAuth.ResList) == 0 || deleteAuth.OwnAddr == "" {
	// 	log.Logger.Debug("DeleteResAuth params:", deleteAuth)
	// 	proto.ResponseAbortLogic(c, proto.INVALID_PARAMS, public.ErrParameter)
	// 	return
	// }

	// if orgUserObj, ok := c.Get(proto.FabUserOrgKey); ok {
	// 	fabUser = orgUserObj.([]string)[0]
	// 	fabOrg = orgUserObj.([]string)[1]
	// }

	// response, err := v1.SVC.FabricIns.ExecuteCC(fabOrg, fabUser, common.ResAuthRoutes, common.DeleteResAuthFunc, deleteAuth)
	// if err != nil {
	// 	log.Logger.Debugf("DeleteResAuth QueryCCByOrgUser Err:%v", err)
	// 	proto.ResponseAbortLogic(c, proto.ERROR, public.ErrMsgFormat(err))
	// 	return
	// } else {
	// 	log.Logger.Debug("Query Tx Response: \n",
	// 		"txhash: ", response.TransactionID, "\n",
	// 		"valiClode: ", response.TxValidationCode, "\n",
	// 		"ccstatus: ", response.ChaincodeStatus, "\n",
	// 		"payload: ", string(response.Payload))
	// }

	// if err != nil {
	// 	log.Logger.Debugf("DeleteResAuth Unmarshal Err:%v", err)
	// 	proto.ResponseAbortLogic(c, proto.ERROR, err)
	// 	return
	// }

	proto.ResponseSuccessLogic(c, proto.SUCCESS, true)
}
