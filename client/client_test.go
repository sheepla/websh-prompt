package client

import (
	"testing"
)

func TestNewURL(t *testing.T) {
	u := newURL("shellgei")
	if u != "https://websh.jiro4989.com/api/shellgei" {
		t.Errorf("Invalid URL: %s\n", u)
	}
}

func TestPing(t *testing.T) {
	result, err := Ping()
	if err != nil {
		t.Errorf("Failed to ping to the websh server: %s\n", err)
	}
	if result.Status != "ok" {
		t.Errorf("Invalid status: %s\n", result.Status)
	}
}

func TestPost(t *testing.T) {
	var p Param
	p.Set("echo Hello")
	result, err := Post(p)
	if err != nil {
		t.Errorf("Failed to post data: %s\n", err)
	}
	if result.Stdout != "Hello\n" {
		t.Errorf("Failed to post data.\nstdout: %s\n stderr: %s\n", result.Stdout, result.Stderr)
	}
}
