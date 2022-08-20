package template

func InternalTpl(tpl string) bool {
	if _, ok := _bindata[tpl]; ok {
		return true
	}

	return false
}
