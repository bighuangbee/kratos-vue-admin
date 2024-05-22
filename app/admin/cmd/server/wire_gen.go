// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/biz"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/conf"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/data"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/pkg/oss"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/server"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, casbin *conf.Casbin, confOss *conf.Oss, logger log.Logger, data_Redis *conf.Data_Redis) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	universalClient := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(db, logger, universalClient)
	if err != nil {
		return nil, nil, err
	}
	casbinRuleRepo := data.NewCasbinRuleRepo(dataData, db, casbin, logger)
	sysUserRepo := data.NewSysUserRepo(dataData, logger)
	ossRepo := oss.NewOssRepo(confOss, logger)
	sysUserUseCase := biz.NewSysUserUseCase(sysUserRepo, ossRepo, confServer, logger)
	sysRoleRepo := data.NewSysRoleRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(auth, sysUserRepo, sysRoleRepo, logger)
	sysRoleMenuRepo := data.NewSysRoleMenuRepo(dataData, logger)
	sysRoleMenuUseCase := biz.NewSysRoleMenuUseCase(sysRoleMenuRepo, logger)
	casbinRuleUseCase := biz.NewCasbinRuleUseCase(casbinRuleRepo, logger)
	transaction := data.NewTransaction(dataData)
	sysMenuRepo := data.NewSysMenuRepo(dataData, logger)
	sysRoleUseCase := biz.NewSysRoleUseCase(sysRoleRepo, logger, sysRoleMenuUseCase, casbinRuleUseCase, sysUserUseCase, transaction, sysMenuRepo)
	sysPostRepo := data.NewSysPostRepo(dataData, logger)
	sysPostUseCase := biz.NewSysPostUseCase(sysPostRepo, logger, sysUserUseCase)
	sysDeptRepo := data.NewSysDeptRepo(dataData, logger)
	sysDeptUseCase := biz.NewSysDeptUseCase(sysDeptRepo, logger)
	sysuserService := service.NewSysuserService(confServer, sysUserUseCase, authUseCase, sysRoleUseCase, sysRoleMenuUseCase, sysPostUseCase, sysDeptUseCase, logger)
	sysApiRepo := data.NewSysApiRepo(dataData, logger)
	sysApiUseCase := biz.NewSysApiUseCase(sysApiRepo, casbinRuleRepo, logger)
	apiService := service.NewApiService(sysApiUseCase, logger, casbinRuleUseCase)
	deptService := service.NewDeptService(sysDeptUseCase, logger)
	sysMenuUseCase := biz.NewSysMenusUseCase(sysMenuRepo, logger)
	menusService := service.NewMenusService(sysMenuUseCase, sysRoleMenuUseCase, logger)
	postService := service.NewPostService(sysPostUseCase, logger)
	sysDictTypeRepo := data.NewSysDictTypeRepo(dataData, logger)
	sysDictTypeUseCase := biz.NewSysDictTypeUseCase(sysDictTypeRepo, logger)
	dictTypeService := service.NewDictTypeService(sysDictTypeUseCase, logger)
	sysDictDatumRepo := data.NewSysDictDataRepo(dataData, logger)
	sysDictDatumUseCase := biz.NewSysDictDatumUseCase(sysDictDatumRepo, logger)
	dictDataService := service.NewDictDataService(sysDictDatumUseCase, logger)
	rolesService := service.NewRolesService(sysRoleUseCase, logger, casbinRuleUseCase)
	sysLogRepo := data.NewSysLogRepo(dataData)
	sysLogCase := biz.NewSysLogCase(sysLogRepo)
	logService := service.NewLogService(sysLogCase)
	httpServer := server.NewHTTPServer(confServer, auth, casbinRuleRepo, logger, sysuserService, apiService, deptService, menusService, postService, dictTypeService, dictDataService, rolesService, logService)
	grpcServer := server.NewGRPCServer(confServer, sysuserService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
