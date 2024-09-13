package util

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func ConvertType(data, dest interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		logrus.Error("convertType body marshal error :: ", err)
		return err
	}
	err = json.Unmarshal(body, dest)
	if err != nil {
		logrus.Error("convertType body unmarshal error :: ", err)
		return err
	}
	return nil
}
