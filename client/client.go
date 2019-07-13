// Package client provides the interface to communicate with FIST server
package client

import (
	"bufio"
	"fmt"
	"go-fist/fisttp"
	"net"
	"os"
)

// FistClient will carry the wire connected to the server. It will also implement
// all the command an user can issue to server
type FistClient struct {
	socket net.Conn
}

// NewFistClient initialises the connection based on program arguments. If the
// dial cannot be setup, an error will be returned. Otherwise, a pointer to instance
// of it.
func NewFistClient(host string, port string) (*FistClient, error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	// TODO: Add support to configure timeout
	// TODO: Add support to configure retries
	if err != nil {
		return nil, err
	}
	client := &FistClient{socket: conn}
	return client, nil
}

func (fc *FistClient) dispatchRequest(request fisttp.Request) fisttp.Response {
	_, err := fc.socket.Write([]byte(request.String()))
	if err != nil {
		fmt.Print(err)
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

// Index command will issue an INDEX request to the server
func (fc *FistClient) Index(doc string, payload string) bool {
	request := fisttp.NewIndexRequest(doc, payload)
	response := fc.dispatchRequest(request)
	return response.IsOk()
}

// Search command will search into the server for matching documents
// for a given string
func (fc *FistClient) Search(payload string) []string {
	request := fisttp.NewSearchRequest(payload)
	response := fc.dispatchRequest(request)
	if response.IsOk() {
		sResponse := response.(*fisttp.SearchResponse)
		return sResponse.Documents
	}
	return nil
}

// Delete command will remove the given keyword set from all documents
func (fc *FistClient) Delete(payload string) bool {
	request := fisttp.NewDeleteRequest(payload)
	response := fc.dispatchRequest(request)
	return response.IsOk()
}

// Version command will pull the server version
func (fc *FistClient) Version() (string, error) {
	request := fisttp.NewVersionRequest()
	response := fc.dispatchRequest(request)
	if response.IsOk() {
		vResponse := response.(*fisttp.VersionResponse)
		return vResponse.GetVersion(), nil
	}
	return "", fmt.Errorf("version could not be retrieved")
}

// Exit command will terminate a connection
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
