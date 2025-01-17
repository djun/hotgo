// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysDictType interface {
		Tree(ctx context.Context) (list []*sysin.DictTypeTree, err error)
		Delete(ctx context.Context, in sysin.DictTypeDeleteInp) (err error)
		Edit(ctx context.Context, in sysin.DictTypeEditInp) (err error)
		TreeSelect(ctx context.Context, in sysin.DictTreeSelectInp) (list []*sysin.DictTypeTree, err error)
	}
	ISysEmsLog interface {
		Delete(ctx context.Context, in sysin.EmsLogDeleteInp) (err error)
		Edit(ctx context.Context, in sysin.EmsLogEditInp) (err error)
		Status(ctx context.Context, in sysin.EmsLogStatusInp) (err error)
		View(ctx context.Context, in sysin.EmsLogViewInp) (res *sysin.EmsLogViewModel, err error)
		List(ctx context.Context, in sysin.EmsLogListInp) (list []*sysin.EmsLogListModel, totalCount int, err error)
		Send(ctx context.Context, in sysin.SendEmsInp) (err error)
		GetTemplate(ctx context.Context, template string, config *model.EmailConfig) (val string, err error)
		AllowSend(ctx context.Context, models *entity.SysEmsLog, config *model.EmailConfig) (err error)
		VerifyCode(ctx context.Context, in sysin.VerifyEmsCodeInp) (err error)
	}
	ISysProvinces interface {
		Tree(ctx context.Context) (list []*sysin.ProvincesTree, err error)
		Delete(ctx context.Context, in sysin.ProvincesDeleteInp) (err error)
		Edit(ctx context.Context, in sysin.ProvincesEditInp) (err error)
		Status(ctx context.Context, in sysin.ProvincesStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.ProvincesMaxSortInp) (res *sysin.ProvincesMaxSortModel, err error)
		View(ctx context.Context, in sysin.ProvincesViewInp) (res *sysin.ProvincesViewModel, err error)
		List(ctx context.Context, in sysin.ProvincesListInp) (list []*sysin.ProvincesListModel, totalCount int, err error)
		ChildrenList(ctx context.Context, in sysin.ProvincesChildrenListInp) (list []*sysin.ProvincesChildrenListModel, totalCount int, err error)
		UniqueId(ctx context.Context, in sysin.ProvincesUniqueIdInp) (res *sysin.ProvincesUniqueIdModel, err error)
		Select(ctx context.Context, in sysin.ProvincesSelectInp) (res *sysin.ProvincesSelectModel, err error)
	}
	ISysSmsLog interface {
		Delete(ctx context.Context, in sysin.SmsLogDeleteInp) (err error)
		Edit(ctx context.Context, in sysin.SmsLogEditInp) (err error)
		Status(ctx context.Context, in sysin.SmsLogStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.SmsLogMaxSortInp) (res *sysin.SmsLogMaxSortModel, err error)
		View(ctx context.Context, in sysin.SmsLogViewInp) (res *sysin.SmsLogViewModel, err error)
		List(ctx context.Context, in sysin.SmsLogListInp) (list []*sysin.SmsLogListModel, totalCount int, err error)
		SendCode(ctx context.Context, in sysin.SendCodeInp) (err error)
		GetTemplate(ctx context.Context, template string, config *model.SmsConfig) (val string, err error)
		AllowSend(ctx context.Context, models *entity.SysSmsLog, config *model.SmsConfig) (err error)
		VerifyCode(ctx context.Context, in sysin.VerifyCodeInp) (err error)
	}
	ISysAddons interface {
		List(ctx context.Context, in sysin.AddonsListInp) (list []*sysin.AddonsListModel, totalCount int, err error)
		Selects(ctx context.Context, in sysin.AddonsSelectsInp) (res *sysin.AddonsSelectsModel, err error)
		Build(ctx context.Context, in sysin.AddonsBuildInp) (err error)
		Install(ctx context.Context, in sysin.AddonsInstallInp) (err error)
		Upgrade(ctx context.Context, in sysin.AddonsUpgradeInp) (err error)
		UnInstall(ctx context.Context, in sysin.AddonsUnInstallInp) (err error)
	}
	ISysAttachment interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		Delete(ctx context.Context, in sysin.AttachmentDeleteInp) (err error)
		View(ctx context.Context, in sysin.AttachmentViewInp) (res *sysin.AttachmentViewModel, err error)
		List(ctx context.Context, in sysin.AttachmentListInp) (list []*sysin.AttachmentListModel, totalCount int, err error)
	}
	ISysBlacklist interface {
		Delete(ctx context.Context, in sysin.BlacklistDeleteInp) (err error)
		Edit(ctx context.Context, in sysin.BlacklistEditInp) (err error)
		Status(ctx context.Context, in sysin.BlacklistStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.BlacklistMaxSortInp) (res *sysin.BlacklistMaxSortModel, err error)
		View(ctx context.Context, in sysin.BlacklistViewInp) (res *sysin.BlacklistViewModel, err error)
		List(ctx context.Context, in sysin.BlacklistListInp) (list []*sysin.BlacklistListModel, totalCount int, err error)
		VariableLoad(ctx context.Context, err error)
		Load(ctx context.Context)
	}
	ISysCurdDemo interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in sysin.CurdDemoListInp) (list []*sysin.CurdDemoListModel, totalCount int, err error)
		Export(ctx context.Context, in sysin.CurdDemoListInp) (err error)
		Edit(ctx context.Context, in sysin.CurdDemoEditInp) (err error)
		Delete(ctx context.Context, in sysin.CurdDemoDeleteInp) (err error)
		MaxSort(ctx context.Context, in sysin.CurdDemoMaxSortInp) (res *sysin.CurdDemoMaxSortModel, err error)
		View(ctx context.Context, in sysin.CurdDemoViewInp) (res *sysin.CurdDemoViewModel, err error)
		Status(ctx context.Context, in sysin.CurdDemoStatusInp) (err error)
		Switch(ctx context.Context, in sysin.CurdDemoSwitchInp) (err error)
	}
	ISysLoginLog interface {
		Model(ctx context.Context) *gdb.Model
		List(ctx context.Context, in sysin.LoginLogListInp) (list []*sysin.LoginLogListModel, totalCount int, err error)
		Export(ctx context.Context, in sysin.LoginLogListInp) (err error)
		Delete(ctx context.Context, in sysin.LoginLogDeleteInp) (err error)
		View(ctx context.Context, in sysin.LoginLogViewInp) (res *sysin.LoginLogViewModel, err error)
		Push(ctx context.Context, in sysin.LoginLogPushInp)
		RealWrite(ctx context.Context, models entity.SysLoginLog) (err error)
	}
	ISysAddonsConfig interface {
		GetConfigByGroup(ctx context.Context, in sysin.GetAddonsConfigInp) (res *sysin.GetAddonsConfigModel, err error)
		ConversionType(ctx context.Context, models *entity.SysAddonsConfig) (value interface{}, err error)
		UpdateConfigByGroup(ctx context.Context, in sysin.UpdateAddonsConfigInp) (err error)
	}
	ISysServeLog interface {
		Model(ctx context.Context) *gdb.Model
		List(ctx context.Context, in sysin.ServeLogListInp) (list []*sysin.ServeLogListModel, totalCount int, err error)
		Export(ctx context.Context, in sysin.ServeLogListInp) (err error)
		Delete(ctx context.Context, in sysin.ServeLogDeleteInp) (err error)
		View(ctx context.Context, in sysin.ServeLogViewInp) (res *sysin.ServeLogViewModel, err error)
		RealWrite(ctx context.Context, models entity.SysServeLog) (err error)
	}
	ISysLog interface {
		Export(ctx context.Context, in sysin.LogListInp) (err error)
		RealWrite(ctx context.Context, log entity.SysLog) (err error)
		AutoLog(ctx context.Context) error
		AnalysisLog(ctx context.Context) entity.SysLog
		View(ctx context.Context, in sysin.LogViewInp) (res *sysin.LogViewModel, err error)
		Delete(ctx context.Context, in sysin.LogDeleteInp) (err error)
		List(ctx context.Context, in sysin.LogListInp) (list []*sysin.LogListModel, totalCount int, err error)
	}
	ISysConfig interface {
		InitConfig(ctx context.Context)
		GetLogin(ctx context.Context) (conf *model.LoginConfig, err error)
		GetWechat(ctx context.Context) (conf *model.WechatConfig, err error)
		GetPay(ctx context.Context) (conf *model.PayConfig, err error)
		GetSms(ctx context.Context) (conf *model.SmsConfig, err error)
		GetGeo(ctx context.Context) (conf *model.GeoConfig, err error)
		GetUpload(ctx context.Context) (conf *model.UploadConfig, err error)
		GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error)
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		GetLoadTCP(ctx context.Context) (conf *model.TCPConfig, err error)
		GetLoadCache(ctx context.Context) (conf *model.CacheConfig, err error)
		GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error)
		GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error)
		GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error)
		GetLoadServeLog(ctx context.Context) (conf *model.ServeLogConfig, err error)
		GetConfigByGroup(ctx context.Context, in sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error)
		UpdateConfigByGroup(ctx context.Context, in sysin.UpdateConfigInp) (err error)
	}
	ISysCron interface {
		StartCron(ctx context.Context)
		Delete(ctx context.Context, in sysin.CronDeleteInp) (err error)
		Edit(ctx context.Context, in sysin.CronEditInp) (err error)
		Status(ctx context.Context, in sysin.CronStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.CronMaxSortInp) (res *sysin.CronMaxSortModel, err error)
		View(ctx context.Context, in sysin.CronViewInp) (res *sysin.CronViewModel, err error)
		List(ctx context.Context, in sysin.CronListInp) (list []*sysin.CronListModel, totalCount int, err error)
		OnlineExec(ctx context.Context, in sysin.OnlineExecInp) (err error)
	}
	ISysCronGroup interface {
		Delete(ctx context.Context, in sysin.CronGroupDeleteInp) (err error)
		Edit(ctx context.Context, in sysin.CronGroupEditInp) (err error)
		Status(ctx context.Context, in sysin.CronGroupStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.CronGroupMaxSortInp) (res *sysin.CronGroupMaxSortModel, err error)
		View(ctx context.Context, in sysin.CronGroupViewInp) (res *sysin.CronGroupViewModel, err error)
		List(ctx context.Context, in sysin.CronGroupListInp) (list []*sysin.CronGroupListModel, totalCount int, err error)
		Select(ctx context.Context, in sysin.CronGroupSelectInp) (res *sysin.CronGroupSelectModel, err error)
	}
	ISysDictData interface {
		Delete(ctx context.Context, in sysin.DictDataDeleteInp) error
		Edit(ctx context.Context, in sysin.DictDataEditInp) (err error)
		List(ctx context.Context, in sysin.DictDataListInp) (list []*sysin.DictDataListModel, totalCount int, err error)
		Select(ctx context.Context, in sysin.DataSelectInp) (list sysin.DataSelectModel, err error)
	}
	ISysGenCodes interface {
		Delete(ctx context.Context, in sysin.GenCodesDeleteInp) (err error)
		Edit(ctx context.Context, in sysin.GenCodesEditInp) (res *sysin.GenCodesEditModel, err error)
		Status(ctx context.Context, in sysin.GenCodesStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.GenCodesMaxSortInp) (res *sysin.GenCodesMaxSortModel, err error)
		View(ctx context.Context, in sysin.GenCodesViewInp) (res *sysin.GenCodesViewModel, err error)
		List(ctx context.Context, in sysin.GenCodesListInp) (list []*sysin.GenCodesListModel, totalCount int, err error)
		Selects(ctx context.Context, in sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error)
		TableSelect(ctx context.Context, in sysin.GenCodesTableSelectInp) (res []*sysin.GenCodesTableSelectModel, err error)
		ColumnSelect(ctx context.Context, in sysin.GenCodesColumnSelectInp) (res []*sysin.GenCodesColumnSelectModel, err error)
		ColumnList(ctx context.Context, in sysin.GenCodesColumnListInp) (res []*sysin.GenCodesColumnListModel, err error)
		Preview(ctx context.Context, in sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error)
		Build(ctx context.Context, in sysin.GenCodesBuildInp) (err error)
	}
)

