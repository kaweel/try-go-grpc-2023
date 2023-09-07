package menu

import (
	"context"
	pb "try-go-grpc/menu/proto"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type MenuServer interface {
	ListBeers(ctx context.Context, request *emptypb.Empty) (*pb.ListBeersResponse, error)
}

type Menu struct {
	pb.UnimplementedMenuServer
	repo MenuRepository
}

func NewMenuServer(repo MenuRepository) pb.MenuServer {
	menu := &Menu{
		repo: repo,
	}
	return menu
}

func (menu *Menu) ListBeers(ctx context.Context, request *emptypb.Empty) (*pb.ListBeersResponse, error) {
	beers, err := menu.repo.ListBeers()
	if err != nil {
		return nil, err
	}
	pbBeers := []*pb.Beer{}
	for _, beer := range beers {
		pbBeer := &pb.Beer{
			Id:   beer.Id,
			Name: beer.Name,
		}
		pbBeers = append(pbBeers, pbBeer)
	}
	return &pb.ListBeersResponse{
		Beers: pbBeers,
	}, nil
}
