// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/api/admin/role"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/payin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IAdminPost interface {
		Delete(ctx context.Context, in adminin.PostDeleteInp) (err error)
		Edit(ctx context.Context, in adminin.PostEditInp) (err error)
		MaxSort(ctx context.Context, in adminin.PostMaxSortInp) (res *adminin.PostMaxSortModel, err error)
		NameUnique(ctx context.Context, in adminin.PostNameUniqueInp) (res *adminin.PostNameUniqueModel, err error)
		CodeUnique(ctx context.Context, in adminin.PostCodeUniqueInp) (res *adminin.PostCodeUniqueModel, err error)
		View(ctx context.Context, in adminin.PostViewInp) (res *adminin.PostViewModel, err error)
		List(ctx context.Context, in adminin.PostListInp) (list []*adminin.PostListModel, totalCount int, err error)
		GetMemberByStartName(ctx context.Context, memberId int64) (name string, err error)
		Status(ctx context.Context, in adminin.PostStatusInp) (err error)
	}
	IAdminRole interface {
		Verify(ctx context.Context, path, method string) bool
		List(ctx context.Context, in adminin.RoleListInp) (res *adminin.RoleListModel, totalCount int, err error)
		GetName(ctx context.Context, id int64) (name string, err error)
		GetMemberList(ctx context.Context, id int64) (list []*adminin.RoleListModel, err error)
		GetPermissions(ctx context.Context, in adminin.GetPermissionsInp) (res *adminin.GetPermissionsModel, err error)
		UpdatePermissions(ctx context.Context, in adminin.UpdatePermissionsInp) (err error)
		Edit(ctx context.Context, in adminin.RoleEditInp) (err error)
		Delete(ctx context.Context, in adminin.RoleDeleteInp) (err error)
		DataScopeSelect() (res form.Selects)
		DataScopeEdit(ctx context.Context, in *adminin.DataScopeEditInp) (err error)
	}
	IAdminCreditsLog interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		SaveBalance(ctx context.Context, in adminin.CreditsLogSaveBalanceInp) (res *adminin.CreditsLogSaveBalanceModel, err error)
		SaveIntegral(ctx context.Context, in adminin.CreditsLogSaveIntegralInp) (res *adminin.CreditsLogSaveIntegralModel, err error)
		List(ctx context.Context, in adminin.CreditsLogListInp) (list []*adminin.CreditsLogListModel, totalCount int, err error)
		Export(ctx context.Context, in adminin.CreditsLogListInp) (err error)
	}
	IAdminDept interface {
		Delete(ctx context.Context, in adminin.DeptDeleteInp) (err error)
		Edit(ctx context.Context, in adminin.DeptEditInp) (err error)
		Status(ctx context.Context, in adminin.DeptStatusInp) (err error)
		MaxSort(ctx context.Context, in adminin.DeptMaxSortInp) (res *adminin.DeptMaxSortModel, err error)
		View(ctx context.Context, in adminin.DeptViewInp) (res *adminin.DeptViewModel, err error)
		Option(ctx context.Context, in adminin.DeptOptionInp) (res *adminin.DeptOptionModel, totalCount int, err error)
		List(ctx context.Context, in adminin.DeptListInp) (res *adminin.DeptListModel, err error)
		GetName(ctx context.Context, id int64) (name string, err error)
	}
	IAdminMember interface {
		AddBalance(ctx context.Context, in adminin.MemberAddBalanceInp) (err error)
		AddIntegral(ctx context.Context, in adminin.MemberAddIntegralInp) (err error)
		UpdateCash(ctx context.Context, in adminin.MemberUpdateCashInp) (err error)
		UpdateEmail(ctx context.Context, in adminin.MemberUpdateEmailInp) (err error)
		UpdateMobile(ctx context.Context, in adminin.MemberUpdateMobileInp) (err error)
		UpdateProfile(ctx context.Context, in adminin.MemberUpdateProfileInp) (err error)
		UpdatePwd(ctx context.Context, in adminin.MemberUpdatePwdInp) (err error)
		ResetPwd(ctx context.Context, in adminin.MemberResetPwdInp) (err error)
		VerifyUnique(ctx context.Context, in adminin.VerifyUniqueInp) (err error)
		Delete(ctx context.Context, in adminin.MemberDeleteInp) (err error)
		Edit(ctx context.Context, in adminin.MemberEditInp) (err error)
		View(ctx context.Context, in adminin.MemberViewInp) (res *adminin.MemberViewModel, err error)
		List(ctx context.Context, in adminin.MemberListInp) (list []*adminin.MemberListModel, totalCount int, err error)
		Status(ctx context.Context, in adminin.MemberStatusInp) (err error)
		GenTree(ctx context.Context, pid int64) (level int, newTree string, err error)
		LoginMemberInfo(ctx context.Context) (res *adminin.LoginMemberInfoModel, err error)
		MemberLoginStat(ctx context.Context, in adminin.MemberLoginStatInp) (res *adminin.MemberLoginStatModel, err error)
		GetIdByCode(ctx context.Context, in adminin.GetIdByCodeInp) (res *adminin.GetIdByCodeModel, err error)
		Select(ctx context.Context, in adminin.MemberSelectInp) (res []*adminin.MemberSelectModel, err error)
		VerifySuperId(ctx context.Context, verifyId int64) bool
		FilterAuthModel(ctx context.Context, memberId int64) *gdb.Model
	}
	IAdminMemberPost interface {
		UpdatePostIds(ctx context.Context, memberId int64, postIds []int64) (err error)
	}
	IAdminMenu interface {
		Delete(ctx context.Context, in adminin.MenuDeleteInp) (err error)
		VerifyUnique(ctx context.Context, in adminin.VerifyUniqueInp) (err error)
		Edit(ctx context.Context, in adminin.MenuEditInp) (err error)
		List(ctx context.Context, in adminin.MenuListInp) (res *adminin.MenuListModel, err error)
		GetMenuList(ctx context.Context, memberId int64) (res *role.DynamicRes, err error)
		LoginPermissions(ctx context.Context, memberId int64) (lists adminin.MemberLoginPermissions, err error)
	}
	IAdminMonitor interface {
		StartMonitor(ctx context.Context)
		GetMeta(ctx context.Context) *model.MonitorData
	}
	IAdminNotice interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		Delete(ctx context.Context, in adminin.NoticeDeleteInp) (err error)
		Edit(ctx context.Context, in adminin.NoticeEditInp) (err error)
		Status(ctx context.Context, in adminin.NoticeStatusInp) (err error)
		MaxSort(ctx context.Context, in adminin.NoticeMaxSortInp) (res *adminin.NoticeMaxSortModel, err error)
		View(ctx context.Context, in adminin.NoticeViewInp) (res *adminin.NoticeViewModel, err error)
		List(ctx context.Context, in adminin.NoticeListInp) (list []*adminin.NoticeListModel, totalCount int, err error)
		PullMessages(ctx context.Context, in adminin.PullMessagesInp) (res *adminin.PullMessagesModel, err error)
		UnreadCount(ctx context.Context, in adminin.NoticeUnreadCountInp) (res *adminin.NoticeUnreadCountModel, err error)
		UpRead(ctx context.Context, in adminin.NoticeUpReadInp) (err error)
		ReadAll(ctx context.Context, in adminin.NoticeReadAllInp) (err error)
		MessageList(ctx context.Context, in adminin.NoticeMessageListInp) (list []*adminin.NoticeMessageListModel, totalCount int, err error)
	}
	IAdminOrder interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		AcceptRefund(ctx context.Context, in adminin.OrderAcceptRefundInp) (err error)
		ApplyRefund(ctx context.Context, in adminin.OrderApplyRefundInp) (err error)
		PayNotify(ctx context.Context, in payin.NotifyCallFuncInp) (err error)
		Create(ctx context.Context, in adminin.OrderCreateInp) (res *adminin.OrderCreateModel, err error)
		List(ctx context.Context, in adminin.OrderListInp) (list []*adminin.OrderListModel, totalCount int, err error)
		Export(ctx context.Context, in adminin.OrderListInp) (err error)
		Edit(ctx context.Context, in adminin.OrderEditInp) (err error)
		Delete(ctx context.Context, in adminin.OrderDeleteInp) (err error)
		View(ctx context.Context, in adminin.OrderViewInp) (res *adminin.OrderViewModel, err error)
		Status(ctx context.Context, in adminin.OrderStatusInp) (err error)
	}
	IAdminCash interface {
		View(ctx context.Context, in adminin.CashViewInp) (res *adminin.CashViewModel, err error)
		List(ctx context.Context, in adminin.CashListInp) (list []*adminin.CashListModel, totalCount int, err error)
		Apply(ctx context.Context, in adminin.CashApplyInp) (err error)
		Payment(ctx context.Context, in adminin.CashPaymentInp) (err error)
	}
	IAdminSite interface {
		Register(ctx context.Context, in adminin.RegisterInp) (err error)
		AccountLogin(ctx context.Context, in adminin.AccountLoginInp) (res *adminin.LoginModel, err error)
		MobileLogin(ctx context.Context, in adminin.MobileLoginInp) (res *adminin.LoginModel, err error)
	}
)

