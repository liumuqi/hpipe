/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Test for Flow
 *
 * @file ast_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 10 19:50:59 2014
 *
 **/

package ast

import (
	"./parser"
	"fmt"
	"testing"
)

func TestParseXML(t *testing.T) {
	p := parser.NewXMLParser()
	f, err := p.ParseFile("flow1.xml", "../test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f.DebugString())
}
