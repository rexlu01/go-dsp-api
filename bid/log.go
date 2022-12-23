package bid

import (
	openrtb "github.com/mxmCherry/openrtb/openrtb2"
)

func makeRequestLog(r *openrtb.BidRequest, br *openrtb.BidResponse) map[string]interface{} {
	// TODO: add more request data (User, Device, etc...)
	return map[string]interface{}{
		"request_id": r.ID,
		"cur":        br.Cur,
		"bid_num":    len(br.SeatBid[0].Bid),
	}
}

func makeBidLog(bid *openrtb.Bid) map[string]interface{} {
	return map[string]interface{}{
		"bid_id":      bid.ID,
		"imp_id":      bid.ImpID,
		"price":       bid.Price,
		"ad_id":       bid.AdID,
		"campaign_id": bid.CID,
		"creative_id": bid.CrID,
	}
}
