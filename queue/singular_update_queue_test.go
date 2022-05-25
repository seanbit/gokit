package queue

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

type TestUpdateRequest struct {
	UserName string `json:"user_name"`
	Remark string `json:"remark"`
}

type TestGiftRecord struct {
	GiftName string `json:"gift_name"`
	Receiver string `json:"receiver"`
	Remark string `json:"remark"`
	CanReceive bool `json:"can_receive"`
}

func testGifts(count int) []*TestGiftRecord {
	var gifts []*TestGiftRecord
	for idx := 1; idx <= count; idx ++ {
		gifts = append(gifts, &TestGiftRecord{
			GiftName:   fmt.Sprintf("Gift-%04d", idx),
			Receiver:   "",
			Remark:     "",
			CanReceive: true,
		})
	}
	return gifts
}

func testUpdateRequests(count int) []*TestUpdateRequest {
	var updateRequests []*TestUpdateRequest
	for idx := 1; idx <= count; idx ++ {
		updateRequests = append(updateRequests, &TestUpdateRequest{
			UserName: fmt.Sprintf("User-%04d", idx),
			Remark:   fmt.Sprintf("User-%04d-Remark", idx),
		})
	}
	return updateRequests
}

func TestSingularUpdateQueue_Push(t *testing.T) {

	gifts := testGifts(1000000)
	requests := testUpdateRequests(1000000)

	var filterCanReceiveGift = func(gifts []*TestGiftRecord) int {
		for idx := 0; idx < len(gifts); idx ++ {
			if !gifts[idx].CanReceive {
				continue
			}
			return idx
		}
		return -1
	}

	queue := NewSingularUpdateQueue(1000000, func(request interface{}) UpdateResponse {
		req := request.(*TestUpdateRequest)
		resp := UpdateResponse{}
		findGiftIndex := filterCanReceiveGift(gifts)
		if findGiftIndex == -1 {
			resp.Err = errors.New("gifts received all")
			return resp
		}
		receivedGift := gifts[findGiftIndex]
		receivedGift.CanReceive = false
		receivedGift.Receiver = req.UserName
		receivedGift.Remark = req.Remark
		resp.Val = receivedGift
		return resp
	})

	var wg sync.WaitGroup
	wg.Add(len(requests))
	for idx := 0; idx < len(requests); idx ++ {
		go func(request *TestUpdateRequest) {
			defer wg.Done()
			queue.Push(request)
		}(requests[idx])
	}
	wg.Wait()
	time.Sleep(time.Second)
	t.Log("Finished")
}

func TestSingularUpdateQueue_PushForResponse(t *testing.T) {

	gifts := testGifts(10000)
	requests := testUpdateRequests(10000)

	var filterCanReceiveGift = func(gifts []*TestGiftRecord) int {
		for idx := 0; idx < len(gifts); idx ++ {
			if !gifts[idx].CanReceive {
				continue
			}
			return idx
		}
		return -1
	}

	queue := NewSingularUpdateQueue(10000, func(request interface{}) UpdateResponse {
		req := request.(*TestUpdateRequest)
		resp := UpdateResponse{}
		findGiftIndex := filterCanReceiveGift(gifts)
		if findGiftIndex == -1 {
			resp.Err = errors.New("gifts received all")
			return resp
		}
		receivedGift := gifts[findGiftIndex]
		receivedGift.CanReceive = false
		receivedGift.Receiver = req.UserName
		receivedGift.Remark = req.Remark
		resp.Val = receivedGift
		return resp
	})

	var wg sync.WaitGroup
	wg.Add(len(requests))
	for idx := 0; idx < len(requests); idx ++ {
		go func(request *TestUpdateRequest) {
			resp := queue.PushForResponse(request)
			if resp.Err != nil {
				t.Logf("request for: %s receive failed: %s\n", request.UserName, resp.Err.Error())
			} else {
				gift := resp.Val.(*TestGiftRecord)
				t.Logf("request for: %s receive success: %+v\n", request.UserName, *gift)
			}
			wg.Done()
		}(requests[idx])
	}
	wg.Wait()
	time.Sleep(time.Second)
	t.Log("Finished")
}