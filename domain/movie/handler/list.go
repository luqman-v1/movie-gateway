package handler

import (
	"context"
	"movie-gateway/domain/movie"
	"movie-gateway/domain/movie/client"
	pb "movie-gateway/proto/movie"
	"movie-gateway/util"

	"google.golang.org/grpc/status"

	"github.com/labstack/echo/v4"
)

type List struct{}

func NewList() *List {
	return &List{}
}

func (h *List) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req := &pb.ListRequest{
		Page:   util.StoI(c.QueryParam("pagination")),
		Search: c.QueryParam("searchword"),
	}
	respList, err := client.List(ctx, req)
	if err != nil {
		return err
	}
	if err != nil {
		if st, ok := status.FromError(err); ok {
			resp := &util.Response{
				Code:    st.Code(),
				Message: util.StatusMessage[st.Code()],
			}
			return resp.JSON(c)
		}
		return err
	}
	resp, err := h.buildResponse(respList)
	if err != nil {
		return err
	}
	return resp.JSON(c)
}

func (h *List) buildResponse(res *pb.ListResponse) (*util.Response, error) {
	code := util.Success
	data := movie.Data{}
	resp := &util.Response{
		Code:    code,
		Message: util.StatusMessage[code],
		Data: map[string]interface{}{
			"movies": data.Map(res),
		},
	}
	resp.Pagination = &util.Pg{
		CurrentPage: res.Pagination.CurrentPage,
		PageSize:    res.Pagination.PageSize,
		TotalPage:   res.Pagination.TotalPage,
		TotalResult: res.Pagination.TotalResult,
	}
	return resp, nil
}
