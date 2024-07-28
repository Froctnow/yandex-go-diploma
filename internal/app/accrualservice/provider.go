package accrualservice

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type accrualsServiceAPI struct {
	baseURL    string
	httpClient HTTPClient
	logger     logger.LogClient
}

//go:generate mockery --with-expecter --case=underscore --name=Service

type Service interface {
	GetOrder(ctx context.Context, orderNumber string) (*GetOrderResponse, error)
}

func NewService(httpClient HTTPClient, accrualSystemAddress string, logger logger.LogClient) Service {
	return &accrualsServiceAPI{
		baseURL:    accrualSystemAddress,
		httpClient: httpClient,
		logger:     logger,
	}
}

func (f *accrualsServiceAPI) doGet(ctx context.Context, path string) ([]byte, error) {
	url := f.baseURL + path
	reqWithCtx, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	respRaw, err := f.httpClient.Do(reqWithCtx)
	if err != nil {
		return nil, err
	}

	if respRaw == nil {
		return nil, errors.New(`nil response value`)
	}

	if respRaw.StatusCode == http.StatusNoContent {
		f.logger.InfoCtx(ctx, "no content response", "url", url)
		return nil, nil
	}

	buf, err := io.ReadAll(respRaw.Body)
	if err != nil {
		return nil, err
	}
	return buf, respRaw.Body.Close()
}

func (f *accrualsServiceAPI) makeURL(path string) string {
	return f.baseURL + path
}
