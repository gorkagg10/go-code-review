package service

import (
	repo "coupon_service/pkg/repository"
	"coupon_service/pkg/service/entity"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		repo repo.Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"initialize service", args{repo: *repo.NewRepository()}, Service{repository: repo.NewRepository()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(&tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ApplyCoupon(t *testing.T) {
	type fields struct {
		repo repo.Repository
	}
	type args struct {
		basket entity.Basket
		code   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   *entity.Basket
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repository: &tt.fields.repo,
			}
			gotB, err := s.ApplyCoupon(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ApplyCoupon() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	type fields struct {
		repo repo.Repository
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{"Apply 10%", fields{*repo.NewRepository()}, args{10, "Superdiscount", 55}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repository: &tt.fields.repo,
			}

			s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
		})
	}
}
