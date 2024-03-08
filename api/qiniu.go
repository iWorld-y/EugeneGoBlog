package api

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/config"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/http"
	"time"
)

func (*ApiHandler) QiNiuToken(w http.ResponseWriter, r *http.Request) {
	bucket := "eugenegoblog"
	putPolicy := storage.PutPolicy{Scope: bucket}
	putPolicy.Expires = uint64(time.Hour * 2) // 两小时有效期
	mac := qbox.NewMac(config.Cfg.System.QinniuAccessKey, config.Cfg.System.QinniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(w, upToken)
}
