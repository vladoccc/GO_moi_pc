package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRune = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ1234567890-*!")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassord(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRune[rand.IntN(len(letterRune))]
	}
	acc.password = string(result)
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("Invalid login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.generatePassord(10)
	}
	return newAcc, nil
}
