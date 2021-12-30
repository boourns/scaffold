package model

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/boourns/scaffold/ast"
	"github.com/boourns/scaffold/sqlgen"
)

type ModelViewData struct {
	Model *ast.Model
}

func (mvd ModelViewData) FieldString(prefix string) template.HTML {
	qualified := make([]string, len(mvd.Model.Fields))
	for i, f := range mvd.Model.Fields {
		qualified[i] = fmt.Sprintf("&%s.%s", prefix, f.Name)
	}
	joined := strings.Join(qualified, ", ")
	return template.HTML(joined)
}

func (mvd ModelViewData) SelectStatement() string {
	fields := strings.Join(mvd.Model.FieldSlice(), ", ")
	return fmt.Sprintf("SELECT %s from %s ", fields, mvd.Model.Name)
}

func (mvd ModelViewData) UpdateStatement() string {
	fs := []string{}
	for _, f := range mvd.Model.Fields {
		if f.Name != "ID" {
			fs = append(fs, f.Name)
		}
	}
	fields := strings.Join(fs, ", ")

	return fmt.Sprintf("UPDATE %s SET %s WHERE ID = ?", mvd.Model.Name, fields)
}

func (mvd ModelViewData) UpdateParams() template.HTML {
	params := []string{}
	for _, f := range mvd.Model.Fields {
		if f.Name != "ID" {
			params = append(params, fmt.Sprintf("m.%s", f.Name))
		}
	}

	params = append(params, "m.ID")
	return template.HTML(fmt.Sprintf("[]interface{}{%s}", strings.Join(params, ", ")))
}

func (mvd ModelViewData) InsertStatement() string {
	fs := []string{}
	qs := []string{}
	for _, f := range mvd.Model.Fields {
		if f.Name != "ID" {
			fs = append(fs, f.Name)
			qs = append(qs, "?")
		}
	}
	fields := strings.Join(fs, ", ")
	questionsMarks := strings.Join(qs, ", ")

	return fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", mvd.Model.Name, fields, questionsMarks)
}

func (mvd ModelViewData) InsertParams() string {
	fs := []string{}

	for _, f := range mvd.Model.Fields {
		if f.Name != "ID" {
			fs = append(fs, fmt.Sprintf("m.%s", f.Name))

		}
	}
	fields := strings.Join(fs, ", ")

	return fmt.Sprintf("[]interface{}{%s}", fields)
}

func (mvd ModelViewData) CreateTable() string {
	return sqlgen.CreateTable(mvd.Model)
}
