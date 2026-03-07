package contract

import "net/http"

type GetDashboardStatsHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
