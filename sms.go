package sms // import "github.com/kihamo/shadow-sms"

//go:generate /bin/bash -c "find components/sms/internal/locales/ -name \\*.po -execdir /bin/bash -c 'msgfmt {} -o `basename {} .po`.mo' '{}' \\;"
//go:generate /bin/bash -c "goimports -w `find . -type f -name '*.go' -not -path './vendor/*'`"
//go:generate /bin/bash -c "cd components/sms/internal && go-bindata-assetfs -ignore='(.*?[.]po$)' -pkg=internal templates/... locales/..."
