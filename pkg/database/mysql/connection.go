package database

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"simple_api/pkg/config"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var MySQLConfig = config.AppConfig.DB_MYSQL

	// Se CertPath está definido, registre a configuração TLS personalizada
	if MySQLConfig.CertPath != "" {
		rootCertPool := x509.NewCertPool()
		pem, err := os.ReadFile(MySQLConfig.CertPath)
		if err != nil {
			panic("failed to read certificate file: " + err.Error())
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			panic("failed to append certificate")
		}
		tlsConfig := &tls.Config{
			RootCAs: rootCertPool,
		}
		err = mysqlDriver.RegisterTLSConfig("custom", tlsConfig)
		if err != nil {
			panic("failed to register custom TLS config: " + err.Error())
		}
	}

	dsn := MySQLConfig.Username + ":" + MySQLConfig.Password + "@tcp(" + MySQLConfig.Host + ":" + MySQLConfig.Port + ")/" + MySQLConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to Azure MySQL: " + err.Error())
	}
	return db
}
