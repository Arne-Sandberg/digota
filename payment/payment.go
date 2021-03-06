// Digota <http://digota.com> - eCommerce microservice
// Copyright (c) 2018 Yaron Sumel <yaron@digota.com>
//
// MIT License
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package payment

import (
	"github.com/digota/digota/payment/paymentpb"
	"google.golang.org/grpc"
	"regexp"
)

const baseMethod = "^(.paymentpb.PaymentService/)"

var s Interface

// Interface defines the functionality of the payment service
type Interface interface {
	paymentpb.PaymentServiceServer
}

// RegisterService register p as the service provider
func RegisterService(p Interface) {
	s = p
}

// RegisterPaymentServer register service to the grpc server
func RegisterPaymentServer(server *grpc.Server) {
	paymentpb.RegisterPaymentServiceServer(server, Service())
}

// Service return the registered service
func Service() Interface {
	if s == nil {
		panic("PaymentService is not registered")
	}
	return s
}

// ReadMethods returns regexp slice of readable methods, mostly used by the acl
func ReadMethods() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(baseMethod + "List"),
		regexp.MustCompile(baseMethod + "Get"),
	}
}

// WriteMethods returns regexp slice of writable methods, mostly used by the acl
func WriteMethods() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(baseMethod + "NewCharge"),
		regexp.MustCompile(baseMethod + "RefundCharge"),
	}
}
