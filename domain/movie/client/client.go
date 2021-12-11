package client

import (
	"context"
	pb "movie-gateway/proto/movie"
	"movie-gateway/util"
	"os"
)

func List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	conn := util.Dial(os.Getenv("movie_grpc"))
	defer conn.Close()
	client := pb.NewMovieClient(conn)
	return client.List(ctx, request)
}

func Detail(ctx context.Context, request *pb.DetailRequest) (*pb.DetailMovie, error) {
	conn := util.Dial(os.Getenv("movie_grpc"))
	defer conn.Close()
	client := pb.NewMovieClient(conn)
	return client.Detail(ctx, request)
}
