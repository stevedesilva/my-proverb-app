package proverb_test

import (
	"testing"

	"github.com/stevedesilva/my-proverb-app/pkg/proverb"
	"github.com/stretchr/testify/assert"
)

func TestProverb_GetProverbs(t *testing.T) {
	p := proverb.New()
	v, err := p.GetProverbs()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, 8, len(v))
}

//func TestProverb_GetProverbs(t *testing.T) {
//	type fields struct {
//		Quote  string
//		Author string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   *[]string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &Proverb{
//				Quote:  tt.fields.Quote,
//				Author: tt.fields.Author,
//			}
//			if got := p.GetProverbs(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetProverbs() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
