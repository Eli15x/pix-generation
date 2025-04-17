package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	UserLoginSuccess = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "user_login_success_total",
		Help: "Total de logins bem-sucedidos",
	})

	UserLoginFailure = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "user_login_failure_total",
		Help: "Total de falhas de login",
	})
)

func InitBusinessMetrics() {
	prometheus.MustRegister(UserLoginSuccess, UserLoginFailure)
}
