package http

import authUseCase "hms-backend-go-gin/app/auth/usecase"

type AuthDelivery struct {
	authUseCase authUseCase.AuthUseCase
}

func NewAuthHandler(authUseCase authUseCase.AuthUseCase) *AuthDelivery {
	return &AuthDelivery{authUseCase}
}
