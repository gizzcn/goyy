// generated by xgen -- DO NOT EDIT
package blacklist

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	ENTITY           = schema.TABLE("sys_blacklist", "BLACKLIST")
	ENTITY_ID        = ENTITY.PRIMARY("id", "ID")
	ENTITY_MEMO      = ENTITY.COLUMN("memo", "MEMO")
	ENTITY_CREATES   = ENTITY.COLUMN("creates", "CREATES")
	ENTITY_CREATER   = ENTITY.CREATER("creater", "CREATER")
	ENTITY_CREATED   = ENTITY.CREATED("created", "CREATED")
	ENTITY_MODIFIER  = ENTITY.MODIFIER("modifier", "MODIFIER")
	ENTITY_MODIFIED  = ENTITY.MODIFIED("modified", "MODIFIED")
	ENTITY_VERSION   = ENTITY.VERSION("version", "VERSION")
	ENTITY_DELETION  = ENTITY.DELETION("deletion", "DELETION")
	ENTITY_ARTIFICAL = ENTITY.COLUMN("artifical", "ARTIFICAL")
	ENTITY_HISTORY   = ENTITY.COLUMN("history", "HISTORY")
	ENTITY_NAME      = ENTITY.COLUMN("name", "NAME")
	ENTITY_GENRE     = ENTITY.COLUMN("genre", "GENRE")
	ENTITY_USABLE    = ENTITY.COLUMN("usable", "USABLE")
)

func NewEntity() *Entity {
	e := &Entity{}
	e.init()
	return e
}

func (me *Entity) Name() string {
	return me.name.Value()
}

func (me *Entity) SetName(v string) {
	me.name.SetValue(v)
}

func (me *Entity) Genre() string {
	return me.genre.Value()
}

func (me *Entity) SetGenre(v string) {
	me.genre.SetValue(v)
}

func (me *Entity) Usable() string {
	return me.usable.Value()
}

func (me *Entity) SetUsable(v string) {
	me.usable.SetValue(v)
}

func (me *Entity) init() {
	me.table = ENTITY
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
	me.initSetExcel()
	me.initSetJson()
	me.initSetXml()
}

func (me *Entity) initSetDict() {
}

func (me *Entity) initSetColumn() {
	if t, ok := me.Sys.Type("id"); ok {
		t.SetColumn(ENTITY_ID)
	}
	if t, ok := me.Sys.Type("memo"); ok {
		t.SetColumn(ENTITY_MEMO)
	}
	if t, ok := me.Sys.Type("creates"); ok {
		t.SetColumn(ENTITY_CREATES)
	}
	if t, ok := me.Sys.Type("creater"); ok {
		t.SetColumn(ENTITY_CREATER)
	}
	if t, ok := me.Sys.Type("created"); ok {
		t.SetColumn(ENTITY_CREATED)
	}
	if t, ok := me.Sys.Type("modifier"); ok {
		t.SetColumn(ENTITY_MODIFIER)
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetColumn(ENTITY_MODIFIED)
	}
	if t, ok := me.Sys.Type("version"); ok {
		t.SetColumn(ENTITY_VERSION)
	}
	if t, ok := me.Sys.Type("deletion"); ok {
		t.SetColumn(ENTITY_DELETION)
	}
	if t, ok := me.Sys.Type("artifical"); ok {
		t.SetColumn(ENTITY_ARTIFICAL)
	}
	if t, ok := me.Sys.Type("history"); ok {
		t.SetColumn(ENTITY_HISTORY)
	}
	me.name.SetColumn(ENTITY_NAME)
	me.genre.SetColumn(ENTITY_GENRE)
	me.usable.SetColumn(ENTITY_USABLE)
}

func (me *Entity) initSetDefault() {
	if t, ok := me.Sys.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}
}

func (me *Entity) initSetField() {
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}
	me.name.SetField(entity.DefaultField())
	me.genre.SetField(entity.DefaultField())
	me.usable.SetField(entity.DefaultField())
}

func (me *Entity) initSetExcel() {
}

func (me *Entity) initSetJson() {
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.Field().SetJson(entity.NewJsonBy(c))
		}
	}
	me.name.Field().SetJson(entity.NewJsonBy("name"))
	me.genre.Field().SetJson(entity.NewJsonBy("genre"))
	me.usable.Field().SetJson(entity.NewJsonBy("usable"))
}

