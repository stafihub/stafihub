package ledger

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibcporttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/stafihub/stafihub/x/ledger/keeper"
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
) error {
	ctx.Logger().Info("OnChanOpenInit", "connectionHops", connectionHops, "portId", portID, "channelId", channelID, "conterparty", counterparty, "channelCap", channelCap.String(), "version", version)
	if err := im.keeper.ClaimCapability(ctx, channelCap, ibchost.ChannelCapabilityPath(portID, channelID)); err != nil {
		return err
	}
	return nil
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
	var metadata icatypes.Metadata
	if err := icatypes.ModuleCdc.UnmarshalJSON([]byte(counterpartyVersion), &metadata); err != nil {
		return sdkerrors.Wrapf(icatypes.ErrUnknownDataType, "cannot unmarshal ICS-27 interchain accounts metadata")
	}

	controllerConnectionId := metadata.ControllerConnectionId
	hostConnectionId := metadata.HostConnectionId
	interchainAddress, found := im.keeper.ICAControllerKeeper.GetInterchainAccountAddress(ctx, controllerConnectionId, portID)
	if !found {
		ctx.Logger().Error(fmt.Sprintf("Expected to find an address for %s/%s", controllerConnectionId, portID))
		return nil
	}

	portIdSlice := strings.Split(portID, "-")
	if len(portIdSlice) != 4 {
		ctx.Logger().Error(fmt.Sprintf("portId format err %s/%s", controllerConnectionId, portID))
		return nil
	}
	if fmt.Sprint(portIdSlice[0], "-") != icatypes.PortPrefix {
		ctx.Logger().Error(fmt.Sprintf("portId prefix err %s/%s", controllerConnectionId, portID))
		return nil
	}

	denom := portIdSlice[1]
	sequence := portIdSlice[2]
	isDelegationAddr := portIdSlice[3] == "delegation"

	icaPoolDetail, found := im.keeper.GetIcaPoolDetail(ctx, denom, sequence)
	if !found {
		ctx.Logger().Error(fmt.Sprintf("ica pool detail not found %s/%s", controllerConnectionId, portID))
		return nil
	}

	if isDelegationAddr {
		icaPoolDetail.Status = icaPoolDetail.Status + 1
		icaPoolDetail.DelegationAccount.Address = interchainAddress
		icaPoolDetail.DelegationAccount.CtrlPortId = portID
		icaPoolDetail.DelegationAccount.HostConnectionId = hostConnectionId
		icaPoolDetail.DelegationAccount.HostPortId = icatypes.PortID

		im.keeper.SetIcaPoolDetail(ctx, icaPoolDetail)
		im.keeper.SetIcaPoolIndex(ctx, icaPoolDetail)
	} else {
		icaPoolDetail.Status = icaPoolDetail.Status + 1
		icaPoolDetail.WithdrawAccount.Address = interchainAddress
		icaPoolDetail.WithdrawAccount.CtrlPortId = portID
		icaPoolDetail.WithdrawAccount.HostConnectionId = hostConnectionId
		icaPoolDetail.WithdrawAccount.HostPortId = icatypes.PortID

		im.keeper.SetIcaPoolDetail(ctx, icaPoolDetail)
	}

	ctx.Logger().Error(fmt.Sprintf("OnChanOpenAck  end %s/%s", controllerConnectionId, portID))
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
