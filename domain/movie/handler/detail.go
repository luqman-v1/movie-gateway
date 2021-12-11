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

type Detail struct{}

func NewDetail() *Detail {
	return &Detail{}
}

func (h *Detail) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req := &pb.DetailRequest{
		ImdbId: c.Param("imdb_id"),
	}
	respList, err := client.Detail(ctx, req)
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

func (h *Detail) buildResponse(res *pb.DetailMovie) (*util.Response, error) {
	code := util.Success
	data := movie.DetailMovie{}
	resp := &util.Response{
		Code:    code,
		Message: util.StatusMessage[code],
		Data: map[string]interface{}{
			"movie": data.Map(res),
		},
	}

	return resp, nil
}
