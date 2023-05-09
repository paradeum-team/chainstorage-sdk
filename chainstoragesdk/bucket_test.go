package chainstoragesdk

import (
	"chainstorage-sdk/model"
	"reflect"
	"testing"
)

func TestBucket_CreateBucket(t *testing.T) {
	type fields struct {
		Config *Configuration
		Client *RestyClient
	}
	type args struct {
		bucketName          string
		storageNetworkCode  int
		bucketPrincipleCode int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.BucketCreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bucket{
				Config: tt.fields.Config,
				Client: tt.fields.Client,
			}
			got, err := b.CreateBucket(tt.args.bucketName, tt.args.storageNetworkCode, tt.args.bucketPrincipleCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucket_EmptyBucket(t *testing.T) {
	type fields struct {
		Config *Configuration
		Client *RestyClient
	}
	type args struct {
		bucketId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.BucketEmptyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bucket{
				Config: tt.fields.Config,
				Client: tt.fields.Client,
			}
			got, err := b.EmptyBucket(tt.args.bucketId)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmptyBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmptyBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucket_GetBucketList(t *testing.T) {
	type fields struct {
		Config *Configuration
		Client *RestyClient
	}
	type args struct {
		bucketName string
		pageSize   int
		pageIndex  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.BucketPageResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bucket{
				Config: tt.fields.Config,
				Client: tt.fields.Client,
			}
			got, err := b.GetBucketList(tt.args.bucketName, tt.args.pageSize, tt.args.pageIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBucketList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBucketList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucket_RemoveBucket(t *testing.T) {
	type fields struct {
		Config *Configuration
		Client *RestyClient
	}
	type args struct {
		bucketId            int
		autoEmptyBucketData bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.BucketRemoveResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bucket{
				Config: tt.fields.Config,
				Client: tt.fields.Client,
			}
			got, err := b.RemoveBucket(tt.args.bucketId, tt.args.autoEmptyBucketData)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkBucketName(t *testing.T) {
	type args struct {
		bucketName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkBucketName(tt.args.bucketName); (err != nil) != tt.wantErr {
				t.Errorf("checkBucketName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkBucketPrincipleCode(t *testing.T) {
	type args struct {
		bucketPrincipleCode int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkBucketPrincipleCode(tt.args.bucketPrincipleCode); (err != nil) != tt.wantErr {
				t.Errorf("checkBucketPrincipleCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkStorageNetworkCode(t *testing.T) {
	type args struct {
		storageNetworkCode int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkStorageNetworkCode(tt.args.storageNetworkCode); (err != nil) != tt.wantErr {
				t.Errorf("checkStorageNetworkCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
