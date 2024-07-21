package config

func init() {
	Default = default_()
}

var Default *Cfg

type Cfg struct {
	*Server
}

type Server struct {
	Port int // 启动端口
}

func default_() *Cfg {
	return &Cfg{
		Server: &Server{
			Port: 8081,
		},
	}
}
