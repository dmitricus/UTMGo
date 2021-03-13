package models

type EmailServers struct {
	ID                uint      `json:"id" gorm:"primary_key"`                                             // ID
	EmailDefaultFrom  string    `json:"email_default_from"`                                                // FROM:
	EmailHost         string    `json:"email_host"`                                                        // Хост
	EmailPort         int32     `json:"email_port"`                                                        // Порт
	EmailUsername     string    `json:"email_username"`                                                    // Имя пользователя
	EmailPassword     string    `json:"email_password"`                                                    // Пароль
	EmailUseSSL       bool      `json:"email_use_ssl"`                                                     // Использовать SSL
	EmailUseTLS       bool      `json:"email_use_tls"`                                                     // Использовать TLS
	EmailFailSilently string    `json:"email_fail_silently"`                                               // Тихое подавление ошибок
	EmailTimeout      int32     `json:"email_timeout"`                                                     // Тайм-аут
	EmailSSLCertFile  string    `json:"email_ssl_cert_file" gorm:filepat`                                  // Сертификат EmailServers/certfile/%Y/%m/%d
	EmailSSLKeyfile   string    `json:"email_ssl_key_file"`                                                // Файл ключа EmailServers/keyfile/%Y/%m/%d
	ApiKey            string    `json:"api_key"`                                                           // API KEY
	ApiUsername       string    `json:"api_username"`                                                      // Имя пользователя для авторизации в API
	ApiFromEmail      string    `json:"api_from_email"`                                                    // Email адрес для отправки через API
	ApiFromName       string    `json:"api_from_name"`                                                     // Имя перед адресом для отправки через API
	SendingMethod     string    `json:"sending_method"`                                                    // Способ отправки писем
	MainServer        string    `json:"main_server"`                                                       // Основной сервер
	IsActive          string    `json:"is_active"`                                                         // Сервер активен
	PreferredDomains  []Domains `json:"preferred_domains" gorm:"many2many:emailservers_preferreddomains;"` // Предпочтительней для доменов
}
