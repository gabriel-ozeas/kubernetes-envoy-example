// Code generated by protoc-gen-gogo.
// source: order/order.proto
// DO NOT EDIT!

/*
Package order is a generated protocol buffer package.

It is generated from these files:
	order/order.proto

It has these top-level messages:
	Order
	CreateOrderRequest
	GetOrderRequest
	GetOrderDetailRequest
	GetOrderDetailResponse
	ListOrdersRequest
	ListOrdersResponse
	DeleteOrderRequest
*/
package order

import regexp "regexp"
import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/bakins/kubernetes-envoy-example/api/item"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Order) Validate() error {
	return nil
}

var _regex_CreateOrderRequest_User = regexp.MustCompile("^[A-Za-z0-9]+")

func (this *CreateOrderRequest) Validate() error {
	if !_regex_CreateOrderRequest_User.MatchString(this.User) {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z0-9]+"`, this.User))
	}
	return nil
}

var _regex_GetOrderRequest_Id = regexp.MustCompile("^[A-Za-z0-9]+")

func (this *GetOrderRequest) Validate() error {
	if !_regex_GetOrderRequest_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z0-9]+"`, this.Id))
	}
	return nil
}

var _regex_GetOrderDetailRequest_Id = regexp.MustCompile("^[A-Za-z0-9]+")

func (this *GetOrderDetailRequest) Validate() error {
	if !_regex_GetOrderDetailRequest_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z0-9]+"`, this.Id))
	}
	return nil
}
func (this *GetOrderDetailResponse) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *ListOrdersRequest) Validate() error {
	return nil
}
func (this *ListOrdersResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}

var _regex_DeleteOrderRequest_Id = regexp.MustCompile("^[A-Za-z0-9]+")

func (this *DeleteOrderRequest) Validate() error {
	if !_regex_DeleteOrderRequest_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Za-z0-9]+"`, this.Id))
	}
	return nil
}