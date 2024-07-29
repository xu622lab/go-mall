/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 10:07:17
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 11:07:31
 * @FilePath: /go-mall/service/address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"form"`
}

func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Show(ctx context.Context, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(uint(addressId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) Delete(ctx context.Context, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	_, err := addressDao.GetAddressByAid(uint(addressId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = addressDao.DeleteAddressByAid(uint(addressId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.UpdateAddressByUserId(uint(addressId), address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)

	addressList, err := addressDao.ListAddressByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddreses(addressList),
	}
}
