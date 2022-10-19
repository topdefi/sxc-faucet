package server

type Config struct {
	network    string
	httpPort   int
	interval   int
	payout     float64
	proxyCount int
	queueCap   int
}

func NewConfig(network string, httpPort int, interval int, payout float64, proxyCount int, queueCap int) *Config {
	return &Config{
		network:    network,
		httpPort:   httpPort,
		interval:   interval,
		payout:     payout,
		proxyCount: proxyCount,
		queueCap:   queueCap,
	}
}
