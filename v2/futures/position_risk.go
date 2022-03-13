package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetPositionRiskService get account balance
type GetPositionRiskService struct {
	c      *Client
	symbol string
}

type GetPositionRiskResponse struct {
	PositionRisks     []*PositionRisk `json:"positionRisks"`
	RateLimitWeight1m string          `json:"rateLimitWeight1m,omitempty"`
}

// Symbol set symbol
func (s *GetPositionRiskService) Symbol(symbol string) *GetPositionRiskService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res *GetPositionRiskResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	res = new(GetPositionRiskResponse)
	var header *http.Header
	data, header, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res.PositionRisks = make([]*PositionRisk, 0)
	res.RateLimitWeight1m = header.Get("X-Mbx-Used-Weight-1m")

	err = json.Unmarshal(data, &res.PositionRisks)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
	EntryPrice       string `json:"entryPrice"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Leverage         string `json:"leverage"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	PositionSide     string `json:"positionSide"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
}
