package main

import (
	"reflect"
	"testing"
)

func Test_saveAndLoad(t *testing.T) {
	t.Run("Test that load and store works", func(t *testing.T) {
		page := &Page{"TestPage", []byte("this is a test")}

		err := page.save()
		if err != nil {
			t.Errorf("Error saving test page: %v", err)
		}

		pageLoaded, err := loadPage(page.Title)
		if err != nil || !reflect.DeepEqual(page, pageLoaded) {
			t.Errorf("Error: %v\nSaved: %v\nRead: %v", err, page, pageLoaded)
		}
	})
}

func Test_loadPage(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		want    *Page
		wantErr bool
	}{
		{"Load page that doesn't exist", args{"TestPage2"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadPage(tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
