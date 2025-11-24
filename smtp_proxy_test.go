package gomail

import (
	"errors"
	"net"
	"testing"
	"time"
)

// TestDialerWithInvalidProxy tests that the dialer handles invalid proxy gracefully
func TestDialerWithInvalidProxy(t *testing.T) {
	d := NewDialer(testHost, testPort, "user", "pwd")
	d.SetSocks5Proxy("invalid-proxy-host:9999")

	_, err := d.Dial()
	if err == nil {
		t.Error("Expected error when dialing with invalid proxy, got nil")
	}
}

// TestDialerWithNilConnection tests that the dialer handles nil connection gracefully
func TestDialerWithNilConnection(t *testing.T) {
	d := NewDialer(testHost, testPort, "user", "pwd")
	d.SetSocks5Proxy("test-proxy:1080")

	// Mock the netDialTimeout to return nil connection but no error
	// This simulates a buggy proxy implementation
	originalNetDialTimeout := netDialTimeout
	defer func() { netDialTimeout = originalNetDialTimeout }()

	// We can't easily mock proxy.SOCKS5 dial, so let's test the regular dial path
	d.Socks5Proxy = "" // Reset proxy to test regular path

	netDialTimeout = func(network, address string, timeout time.Duration) (net.Conn, error) {
		// Return nil connection with nil error - this should not happen in real code
		// but we want to be defensive
		return nil, nil
	}

	_, err := d.Dial()
	if err == nil {
		t.Error("Expected error when connection is nil, got nil")
	}
	// Check that it's not a panic by having a meaningful error message
	expectedErr := "dial failed: connection is nil"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %q, got %q", expectedErr, err.Error())
	}
}

// TestDialerProxyConnectionFailure tests that connection failure with proxy returns proper error
func TestDialerProxyConnectionFailure(t *testing.T) {
	d := NewDialer(testHost, testPort, "user", "pwd")
	d.SetSocks5Proxy("127.0.0.1:9999") // Non-existent proxy

	_, err := d.Dial()
	if err == nil {
		t.Error("Expected error when proxy is unreachable, got nil")
	}
}

// TestDialerRegularConnectionFailure tests that connection failure without proxy returns proper error
func TestDialerRegularConnectionFailure(t *testing.T) {
	d := NewDialer("invalid-host-that-does-not-exist.example.com", 9999, "user", "pwd")

	_, err := d.Dial()
	if err == nil {
		t.Error("Expected error when host is unreachable, got nil")
	}
}

// TestDialerSSLWithNilConnection tests that SSL connection handles nil gracefully
func TestDialerSSLWithNilConnection(t *testing.T) {
	d := NewDialer(testHost, testSSLPort, "user", "pwd")
	d.SSL = true

	// Mock the netDialTimeout to return nil connection but no error
	originalNetDialTimeout := netDialTimeout
	defer func() { netDialTimeout = originalNetDialTimeout }()

	netDialTimeout = func(network, address string, timeout time.Duration) (net.Conn, error) {
		return nil, nil
	}

	_, err := d.Dial()
	if err == nil {
		t.Error("Expected error when connection is nil with SSL, got nil")
	}
}

// TestDialerWithErrorConnection tests that errors from dial are properly propagated
func TestDialerWithErrorConnection(t *testing.T) {
	d := NewDialer(testHost, testPort, "user", "pwd")

	// Mock the netDialTimeout to return an error
	originalNetDialTimeout := netDialTimeout
	defer func() { netDialTimeout = originalNetDialTimeout }()

	expectedError := errors.New("connection refused")
	netDialTimeout = func(network, address string, timeout time.Duration) (net.Conn, error) {
		return nil, expectedError
	}

	_, err := d.Dial()
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}
