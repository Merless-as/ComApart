package handler

import (
	ComApart "HelpWithComm"
	"HelpWithComm/pkg/repository"
	"strconv"
	"strings"
)

type Num struct {
	r       *repository.Repository
}

func NewNum(r *repository.Repository) *Num {
	return &Num{r: r}
}

func (n *Num) Count(id int, message string) (int, error) {
	stC, err := n.r.SearchCounter(id)
	if err != nil {
		return 0, err
	}

	stP, err := n.r.SearchPrice(id)
	if err != nil {
		return 0, err
	}

	NewS, err := n.GetCount(message)
	if err != nil {
		return 0, err
	}

	SumP, err := n.SumOfPrice(stP)
	if err != nil {
		return 0, err
	}

	sum := SumP + int(float64(NewS.Gas-stC.Gas)*stP.CostGas+
		float64(NewS.Water-stC.Water)*stP.CostWater+float64(NewS.Electricity-stC.Electricity)*stP.CostElectricity)

	return sum, nil
}

func (n *Num) GetCount(message string) (ComApart.SearchCounter, error) {
	text := strings.Split(message, " ")
	gas, err := strconv.Atoi(text[0])
	if err != nil {
		return ComApart.SearchCounter{}, err
	}
	water, err := strconv.Atoi(text[0])
	if err != nil {
		return ComApart.SearchCounter{}, err
	}
	el, err := strconv.Atoi(text[0])
	if err != nil {
		return ComApart.SearchCounter{}, err
	}
	return ComApart.SearchCounter{Gas: gas, Water: water, Electricity: el}, nil
}

func (n *Num) SumOfPrice(price ComApart.SearchPrice) (int, error) {
	return price.Repair + price.Intercom + price.CommunalServices, nil
}
