package main

import (
    "testing"
)

func Test_hello_1(t *testing.T){
    var a=8
    var b=8
    if a==b{
        t.Log("hello test pass")
    }else{
        t.Error("hello test fail")
    }
}
