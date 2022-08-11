package grpcclients

import (
	"github.com/pkg/errors"
	"github.com/xiaoJack/GraphQL-Golang/api/proto"
	"github.com/xiaoJack/GraphQL-Golang/internal/pkg/transports/grpc"
)

func NewRatingsClient(client *grpc.Client) (proto.RatingsClient, error) {
	conn, err := client.Dial("Ratings")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewRatingsClient(conn)

	return c, nil
}
