package service

import (
	"context"
	pb "github.com/byteflowteam/kratos-vue-admin/api/admin/v1"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/biz"
	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/data/dal/model"
	"github.com/byteflowteam/kratos-vue-admin/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type ApiService struct {
	pb.UnimplementedApiServer
	apiUseCase    *biz.SysApiUseCase
	casbinUseCase *biz.CasbinRuleUseCase
	log           *log.Helper
	logger        log.Logger
	sysApis       map[string]*pb.ApiData
}

func NewApiService(ac *biz.SysApiUseCase, logger log.Logger, casbinUseCase *biz.CasbinRuleUseCase) *ApiService {
	apiService := &ApiService{
		apiUseCase:    ac,
		log:           log.NewHelper(log.With(logger, "module", "service/api")),
		casbinUseCase: casbinUseCase,
		logger:        logger,
		sysApis:       make(map[string]*pb.ApiData, 0),
	}

	apiService.ReloadSysApi()
	return apiService
}

func (a *ApiService) GetApi(ctx context.Context, req *pb.GetApiRequest) (*pb.GetApiReply, error) {
	api, err := a.apiUseCase.GetApiByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	data := &pb.ApiData{
		Id:          int32(api.ID),
		Path:        api.Path,
		Description: api.Description,
		ApiGroup:    api.APIGroup,
		Method:      api.Method,
		CreateTime:  util.NewTimestamp(api.CreatedAt),
		UpdateTime:  util.NewTimestamp(api.UpdatedAt),
	}
	return &pb.GetApiReply{
		Api: data,
	}, nil
}

func (a *ApiService) ListApi(ctx context.Context, req *pb.ListApiRequest) (*pb.ListApiReply, error) {
	apiList, total, err := a.apiUseCase.ListPage(ctx, req.PageNum, req.PageSize, req.Path, req.Description, req.ApiGroup, req.Method)
	if err != nil {
		return nil, err
	}
	data := ConvertApiDataFromApiList(apiList)
	return &pb.ListApiReply{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    total,
		Data:     data,
	}, nil
}
func (a *ApiService) CreateApi(ctx context.Context, req *pb.CreateApiRequest) (*pb.CreateApiReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := a.apiUseCase.CreateApi(ctx, &model.SysAPI{
		Path:        req.Path,
		Description: req.Description,
		APIGroup:    req.ApiGroup,
		Method:      req.Method,
	})

	a.ReloadSysApi()

	return &pb.CreateApiReply{}, err
}
func (a *ApiService) UpdateApi(ctx context.Context, req *pb.UpdateApiRequest) (*pb.UpdateApiReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := a.apiUseCase.UpdateApi(ctx, &model.SysAPI{
		ID:          req.Id,
		Path:        req.Path,
		Description: req.Description,
		APIGroup:    req.ApiGroup,
		Method:      req.Method,
	})

	a.ReloadSysApi()

	return &pb.UpdateApiReply{}, err
}
func (a *ApiService) DeleteApi(ctx context.Context, req *pb.DeleteApiRequest) (*pb.DeleteApiReply, error) {
	err := a.apiUseCase.DeleteApi(ctx, req.Id)

	a.ReloadSysApi()

	return &pb.DeleteApiReply{}, err
}

func (a *ApiService) AllApi(ctx context.Context, _ *pb.AllApiRequest) (*pb.AllApiReply, error) {
	apiList, err := a.apiUseCase.AllApi(ctx)
	if err != nil {
		return nil, err
	}

	data := ConvertApiDataFromApiList(apiList)

	return &pb.AllApiReply{
		Data: data,
	}, nil
}

func (a *ApiService) GetPolicyPathByRoleKey(ctx context.Context, req *pb.GetPolicyPathByRoleKeyRequest) (*pb.GetPolicyPathByRoleKeyReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	policyList := a.casbinUseCase.GetPolicyPathByRoleId(req.RoleKey)

	apis := ConvertApiBaseFromList(policyList)

	return &pb.GetPolicyPathByRoleKeyReply{
		Apis: apis,
	}, nil
}

func (a *ApiService) ReloadSysApi() {
	reply, err := a.AllApi(context.Background(), &pb.AllApiRequest{})
	if err != nil {
		log.Error("ReloadSysApi", err)
		return
	}

	for _, item := range reply.Data {
		a.sysApis[item.Path] = item
	}
}

func (a *ApiService) GetSysApi(path string) *pb.ApiData {
	return a.sysApis[path]
}
