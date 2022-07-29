package handler

import (
	"eureka/model"
	"reflect"
	"testing"
)

func TestRunAdvisoryLayerSet(t *testing.T) {
	type args struct {
		input *model.AdvisoryQueryModel
	}
	tests := []struct {
		name string
		args args
		want *model.AdvisoryResponseModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunAdvisoryLayerSet(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunAdvisoryLayerSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetASGeojson(t *testing.T) {
	type args struct {
		layer    string
		location string
	}
	tests := []struct {
		name    string
		args    args
		want    *map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetASGeojson(tt.args.layer, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetASGeojson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetASGeojson() = %v, want %v", got, tt.want)
			}
		})
	}
}
