package impl

import (
	"context"
	"github.com/ardikapras/owldetect/internal/app/usecase/analyze/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyzeUsecase_DoAnalyze(t *testing.T) {
	t.Run("AnalyzeUsecase", func(t *testing.T) {
		type expectedFields struct {
			matchs []model.Match
		}

		type args struct {
			ctx     context.Context
			payload model.AnalyzeReqBody
		}
		tests := []struct {
			name           string
			expectedFields expectedFields
			args           args
			mocks          func()
			wantErr        bool
		}{
			{
				name: "Test 1: found plagiarism",
				args: args{
					ctx: context.Background(),
					payload: model.AnalyzeReqBody{
						InputText: "Hello World! I'm a cat!",
						RefText:   "Welcome to my world! Hello World! I'm a cat! But I eat pizza!",
					},
				},
				mocks:   nil,
				wantErr: false,
				expectedFields: expectedFields{
					matchs: []model.Match{
						{
							Input: model.MatchDetails{
								Text:     "Hello World! I'm a cat!",
								StartIdx: 0,
								EndIdx:   22,
							},
							Reference: model.MatchDetails{
								Text:     "Hello World! I'm a cat!",
								StartIdx: 21,
								EndIdx:   43,
							},
						},
					},
				},
			},
			{
				name: "Test 2: not found plagiarism",
				args: args{
					ctx: context.Background(),
					payload: model.AnalyzeReqBody{
						InputText: "Hello to my world!",
						RefText:   "Welcome to my world! Hello World! I'm a cat! But I eat pizza!",
					},
				},
				mocks:   nil,
				wantErr: false,
				expectedFields: expectedFields{
					matchs: nil,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if tt.mocks != nil {
					tt.mocks()
				}
				aimpl := New()
				match, err := aimpl.DoAnalyze(tt.args.ctx, tt.args.payload)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.expectedFields.matchs, match)
				}
			})
		}
	})
}
