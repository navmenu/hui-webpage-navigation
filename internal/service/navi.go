package service

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/biz"
	"hui-webpage-navigation/internal/models"
	"hui-webpage-navigation/internal/utils/utils_kratos/utils_kratos"
	"unicode/utf8"

	"github.com/barweiss/go-tuple"
	"github.com/go-kratos/kratos/v2/log"
)

type NaviService struct {
	pb.UnimplementedNaviServer
	uc   *biz.NaviUsecase
	xlog *log.Helper
}

func NewNaviService(uc *biz.NaviUsecase, logger log.Logger) *NaviService {
	return &NaviService{
		uc:   uc,
		xlog: log.NewHelper(logger),
	}
}

func (s *NaviService) CreateNavi(ctx context.Context, req *pb.CreateNaviRequest) (*pb.CreateNaviReply, error) {
	if req.Name == "" {
		return nil, pb.ErrorMissingParam("navi name is missing")
	}
	if utf8.RuneCountInString(req.Name) > 100 {
		return nil, pb.ErrorBadParam("navi name is too long")
	}
	res, erk := s.uc.CreateNavi(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.CreateNaviReply{
		Name: res.Name,
	}, nil
}
func (s *NaviService) DeleteNavi(ctx context.Context, req *pb.DeleteNaviRequest) (*pb.DeleteNaviReply, error) {
	if req.Name == "" {
		return nil, pb.ErrorMissingParam("navi name is missing")
	}
	if utf8.RuneCountInString(req.Name) > 100 {
		return nil, pb.ErrorBadParam("navi name is too long")
	}
	erk := s.uc.DeleteNavi(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.DeleteNaviReply{}, nil
}
func (s *NaviService) SortNavi(ctx context.Context, req *pb.SortNaviRequest) (*pb.SortNaviReply, error) {
	if len(req.Items) == 0 {
		return nil, pb.ErrorMissingParam("navi items is missing")
	}
	for _, item := range req.Items {
		if utf8.RuneCountInString(item.Name) > 100 {
			return nil, pb.ErrorBadParam("navi name is too long")
		}
		if item.Sort < 0 {
			return nil, pb.ErrorBadParam("navi sort is negative")
		}
	}
	erk := s.uc.SortNavi(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.SortNaviReply{}, nil
}
func (s *NaviService) GetNaviOrders(ctx context.Context, req *pb.GetNaviOrdersRequest) (*pb.GetNaviOrdersReply, error) {
	navis, erk := s.uc.GetNaviOrders(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.GetNaviOrdersReply{
		Navis: navis2pb(navis),
	}, nil
}
func (s *NaviService) ListNavi(ctx context.Context, req *pb.ListNaviRequest) (*pb.ListNaviReply, error) {
	res, erk := s.uc.ListNavi(ctx, req)
	if erk != nil {
		return nil, erk
	}
	//这里还是平铺结构的
	slice := naviTuple2pbs(res)
	//转换成带层级结构的
	items := navisToLayers(slice)
	return &pb.ListNaviReply{Items: items}, nil
}

func naviTuple2pbs(a []*tuple.T2[*models.Navi, []*models.NaviLvl2]) []*pb.NaviListItemType {
	var res = make([]*pb.NaviListItemType, 0, len(a))
	for _, x := range a {
		res = append(res, naviTuple2pb(x))
	}
	return res
}

func naviTuple2pb(x *tuple.T2[*models.Navi, []*models.NaviLvl2]) *pb.NaviListItemType {
	return &pb.NaviListItemType{
		Navi:      naviTv1x2pb(x),
		NaviLvl2S: naviLvl2resp2pbs(x.V2),
	}
}

func naviTv1x2pb(x *tuple.T2[*models.Navi, []*models.NaviLvl2]) *pb.NaviType {
	return navi2pb(x.V1)
}

func navi2pb(navi *models.Navi) *pb.NaviType {
	return &pb.NaviType{
		Id:         navi.ID,
		Name:       navi.Name,
		Sort:       navi.Sort,
		ParentNvid: navi.ParentNvid,
	}
}

func navis2pb(navis []*models.Navi) []*pb.NaviType {
	var res = make([]*pb.NaviType, 0, len(navis))
	for _, a := range navis {
		res = append(res, navi2pb(a))
	}
	return res
}

func naviLvl2resp2pbs(a []*models.NaviLvl2) []*pb.NaviLvl2Item {
	var res = make([]*pb.NaviLvl2Item, 0, len(a))
	for _, x := range a {
		res = append(res, naviLvl2item2pb(x))
	}
	return res
}

func naviLvl2item2pb(x *models.NaviLvl2) *pb.NaviLvl2Item {
	return &pb.NaviLvl2Item{
		NaviName: x.NaviName,
		Name:     x.Name,
		Text:     x.Text,
		Link:     x.Link,
		IsEscrow: x.IsEscrow,
		Sort:     x.Sort,
	}
}

func navisToLayers(slice []*pb.NaviListItemType) []*pb.NaviListItemType {
	var mp = make(map[uint64][]*pb.NaviListItemType)
	for _, v := range slice {
		if k := v.Navi.ParentNvid; k != 0 {
			mp[k] = append(mp[k], v)
		}
	}
	for _, v := range slice {
		v.NextNavis = mp[v.Navi.Id]
	}
	//这里最好是像这样，三个循环，第一个组成map，第二个收集结果，第三个是收集树根，否的会出错
	var resTopLayers []*pb.NaviListItemType
	for _, v := range slice {
		if v.Navi.ParentNvid == 0 {
			v.NextNavis = mp[v.Navi.Id]
			resTopLayers = append(resTopLayers, v)
		}
	}
	return resTopLayers
}

func (s *NaviService) GetGuestSettings(ctx context.Context, req *pb.GetGuestSettingsRequest) (*pb.GetGuestSettingsReply, error) {
	ip, err := utils_kratos.GetIp(ctx)
	if err != nil {
		return nil, pb.ErrorBadParam("no ip error=%v", err)
	}
	cookie, err := utils_kratos.GetCookie(ctx)
	if err != nil {
		return nil, pb.ErrorBadParam("no cookie error=%v", err)
	}
	param := &biz.GuestSettingPK{
		Ip:     ip,
		Cookie: cookie,
	}
	res, erk := s.uc.GetGuestSettings(ctx, param)
	if erk != nil {
		if erk.Reason != pb.ErrorReason_NOT_EXIST.String() {
			return nil, erk
		}
	}
	resp := &pb.GetGuestSettingsReply{}
	if res != nil {
		resp.IsNotRemindInfo = res.IsNotRemindInfo
	}
	return resp, nil
}
func (s *NaviService) SetNotRemindInfo(ctx context.Context, req *pb.SetNotRemindInfoRequest) (*pb.SetNotRemindInfoReply, error) {
	ip, err := utils_kratos.GetIp(ctx)
	if err != nil {
		return nil, pb.ErrorBadParam("no ip error=%v", err)
	}
	cookie, err := utils_kratos.GetCookie(ctx)
	if err != nil {
		return nil, pb.ErrorBadParam("no cookie error=%v", err)
	}
	param := &biz.GuestSettingPK{
		Ip:     ip,
		Cookie: cookie,
	}
	erk := s.uc.SetNotRemindInfo(ctx, param)
	if erk != nil {
		return nil, erk
	}
	return &pb.SetNotRemindInfoReply{}, nil
}
