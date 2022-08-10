package cas

import (
	"testing"
	"time"
)

func TestParseLogoutRequest(t *testing.T) {
	xml := `<samlp:LogoutRequest xmlns:samlp="urn:oasis:names:tc:SAML:2.0:protocol"
  ID="8r7834d6r78346s7823d46678235d" Version="2.0" IssueInstant="Fri, 27 Feb 2015 13:31:34 -0000">
  <saml:NameID xmlns:saml="urn:oasis:names:tc:SAML:2.0:assertion">
    @NOT_USED@
  </saml:NameID>
  <samlp:SessionIndex>ST-io34f34vr7823vcr82346r782c4b78i2364i76cvr72364rv7263</samlp:SessionIndex>
</samlp:LogoutRequest>`

	l, err := ParseLogoutRequest([]byte(xml))
	if err != nil {
		t.Errorf("parseLogoutRequest returned error: %v", err)
	}

	if l.Version != "2.0" {
		t.Errorf("Expected Version to be %q, got %q", "2.0", l.Version)
	}

	if l.ID != "8r7834d6r78346s7823d46678235d" {
		t.Errorf("Expected ID to be %q, got %q",
			"8r7834d6r78346s7823d46678235d", l.ID)
	}

	if l.NameID != "@NOT_USED@" {
		t.Errorf("Expected NameID to be %q, got %q", "@NOT_USED@", l.NameID)
	}

	ticketName := "ST-io34f34vr7823vcr82346r782c4b78i2364i76cvr72364rv7263"
	if l.SessionIndex != ticketName {
		t.Errorf("Expected SessionIndex to be %q, got %q", ticketName, l.SessionIndex)
	}

	instant := time.Date(2015, 02, 27, 13, 31, 34, 0, time.UTC)
	if !instant.Equal(l.IssueInstant) {
		t.Errorf("Expected IssueInstant to be <%v>, got <%v>",
			instant, l.IssueInstant)
	}
}
