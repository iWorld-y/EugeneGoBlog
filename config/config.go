package config

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QinniuAccessKey string
	QinniuSecretKey string
	Valine          bool
	ValineAppid     string
	ValineAppKey    string
	ValineServerURL string
}
