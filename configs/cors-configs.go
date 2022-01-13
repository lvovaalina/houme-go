package configs

type CorsConfigs struct {
	Domain    string
	SubDomain string
	IsProd    bool
}

func GetCorsConfigs(env string) *CorsConfigs {
	if env == "qa" {
		return &CorsConfigs{
			Domain: "https://houmly-dev.herokuapp.com",
			IsProd: true}
	}

	if env == "prod" {
		return &CorsConfigs{
			Domain:    "https://houmly.herokuapp.com",
			SubDomain: "https://demo.houmly.com",
			IsProd:    true}
	}

	return &CorsConfigs{
		Domain: "http://localhost:5001",
		IsProd: false}
}
