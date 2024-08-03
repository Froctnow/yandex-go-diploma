package bootstrap

import (
	"time"

	"github.com/Froctnow/yandex-go-diploma/internal/app/accrualservice"
	"github.com/Froctnow/yandex-go-diploma/internal/app/client/http"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

func NewAccrualsService(accrualsSystemAddress string, logger logger.LogClient) accrualservice.Service {
	transport := http.NewHTTPClient(
		accrualservice.ServiceName,
		nil,
		time.Millisecond*time.Duration(10000),
		logger,
	)
	service := accrualservice.NewService(transport, accrualsSystemAddress, logger)
	return service
}
