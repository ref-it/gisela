package common

// DatabaseConfiguration serves the database settings declared in config/config.json
type DatabaseConfiguration struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type SmtpConfiguration struct {
	Host          string `json:"host"`
	Port          int    `json:"port"`
	User          string `json:"user"`
	Password      string `json:"password"`
	TLS           bool   `json:"tls"`
	StartTLS      bool   `json:"starttls"`
	SenderAddress string `json:"sender_address"`
	SenderName    string `json:"sender_name"`
}

type EmailConfiguration struct {
	SMTP SmtpConfiguration
}

// WebserverConfiguration serves the webserver settings declared in config/config.json
type WebserverConfiguration struct {
	Protocol string `json:"protocol"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
}

// SamlConfiguration serves the settings for SAML Authentification declared in config/config.json
type SamlConfiguration struct {
	MetdataUrl string `json:"metdataUrl"`
	SpCert     string `json:"spCert"`
	SpKey      string `json:"spKey"`
	SpBaseUrl  string `json:"spBaseUrl"`
}

type LatexConfiguration struct {
	Path string `json:"path"`
}

// Configuration serves all program settings declared in config/config.json
type Configuration struct {
	Database DatabaseConfiguration  `json:"database"`
	Email    EmailConfiguration     `json:"email"`
	Web      WebserverConfiguration `json:"web"`
	Saml     SamlConfiguration      `json:"saml"`
	Latex    LatexConfiguration     `json:"latex"`
}

type SpaHandler struct {
	StaticPath string
	IndexPath  string
}
