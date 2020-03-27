package model

type Collect struct {
	MysqlConf `ini:"mysql"`
	AddrConf  `ini:"addr"`
	ModeConf  `ini:"mode"`
}

type MysqlConf struct {
	Alias   string `ini:"db_alias"`
	Name    string `ini:"db_name"`
	User    string `ini:"db_user"`
	Pwd     string `ini:"db_pwd"`
	Host    string `ini:"db_host"`
	Port    string `ini:"db_port"`
	Charset string `ini:"db_charset"`
}

type AddrConf struct {
	GateAddr        string `ini:"gatewayaddr"`
	OrderserverAddr string `ini:"orderserveraddr"`
}

type ModeConf struct {
	RunMode string `ini:"runmode"`
}
