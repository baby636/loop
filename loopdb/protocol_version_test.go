package loopdb

import (
	"testing"

	looprpc "github.com/lightninglabs/loop/swapserverrpc"
	"github.com/stretchr/testify/require"
)

// TestProtocolVersionSanity tests that protocol versions are sane, meaning
// we always keep our stored protocol version in sync with the RPC protocol
// version except for the unrecorded version.
func TestProtocolVersionSanity(t *testing.T) {
	t.Parallel()

	versions := [...]ProtocolVersion{
		ProtocolVersionLegacy,
		ProtocolVersionMultiLoopOut,
		ProtocolVersionSegwitLoopIn,
		ProtocolVersionPreimagePush,
		ProtocolVersionUserExpiryLoopOut,
		ProtocolVersionHtlcV2,
		ProtocolVersionMultiLoopIn,
		ProtocolVersionLoopOutCancel,
		ProtocolVersionProbe,
		ProtocolVersionRoutingPlugin,
		ProtocolVersionHtlcV3,
	}

	rpcVersions := [...]looprpc.ProtocolVersion{
		looprpc.ProtocolVersion_LEGACY,
		looprpc.ProtocolVersion_MULTI_LOOP_OUT,
		looprpc.ProtocolVersion_NATIVE_SEGWIT_LOOP_IN,
		looprpc.ProtocolVersion_PREIMAGE_PUSH_LOOP_OUT,
		looprpc.ProtocolVersion_USER_EXPIRY_LOOP_OUT,
		looprpc.ProtocolVersion_HTLC_V2,
		looprpc.ProtocolVersion_MULTI_LOOP_IN,
		looprpc.ProtocolVersion_LOOP_OUT_CANCEL,
		looprpc.ProtocolVersion_PROBE,
		looprpc.ProtocolVersion_ROUTING_PLUGIN,
		looprpc.ProtocolVersion_HTLC_V3,
	}

	require.Equal(t, len(versions), len(rpcVersions))
	for i, version := range versions {
		require.Equal(t, uint32(version), uint32(rpcVersions[i]))
	}

	// Finally test that the current version contants are up to date
	require.Equal(t,
		CurrentProtocolVersion(),
		versions[len(versions)-2],
	)

	require.Equal(t,
		uint32(CurrentProtocolVersion()),
		uint32(CurrentRPCProtocolVersion()),
	)

	EnableExperimentalProtocol()

	require.Equal(t,
		CurrentProtocolVersion(),
		ProtocolVersion(experimentalRPCProtocolVersion),
	)
}
