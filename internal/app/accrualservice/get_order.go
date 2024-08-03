package accrualservice

import (
	"context"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func (f *accrualsServiceAPI) GetOrder(ctx context.Context, orderNumber string) (*GetOrderResponse, error) {
	path := fmt.Sprintf(GetOrderURL, orderNumber)
	buf, err := f.doGet(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to do get: %w", err)
	}

	if buf == nil {
		return nil, nil
	}

	response := &GetOrderResponse{}

	err = jsoniter.Unmarshal(buf, response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	f.logger.InfoCtx(ctx, "get order response", response)

	return response, nil
}
