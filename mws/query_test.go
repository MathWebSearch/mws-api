package mws

import (
	"reflect"
	"testing"

	"github.com/MathWebSearch/mwsapi/utils"
)

func TestQuery_ToXML(t *testing.T) {
	type fields struct {
		From         int64
		Size         int64
		ReturnTotal  utils.BooleanYesNo
		OutputFormat string
		Expressions  []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"empty query", fields{}, "<mws:query limitmin=\"0\" answsize=\"0\" totalreq=\"no\" output=\"\" xmlns:mws=\"http://www.mathweb.org/mws/ns\" xmlns:m=\"http://www.w3.org/1998/Math/MathML\"></mws:query>", false},
		{"simple query", fields{5, 10, true, "xml", []string{"<m:limit/>"}}, "<mws:query limitmin=\"5\" answsize=\"10\" totalreq=\"yes\" output=\"xml\" xmlns:mws=\"http://www.mathweb.org/mws/ns\" xmlns:m=\"http://www.w3.org/1998/Math/MathML\"><mws:expr><m:limit/></mws:expr></mws:query>", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := make([]*Expression, len(tt.fields.Expressions))
			for i, s := range tt.fields.Expressions {
				e[i] = &Expression{Term: s}
			}
			q := &RawQuery{
				From:         tt.fields.From,
				Size:         tt.fields.Size,
				ReturnTotal:  tt.fields.ReturnTotal,
				OutputFormat: tt.fields.OutputFormat,
				Expressions:  e,
			}
			got, err := q.ToXML()
			gotString := string(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query.ToXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotString, tt.want) {
				t.Errorf("Query.ToXML() = %v, want %v", gotString, tt.want)
			}
		})
	}
}