var (
	localAdminDept       IAdminDept
	localAdminMember     IAdminMember
	localAdminMemberPost IAdminMemberPost
	localAdminMenu       IAdminMenu
	localAdminMonitor    IAdminMonitor
	localAdminNotice     IAdminNotice
	localAdminOrder      IAdminOrder
	localAdminCash       IAdminCash
	localAdminSite       IAdminSite
	localAdminPost       IAdminPost
	localAdminRole       IAdminRole
	localAdminCreditsLog IAdminCreditsLog
)

func AdminMonitor() IAdminMonitor {
	if localAdminMonitor == nil {
		panic("implement not found for interface IAdminMonitor, forgot register?")
	}
	return localAdminMonitor
}

func RegisterAdminMonitor(i IAdminMonitor) {
	localAdminMonitor = i
}

func AdminNotice() IAdminNotice {
	if localAdminNotice == nil {
		panic("implement not found for interface IAdminNotice, forgot register?")
	}
	return localAdminNotice
}

func RegisterAdminNotice(i IAdminNotice) {
	localAdminNotice = i
}

func AdminOrder() IAdminOrder {
	if localAdminOrder == nil {
		panic("implement not found for interface IAdminOrder, forgot register?")
	}
	return localAdminOrder
}

func RegisterAdminOrder(i IAdminOrder) {
	localAdminOrder = i
}

