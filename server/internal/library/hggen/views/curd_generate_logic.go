// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/sysin"
)

const (
	LogicWhereComments      = "\n\t// 查询%s\n"
	LogicWhereNoSupport     = "\t// TODO 暂不支持生成[ %s ]查询方式，请自行补充此处代码！"
	LogicListSimpleSelect   = "\tfields, err := hgorm.GenSelect(ctx, sysin.%sListModel{}, dao.%s)\n\tif err != nil {\n\t\treturn\n\t}"
	LogicListJoinSelect     = "\t//关联表select\n\tfields, err := hgorm.GenJoinSelect(ctx, %sin.%sListModel{}, &dao.%s, []*hgorm.Join{\n%v\t})\n\n\tif err != nil {\n\t\terr = gerror.Wrap(err, \"获取%s关联字段失败，请稍后重试！\")\n\t\treturn\n\t}"
	LogicListJoinOnRelation = "\t// 关联表%s\n\tmod = mod.%s(hgorm.GenJoinOnRelation(\n\t\tdao.%s.Table(), dao.%s.Columns().%s, // 主表表名,关联字段\n\t\tdao.%s.Table(), \"%s\", dao.%s.Columns().%s, // 关联表表名,别名,关联字段\n\t)...)\n\n"
	LogicEditUpdate         = "\tif _, err = s.Model(ctx%s).\n\t\t\tFields(%sin.%sUpdateFields{}).\n\t\t\tWherePri(in.%s).Data(in).Update(); err != nil {\n\t\t\terr = gerror.Wrap(err, \"修改%s失败，请稍后重试！\")\n\t\t}\n\t\treturn"
	LogicEditInsert         = "\tif _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).\n\t\tFields(%sin.%sInsertFields{}).\n\t\tData(in).Insert(); err != nil {\n\t\terr = gerror.Wrap(err, \"新增%s失败，请稍后重试！\")\n\t}"
	LogicEditUnique         = "\t// 验证'%s'唯一\n\tif err = hgorm.IsUnique(ctx, &dao.%s, g.Map{dao.%s.Columns().%s: in.%s}, \"%s已存在\", in.Id); err != nil {\n\t\treturn\n\t}\n"
	LogicSwitchUpdate       = "g.Map{\n\t\tin.Key:                       in.Value,\n%s}"
	LogicStatusUpdate       = "g.Map{\n\t\tdao.%s.Columns().Status:    in.Status,\n%s}"
)

func (l *gCurd) logicTplData(ctx context.Context, in *CurdPreviewInput) (data g.Map, err error) {
	data = make(g.Map)
	data["listWhere"] = l.generateLogicListWhere(ctx, in)
	data["listJoin"] = l.generateLogicListJoin(ctx, in)
	data["listOrder"] = l.generateLogicListOrder(ctx, in)
	data["edit"] = l.generateLogicEdit(ctx, in)
	data["switchFields"] = l.generateLogicSwitchFields(ctx, in)
	data["switchUpdate"] = l.generateLogicSwitchUpdate(ctx, in)
	data["statusUpdate"] = l.generateLogicStatusUpdate(ctx, in)
	return
}

func (l *gCurd) generateLogicStatusUpdate(ctx context.Context, in *CurdPreviewInput) string {
	var update string
	for _, field := range in.masterFields {
		if field.GoName == "UpdatedBy" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().UpdatedBy: contexts.GetUserId(ctx),\n"
		}
	}

	update += "\t"
	return fmt.Sprintf(LogicStatusUpdate, in.In.DaoName, update)
}

func (l *gCurd) generateLogicSwitchUpdate(ctx context.Context, in *CurdPreviewInput) string {
	var update string
	for _, field := range in.masterFields {
		if field.GoName == "UpdatedBy" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().UpdatedBy: contexts.GetUserId(ctx),\n"
		}
	}

	update += "\t"
	return fmt.Sprintf(LogicSwitchUpdate, update)
}

func (l *gCurd) generateLogicSwitchFields(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	if in.options.Step.HasSwitch {
		for _, field := range in.masterFields {
			if field.FormMode == "Switch" {
				buffer.WriteString("\t\tdao." + in.In.DaoName + ".Columns()." + field.GoName + ",\n")
			}
		}
	}
	return buffer.String()
}

func (l *gCurd) generateLogicEdit(ctx context.Context, in *CurdPreviewInput) g.Map {
	var (
		data         = make(g.Map)
		updateBuffer = bytes.NewBuffer(nil)
		insertBuffer = bytes.NewBuffer(nil)
		uniqueBuffer = bytes.NewBuffer(nil)
	)

	for _, field := range in.masterFields {
		if field.GoName == "UpdatedBy" {
			updateBuffer.WriteString("\t\tin.UpdatedBy = contexts.GetUserId(ctx)\n")
		}

		if field.GoName == "CreatedBy" {
			insertBuffer.WriteString("\tin.CreatedBy = contexts.GetUserId(ctx)\n")
		}

		if field.Unique {
			uniqueBuffer.WriteString(fmt.Sprintf(LogicEditUnique, field.GoName, in.In.DaoName, in.In.DaoName, field.GoName, field.GoName, field.Dc))
		}
	}

	notFilterAuth := ""
	if gstr.InArray(in.options.ColumnOps, "notFilterAuth") {
		notFilterAuth = ", &handler.Option{FilterAuth: false}"
	}

	updateBuffer.WriteString(fmt.Sprintf(LogicEditUpdate, notFilterAuth, in.options.TemplateGroup, in.In.VarName, in.pk.GoName, in.In.TableComment))
	insertBuffer.WriteString(fmt.Sprintf(LogicEditInsert, in.options.TemplateGroup, in.In.VarName, in.In.TableComment))

	data["update"] = updateBuffer.String()
	data["insert"] = insertBuffer.String()
	data["unique"] = uniqueBuffer.String()
	return data
}

