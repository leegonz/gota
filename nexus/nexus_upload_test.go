package nexus

import (
	"fmt"
	"testing"
	"time"

	"github.com/bzon/gota/parser"
	"github.com/bzon/ipapk"
)

var nexus = Nexus{
	HostURL:        "http://localhost:8081",
	Username:       "admin",
	Password:       "admin123",
	SiteRepository: "site",
}

func TestNexusUpload(t *testing.T) {
	var testComponent = NexusComponent{
		File:      "../resources/index.html",
		Filename:  "index.html",
		Directory: "go_upload_test",
	}
	uri, err := nexus.NexusUpload(testComponent)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nexus url:", uri)
}

func TestNexusUploadAssets(t *testing.T) {
	type tc struct {
		name    string
		destDir string
		file    string
	}
	tt := []tc{
		{"upload ios assets", "nexus_ios_repo", "../parser/testdata/sample.ipa"},
		{"upload android assets", "nexus_android_repo", "../parser/testdata/sample.apk"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			appInfo, err := ipapk.NewAppParser(tc.file)
			if err != nil {
				t.Fatal(err)
			}
			var app parser.MobileApp
			app.UploadDate = time.Now().Format(time.RFC1123)
			app.AppInfo = appInfo
			app.File = tc.file
			if err := app.GenerateAssets(); err != nil {
				t.Fatal(err)
			}
			assets, err := nexus.NexusUploadAssets(&app, tc.destDir)
			if err != nil {
				t.Fatal(err)
			}
			for _, v := range assets {
				fmt.Println("nexus url:", v)
			}
		})
	}
}
