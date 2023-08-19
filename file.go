package echoopenai

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
)

type FileEntry struct {
	ID            string `json:"id,omitempty"`
	Object        string `json:"object,omitempty"`
	Bytes         int64  `json:"bytes,omitempty"`
	CreateAt      int64  `json:"create_at,omitempty"`
	FileName      string `json:"file_name,omitempty"`
	Purpose       string `json:"purpose,omitempty"`
	Status        string `json:"status,omitempty"`
	StatusDetails string `json:"status_details,omitempty"`
}

type FileResponse struct {
	Object string      `json:"object,omitempty"`
	Data   []FileEntry `json:"data,omitempty"`
}

type FileUploadRequest struct {
	File    *os.File `json:"file"`
	Purpose string   `json:"purpose"`
}

type FileDeletionStatus struct {
	Object  string `json:"object,omitempty"`
	ID      string `json:"id,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

func (c *Client) ListFiles() (response FileResponse, err error) {
	return c.ListFilesWithContext(context.Background())
}

func (c *Client) ListFilesWithContext(ctx context.Context) (response FileResponse, err error) {
	urlSuffix := "files"

	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodGet, c.getFullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	err = c.sendRequestWithContext(ctx, req, &response)
	return
}

func (c *Client) RetrieveFile(id string) (response FileEntry, err error) {
	return c.RetrieveFileWithContext(context.Background(), id)
}

func (c *Client) RetrieveFileWithContext(ctx context.Context, id string) (response FileEntry, err error) {
	urlSuffix := fmt.Sprintf("files/%v", id)

	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodGet, c.getFullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	err = c.sendRequestWithContext(ctx, req, &response)
	return
}

func (c *Client) UploadFile(request FileUploadRequest) (response FileResponse, err error) {
	return c.UploadFileWithContext(context.Background(), request)
}

func (c *Client) UploadFileWithContext(ctx context.Context, request FileUploadRequest) (response FileResponse, err error) {
	body := &bytes.Buffer{}
	builder := c.createFormBuilder(body)

	err = builder.CreateFormFile("file", request.File)
	if err != nil {
		return
	}

	err = builder.WriteField("purpose", request.Purpose)
	if err != nil {
		return
	}

	err = builder.Close()
	if err != nil {
		return
	}

	urlSuffix := "files"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getFullURL(urlSuffix), body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", builder.FormDataContentType())
	c.setCommonHeader(req)
	err = c.sendRequestWithContext(ctx, req, &response)
	return
}

func (c *Client) DeleteFile(id string) (response FileDeletionStatus, err error) {
	return c.DeleteFileWithContext(context.Background(), id)
}

func (c *Client) DeleteFileWithContext(ctx context.Context, id string) (response FileDeletionStatus, err error) {
	urlSuffix := fmt.Sprintf("files/%v", id)

	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodDelete, c.getFullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	err = c.sendRequestWithContext(ctx, req, &response)
	return
}

func (c *Client) RetrieveFileContent(id string) (response string, err error) {
	return c.RetrieveFileContentWithContext(context.Background(), id)
}

func (c *Client) RetrieveFileContentWithContext(ctx context.Context, id string) (response string, err error) {
	urlSuffix := fmt.Sprintf("files/%v/content", id)

	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodGet, c.getFullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	err = c.sendRequestWithContext(ctx, req, &response)
	return
}