var (
	localSysAddonsConfig ISysAddonsConfig
	localSysServeLog     ISysServeLog
	localSysLog          ISysLog
	localSysConfig       ISysConfig
	localSysCron         ISysCron
	localSysCronGroup    ISysCronGroup
	localSysDictData     ISysDictData
	localSysGenCodes     ISysGenCodes
	localSysDictType     ISysDictType
	localSysEmsLog       ISysEmsLog
	localSysProvinces    ISysProvinces
	localSysSmsLog       ISysSmsLog
	localSysAddons       ISysAddons
	localSysAttachment   ISysAttachment
	localSysBlacklist    ISysBlacklist
	localSysCurdDemo     ISysCurdDemo
	localSysLoginLog     ISysLoginLog
)

func SysBlacklist() ISysBlacklist {
	if localSysBlacklist == nil {
		panic("implement not found for interface ISysBlacklist, forgot register?")
	}
	return localSysBlacklist
}

func RegisterSysBlacklist(i ISysBlacklist) {
	localSysBlacklist = i
}

func SysCurdDemo() ISysCurdDemo {
	if localSysCurdDemo == nil {
		panic("implement not found for interface ISysCurdDemo, forgot register?")
	}
	return localSysCurdDemo
}

func RegisterSysCurdDemo(i ISysCurdDemo) {
	localSysCurdDemo = i
}

