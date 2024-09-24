package utils

import (
	"fmt"
	"middleware-experience/constants"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
)

var (
	servertls12client string
	serverfileca      string

	//tlsserverconfig = true

	httpEnv Env
	log     = logrus.New()
	logDb   = logrus.New()
	year    int
	month   time.Month
	day     int
	PkgName string = "Utils - "
	LogHook *graylog.GraylogHook
	//GLog *graylog.Graylog
)

func init() {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "localhost"
	}
	//log.SetFormatter(formatters.NewGelf(hostname))
	log.SetFormatter(&logrus.JSONFormatter{})
	logDb.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "01-02-2006 15:04:05",
	})

	LogHook = graylog.NewGraylogHook("127.0.0.1:7777", map[string]interface{}{"Service": "Middleware Experience Services"})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stderr)
	logDb.SetOutput(os.Stderr)

	// Only log the warning severity or above.

	year, month, day = time.Now().Local().Date()

	if err := envconfig.Process("HTTP", &httpEnv); err != nil {
		LogData("Utils - HTTP", "init", constants.LEVEL_LOG_WARNING, err.Error())
	}

	viper.SetConfigName("development-config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		LogData("Viper - Init", "viper.ReadInConfig", constants.LEVEL_LOG_WARNING, fmt.Sprintf("Error while reading config file %s", err))
	}
}

const (
	CONSULREGPREFIX     = "consulreg"
	SERVERPREFIX        = "server"
	KAFKAPRODUCERPREFIX = "kafkapro"
	KAFKACONSUMERPREFIX = "kafkaconsumer"
	OTTOHTTPCLIENT      = "ottohttpclient"
)

type (
	Env struct {
		DebugClient bool   `envconfig:"DEBUG_CLIENT" default:"true"`
		Timeout     string `envconfig:"TIMEOUT_CLIENT" default:"60s"`
		RetryBad    int    `envconfig:"RETRY_BAD_CLIENT" default:"1"`
	}

	EnvHttpClient struct {
		HttpClientTimeout string `envconfig:"timeout"`
		HttpClientRetry   int    `envconfig:"retry"`
		HttpClientTracing bool   `envconfig:"enabletracing"`
		HttpClientDebug   bool   `envconfig:"debug"`
	}
	EnvKafkaProcedureConfig struct {
		KafkaBrokerUrl            string        `envconfig:"brokerurl"`
		KafkaClient               string        `envconfig:"clientid"`
		KafkaProducerTimeout      time.Duration `envconfig:"to"`
		KafkaProducerDialTimeout  time.Duration `envconfig:"dialto"`
		KafkaProducerReadTimeout  time.Duration `envconfig:"readto"`
		KafkaProducerWriteTimeout time.Duration `envconfig:"writeto"`
		KafkaProducerMaxmsgbyte   int           `envconfig:"maxmsgbyte"`
	}
	EnvKafkaConsumerConfig struct {
		KafkaZookeeper string `envconfig:"kafkazookeeper"`
		KafkaBroker    string `envconfig:"kafkabroker"`
	}

	ConsulConfig struct {
		HostAddres string `envconfig:"hostaddres"`
		IdServer   string
		Name       string
		Server     string
		Port       int
		Hcurl      string
		HcInterval string
		HcTimeout  string
		Status     bool
	}

	ServerConfig struct {

		// ReadTimeout is the maximum duration for reading the entire
		// request, including the body.
		//
		// Because ReadTimeout does not let Handlers make per-request
		// decisions on each request body's acceptable deadline or
		// upload rate, most users will prefer to use
		// ReadHeaderTimeout. It is valid to use them both.
		ReadTimeout time.Duration

		// ReadHeaderTimeout is the amount of time allowed to read
		// request headers. The connection's read deadline is reset
		// after reading the headers and the Handler can decide what
		// is considered too slow for the body.
		ReadHeaderTimeout time.Duration `envconfig:"rhtimeout"`

		// WriteTimeout is the maximum duration before timing out
		// writes of the response. It is reset whenever a new
		// request's header is read. Like ReadTimeout, it does not
		// let Handlers make decisions on a per-request basis.
		WriteTimeout time.Duration

		// IdleTimeout is the maximum amount of time to wait for the
		// next request when keep-alives are enabled. If IdleTimeout
		// is zero, the value of ReadTimeout is used. If both are
		// zero, ReadHeaderTimeout is used.
		IdleTimeout time.Duration

		// MaxHeaderBytes controls the maximum number of bytes the
		// server will read parsing the request header's keys and
		// values, including the request line. It does not limit the
		// size of the request body.
		// If zero, DefaultMaxHeaderBytes is used.
		MaxHeaderBytes int `envconfig:"maxbytes"`

		Serverfileca         string `envconfig:"FILECA"`
		Serverfileprivatekey string `envconfig:"PRIVATEKEY"`
		Serverfilepubkey     string `envconfig:"PUBLICKEY"`
		Servertls12client    string `envconfig:"TLS12STATUS"`
	}
)
