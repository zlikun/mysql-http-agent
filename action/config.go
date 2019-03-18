package action

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 加载数据源（配置文件或MySQL）
func LoadSettings(filePath string) (*SettingsConf, error) {
	var settings *SettingsConf
	if data, err := ioutil.ReadFile(filePath); err != nil {
		return settings, err
	} else {
		if err := yaml.Unmarshal(data, &settings); err != nil {
			return settings, err
		}
	}

	return settings, nil
}

// 加载查询配置（配置文件或MySQL）
func LoadQueries(filePath string) (map[string]*QueryConf, error) {
	queryConfMap := make(map[string]*QueryConf)
	if data, err := ioutil.ReadFile(filePath); err != nil {
		return queryConfMap, err
	} else {
		if err := yaml.Unmarshal(data, &queryConfMap); err != nil {
			return queryConfMap, err
		}
	}
	return queryConfMap, nil
}

//// 检查数据源
//func checkDataSource(ds *DataSourceConf) error {
//
//	return nil
//}
//
//// 检查查询参数
//func checkQuery(query *QueryConf) error {
//
//	return nil
//}
