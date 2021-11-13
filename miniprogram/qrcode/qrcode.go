package qrcode

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/polichan/qq/miniprogram/context"
	"github.com/polichan/qq/util"
)

const (
	createMiniCodeURL = "https://api.q.qq.com/api/json/qqa/CreateMiniCode?access_token=%s"
)

// QRCode struct
type QRCode struct {
	*context.Context
}

// NewQRCode 实例
func NewQRCode(context *context.Context) *QRCode {
	qrCode := new(QRCode)
	qrCode.Context = context
	return qrCode
}

// QRCoder 小程序码参数
type QRCoder struct {
	// path 扫码进入的小程序页面路径
	Path string `json:"path,omitempty"`
}

// fetchCode 请求并返回二维码二进制数据
func (qrCode *QRCode) fetchCode(urlStr string, body interface{}) (response []byte, err error) {
	var accessToken string
	accessToken, err = qrCode.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(urlStr, accessToken)
	var contentType string
	response, contentType, err = util.PostJSONWithRespContentType(urlStr, body)
	if err != nil {
		return
	}
	if strings.HasPrefix(contentType, "application/json") {
		// 返回错误信息
		var result util.CommonError
		err = json.Unmarshal(response, &result)
		if err == nil && result.ErrCode != 0 {
			err = fmt.Errorf("fetchCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
			return nil, err
		}
	}
	if contentType == "image/jpeg" {
		// 返回文件
		return response, nil
	}
	err = fmt.Errorf("fetchCode error : unknown response content type - %v", contentType)
	return nil, err
}

// CreateQRCode 获取小程序二维码
// 文档地址： https://q.qq.com/wiki/develop/game/server/open-port/qr-code.html
func (qrCode *QRCode) CreateQRCode(coderParams QRCoder) (response []byte, err error) {
	return qrCode.fetchCode(createMiniCodeURL, coderParams)
}
