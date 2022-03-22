package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"your/path/project/domain_myspending/usecase/runexpensecreate"
	"your/path/project/shared/infrastructure/logger"
	"your/path/project/shared/infrastructure/util"
	"your/path/project/shared/model/payload"
)

// runExpenseCreateHandler ...
func (r *Controller) runExpenseCreateHandler(inputPort runexpensecreate.Inport) gin.HandlerFunc {

	type request struct {
		Nilai     int    `json:"nilai"`
		Deskripsi string `json:"deskripsi"`
		Tanggal   string `json:"tanggal"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req runexpensecreate.InportRequest
		req.Nilai = jsonReq.Nilai
		req.Tanggal = jsonReq.Tanggal
		req.Deskripsi = jsonReq.Deskripsi

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
