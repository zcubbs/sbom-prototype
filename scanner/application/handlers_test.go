package application

import "testing"

func TestRunScan(t *testing.T) {
	type args struct {
		image string
	}
	tests := []struct {
		name               string
		args               args
		wantSbomJson       string
		wantErr            bool
		wantScanReportJson string
	}{
		{
			name: "test",
			args: args{
				image: "docker.io/library/alpine:latest",
			},
			wantSbomJson:       "",
			wantErr:            false,
			wantScanReportJson: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSbomJson, gotScanReportJson, err := RunScan(tt.args.image)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunScan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSbomJson != tt.wantSbomJson {
				t.Errorf("RunScan() gotSbomJson = %v, want %v", gotSbomJson, tt.wantSbomJson)
			}
			if gotScanReportJson != tt.wantScanReportJson {
				t.Errorf("RunScan() gotScanReportJson = %v, want %v", gotScanReportJson, tt.wantScanReportJson)
			}
		})
	}
}