func (me *Entity) initSetXml() {
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.Field().SetXml(entity.NewXmlBy(c))
		}
	}
	me.name.Field().SetXml(entity.NewXmlBy("name"))
	me.genre.Field().SetXml(entity.NewXmlBy("genre"))
	me.usable.Field().SetXml(entity.NewXmlBy("usable"))
}

func (me Entity) New() entity.Interface {
	return NewEntity()
}

func (me *Entity) Get(column string) interface{} {
	switch column {
	case ENTITY_NAME.Name():
		return me.name.Value()
	case ENTITY_GENRE.Name():
		return me.genre.Value()
	case ENTITY_USABLE.Name():
		return me.usable.Value()
	}
	return me.Sys.Get(column)
}

func (me *Entity) GetPtr(column string) interface{} {
	switch column {
	case ENTITY_NAME.Name():
		return me.name.ValuePtr()
	case ENTITY_GENRE.Name():
		return me.genre.ValuePtr()
	case ENTITY_USABLE.Name():
		return me.usable.ValuePtr()
	}
	return me.Sys.GetPtr(column)
}

func (me *Entity) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "name":
		return me.name.String()
	case "genre":
		return me.genre.String()
	case "usable":
		return me.usable.String()
	}
	return me.Sys.GetString(field)
}

func (me *Entity) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "name":
		return me.name.SetString(value)
	case "genre":
		return me.genre.SetString(value)
	case "usable":
		return me.usable.SetString(value)
	}
	return me.Sys.SetString(field, value)
}

func (me *Entity) Table() schema.Table {
	return me.table
}

func (me *Entity) Type(column string) (entity.Type, bool) {
	switch column {
	case ENTITY_NAME.Name():
		return &me.name, true
	case ENTITY_GENRE.Name():
		return &me.genre, true
	case ENTITY_USABLE.Name():
		return &me.usable, true
	}
	return me.Sys.Type(column)
}

func (me *Entity) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "name":
		return ENTITY_NAME, true
	case "genre":
		return ENTITY_GENRE, true
	case "usable":
		return ENTITY_USABLE, true
	}
	return me.Sys.Column(field)
}

func (me *Entity) Columns() []schema.Column {
	return []schema.Column{
		ENTITY_ID,
		ENTITY_MEMO,
		ENTITY_CREATES,
		ENTITY_CREATER,
		ENTITY_CREATED,
		ENTITY_MODIFIER,
		ENTITY_MODIFIED,
		ENTITY_VERSION,
		ENTITY_DELETION,
		ENTITY_ARTIFICAL,
		ENTITY_HISTORY,
		ENTITY_NAME,
		ENTITY_GENRE,
		ENTITY_USABLE,
	}
}

func (me *Entity) Names() []string {
	return []string{
		"id",
		"memo",
		"creates",
		"creater",
		"created",
		"modifier",
		"modified",
		"version",
		"deletion",
		"artifical",
		"history",
		"name",
		"genre",
		"usable",
	}
}

func (me *Entity) Value() *Entity {
	return me
}

func (me *Entity) Validate() error {
	return nil
}

func (me *Entity) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(`"id":"` + jsons.Format(me.GetString("id")) + `"`)
	b.WriteString(`,"memo":"` + jsons.Format(me.GetString("memo")) + `"`)
	b.WriteString(`,"creates":"` + jsons.Format(me.GetString("creates")) + `"`)
	b.WriteString(`,"creater":"` + jsons.Format(me.GetString("creater")) + `"`)
	b.WriteString(`,"created":` + me.GetString("created"))
	b.WriteString(`,"modifier":"` + jsons.Format(me.GetString("modifier")) + `"`)
	b.WriteString(`,"modified":` + me.GetString("modified"))
	b.WriteString(`,"version":` + me.GetString("version"))
	b.WriteString(`,"deletion":` + me.GetString("deletion"))
	b.WriteString(`,"artifical":` + me.GetString("artifical"))
	b.WriteString(`,"history":` + me.GetString("history"))
	b.WriteString(`,"name":"` + jsons.Format(me.GetString("name")) + `"`)
	b.WriteString(`,"genre":"` + jsons.Format(me.GetString("genre")) + `"`)
	b.WriteString(`,"usable":"` + jsons.Format(me.GetString("usable")) + `"`)
	b.WriteString("}")
	return b.String()
}

func (me *Entity) ExcelColumns() []string {
	return nil
}
