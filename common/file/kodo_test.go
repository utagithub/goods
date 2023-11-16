package file

import (
	"testing"
)

func TestKODOUpload(t *testing.T) {
	e := OXS{"", "", "", ""}
	var oxs = e.Setup(QiNiuKodo, map[string]interface{}{"Zone": "华东"})
	url, err := oxs.UpLoad("test.png", "./test.png")
	if err != nil {
		t.Error(err)
	}
	t.Log("ok")
	t.Log("url" + url)
}

func TestKODOGetTempToken(t *testing.T) {
	e := OXS{"", "", "", ""}
	var oxs = e.Setup(QiNiuKodo, map[string]interface{}{"Zone": "华东"})
	token, _ := oxs.GetTempToken()
	t.Log(token)
	t.Log("ok")
}
