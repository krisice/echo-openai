package echoopenai

import (
	"bytes"
	"context"
	"net/http"
)

type RequestBuilder interface {
	Build(method, url string, request any) (*http.Request, error)
	BuildWithContext(ctx context.Context, method, url string, request any) (*http.Request, error)
}

type HTTPRequestBuilder struct {
	marshaller Marshaller
}

func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		marshaller: &JSONMarshaller{},
	}
}

func (b *HTTPRequestBuilder) BuildWithContext(ctx context.Context, method, url string, request any) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	var reqBytes []byte
	var err error

	var contentType string
	switch request.(type) {
	case ImageEditRequest, ImageVariationRequest:
		reqBytes, contentType, err = generateFormData(request)
	default:
		reqBytes, err = b.marshaller.Marshal(request)
	}

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBytes),
	)
	if err != nil {
		return nil, err
	}

	if len(contentType) != 0 {
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Connection", "keep-alive")
	}

	return req, nil
}

func (b *HTTPRequestBuilder) Build(method, url string, request any) (*http.Request, error) {
	return b.BuildWithContext(context.Background(), method, url, request)
}

func generateFormData(request any) (reqBytes []byte, contentType string, err error) {
	switch req := request.(type) {
	case ImageEditRequest:
		reqBytes, contentType, err = generateImageEditRequestFormData(req)
	case ImageVariationRequest:
		reqBytes, err = generateImageVariationRequestFormData(req)
	default:
		err = ErrUnexpectedImageRequestType
	}
	return
}

func generateImageEditRequestFormData(request ImageEditRequest) (reqBytes []byte, contentType string, err error) {
	return
}

func generateImageVariationRequestFormData(request ImageVariationRequest) (reqBytes []byte, err error) {
	// var bytesBuf bytes.Buffer
	// multiPartWriter := multipart.NewWriter(&bytesBuf)
	// defer multiPartWriter.Close()

	// createFormDataErr := generateFileFormData(multiPartWriter, "image", request.Image, "image.png")
	// if createFormDataErr != nil {
	// 	err = createFormDataErr
	// 	return
	// }

	// err = generateCommonRequestFormData(multiPartWriter, request.Config)
	// reqBytes = bytesBuf.Bytes()
	return
}

// func generateFileFormData(multiPartWriter *multipart.Writer, key string, value string, name string) error {
// 	// fileReader := strings.NewReader("test")

// 	mimeHeader := make(textproto.MIMEHeader)
// 	mimeHeader.Add("Content-Type", "image/png")
// 	mimeHeader.Add("Content-Disposition", fmt.Sprintf("form-data; name=\"%v\"; filename=%v", key, name))
// 	part, createPartErr := multiPartWriter.CreatePart(mimeHeader)

// 	if createPartErr != nil {
// 		return createPartErr
// 	}
// 	_, writeErr := part.Write([]byte("this is an image content"))
// 	if writeErr != nil {
// 		return writeErr
// 	}

// 	// fileWriter, createFormErr := multiPartWriter.CreateFormFile(key, name)
// 	// if createFormErr != nil {
// 	// 	return createFormErr
// 	// }

// 	// _, copyErr := io.Copy(fileWriter, fileReader)
// 	// if copyErr != nil {
// 	// 	return copyErr
// 	// }

// 	return nil
// }

// func generateCommonRequestFormData(multiPartWriter *multipart.Writer, request ImageRequestCommonConfig) error {
// 	writeFieldNErr := multiPartWriter.WriteField("n", strconv.FormatInt(int64(request.N), 10))
// 	writeFieldSizeErr := multiPartWriter.WriteField("size", string(request.Size))
// 	writeFieldResFormatErr := multiPartWriter.WriteField("response_format", string(request.ResponseFormat))
// 	_ = multiPartWriter.WriteField("image", "image")
// 	// writeFieldUserErr := multiPartWriter.WriteField("user", request.User)
// 	if writeFieldNErr != nil || writeFieldSizeErr != nil || writeFieldResFormatErr != nil {
// 		return errors.New("error: multi part writer write field failed")
// 	}
// 	return nil
// }
