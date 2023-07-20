package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Log struct {
	Msg      string
	LoggedAt string
}

func (l Log) showLog() {
	fmt.Printf("Message: %s, At: %v\n", l.Msg, l.LoggedAt)
}
func (l Log) getUnixTime() int64 {
	layout := "01/02/2006 15:04:05"
	datetime, err := time.Parse(layout, l.LoggedAt)
	if err != nil {
		panic(err)
	}
	unixTimestamp := datetime.Unix()
	return unixTimestamp
}
func (l *Log) MarshalJsontToFile(w io.Writer) []byte {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	type Allias Log
	/*
		 struct{
			LoggedAt int64
			Allias
		 }
		 This will populate all the fields inside Allias
	*/
	/*
		struct{
			LoggedAt int64
			Allias Allias
		}
		This will create a field of type Allias
	*/
	customLog := &struct {
		LoggedAt int64
		Allias   // populating remaining fields
	}{
		LoggedAt: l.getUnixTime(),
		Allias:   (Allias)(*l),
	}
	if bs, er := json.Marshal(customLog); er != nil {
		panic(er)
	} else if _, er := w.Write(bs); er != nil {
		panic(er)
	} else {
		fmt.Printf("Wrote:%s\n", string(bs))
		return bs
	}
}

// This current implementation will go into infinite loop
// although we constructed a annoymous struct but it has Log type in it
// meaning it'll have all methods of Log type
// and when we marshal any type.. the marshaller will check if the type has MarshalJSON method or not
// if it has then it calls it to get the json response
// here it'll call itself again and again casusing infinite loop
// func (l *Log) MarshalJSON() ([]byte, error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println(r)
// 		}
// 	}()
// 	bs, er := json.Marshal(&struct {
// 		LoggedAt int64
// 		*Log
// 	}{
// 		LoggedAt: l.getUnixTime(),
// 		Log:      l,
// 	})
// 	if er != nil {
// 		panic(er)
// 	}
// 	return bs, nil
// }

// type Allias Log // NOTE this will not copy methods present in Log
// this method will not go into a infinite loop
func (l *Log) MarshalJSON() ([]byte, error) {
	type AlliasLog Log
	customLog := &struct {
		LoggedAt int64
		*AlliasLog
	}{
		LoggedAt:  l.getUnixTime(),
		AlliasLog: (*AlliasLog)(l),
	}
	return json.Marshal(customLog)
}
func main() {
	log := Log{
		Msg:      "Error occured at line 89",
		LoggedAt: "12/11/2023 11:09:49",
	}

	bs, _ := log.MarshalJSON()
	fmt.Println("JSON: " + string(bs))
	// now see this
	bs2, _ := json.Marshal(log)
	fmt.Println("JSON: " + string(bs2)) // it will not call MarsalJSON internally
	// Recall how marshal works..It goes in each type and check for MarshalJSON method,
	// if its there then it calls otherwise it dont
	// in our log struct, non of the fields have MarshalJSON embeded
	type customLog struct {
		Log
	}
	myCustomLog := &customLog{
		Log: log,
	}
	bs3, _ := json.Marshal(myCustomLog)
	fmt.Println("JSON: " + string(bs3)) // this will work as expected
}

/*
	Marshal traverses the value v recursively.
	If an encountered value implements the Marshaler interface
	and is not a nil pointer, Marshal calls its MarshalJSON method
	to produce JSON.
	Marshaler is the interface implemented by types that
	can marshal themselves into valid JSON.
	type Marshaler interface {
		MarshalJSON() ([]byte, error)
	}
	bs, _ := log.MarshalJSON()
	fmt.Println(string(bs))

	Similarly we can do for Unmarshall as well
	Unmarshaler is the interface implemented by types
	that can unmarshal a JSON description of themselves.
	The input can be assumed to be a valid encoding of
	a JSON value. UnmarshalJSON must copy the JSON data
	if it wishes to retain the data after returning.

	By convention, to approximate the behavior of Unmarshal itself,
	Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
	type Unmarshaler interface {
		UnmarshalJSON([]byte) error
	}
*/
