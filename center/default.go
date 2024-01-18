package center

import (
	"github.com/gin-gonic/gin"
	"github.com/grafana/pyroscope-go"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/pkg/security"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var Default = &DefaultCenter{
	Manager: server.New(),
}

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2024/1/8 09:54:13
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2024/1/8 09:54:13
 */

type DefaultCenter struct {
	NoticeImp
	TenantImp
	UserImp
	VirtualModelImp
	ConfigImp
	server.Manager
	gin.IRouter
	StageImp
	AppConfigImp
	Profiler *pyroscope.Profiler
	StatisticsImp
}

func (d *DefaultCenter) SetNotice(n NoticeImp) {
	d.NoticeImp = n
}

func (d *DefaultCenter) SetTenant(t TenantImp) {
	d.TenantImp = t
}

func (d *DefaultCenter) SetVerify(v UserImp) {
	d.UserImp = v
}

func (d *DefaultCenter) SetConfig(e ConfigImp) {
	d.ConfigImp = e
}

func (d *DefaultCenter) SetVirtualModel(v VirtualModelImp) {
	d.VirtualModelImp = v
}

func (d *DefaultCenter) SetServerManager(m server.Manager) {
	d.Manager = m
}

func (d *DefaultCenter) SetRouter(r gin.IRouter) {
	d.IRouter = r
}

func (d *DefaultCenter) SetAppConfig(a AppConfigImp) {
	d.AppConfigImp = a
}

func (d *DefaultCenter) SetProfiler(p *pyroscope.Profiler) {
	d.Profiler = p
}

func (d *DefaultCenter) SetStatistics(s StatisticsImp) {
	d.StatisticsImp = s
}

func (d *DefaultCenter) GetNotice() NoticeImp {
	return d.NoticeImp
}

func (d *DefaultCenter) GetTenant() TenantImp {
	return d.TenantImp
}

func (d *DefaultCenter) GetVerify() UserImp {
	return d.UserImp
}

func (d *DefaultCenter) GetConfig() ConfigImp {
	return d.ConfigImp
}

func (d *DefaultCenter) GetVirtualModel() VirtualModelImp {
	return d.VirtualModelImp
}

func (d *DefaultCenter) GetServerManager() server.Manager {
	return d.Manager
}

func (d *DefaultCenter) GetRouter() gin.IRouter {
	return d.IRouter
}

func (d *DefaultCenter) GetAppConfig() AppConfigImp {
	return d.AppConfigImp
}

func (d *DefaultCenter) GetProfiler() *pyroscope.Profiler {
	return d.Profiler
}

func (d *DefaultCenter) GetStatistics() StatisticsImp {
	return d.StatisticsImp
}

func (d *DefaultCenter) Stage() string {
	stage := os.Getenv("STAGE")
	if stage == "" {
		stage = os.Getenv("stage")
	}
	if stage == "" {
		stage = "local"
	}
	return stage
}

func GetDB(ctx *gin.Context, table schema.Tabler) *gorm.DB {
	return Default.GetDB(ctx, table)
}

func SetNotice(n NoticeImp) *DefaultCenter {
	Default.SetNotice(n)
	return Default
}

func SetTenant(t TenantImp) *DefaultCenter {
	Default.SetTenant(t)
	return Default
}

func SetVerify(v security.Verifier) *DefaultCenter {
	Default.SetVerify(v)
	return Default
}

func SetConfig(e ConfigImp) *DefaultCenter {
	Default.SetConfig(e)
	return Default
}

func SetVirtualModel(v VirtualModelImp) *DefaultCenter {
	Default.SetVirtualModel(v)
	return Default
}

func SetServerManager(m server.Manager) *DefaultCenter {
	Default.SetServerManager(m)
	return Default
}

func SetAppConfig(a AppConfigImp) *DefaultCenter {
	Default.SetAppConfig(a)
	return Default
}

func SetRouter(r gin.IRouter) *DefaultCenter {
	Default.SetRouter(r)
	return Default
}

func SetProfiler(p *pyroscope.Profiler) *DefaultCenter {
	Default.SetProfiler(p)
	return Default
}

func SetStatistics(s StatisticsImp) *DefaultCenter {
	Default.SetStatistics(s)
	return Default
}

func GetNotice() NoticeImp {
	return Default.GetNotice()
}

func GetTenant() TenantImp {
	return Default.GetTenant()
}

func GetUser() UserImp {
	return Default.GetVerify()
}

func GetConfig() ConfigImp {
	return Default.GetConfig()
}

func GetVirtualModel() VirtualModelImp {
	return Default.GetVirtualModel()
}

func GetServerManager() server.Manager {
	return Default.GetServerManager()
}

func GetRouter() gin.IRouter {
	return Default.GetRouter()
}

func Stage() string {
	return Default.Stage()
}

func GetAppConfig() AppConfigImp {
	return Default.GetAppConfig()
}

func GetProfiler() *pyroscope.Profiler {
	return Default.GetProfiler()
}

func GetStatistics() StatisticsImp {
	return Default.GetStatistics()
}