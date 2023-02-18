package service

import (
	"e-wallet-api/internal/dto"
	"e-wallet-api/internal/model"
	r "e-wallet-api/internal/repository"
	"e-wallet-api/pkg/custom_error"
	"e-wallet-api/pkg/utils"
)

type WalletService interface {
	GetWalletByUserId(input *dto.WalletRequestBody) (*model.Wallet, error)
	CreateWallet(input *dto.WalletRequestBody) (*model.Wallet, error)
}

type walletService struct {
	userRepository   r.UserRepository
	walletRepository r.WalletRepository
}

type WSConfig struct {
	UserRepository   r.UserRepository
	WalletRepository r.WalletRepository
}

func NewWalletService(c *WSConfig) WalletService {
	return &walletService{
		userRepository:   c.UserRepository,
		walletRepository: c.WalletRepository,
	}
}

func (s *walletService) GetWalletByUserId(input *dto.WalletRequestBody) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByUserId(input.UserID)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (s *walletService) CreateWallet(input *dto.WalletRequestBody) (*model.Wallet, error) {
	user, err := s.userRepository.FindById(input.UserID)
	if err != nil {
		return &model.Wallet{}, err
	}
	if user.ID == 0 {
		return &model.Wallet{}, &custom_error.UserNotFoundError{}
	}

	wallet, err := s.walletRepository.FindByUserId(int(user.ID))
	if err != nil {
		return &model.Wallet{}, err
	}
	if wallet.ID != 0 {
		return &model.Wallet{}, &custom_error.WalletAlreadyExistsError{}
	}

	wallet.UserID = user.ID
	wallet.Number = utils.GenerateWalletNumber(user.ID)
	wallet.Balance = 0

	newWallet, err := s.walletRepository.Save(wallet)
	if err != nil {
		return newWallet, err
	}

	return newWallet, nil
}
