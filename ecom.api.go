// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.
// source: douyin-miniapp/ecom.proto

package douyin_miniapp

import (
	context "context"
	fmt "fmt"
	io "io"
	json "encoding/json"
	bytes "bytes"
	http "net/http"
	strings "strings"
	url "net/url"
	multipart "mime/multipart"
)
// Reference imports to suppress errors if they are not otherwise used.
var _ = context.Background
var _ = http.NewRequest
var _ = io.Copy
var _ = bytes.Compare
var _ = json.Marshal
var _ = strings.Compare
var _ = fmt.Errorf
var _ = url.Parse
var _ = multipart.ErrMessageTooLarge


// Client API for Ecom service

type EcomService interface {
	// RegisterInfo  注册小程序预览图 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/e-commerce/registermicroappcoverfigure
	RegisterInfo(ctx context.Context, in *RegisterInfoReq, opts ...Option) (*RegisterInfoRes, error)
	// PointLimitRegister  注册小程序积分阈值 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/e-commerce/register-credits-threshold
	PointLimitRegister(ctx context.Context, in *PointLimitRegisterReq, opts ...Option) (*PointLimitRegisterRes, error)
	// GetCustomMadeGoodsStatus  查询订单的定制完成状态 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/e-commerce/customGoodsOrderDetail
	GetCustomMadeGoodsStatus(ctx context.Context, in *GetCustomMadeGoodsStatusReq, opts ...Option) (*GetCustomMadeGoodsStatusRes, error)
	// LeaveMember  退会 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/e-commerce/withdraw
	LeaveMember(ctx context.Context, in *LeaveMemberReq, opts ...Option) (*LeaveMemberRes, error)
}

type ecomService struct {
	// opts
	opts *Options
}

func NewEcomService(opts ...Option) EcomService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://ma.zijieapi.com"
	}
	return &ecomService {
		opts: opt,
	}
}


func (c *ecomService) RegisterInfo(ctx context.Context, in *RegisterInfoReq, opts ...Option) (*RegisterInfoRes, error) {
	var res RegisterInfoRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/api/apps/industry/ecom/register_info", opt.addr)

	// body
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *ecomService) PointLimitRegister(ctx context.Context, in *PointLimitRegisterReq, opts ...Option) (*PointLimitRegisterRes, error) {
	var res PointLimitRegisterRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/api/apps/industry/ecom/point/limit/register", opt.addr)

	// body
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *ecomService) GetCustomMadeGoodsStatus(ctx context.Context, in *GetCustomMadeGoodsStatusReq, opts ...Option) (*GetCustomMadeGoodsStatusRes, error) {
	var res GetCustomMadeGoodsStatusRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/api/apps/industry/ecom/get_custom_made_goods_status", opt.addr)

	// body
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}

func (c *ecomService) LeaveMember(ctx context.Context, in *LeaveMemberReq, opts ...Option) (*LeaveMemberRes, error) {
	var res LeaveMemberRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/api/apps/industry/ecom/v1/shop/leave_member", opt.addr)

	// body
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}
