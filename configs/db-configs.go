package configs

type DBConfig struct {
	Host     string
	Name     string
	Password string
	User     string
	IsProd   bool
}

func GetDBConfigs(env string) *DBConfig {
	if env == "qa" {
		return &DBConfig{
			Host:     "ec2-3-209-38-221.compute-1.amazonaws.com",
			Name:     "d440nml0a86pe8",
			Password: "610c4d0e7a853449d21c4e1344b432b7ebff5a71466166728ee9dca963958fb8",
			User:     "tqykcafyttxirk",
			IsProd:   true}
	}

	if env == "prod" {
		return &DBConfig{
			Host:     "ec2-52-22-81-147.compute-1.amazonaws.com",
			Name:     "ddnmu64tjqh9ju",
			Password: "0ab277b623defd4ca7a72cba84bc60f06d7cabb6a8b311bc7580250bcef78b69",
			User:     "soxoxijvmbhqiv",
			IsProd:   true}
	}

	return &DBConfig{
		Host:     "localhost",
		Name:     "houmly",
		Password: "l8397040",
		User:     "postgres",
		IsProd:   false}
}
