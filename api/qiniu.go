package api

import (
	"EugeneGoBlog/common"
	"EugeneGoBlog/config"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"net/http"
	"time"
)

func (*ApiHandler) QiNiuToken(w http.ResponseWriter, r *http.Request) {
	log.Println("开始上传文件")
	bucket := "eugenegoblog"
	putPolicy := storage.PutPolicy{Scope: bucket}
	putPolicy.Expires = uint64(time.Hour) // 一小时有效期
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(w, upToken)
}
