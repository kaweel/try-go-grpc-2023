package menu_test

import (
	"context"
	"errors"
	"testing"
	"try-go-grpc/menu"
	"try-go-grpc/menu/mocks"
	pb "try-go-grpc/menu/proto"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestListBeers(t *testing.T) {
	var (
		ctx    context.Context
		repo   *mocks.MenuRepository
		server menu.MenuServer

		listBeersResp []menu.Beer
		listBeersErr  error
	)
	beforeEach := func(t *testing.T) {
		ctx = context.TODO()
		repo = &mocks.MenuRepository{}
		server = menu.NewMenuServer(repo)
		listBeersResp = []menu.Beer{{
			Id:   1,
			Name: "Change",
		}, {
			Id:   2,
			Name: "Singha",
		}}

		repo.On("ListBeers").Return(
			func() []menu.Beer {
				return listBeersResp
			},
			func() error {
				return listBeersErr
			},
		)
	}

	t.Run("should return list of beer when stuff in stock", func(t *testing.T) {
		beforeEach(t)

		actual, _ := server.ListBeers(ctx, &emptypb.Empty{})

		expect := pb.ListBeersResponse{
			Beers: []*pb.Beer{{Id: int32(1), Name: "Change"}, {Id: int32(2), Name: "Singha"}},
		}
		assert.ElementsMatch(t, expect.Beers, actual.Beers)
	})

	t.Run("should return error 'out of stock' when stuff out of stock", func(t *testing.T) {
		beforeEach(t)
		listBeersErr = errors.New("out of stock")

		_, actual := server.ListBeers(ctx, &emptypb.Empty{})

		assert.EqualError(t, actual, "out of stock")
	})

}
