package auth

import (
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

type authUseCase struct {
	logger logger.LogClient
}

func NewUseCase(
	logger logger.LogClient,
) UseCase {
	return &authUseCase{
		logger: logger,
	}
}

type UseCase interface {
}
