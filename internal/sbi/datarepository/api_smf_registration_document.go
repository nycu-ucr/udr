/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"

	"github.com/nycu-ucr/openapi"
	"github.com/nycu-ucr/openapi/models"
	"github.com/free5gc/udr/internal/logger"
	"github.com/free5gc/udr/internal/sbi/producer"
	"github.com/nycu-ucr/util/httpwrapper"
)

// HTTPCreateSmfContextNon3gpp - To create an individual SMF context data of a UE in the UDR
func HTTPCreateSmfContextNon3gpp(c *gin.Context) {
	var smfRegistration models.SmfRegistration

	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.DataRepoLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&smfRegistration, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.DataRepoLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(c.Request, smfRegistration)
	req.Params["ueId"] = c.Params.ByName("ueId")

	rsp := producer.HandleCreateSmfContextNon3gpp(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// HTTPDeleteSmfContext - To remove an individual SMF context data of a UE the UDR
func HTTPDeleteSmfContext(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["pduSessionId"] = c.Params.ByName("pduSessionId")

	rsp := producer.HandleDeleteSmfContext(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// HTTPQuerySmfRegistration - Retrieves the individual SMF registration of a UE
func HTTPQuerySmfRegistration(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["pduSessionId"] = c.Params.ByName("pduSessionId")

	rsp := producer.HandleQuerySmfRegistration(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