func (l *gCurd) generateLogicListOrder(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)
	if in.options.Step.HasMaxSort {
		buffer.WriteString("OrderAsc(dao." + in.In.DaoName + ".Columns().Sort).")
	}
	buffer.WriteString("OrderDesc(dao." + in.In.DaoName + ".Columns()." + in.pk.GoName + ")")
	return buffer.String()
}

func (l *gCurd) generateLogicListJoin(ctx context.Context, in *CurdPreviewInput) g.Map {
	var data = make(g.Map)
	data["link"] = ""
	if hasEffectiveJoins(in.options.Join) {
		var (
			selectBuffer   = bytes.NewBuffer(nil)
			linkBuffer     = bytes.NewBuffer(nil)
			joinSelectRows string
		)

		for _, join := range in.options.Join {
			if isEffectiveJoin(join) {
				joinSelectRows = joinSelectRows + fmt.Sprintf("\t\t{Dao: &dao.%s, Alias: \"%s\"},\n", join.DaoName, join.Alias)
				linkBuffer.WriteString(fmt.Sprintf(LogicListJoinOnRelation, join.Alias, consts.GenCodesJoinLinkMap[join.LinkMode], in.In.DaoName, in.In.DaoName, gstr.CaseCamel(join.MasterField), join.DaoName, join.Alias, join.DaoName, gstr.CaseCamel(join.Field)))
			}
		}

		selectBuffer.WriteString(fmt.Sprintf(LogicListJoinSelect, in.options.TemplateGroup, in.In.VarName, in.In.DaoName, joinSelectRows, in.In.TableComment))

		data["select"] = selectBuffer.String()
		data["fields"] = "fields"
		data["link"] = linkBuffer.String()
	} else {
		data["fields"] = fmt.Sprintf("%sin.%sListModel{}", in.options.TemplateGroup, in.In.VarName)
	}
	return data
}

func (l *gCurd) generateLogicListWhere(ctx context.Context, in *CurdPreviewInput) string {
	buffer := bytes.NewBuffer(nil)

	// 主表
	l.generateLogicListWhereEach(buffer, in.masterFields, in.In.DaoName, "")

	// 关联表
	if hasEffectiveJoins(in.options.Join) {
		for _, v := range in.options.Join {
			if isEffectiveJoin(v) {
				l.generateLogicListWhereEach(buffer, v.Columns, v.DaoName, v.Alias)
			}
		}
	}
	return buffer.String()
}

func (l *gCurd) generateLogicListWhereEach(buffer *bytes.Buffer, fields []*sysin.GenCodesColumnListModel, daoName string, alias string) {
	isLink := false
	if alias != "" {
		alias = `"` + alias + `."+`
		isLink = true
	}
	for _, field := range fields {
		if !field.IsQuery || field.QueryWhere == "" {
			continue
		}

		var (
			linkMode   string
			whereTag   string
			columnName string
		)

		if IsNumberType(field.GoType) {
			linkMode = `in.` + field.GoName + ` > 0`
		} else if field.GoType == GoTypeGTime {
			linkMode = `in.` + field.GoName + ` != nil`
		} else if field.GoType == GoTypeJson {
			linkMode = `!in.` + field.GoName + `.IsNil()`
		} else {
			linkMode = `in.` + field.GoName + ` != ""`
		}

		if field.QueryWhere == WhereModeBetween || field.QueryWhere == WhereModeNotBetween {
			linkMode = `len(in.` + field.GoName + `) == 2`
		}

		buffer.WriteString(fmt.Sprintf(LogicWhereComments, field.Dc))

		// 如果是关联表重新转换字段
		columnName = field.GoName
		if isLink {
			columnName = gstr.CaseCamel(field.Name)
		}

		switch field.QueryWhere {
		case WhereModeEq:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.Where(" + alias + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeNeq:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereNot(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeGt:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereGT(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeGte:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereGTE(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeLt:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereLT(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeLte:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereLTE(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeIn:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereIn(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeNotIn:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereNotIn(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeBetween:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereBetween(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + "[0], in." + field.GoName + "[1])\n\t}"
		case WhereModeNotBetween:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereNotBetween(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + "[0], in." + field.GoName + "[1])\n\t}"
		case WhereModeLike:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereLike(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeLikeAll:
			val := `"%"+in.` + field.GoName + `+"%"`
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereLike(dao." + daoName + ".Columns()." + columnName + ", " + val + ")\n\t}"
		case WhereModeNotLike:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod.WhereNotLike(dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WhereModeJsonContains:
			val := "fmt.Sprintf(`JSON_CONTAINS(%s,'%v')`, dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")"
			whereTag = "\tif in." + field.GoName + linkMode + " {\n\t\tmod = mod.Where(" + val + ")\n\t}"

		default:
			buffer.WriteString(fmt.Sprintf(LogicWhereNoSupport, field.QueryWhere))
		}

		buffer.WriteString(whereTag + "\n")
	}
}
