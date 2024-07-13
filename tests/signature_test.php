<?php

function test_signature(string $token, string $timestamp, string $nonce, string $msg) : string
{
    $data = [
        $token,
        $timestamp,
        $nonce,
        $msg,
    ];
    sort($data, SORT_STRING);
    return sha1(join('', $data));
}
$token     = "A1AB219F3547AA7BC57B78EE89768CC1";
$timestamp = "1720421012";
$nonce     = "797";
$msg       = "{\"appid\":\"tt07e3715e98c9aac0\",\"cp_orderno\":\"out_order_no_1\",\"cp_extra\":\"\",\"way\":\"2\",\"payment_order_no\":\"2021070722001450071438803941\",\"total_amount\":9980,\"status\":\"SUCCESS\",\"seller_uid\":\"69631798443938962290\",\"extra\":\"null\",\"item_id\":\"\",\"order_id\":\"N71016888186626816\"}";
echo test_signature($token, $timestamp, $nonce, $msg) . PHP_EOL;