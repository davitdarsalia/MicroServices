package service

import (
	"context"
	"github.com/davitdarsalia/auth/internal/constants"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/types"
	"github.com/davitdarsalia/auth/internal/utils"
	"time"
)

func (r *RootService) Create(u *entities.User) (string, error) {
	ipChan := make(chan types.IpV16, 1)
	hashChan := make(chan types.Hash512, 1)
	salt := utils.Salt()

	go func() {
		ipChan <- utils.IpAddress()
	}()

	go func() {
		hashChan <- utils.Hash(u.Password, salt)
	}()

	u.CreatedAt = time.Now().Format(constants.RegularFormat)
	u.IpAddress = <-ipChan
	u.Password = <-hashChan

	userID, err := r.repository.Create(u)

	// If err == nil -> Do Redis Stuff
	if err == nil {
		r.redis.Set(context.Background(), constants.RedisSalt, salt, 0)
	}

	close(ipChan)
	close(hashChan)

	return userID, err
}

func (r *RootService) Login(u *entities.UserInput) (types.TokenPair, error) {
	salt, err := r.redis.Get(context.Background(), constants.RedisSalt).Result()

	if err != nil {
		return [2]string{"", ""}, err
	}

	userID, err := r.repository.Login(u.Email, utils.Hash(u.Password, salt))

	return utils.TokenPair(userID), err
}

func (r *RootService) Refresh(refreshToken types.RefreshToken) (types.TokenPair, error) {
	return utils.ParseRefreshToken(refreshToken)
}

func (r *RootService) Verify() {
	//TODO implement me
}

func (r *RootService) Reset() {
	//TODO implement me
}
