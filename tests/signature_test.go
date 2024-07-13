package tests

import (
	`crypto/sha1`
	`fmt`
	`sort`
	`strings`
	`testing`
)

var (
	// timestamp = fmt.Sprintf("%d", time.Now().Unix())
	token     = "A1AB219F3547AA7BC57B78EE89768CC1"
	timestamp = "1720421012"
	nonce     = "797"
	msg       = "{\"appid\":\"tt07e3715e98c9aac0\",\"cp_orderno\":\"out_order_no_1\",\"cp_extra\":\"\",\"way\":\"2\",\"payment_order_no\":\"2021070722001450071438803941\",\"total_amount\":9980,\"status\":\"SUCCESS\",\"seller_uid\":\"69631798443938962290\",\"extra\":\"null\",\"item_id\":\"\",\"order_id\":\"N71016888186626816\"}"
)

func TestSignature(t *testing.T) {
	sortedString := make([]string, 0)
	sortedString = append(sortedString, token)
	sortedString = append(sortedString, timestamp)
	sortedString = append(sortedString, nonce)
	sortedString = append(sortedString, msg)
	sort.Strings(sortedString)
	fmt.Println(strings.Join(sortedString, ""))
	h := sha1.New()
	h.Write([]byte(strings.Join(sortedString, "")))
	bs := h.Sum(nil)
	signature := fmt.Sprintf("%x", bs)
	t.Log(signature)
}