func SysLoginLog() ISysLoginLog {
	if localSysLoginLog == nil {
		panic("implement not found for interface ISysLoginLog, forgot register?")
	}
	return localSysLoginLog
}

func RegisterSysLoginLog(i ISysLoginLog) {
	localSysLoginLog = i
}

func SysProvinces() ISysProvinces {
	if localSysProvinces == nil {
		panic("implement not found for interface ISysProvinces, forgot register?")
	}
	return localSysProvinces
}

func RegisterSysProvinces(i ISysProvinces) {
	localSysProvinces = i
}

func SysSmsLog() ISysSmsLog {
	if localSysSmsLog == nil {
		panic("implement not found for interface ISysSmsLog, forgot register?")
	}
	return localSysSmsLog
}

func RegisterSysSmsLog(i ISysSmsLog) {
	localSysSmsLog = i
}

func SysAddons() ISysAddons {
	if localSysAddons == nil {
		panic("implement not found for interface ISysAddons, forgot register?")
	}
	return localSysAddons
}

func RegisterSysAddons(i ISysAddons) {
	localSysAddons = i
}

func SysAttachment() ISysAttachment {
	if localSysAttachment == nil {
		panic("implement not found for interface ISysAttachment, forgot register?")
	}
	return localSysAttachment
}

