package utils

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
	//"strconv"
)

func GetServerConfig() *ServerConfig {
	var serverCfg ServerConfig
	err := envconfig.Process(SERVERPREFIX, &serverCfg)
	fmt.Println("Error Config Consul : ", err)
	return &serverCfg
}

func GetConsulConfig() ConsulConfig {
	var consulcfg ConsulConfig
	err := envconfig.Process(CONSULREGPREFIX, &consulcfg)
	fmt.Println("Error Config Consul : ", err)
	return consulcfg
}

func GetEnvConfigConsumerKafka() EnvKafkaConsumerConfig {
	var cfg EnvKafkaConsumerConfig
	err := envconfig.Process(KAFKACONSUMERPREFIX, &cfg)
	fmt.Println("Error Config Kafka Consumer : ", err)
	return cfg
}

func GetEnvConfigProcedurKafka() EnvKafkaProcedureConfig {
	var cfg EnvKafkaProcedureConfig
	err := envconfig.Process(KAFKAPRODUCERPREFIX, &cfg)
	fmt.Println("Error Config Consul : ", err)
	return cfg

}

// Simple helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetServerTlsConfig() *tls.Config {
	if servertls12client == "ON" {

		caCert, err := ioutil.ReadFile(serverfileca)
		if err != nil {
			fmt.Println("Error : ", err)

		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
		tlsConfig.BuildNameToCertificate()
		return tlsConfig
	}
	return &tls.Config{InsecureSkipVerify: true}

}

func StrucToMap(in interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

func GetEnvConfigOttoHttpReq() EnvHttpClient {
	var cfg EnvHttpClient
	err := envconfig.Process(OTTOHTTPCLIENT, &cfg)
	fmt.Println("Error Config HttpClient : ", err)
	return cfg
}

func ParsingDateSettlement(dateString string) (time.Time, error) {
	layoutDate := EnvString("MainSetup.DateLayout.Settlement", "02-01-2006 15:04:05")
	date, err := time.Parse(layoutDate, dateString)
	return date, err
}

func GenerateMessageErrorValidate(err error) []string {
	var res []string
	i := 0
	for _, fieldErr := range err.(validator.ValidationErrors) {
		i++
		res = append(res, fieldErr.Field()+" Required")

	}

	return res
}
