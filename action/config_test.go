package action

import (
	"fmt"
	"testing"
)

func Test_LoadSettings(t *testing.T) {
	if settings, err := LoadSettings("../conf/settings.yml"); err != nil {
		t.Errorf("加载系统配置文件失败：%v", err)
		return
	} else {
		fmt.Println(settings)
	}
}

func Test_LoadQueries(t *testing.T) {
	if queries, err := LoadQueries("../conf/queries.yml"); err != nil {
		t.Errorf("加载查询配置文件失败：%v", err)
		return
	} else {
		for key, value := range queries {
			fmt.Println(key, *value)
		}
	}
}
