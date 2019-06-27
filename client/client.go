package client

import (
	"bufio"
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
	return client, nil
}

func (fc *FistClient) dispatchRequest(request fisttp.Request) fisttp.Response {
	_, err := fc.socket.Write([]byte(request.String()))
	if err != nil {
		fmt.Print("Error when writing")
	}

	responseBuffer, err := fc.read()

	if err != nil {
		fmt.Println(err)
	}

	return fisttp.ParseResponse(request.Type(), responseBuffer)
}

func (fc *FistClient) read() ([]byte, error) {
	reader := bufio.NewReader(fc.socket)
	mBytes, err := reader.ReadBytes('\n')
	return mBytes, err
}

func (fc *FistClient) Index(doc string, payload string) bool {
	request := fisttp.NewIndexRequest(doc, payload)
	response := fc.dispatchRequest(request)
	return response.IsOk()
}

func (fc *FistClient) Search(payload string) []string {
	request := fisttp.NewSearchRequest(payload)
	response := fc.dispatchRequest(request)
	if response.IsOk() {
		sResponse := response.(*fisttp.SearchResponse)
		return sResponse.Documents
	}
	return nil
}

func (fc *FistClient) Exit() {
	defer func() {
		_ = fc.socket.Close()
	}()

	request := fisttp.NewExitRequest()
	response := fc.dispatchRequest(request)
	if response.IsOk() {
		os.Exit(0)
		return
	}
	os.Exit(1)
}