func RegisterSysAttachment(i ISysAttachment) {
	localSysAttachment = i
}

func SysAddonsConfig() ISysAddonsConfig {
	if localSysAddonsConfig == nil {
		panic("implement not found for interface ISysAddonsConfig, forgot register?")
	}
	return localSysAddonsConfig
}

func RegisterSysAddonsConfig(i ISysAddonsConfig) {
	localSysAddonsConfig = i
}

func SysServeLog() ISysServeLog {
	if localSysServeLog == nil {
		panic("implement not found for interface ISysServeLog, forgot register?")
	}
	return localSysServeLog
}

func RegisterSysServeLog(i ISysServeLog) {
	localSysServeLog = i
}

func SysCronGroup() ISysCronGroup {
	if localSysCronGroup == nil {
		panic("implement not found for interface ISysCronGroup, forgot register?")
	}
	return localSysCronGroup
}

func RegisterSysCronGroup(i ISysCronGroup) {
	localSysCronGroup = i
}

func SysDictData() ISysDictData {
	if localSysDictData == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return localSysDictData
}

func RegisterSysDictData(i ISysDictData) {
	localSysDictData = i
}

func SysGenCodes() ISysGenCodes {
	if localSysGenCodes == nil {
		panic("implement not found for interface ISysGenCodes, forgot register?")
	}
	return localSysGenCodes
}

func RegisterSysGenCodes(i ISysGenCodes) {
	localSysGenCodes = i
}

func SysLog() ISysLog {
	if localSysLog == nil {
		panic("implement not found for interface ISysLog, forgot register?")
	}
	return localSysLog
}

func RegisterSysLog(i ISysLog) {
	localSysLog = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysCron() ISysCron {
	if localSysCron == nil {
		panic("implement not found for interface ISysCron, forgot register?")
	}
	return localSysCron
}

func RegisterSysCron(i ISysCron) {
	localSysCron = i
}

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}

func SysEmsLog() ISysEmsLog {
	if localSysEmsLog == nil {
		panic("implement not found for interface ISysEmsLog, forgot register?")
	}
	return localSysEmsLog
}

func RegisterSysEmsLog(i ISysEmsLog) {
	localSysEmsLog = i
}
