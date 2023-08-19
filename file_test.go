package echoopenai

import (
	"errors"
	"os"
	"testing"
)

func TestListFiles(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	res, err := client.ListFiles()
	if err != nil {
		t.Errorf("test list files func failed %v", err)
		return
	}

	t.Logf("test list files successed %v", res)
}

func TestRetrieveFile(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	id, err := retrieveFileID(client)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	res, err := client.RetrieveFile(id)
	if err != nil {
		t.Errorf("test retrieve files func failed %v", err)
		return
	}

	t.Logf("test retrieve file successed %v", res)
}

func TestUploadFile(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	file, err := os.Open("files/file_upload_prepared.jsonl")
	if err != nil {
		t.Errorf("open prepared upload files failed %v", err)
		return
	}

	req := FileUploadRequest{
		File:    file,
		Purpose: "fine-tune",
	}

	res, err := client.UploadFile(req)
	if err != nil {
		t.Errorf("test upload file func failed %v", err)
		return
	}

	t.Logf("test upload file func successed %v", res)
}

func TestDeleteFile(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	id, err := retrieveFileID(client)
	if err != nil {
		t.Errorf("test delete file func failed %v", err)
		return
	}

	res, err := client.DeleteFile(id)
	if err != nil {
		t.Errorf("test delete file func failed %v", err)
		return
	}

	t.Logf("test delete file func successed %v", res)
}

func TestRetrieveFileContent(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	id, err := retrieveFileID(client)
	if err != nil {

		return
	}

	res, err := client.RetrieveFileContent(id)
	if err != nil {
		t.Errorf("test retrieve file content func failed %v", err)
		return
	}

	t.Logf("test retrieve file content func successed %v", res)
}

func retrieveFileID(client *Client) (string, error) {
	res, err := client.ListFiles()
	if err != nil {
		return "", err
	}

	if len(res.Data) == 0 {
		return "", errors.New("error: list files len is 0")
	}

	return res.Data[0].ID, nil
}
