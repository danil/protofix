// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package testspeclogon is a format of the test FIX.0.1 Logon message.
package testspeclogon

import (
	"time"

	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	TESTSPECLogonMarshaler   = marshfix.Marshal{Tag: "TESTSPEC", Format: TESTSPECLogon}
	TESTSPECLogonUnmarshaler = marshfix.Unmarshal{Tag: "TESTSPEC", Format: TESTSPECLogon}
)

// TESTSPECLogon is a test FIX.0.1 format of the Logon message which maps the codecs into individual fields.
var TESTSPECLogon = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:           f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.0.1"), f0.Con{7}},
		MsgType35:              f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "5" /* LOGOUT */, "A" /* LOGON */), f0.Var{1, 2}},
		SenderCompID49:         f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:         f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MsgSeqNum34:            f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		PossDupFlag43:          f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:           f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:          f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		OrigSendingTime122:     f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		AgreementDesc913:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		YieldType235:           f0.Fld{Opt, f0.ASCII, f0.String("AFTERTAX" /* AFTERTAXYIELD */), f0.Var{1, 65536}},
		UnderlyingStipType888:  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		HopCompID628:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncryptMethod98:        f0.Fld{Req, f0.ASCII, f0.IntDefault(0 /* NONEOTHER */), f0.Con{1}},
		HeartBtInt108:          f0.Fld{Req, f0.ASCII, f0.SecondsDuration(1*time.Second, 2*time.Second, 3*time.Second, 4*time.Second, 5*time.Second, 6*time.Second, 7*time.Second, 8*time.Second, 9*time.Second, 10*time.Second, 11*time.Second, 12*time.Second, 13*time.Second, 14*time.Second, 15*time.Second, 16*time.Second, 17*time.Second, 18*time.Second, 19*time.Second, 20*time.Second, 21*time.Second, 22*time.Second, 23*time.Second, 24*time.Second, 25*time.Second, 26*time.Second, 27*time.Second, 28*time.Second, 29*time.Second, 30*time.Second, 31*time.Second, 32*time.Second, 33*time.Second, 34*time.Second, 35*time.Second, 36*time.Second, 37*time.Second, 38*time.Second, 39*time.Second, 40*time.Second, 41*time.Second, 42*time.Second, 43*time.Second, 44*time.Second, 45*time.Second, 46*time.Second, 47*time.Second, 48*time.Second, 49*time.Second, 50*time.Second, 51*time.Second, 52*time.Second, 53*time.Second, 54*time.Second, 55*time.Second, 56*time.Second, 57*time.Second, 58*time.Second, 59*time.Second, 60*time.Second), f0.Var{1, 18}},
		ResetSeqNumFlag141:     f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		Password554:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		NewPassword925:         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SessionStatus1409:      f0.Fld{Opt, f0.ASCII, f0.IntDefault(0, 3), f0.Var{1, 19}},
		CancelOnDisconnect6867: f0.Fld{Opt, f0.ASCII, f0.String("A"), f0.Var{1, 65536}},
		LanguageID6936:         f0.Fld{Opt, f0.ASCII, f0.String("R" /* RUSSIAN */, "E" /* ENGLISH */), f0.Var{1, 65536}},
		PosType703:             f0.Fld{Opt, f0.ASCII, f0.String("ALC" /* ALLOCATION_TRADE_QTY */, "AS" /* OPTION_ASSIGNMENT */), f0.Var{1, 65536}},
		NestedPartyID524:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		NestedPartySubID545:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,           // STRING
		BodyLength9,            // LENGTH
		MsgType35,              // STRING
		SenderCompID49,         // STRING
		TargetCompID56,         // STRING
		MsgSeqNum34,            // SEQNUM
		PossDupFlag43,          // BOOLEAN
		PossResend97,           // BOOLEAN
		SendingTime52,          // UTCTIMESTAMP
		OrigSendingTime122,     // UTCTIMESTAMP
		AgreementDesc913,       // STRING
		YieldType235,           // STRING
		UnderlyingStipType888,  // STRING
		HopCompID628,           // STRING
		EncryptMethod98,        // INT
		HeartBtInt108,          // INT
		ResetSeqNumFlag141,     // BOOLEAN
		Password554,            // STRING
		NewPassword925,         // STRING
		SessionStatus1409,      // INT
		CancelOnDisconnect6867, // STRING
		LanguageID6936,         // STRING
		PosType703,             // STRING
		NestedPartyID524,       // STRING
		NestedPartySubID545,    // STRING
		CheckSum10,             // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8           = 8    // STRING
	BodyLength9            = 9    // LENGTH
	MsgType35              = 35   // STRING
	SenderCompID49         = 49   // STRING
	TargetCompID56         = 56   // STRING
	MsgSeqNum34            = 34   // SEQNUM
	PossDupFlag43          = 43   // BOOLEAN
	PossResend97           = 97   // BOOLEAN
	SendingTime52          = 52   // UTCTIMESTAMP
	OrigSendingTime122     = 122  // UTCTIMESTAMP
	AgreementDesc913       = 913  // STRING
	YieldType235           = 235  // STRING
	UnderlyingStipType888  = 888  // STRING
	HopCompID628           = 628  // STRING
	EncryptMethod98        = 98   // INT
	HeartBtInt108          = 108  // INT
	ResetSeqNumFlag141     = 141  // BOOLEAN
	Password554            = 554  // STRING
	NewPassword925         = 925  // STRING
	SessionStatus1409      = 1409 // INT
	CancelOnDisconnect6867 = 6867 // STRING
	LanguageID6936         = 6936 // STRING
	PosType703             = 703  // STRING
	NestedPartyID524       = 524  // STRING
	NestedPartySubID545    = 545  // STRING
	CheckSum10             = 10   // STRING
)
