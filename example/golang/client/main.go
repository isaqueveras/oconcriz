package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/isaqueveras/juazeiro"

	"github.com/isaqueveras/juaz/example/golang/protos/biometry"
	"github.com/isaqueveras/juaz/example/golang/protos/user"
)

func main() {
	conn, err := juazeiro.NewClient("http://localhost:8181")
	if err != nil {
		handling(err)
		return
	}

	input := &user.User{
		Id:    pointer(int64(52)),
		Name:  pointer("Steve"),
		Level: pointer(user.LevelAdmin),
	}

	input.ParamLimit(pointer(89))
	input.ParamOffset(pointer(0))
	input.ParamTickets([]*int64{
		pointer(int64(123213)),
		pointer(int64(121235)),
		pointer(int64(768663)),
	})

	repoUser := user.NewUserClient(conn)
	userData, err := repoUser.GetUser(context.Background(), input)
	if err != nil {
		handling(err)
		return
	}
	showResponse(userData)

	inputBio := &biometry.SearchParams{
		RequestId: userData.Name,
	}

	repoBio := biometry.NewBiometryClient(conn)
	searchData, err := repoBio.Search(context.Background(), inputBio)
	if err != nil {
		handling(err)
		return
	}

	showResponse(searchData)

	if _, err = repoUser.EditUser(context.Background(), input); err != nil {
		handling(err)
		return
	}
}

func showResponse(value any) {
	str, err := json.Marshal(&value)
	if err != nil {
		panic(err)
	}
	log.Println(string(str))
}