func AdminCash() IAdminCash {
	if localAdminCash == nil {
		panic("implement not found for interface IAdminCash, forgot register?")
	}
	return localAdminCash
}

func RegisterAdminCash(i IAdminCash) {
	localAdminCash = i
}

func AdminDept() IAdminDept {
	if localAdminDept == nil {
		panic("implement not found for interface IAdminDept, forgot register?")
	}
	return localAdminDept
}

func RegisterAdminDept(i IAdminDept) {
	localAdminDept = i
}

func AdminMember() IAdminMember {
	if localAdminMember == nil {
		panic("implement not found for interface IAdminMember, forgot register?")
	}
	return localAdminMember
}

func RegisterAdminMember(i IAdminMember) {
	localAdminMember = i
}

func AdminMemberPost() IAdminMemberPost {
	if localAdminMemberPost == nil {
		panic("implement not found for interface IAdminMemberPost, forgot register?")
	}
	return localAdminMemberPost
}

func RegisterAdminMemberPost(i IAdminMemberPost) {
	localAdminMemberPost = i
}

func AdminMenu() IAdminMenu {
	if localAdminMenu == nil {
		panic("implement not found for interface IAdminMenu, forgot register?")
	}
	return localAdminMenu
}

func RegisterAdminMenu(i IAdminMenu) {
	localAdminMenu = i
}

func AdminSite() IAdminSite {
	if localAdminSite == nil {
		panic("implement not found for interface IAdminSite, forgot register?")
	}
	return localAdminSite
}

func RegisterAdminSite(i IAdminSite) {
	localAdminSite = i
}

func AdminCreditsLog() IAdminCreditsLog {
	if localAdminCreditsLog == nil {
		panic("implement not found for interface IAdminCreditsLog, forgot register?")
	}
	return localAdminCreditsLog
}

func RegisterAdminCreditsLog(i IAdminCreditsLog) {
	localAdminCreditsLog = i
}

func AdminPost() IAdminPost {
	if localAdminPost == nil {
		panic("implement not found for interface IAdminPost, forgot register?")
	}
	return localAdminPost
}

func RegisterAdminPost(i IAdminPost) {
	localAdminPost = i
}

func AdminRole() IAdminRole {
	if localAdminRole == nil {
		panic("implement not found for interface IAdminRole, forgot register?")
	}
	return localAdminRole
}

func RegisterAdminRole(i IAdminRole) {
	localAdminRole = i
}
