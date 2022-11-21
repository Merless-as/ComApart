package handler

import (
	ComApart "HelpWithComm"
	"HelpWithComm/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Num struct {
	r repository.Repository
	message tgbotapi.Message
}

func NewNum(r repository.Repository, message tgbotapi.Message) *Num {
	return &Num{r: r, message: message}
}

func (n *Num) Count (id int) (int, error) {
	stC, err := n.r.SearchCounter(id)
	if err != nil {
		return 0, err
	}

	stP, err := n.r.SearchPrice(id)
	if err != nil {
		return 0, err
	}

	NewS, err := n.GetCount(n.message)
	if err != nil {
		return 0, err
	}

	SumP, err := n.SumOfPrice(stP)
	if err != nil {
		 return 0, err
	}

	sum := SumP + int(float64(NewS.Gas-stC.Gas) * stP.CostGas +
		float64(NewS.Water-stC.Water) * stP.CostWater + float64(NewS.Electricity-stC.Electricity) * stP.CostElectricity)

	return sum, nil
}

func (n *Num) GetCount(message tgbotapi.Message) (ComApart.SearchCounter, error)  {
	return ComApart.SearchCounter{}, nil
}

func (n *Num) SumOfPrice(price ComApart.SearchPrice) (int, error) {
	return 0, nil
}