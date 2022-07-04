// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

func TestUpsert(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsUpsert)

	t.Run("Clients", testClientsUpsert)

	t.Run("ConnectionEvents", testConnectionEventsUpsert)

	t.Run("HolePunchAttempts", testHolePunchAttemptsUpsert)

	t.Run("HolePunchResults", testHolePunchResultsUpsert)

	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesUpsert)

	t.Run("IPAddresses", testIPAddressesUpsert)

	t.Run("MultiAddresses", testMultiAddressesUpsert)

	t.Run("MultiAddressesSets", testMultiAddressesSetsUpsert)

	t.Run("PeerLogs", testPeerLogsUpsert)

	t.Run("Peers", testPeersUpsert)
}
