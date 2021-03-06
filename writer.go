/* Copyright 2018 Ron Zhang <ronzxy@mx.aketi.cn>. All rights reserved.
 *
 * Licensed under the Apache License, version 2.0 (the "License").
 * You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package logger

import (
	xormlog "github.com/ronzxy/go-xorm/log"
)

// logger writer interface
type Writer interface {
	Tracef(format string, args ...interface{})

	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	FatalfWithExit(exit bool, format string, args ...interface{})

	Trace(args ...interface{})

	Debug(args ...interface{})

	Info(args ...interface{})

	Warn(args ...interface{})

	Error(args ...interface{})

	FatalWithExit(exit bool, args ...interface{})

	CheckRollingSize()

	/*
	   Include xorm logger
	*/

	BeforeSQL(context xormlog.LogContext) // only invoked when IsShowSQL is true

	AfterSQL(context xormlog.LogContext)  // only invoked when IsShowSQL is true

	Level() xormlog.LogLevel

	SetLevel(l xormlog.LogLevel)

	ShowSQL(show ...bool)

	IsShowSQL() bool
}
