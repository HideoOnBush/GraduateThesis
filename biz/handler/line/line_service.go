// Code generated by hertz generator.

package line

import (
	"GraduateThesis/biz/handler"
	"context"
	"net/http"

	baseModel "GraduateThesis/biz/model/base"
	lineModel "GraduateThesis/biz/model/line"
	"GraduateThesis/biz/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Query .
// @router /api/line/query [GET]
func Query(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lineModel.LineReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	lines, totalCount := service.Line.Query(ctx, c, &req)
	if len(c.Errors) > 0 {
		var errMessage = c.Errors.String()
		c.JSON(http.StatusInternalServerError, lineModel.LineResp{
			Code: 1,
			Data: &lineModel.LineRespSample{
				Lines:      make([]*lineModel.Line, 0),
				TotalCount: 0,
			},
			Msg: errMessage,
		})
	}
	c.JSON(consts.StatusOK, lineModel.LineResp{
		Code: handler.SuccessCode,
		Data: &lineModel.LineRespSample{
			Lines:      lines,
			TotalCount: int64(totalCount),
		},
		Msg: handler.EmptyMessage,
	})
}

// ChangeDependence .
// @router /api/line/change-dependence [POST]
func ChangeDependence(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lineModel.ChangeDependenceWithRelationReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	result := service.Line.ChangeDependenceByRelation(ctx, c, req)
	if len(c.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, baseModel.SampleResp{
			Code:    1,
			Data:    false,
			Message: c.Errors.String(),
		})
	}

	resp := new(baseModel.SampleResp)

	if result == false {
		resp.Code = 1
		resp.Data = false
	} else {
		resp.Data = true
		resp.Code = 0
		resp.Message = ""
	}

	c.JSON(consts.StatusOK, resp)
}

// Bulk .
// @router /api/line/bulk [POST]
func Bulk(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lineModel.LineBulkReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	lines := make([]*lineModel.Line, 0, len(req.Lines))

	resp := new(baseModel.SampleResp)

	err = service.Line.Bulk(ctx, lines)
	if err != nil {
		if err != nil {
			c.String(400, err.Error())
			return
		}
	}
	resp.Code = 0
	resp.Data = true
	resp.Message = ""

	c.JSON(consts.StatusOK, resp)
}

// Delete .
// @router /api/line/delete [POST]
func Delete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lineModel.LineDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	result, err := service.Line.DeleteById(ctx, req.GetScene(), req.GetDocId())
	if err != nil {
		c.String(400, err.Error())
		return
	}
	resp := new(baseModel.SampleResp)
	if result == false {
		resp.Code = 1
		resp.Data = false
	} else {
		resp.Data = true
		resp.Code = 0
		resp.Message = ""
	}
	c.JSON(consts.StatusOK, resp)
}

// TopologyAnalyse .
// @router /api/line/topology_analyse [POST]
func TopologyAnalyse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lineModel.TopologyIndicatorReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	ok, result := service.Line.TopologyAnalyse(ctx, &req)
	resp := new(lineModel.TopologyIndicatorResp)

	if !ok {
		resp.Data = nil
		resp.Msg = ""
		resp.Code = 1
	} else {
		resp.Data = &result
		resp.Msg = ""
		resp.Code = 0
	}

	c.JSON(consts.StatusOK, resp)
}