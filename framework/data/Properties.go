package data

type DBType string

const(
	POSTGRES DBType = "postgres"
	ORACLE DBType = "oracle"
	MONGO DBType = "mongo"
)

type DbProperties struct {
	Type DBType `yaml:"type"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

func (p DbProperties) IsValid() bool {
	return p.Host != "" && p.Database != ""
}

type WithProperties struct {
	Properties *DbProperties
}

func (w WithProperties) IsValid() bool {
	return w.Properties.IsValid()
}