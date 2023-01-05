package mysql

// Config contains the values necessary to establish a connection with
// MySQL.
type Config struct {
	Database string `json:"database" yaml:"database"`
	Host     string `json:"host" yaml:"host"`
	Password string `json:"password" yaml:"password"`
	User     string `json:"user" yaml:"user"`
}
