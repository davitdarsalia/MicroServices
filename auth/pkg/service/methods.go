package service

import (
	"github.com/davitdarsalia/auth/internal/constants"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/types"
	"github.com/davitdarsalia/auth/internal/utils"
	"time"
)

func (r RootService) Create(u *entities.User) (string, error) {
	ipChan := make(chan types.IpV16, 1)
	hashChan := make(chan types.Hash512, 1)

	go func() {
		ipChan <- utils.IpAddress()
	}()

	go func() {
		hashChan <- utils.Hash(utils.Hash(u.Password))
	}()

	u.CreatedAt = time.Now().Format(constants.RegularFormat)
	u.IpAddress = <-ipChan
	u.Password = <-hashChan

	userID, err := r.repository.Create(u)

	// If err == nil -> Do Redis Stuff

	return userID, err
}

func (r RootService) Login() {
	//TODO implement me
	panic("implement me")
}

func (r RootService) Refresh() {
	//TODO implement me
	panic("implement me")
}

func (r RootService) Verify() {
	//TODO implement me
	panic("implement me")
}

func (r RootService) Reset() {
	//TODO implement me
	panic("implement me")
}
