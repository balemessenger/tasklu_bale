package internal

type Bale struct {
	baseUrl string
}

func NewBale() *Bale {
	return &Bale{
		baseUrl: "",
	}
}

func (b *Bale) send(text string) error {
	return nil
}
