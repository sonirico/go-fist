package client

import (
	"fmt"
	"go-fist/fisttp"
	"log"
	"net"
	"os"
)

type FistClient struct {
	socket net.Conn
}

func NewFistClient(host string, port string) (*FistClient, error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	client := &FistClient{socket: conn}
	defer func() { client.Exit() }()
	return client, nil
}

func (fc *FistClient) dispatchRequest(request fisttp.Request) (*fisttp.Response, error) {
	_, err := fc.socket.Write([]byte(request.String()))
	if err != nil {
		fmt.Print("Error when writing")
	}
	responseBuffer := make([]byte, 1024)
	_, err = fc.socket.Read(responseBuffer)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("ok")

	return nil, nil
	// response, err := fisttp.ResponseParser(responseBuffer, request)
	// if err != nil {
	// 	fmt.Println("Error when parsing response")
	// 	return nil, fmt.Errorf("Error when parsing response")
	// }
	// return response, nil
}

func (fc *FistClient) Index(doc string, payload string) error {
	request := fisttp.NewIndexRequest(doc, payload)
	fc.dispatchRequest(request)
	return nil
}

func (fc *FistClient) Search(payload string) error {
	request := fisttp.NewSearchRequest(payload)
	fc.dispatchRequest(request)
	return nil
}

func (fc *FistClient) Exit() {
	request := fisttp.NewExitRequest()
	fc.dispatchRequest(request)
	fc.socket.Close()
	os.Exit(0)
}
