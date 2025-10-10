package service

import (
	"IbtService/internal/config"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type ExternalService interface {
	Send(req interface{}) (interface{}, error)
}

type externalService struct {
	client *http.Client
	cfg    *config.Config
}

func NewExternalService(client *http.Client, cfgLoad *config.Config) ExternalService {
	return &externalService{
		client: client,
		cfg:    cfgLoad,
	}
}

func (s *externalService) Send(reqData interface{}) (interface{}, error) {
	xmlBytes, err := xml.MarshalIndent(reqData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal xml error: %w", err)
	}

	req, err := http.NewRequest("POST", s.cfg.UrlRrebqin, bytes.NewBuffer(xmlBytes))
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

	return body, nil
}
