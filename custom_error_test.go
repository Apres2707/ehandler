package ehandler

import (
	"fmt"
	"testing"
)

const customErrorCode = "customCode"

func TestCustomError(t *testing.T) {
	const incomingErrorMsg = "some error"

	type args struct {
		err        error
		showSystem bool
	}
	tests := []struct {
		name     string
		args     args
		wantCode string
		wantMsg  string
	}{
		{
			name: "hide system with custom error",
			args: args{
				err: customError{
					msg: incomingErrorMsg,
				},
				showSystem: false,
			},
			wantCode: customErrorCode,
			wantMsg:  incomingErrorMsg,
		},
		{
			name: "hide system with wrapped custom error",
			args: args{
				err:        fmt.Errorf("original error: %w", customError{msg: incomingErrorMsg}),
				showSystem: false,
			},
			wantCode: systemError,
			wantMsg:  systemErrorMsg,
		},
		{
			name: "show system with wrapped custom error",
			args: args{
				err:        fmt.Errorf("original error: %w", customError{msg: incomingErrorMsg}),
				showSystem: true,
			},
			wantCode: systemError,
			wantMsg:  fmt.Sprintf("original error: %s", incomingErrorMsg),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotMsg := Handle(tt.args.err, tt.args.showSystem)
			if gotCode != tt.wantCode {
				t.Errorf("Handle() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("Handle() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

type customError struct {
	msg string
}

func (e customError) Error() string {
	return e.msg
}

func (e customError) Code() string {
	return customErrorCode
}
