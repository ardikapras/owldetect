package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyzeReqBody_Validate(t *testing.T) {
	t.Run("TestAnalyzeReqBody", func(t *testing.T) {
		type expectedFields struct {
		}

		type args struct {
			ctx     context.Context
			reqBody AnalyzeReqBody
		}
		tests := []struct {
			name           string
			expectedFields expectedFields
			args           args
			mocks          func()
			wantErr        bool
		}{
			{
				name: "Test 1: error validation 1 - input text is empty",
				args: args{
					ctx: context.Background(),
					reqBody: AnalyzeReqBody{
						InputText: "",
						RefText:   "ref text",
					},
				},
				mocks:          nil,
				wantErr:        true,
				expectedFields: expectedFields{},
			},
			{
				name: "Test 3: error validation 2 - ref text is empty",
				args: args{
					ctx: context.Background(),
					reqBody: AnalyzeReqBody{
						InputText: "input text",
						RefText:   "",
					},
				},
				mocks:          nil,
				wantErr:        true,
				expectedFields: expectedFields{},
			},
			{
				name: "Test 3: error validation 3 - input longer than ref",
				args: args{
					ctx: context.Background(),
					reqBody: AnalyzeReqBody{
						InputText: "input text",
						RefText:   "ref text",
					},
				},
				mocks:          nil,
				wantErr:        true,
				expectedFields: expectedFields{},
			},
			{
				name: "Test 4: success",
				args: args{
					ctx: context.Background(),
					reqBody: AnalyzeReqBody{
						InputText: "input text",
						RefText:   "ref text but should be longer than input text",
					},
				},
				mocks:          nil,
				wantErr:        false,
				expectedFields: expectedFields{},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if tt.mocks != nil {
					tt.mocks()
				}
				err := tt.args.reqBody.Validate()
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
}
