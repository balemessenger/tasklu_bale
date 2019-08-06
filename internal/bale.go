package internal

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type Bale struct {
	baseUrl string
}

func NewBale(token string) *Bale {
	return &Bale{
		baseUrl: "https://api.bale.ai/v1/webhooks/" + token,
	}
}

func (b *Bale) Send(text string) error {
	reader := strings.NewReader(fmt.Sprintf(`{"text":"%s"}`, text))
	resp, err := http.Post(b.baseUrl, "application/json", reader)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil
}
