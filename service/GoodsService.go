package service

import (
	"context"
	"errors"
	"grpc-demo/global"
	model "grpc-demo/model/system"
	"grpc-demo/proto"
	"log"
)

type goodsService struct {
}

func NewGoodsServices() *goodsService {
	return &goodsService{}
}

func (g *goodsService) SaveGoods(ctx context.Context, request *proto.Goods) (*proto.GoodsId, error) {
	price := request.GetPrice()
	goodsName := request.GetGoodsName()
	if err := checkValid(request, false); err != nil {
		return nil, err
	}
	goods := &model.Goods{
		GoodsName: goodsName,
		Price:     price,
	}
	err := global.DB.Create(goods).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("add goods failed")
	}
	return &proto.GoodsId{
		Id: uint64(goods.ID),
	}, nil
}

func (g *goodsService) ModifyGoods(ctx context.Context, request *proto.Goods) (*proto.GoodsId, error) {
	price := request.GetPrice()
	goodsName := request.GetGoodsName()
	id := request.GetId()
	if err := checkValid(request, true); err != nil {
		return nil, err
	}
	err := global.DB.Model(&model.Goods{}).Where("id = ?", id).Updates(model.Goods{GoodsName: goodsName, Price: price}).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("modify goods failed")
	}
	return &proto.GoodsId{
		Id: id,
	}, nil
}

func (g *goodsService) DeleteGoods(ctx context.Context, request *proto.GoodsIds) (*proto.DelGoodsResponse, error) {
	ids := request.GetIds()
	err := global.DB.Delete(&model.Goods{}, ids).Error
	if len(ids) == 0 {
		return nil, errors.New("id least one")
	}
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("deletes goods failed")
	}
	return &proto.DelGoodsResponse{}, nil
}

func (g *goodsService) GetGoods(ctx context.Context, request *proto.GetGoodsRequest) (*proto.GetGoodsResponse, error) {
	var goodsList []*proto.Goods
	err := global.DB.Model(&model.Goods{}).Find(&goodsList).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("find goods failed")
	}
	return &proto.GetGoodsResponse{
		GoodsList: goodsList,
	}, nil
}

func checkValid(goods *proto.Goods, isUpdate bool) error {
	//price := goods.GetPrice()
	goodsName := goods.GetGoodsName()
	if len(goodsName) == 0 {
		return errors.New("bad goods name")
	}
	if isUpdate && goods.GetId() == 0 {
		return errors.New("bad goods id")
	}
	return nil
}
