package sms // import "github.com/kihamo/shadow-sms"

//go:generate goimports -w ./
//go:generate /bin/bash -c "cd components/sms/internal && go-bindata-assetfs -pkg=internal templates/..."
