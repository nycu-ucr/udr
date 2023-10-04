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

// HTTPApplicationDataInfluenceDataGet -
func HTTPApplicationDataInfluenceDataGet(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Query["influence-Ids"] = c.QueryArray("influence-Ids")
	req.Query["dnns"] = c.QueryArray("dnns")
	req.Query["snssais"] = c.QueryArray("snssais")
	req.Query["internal-Group-Ids"] = c.QueryArray("internal-Group-Id")
	req.Query["supis"] = c.QueryArray("supis")
	rsp := producer.HandleApplicationDataInfluenceDataGet(req)
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
