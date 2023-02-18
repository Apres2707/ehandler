package ehandler

import (
	"errors"
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
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
			name: "hide system with std golang error",
			args: args{
				err:        errors.New(incomingErrorMsg),
				showSystem: false,
			},
			wantCode: systemError,
			wantMsg:  systemErrorMsg,
		},
		{
			name: "hide system with process error",
			args: args{
				err:        NewProcessError(incomingErrorMsg, nil),
				showSystem: false,
			},
			wantCode: processError,
			wantMsg:  incomingErrorMsg,
		},
		{
			name: "hide system with wrapped process error",
			args: args{
				err:        fmt.Errorf("original error: %w", NewProcessError(incomingErrorMsg, nil)),
				showSystem: false,
			},
			wantCode: systemError,
			wantMsg:  systemErrorMsg,
		},
		{
			name: "hide system with wrapped std error in process error",
			args: args{
				err: NewProcessError(incomingErrorMsg, NewProcessError(incomingErrorMsg,
					errors.New("std error"))),
				showSystem: false,
			},
			wantCode: systemError,
			wantMsg:  systemErrorMsg,
		},
		{
			name: "show system with wrapped process error",
			args: args{
				err:        fmt.Errorf("original error: %w", NewProcessError(incomingErrorMsg, nil)),
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
