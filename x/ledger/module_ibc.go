package ledger

import (
	"fmt"
	"strings"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	ibcporttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"
	"github.com/stafihub/stafihub/x/ledger/keeper"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ ibcporttypes.IBCModule = (*IBCModule)(nil)

// IBCModule implements the ICS26 interface for interchain accounts controller chains
type IBCModule struct {
	keeper keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	ctx.Logger().Info("OnChanOpenInit", "connectionHops", connectionHops, "portId", portID, "channelId", channelID, "conterparty", counterparty, "channelCap", channelCap.String(), "version", version)
	// Note: The channel capability must be claimed by the authentication module in OnChanOpenInit otherwise the
	// authentication module will not be able to send packets on the channel created for the associated interchain account.
	if err := im.keeper.ClaimCapability(ctx, channelCap, ibchost.ChannelCapabilityPath(portID, channelID)); err != nil {
		return version, err
	}
	return version, nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	ctx.Logger().Info("OnChanOpenAck", "portId", portID, "channelId", channelID, "counterpartyChannelID", counterpartyChannelID, "counterpartyVersion", counterpartyVersion)
	connectionId, connectionEnd, err := im.keeper.IBCKeeper.ChannelKeeper.GetChannelConnection(ctx, portID, channelID)
	if err != nil {
		return fmt.Errorf("cannot get channel connection, portId: %s channelId: %s", portID, channelID)
	}

	controllerConnectionId := connectionId
	hostConnectionId := connectionEnd.GetCounterparty().GetConnectionID()

	interchainAddress, found := im.keeper.ICAControllerKeeper.GetInterchainAccountAddress(ctx, controllerConnectionId, portID)
	if !found {
		return fmt.Errorf("GetInterchainAccountAddress failed for %s/%s", controllerConnectionId, portID)
	}

	portIdSlice := strings.Split(portID, "-")
	if len(portIdSlice) != 4 {
		return fmt.Errorf("portId format err %s/%s", controllerConnectionId, portID)
	}
	if fmt.Sprint(portIdSlice[0], "-") != icatypes.PortPrefix {
		return fmt.Errorf("portId prefix err %s/%s", controllerConnectionId, portID)
	}

	denom := portIdSlice[1]
	index, err := math.ParseUint(portIdSlice[2])
	if err != nil {
		return err
	}
	addressTail := portIdSlice[3]

	icaPoolDetail, found := im.keeper.GetIcaPoolDetail(ctx, denom, uint32(index.Uint64()))
	if !found {
		return fmt.Errorf("ica pool detail not found %s/%s", controllerConnectionId, portID)
	}

	err = im.keeper.CheckAddress(ctx, icaPoolDetail.Denom, interchainAddress)
	if err != nil {
		return fmt.Errorf("check interchainAddress failed, err: %s", err)
	}
	switch addressTail {
	case types.DelegationOwnerTail:
		if icaPoolDetail.Status >= types.IcaPoolStatusSetWithdrawal {
			icaPoolDetail.Status = types.IcaPoolStatusSetWithdrawal
		} else {
			icaPoolDetail.Status = icaPoolDetail.Status + 1
		}
		icaPoolDetail.DelegationAccount.Address = interchainAddress
		icaPoolDetail.DelegationAccount.CtrlPortId = portID
		icaPoolDetail.DelegationAccount.CtrlChannelId = channelID
		icaPoolDetail.DelegationAccount.HostConnectionId = hostConnectionId
		icaPoolDetail.DelegationAccount.HostPortId = icatypes.PortID
		icaPoolDetail.DelegationAccount.HostChannelId = counterpartyChannelID

		im.keeper.SetIcaPoolDetail(ctx, icaPoolDetail)
		im.keeper.SetIcaPoolDelegationAddrIndex(ctx, icaPoolDetail)
	case types.WithdrawalOwnerTail:
		if icaPoolDetail.Status >= types.IcaPoolStatusSetWithdrawal {
			icaPoolDetail.Status = types.IcaPoolStatusSetWithdrawal
		} else {
			icaPoolDetail.Status = icaPoolDetail.Status + 1
		}
		icaPoolDetail.WithdrawalAccount.Address = interchainAddress
		icaPoolDetail.WithdrawalAccount.CtrlPortId = portID
		icaPoolDetail.WithdrawalAccount.CtrlChannelId = channelID
		icaPoolDetail.WithdrawalAccount.HostConnectionId = hostConnectionId
		icaPoolDetail.WithdrawalAccount.HostPortId = icatypes.PortID
		icaPoolDetail.WithdrawalAccount.HostChannelId = counterpartyChannelID

		im.keeper.SetIcaPoolDetail(ctx, icaPoolDetail)
	default:
		return fmt.Errorf("unknown address tail")
	}

	ctx.Logger().Info(fmt.Sprintf("OnChanOpenAck  end %s/%s", controllerConnectionId, portID))
	return nil
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return im.keeper.OnAcknowledgement(ctx, modulePacket, acknowledgement)
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

func (im IBCModule) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (version string, err error) {
	return proposedVersion, nil
}

// ###################################################################################
// 	Required functions to satisfy interface but not implemented for ICA auth modules
// ###################################################################################

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	panic("UNIMPLEMENTED")
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	panic("UNIMPLEMENTED")
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	panic("UNIMPLEMENTED")
}

// OnRecvPacket implements the IBCModule interface
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	panic("UNIMPLEMENTED")
}
