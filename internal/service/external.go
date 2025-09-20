package service

import (
	"IbtService/internal/domain"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type ExternalService interface {
	Send(req *domain.Request) (*domain.Response, error)
}

type externalService struct {
	client *http.Client
}

func NewExternalService(client *http.Client) ExternalService {
	return &externalService{client: client}
}

func (s *externalService) Send(reqData *domain.Request) (*domain.Response, error) {
	xmlBytes, err := xml.MarshalIndent(reqData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal xml error: %w", err)
	}

	req, err := http.NewRequest("POST", "https://reqbin.com/echo/post/xml", bytes.NewBuffer(xmlBytes))
	if err != nil {
		return nil, fmt.Errorf("new request error: %w", err)
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("User-Agent", "Go-Client")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %w", err)
	}

	xmlResp := &domain.Response{}
	if err := xml.Unmarshal(body, xmlResp); err != nil {
		return nil, fmt.Errorf("unmarshal xml error: %w", err)
	}

	return xmlResp, nil
}
