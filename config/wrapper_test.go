package config

import (
	"testing"

	_kiosk "github.com/immobiliare/peephole/kiosk"
	_mold "github.com/immobiliare/peephole/mold"
	_spy "github.com/immobiliare/peephole/spy"
)

var (
	spy = _spy.Config{
		API:    "http://localhost:8080",
		User:   "user",
		Pass:   "pass",
		Client: "pam",
	}
	kiosk = _kiosk.Config{
		Bind:      ":8080",
		BasicAuth: map[string]string{},
	}
	mold = _mold.Config{
		Spool:     "/var/spool/peephole",
		Retention: "7d",
	}
)

func TestParseValid(t *testing.T) {
	cfg, err := parseYaml([]byte(`
spy:
  - api: ` + spy.API + `
    user: ` + spy.User + `
    pass: ` + spy.Pass + `
    client: ` + spy.Client + `
kiosk:
  bind: ` + kiosk.Bind + `
mold:
  spool: ` + mold.Spool + `
  retention: ` + mold.Retention,
	))

	if err != nil {
		t.Errorf("Supplied config is not supposed to return an error")
	}

	cond := len(cfg.Spy) == 1
	cond = cond && cfg.Spy[0].API == spy.API
	cond = cond && cfg.Spy[0].User == spy.User
	cond = cond && cfg.Spy[0].Pass == spy.Pass
	cond = cond && cfg.Spy[0].Client == spy.Client
	cond = cond && cfg.Kiosk.Bind == kiosk.Bind
	cond = cond && len(cfg.Kiosk.BasicAuth) == 0
	cond = cond && cfg.Mold.Spool == mold.Spool
	cond = cond && cfg.Mold.Retention == mold.Retention
	if !cond {
		t.Errorf("Parsed config is not corresponding to supplied data")
	}
}

func TestParseEmpty(t *testing.T) {
	if _, err := parseYaml([]byte("")); err == nil {
		t.Errorf("Empty config is supposed to return an error")
	}
}

func TestParseUnexisting(t *testing.T) {
	if _, err := Parse("/tmp/not-existing-file"); err == nil {
		t.Errorf("Non-existing config file is supposed to not return a valid config instance")
	}
}
