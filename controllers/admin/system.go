package admin

import (
	"github.com/lisijie/goblog/models/option"
)

type SystemController struct {
	baseController
}

//系统设置
func (this *SystemController) Setting() {
	var result []*option.Option
	new(option.Option).Query().All(&result)

	options := make(map[string]string)
	mp := make(map[string]*option.Option)
	for _, v := range result {
		options[v.Name] = v.Value
		mp[v.Name] = v
	}

	if this.isPost() {
		keys := []string{"sitename", "siteurl", "subtitle", "pagesize", "keywords", "description", "email", "theme", "timezone", "stat"}
		for _, key := range keys {
			val := this.GetString(key)
			if _, ok := mp[key]; !ok {
				opt := new(option.Option)
				opt.Name = key
				opt.Value = val
				options[key] = val
				opt.Insert()
			} else {
				opt := mp[key]
				opt.Value = val
				opt.Update("Value")
			}
		}
		this.Redirect("/admin/system/setting", 302)
	}

	this.Data["now"] = this.getTime()
	this.Data["options"] = options
	this.display()
}
