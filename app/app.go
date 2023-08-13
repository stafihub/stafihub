package app

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"

	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/spf13/cast"

	"github.com/stafihub/stafihub/cosmoscmd"
	"github.com/tendermint/spm/openapiconsole"

	"github.com/stafihub/stafihub/docs"

	bridgemodule "github.com/stafihub/stafihub/x/bridge"
	bridgemodulekeeper "github.com/stafihub/stafihub/x/bridge/keeper"
	bridgemoduletypes "github.com/stafihub/stafihub/x/bridge/types"
	"github.com/stafihub/stafihub/x/claim"
	claimmodulekeeper "github.com/stafihub/stafihub/x/claim/keeper"
	claimmoduletypes "github.com/stafihub/stafihub/x/claim/types"
	"github.com/stafihub/stafihub/x/ledger"
	ledgerkeeper "github.com/stafihub/stafihub/x/ledger/keeper"
	ledgertypes "github.com/stafihub/stafihub/x/ledger/types"
	miningmodule "github.com/stafihub/stafihub/x/mining"
	miningmodulekeeper "github.com/stafihub/stafihub/x/mining/keeper"
	miningmoduletypes "github.com/stafihub/stafihub/x/mining/types"
	rbankmodule "github.com/stafihub/stafihub/x/rbank"
	rbankmodulekeeper "github.com/stafihub/stafihub/x/rbank/keeper"
	rbankmoduletypes "github.com/stafihub/stafihub/x/rbank/types"
	rdexmodule "github.com/stafihub/stafihub/x/rdex"
	rdexmodulekeeper "github.com/stafihub/stafihub/x/rdex/keeper"
	rdexmoduletypes "github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stafihub/stafihub/x/relayers"
	relayerskeeper "github.com/stafihub/stafihub/x/relayers/keeper"
	relayerstypes "github.com/stafihub/stafihub/x/relayers/types"
	rmintrewardmodule "github.com/stafihub/stafihub/x/rmintreward"
	rmintrewardmodulekeeper "github.com/stafihub/stafihub/x/rmintreward/keeper"
	rmintrewardmoduletypes "github.com/stafihub/stafihub/x/rmintreward/types"
	rstakingmodule "github.com/stafihub/stafihub/x/rstaking"
	rstakingmodulekeeper "github.com/stafihub/stafihub/x/rstaking/keeper"
	rstakingmoduletypes "github.com/stafihub/stafihub/x/rstaking/types"
	"github.com/stafihub/stafihub/x/rvalidator"
	rvalidatormodulekeeper "github.com/stafihub/stafihub/x/rvalidator/keeper"
	rvalidatormoduletypes "github.com/stafihub/stafihub/x/rvalidator/types"
	"github.com/stafihub/stafihub/x/rvote"
	rvotekeeper "github.com/stafihub/stafihub/x/rvote/keeper"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
	"github.com/stafihub/stafihub/x/sudo"
	sudokeeper "github.com/stafihub/stafihub/x/sudo/keeper"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"

	// this line is used by starport scaffolding # stargate/app/moduleImport
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	custombank "github.com/stafihub/stafihub/custom/bank"
	customcrisis "github.com/stafihub/stafihub/custom/crisis"
	customgov "github.com/stafihub/stafihub/custom/gov"
	custommint "github.com/stafihub/stafihub/custom/mint"
	customstaking "github.com/stafihub/stafihub/custom/staking"
)

const (
	AccountAddressPrefix = "stafi"
	Name                 = "stafihub"
)

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		// distrclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		custombank.AppModuleBasic{},
		capability.AppModuleBasic{},
		customstaking.AppModuleBasic{},
		custommint.AppModuleBasic{},
		distr.AppModuleBasic{},
		customgov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		customcrisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ica.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		sudo.AppModuleBasic{},
		relayers.AppModuleBasic{},
		ledger.AppModuleBasic{},
		rvote.AppModuleBasic{},
		rstakingmodule.AppModuleBasic{},
		bridgemodule.AppModuleBasic{},
		rmintrewardmodule.AppModuleBasic{},
		rbankmodule.AppModuleBasic{},
		rdexmodule.AppModuleBasic{},
		miningmodule.AppModuleBasic{},
		rvalidator.AppModuleBasic{},
		claim.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:        {authtypes.Burner},
		distrtypes.ModuleName:             nil,
		minttypes.ModuleName:              {authtypes.Minter},
		stakingtypes.BondedPoolName:       {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:    {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:               {authtypes.Burner},
		ibctransfertypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		ledgertypes.ModuleName:            {authtypes.Minter, authtypes.Burner},
		rstakingmoduletypes.ModuleName:    {authtypes.Burner, authtypes.Minter},
		bridgemoduletypes.ModuleName:      {authtypes.Burner, authtypes.Minter},
		rmintrewardmoduletypes.ModuleName: nil,
		rdexmoduletypes.ModuleName:        {authtypes.Burner, authtypes.Minter},
		miningmoduletypes.ModuleName:      nil,
		icatypes.ModuleName:               nil,
		claimmoduletypes.ModuleName:       nil,
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}
)

var (
	_ cosmoscmd.CosmosApp     = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper     authkeeper.AccountKeeper
	BankKeeper        bankkeeper.Keeper
	CapabilityKeeper  *capabilitykeeper.Keeper
	StakingKeeper     stakingkeeper.Keeper
	SlashingKeeper    slashingkeeper.Keeper
	MintKeeper        mintkeeper.Keeper
	DistrKeeper       distrkeeper.Keeper
	GovKeeper         govkeeper.Keeper
	CrisisKeeper      crisiskeeper.Keeper
	UpgradeKeeper     upgradekeeper.Keeper
	ParamsKeeper      paramskeeper.Keeper
	IBCKeeper         *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper    evidencekeeper.Keeper
	IBCTransferKeeper ibctransferkeeper.Keeper
	FeeGrantKeeper    feegrantkeeper.Keeper

	ConsensusParamsKeeper consensuskeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper         capabilitykeeper.ScopedKeeper
	ScopedIBCTransferKeeper capabilitykeeper.ScopedKeeper

	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper

	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper

	SudoKeeper sudokeeper.Keeper

	RelayersKeeper relayerskeeper.Keeper

	LedgerKeeper ledgerkeeper.Keeper

	RvoteKeeper rvotekeeper.Keeper

	RStakingKeeper rstakingmodulekeeper.Keeper

	BridgeKeeper bridgemodulekeeper.Keeper

	RmintrewardKeeper rmintrewardmodulekeeper.Keeper

	RbankKeeper rbankmodulekeeper.Keeper

	RdexKeeper rdexmodulekeeper.Keeper

	MiningKeeper miningmodulekeeper.Keeper

	RValidatorKeeper rvalidatormodulekeeper.Keeper

	ClaimKeeper claimmodulekeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// the module manager
	mm           *module.Manager
	configurator module.Configurator
}

// RegisterNodeService implements types.Application.
func (*App) RegisterNodeService(client.Context) {
	panic("unimplemented")
}

// New returns a reference to an initialized Gaia.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig cosmoscmd.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) cosmoscmd.App {
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		consensustypes.StoreKey,
		ibchost.StoreKey,
		icacontrollertypes.StoreKey,
		icahosttypes.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,
		sudotypes.StoreKey,
		relayerstypes.StoreKey,
		ledgertypes.StoreKey,
		rvotetypes.StoreKey,
		rstakingmoduletypes.StoreKey,
		bridgemoduletypes.StoreKey,
		rmintrewardmoduletypes.StoreKey,
		rbankmoduletypes.StoreKey,
		rdexmoduletypes.StoreKey,
		miningmoduletypes.StoreKey,
		rvalidatormoduletypes.StoreKey,
		claimmoduletypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])
	app.ConsensusParamsKeeper = consensuskeeper.NewKeeper(appCodec, keys[consensustypes.StoreKey], authtypes.NewModuleAddress(govtypes.ModuleName).String())
	// set the BaseApp's parameter store
	bApp.SetParamStore(&app.ConsensusParamsKeeper)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedIBCTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)

	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		maccPerms,
		sdk.Bech32MainPrefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.ModuleAccountAddrs(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.SudoKeeper = sudokeeper.NewKeeper(
		appCodec,
		keys[sudotypes.StoreKey],
		keys[sudotypes.MemStoreKey],
	)

	rstakingKeeper := rstakingmodulekeeper.NewKeeper(
		appCodec,
		keys[rstakingmoduletypes.StoreKey],
		keys[rstakingmoduletypes.MemStoreKey],
		app.GetSubspace(rstakingmoduletypes.ModuleName),
		app.BankKeeper,
		app.SudoKeeper,
		authtypes.FeeCollectorName,
	)

	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		rstakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// NOTE: stakingKeeper is passed by reference, so that it will contain the hooks setted below
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		stakingKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		cdc,
		keys[slashingtypes.StoreKey],
		stakingKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.CrisisKeeper = *crisiskeeper.NewKeeper(
		appCodec,
		keys[crisistypes.StoreKey],
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper)
	app.UpgradeKeeper = *upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String())

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(rstakingKeeper.Hooks(), app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)
	app.StakingKeeper = *stakingKeeper

	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibchost.StoreKey],
		app.GetSubspace(ibchost.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

	// Create IBCTransfer Keeper
	app.IBCTransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedIBCTransferKeeper,
	)
	ibcTransferModule := ibctransfer.NewAppModule(app.IBCTransferKeeper)
	ibcTransferIBCModule := ibctransfer.NewIBCModule(app.IBCTransferKeeper)

	// Create ICAController keeper
	app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec,
		keys[icacontrollertypes.StoreKey],
		app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper,
		app.MsgServiceRouter(),
	)

	// Create ICAController keeper
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec,
		keys[icahosttypes.StoreKey],
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		scopedICAHostKeeper,
		app.MsgServiceRouter(),
	)
	icaModule := ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper)
	// create ica host ibcmodule
	icaHostIBCModule := icahost.NewIBCModule(app.ICAHostKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	// register the proposal types
	govRouter := govtypesv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govtypesv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(&app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	govConfig := govtypes.DefaultConfig()

	app.GovKeeper = *govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		stakingKeeper,
		app.MsgServiceRouter(),
		govConfig,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.RbankKeeper = *rbankmodulekeeper.NewKeeper(
		appCodec,
		keys[rbankmoduletypes.StoreKey],
		keys[rbankmoduletypes.MemStoreKey],
		app.GetSubspace(rbankmoduletypes.ModuleName),
		app.SudoKeeper,
		app.BankKeeper,
	)
	rbankModule := rbankmodule.NewAppModule(appCodec, app.RbankKeeper)

	app.RelayersKeeper = *relayerskeeper.NewKeeper(
		appCodec,
		keys[relayerstypes.StoreKey],
		keys[relayerstypes.MemStoreKey],
		app.SudoKeeper,
		app.BankKeeper,
	)

	app.RmintrewardKeeper = *rmintrewardmodulekeeper.NewKeeper(
		appCodec,
		keys[rmintrewardmoduletypes.StoreKey],
		keys[rmintrewardmoduletypes.MemStoreKey],
		app.GetSubspace(rmintrewardmoduletypes.ModuleName),
		app.SudoKeeper,
		app.BankKeeper,
	)

	// todo recheck scopedLedgerKeeper
	scopedLedgerKeeper := app.CapabilityKeeper.ScopeToModule(ledgertypes.ModuleName)
	app.LedgerKeeper = *ledgerkeeper.NewKeeper(
		appCodec,
		keys[ledgertypes.StoreKey],
		keys[ledgertypes.MemStoreKey],
		app.SudoKeeper,
		app.BankKeeper,
		app.RelayersKeeper,
		app.RmintrewardKeeper,
		app.RbankKeeper,
		app.ICAControllerKeeper,
		scopedLedgerKeeper,
		*app.IBCKeeper,
	)
	ledgerIBCModule := ledger.NewIBCModule(app.LedgerKeeper)
	// create ica controller ibcmodule
	icaControllerIBCModule := icacontroller.NewIBCMiddleware(ledgerIBCModule, app.ICAControllerKeeper)

	app.RValidatorKeeper = *rvalidatormodulekeeper.NewKeeper(
		appCodec,
		keys[rvalidatormoduletypes.StoreKey],
		keys[rvalidatormoduletypes.MemStoreKey],
		app.GetSubspace(rvalidatormoduletypes.ModuleName),
		app.SudoKeeper,
		app.RbankKeeper,
		app.LedgerKeeper,
	)
	rvalidatorModule := rvalidator.NewAppModule(appCodec, app.RValidatorKeeper, app.AccountKeeper, app.BankKeeper)

	app.ClaimKeeper = *claimmodulekeeper.NewKeeper(
		appCodec,
		keys[claimmoduletypes.StoreKey],
		keys[claimmoduletypes.MemStoreKey],
		app.GetSubspace(claimmoduletypes.ModuleName),
		app.SudoKeeper,
		app.BankKeeper,
	)
	claimModule := claim.NewAppModule(appCodec, app.ClaimKeeper, app.AccountKeeper, app.BankKeeper)

	rvoteRouter := rvotetypes.NewRouter()
	rvoteRouter.AddRoute(ledgertypes.RouterKey, ledger.NewProposalHandler(app.LedgerKeeper))
	rvoteRouter.AddRoute(rvalidatormoduletypes.RouterKey, rvalidator.NewProposalHandler(app.RValidatorKeeper))

	app.RvoteKeeper = *rvotekeeper.NewKeeper(
		appCodec,
		keys[rvotetypes.StoreKey],
		keys[rvotetypes.MemStoreKey],
		app.SudoKeeper,
		app.RelayersKeeper,
		rvoteRouter,
	)

	app.RStakingKeeper = *rstakingKeeper
	rstakingModule := rstakingmodule.NewAppModule(appCodec, app.RStakingKeeper, app.MintKeeper)

	app.BridgeKeeper = *bridgemodulekeeper.NewKeeper(
		appCodec,
		keys[bridgemoduletypes.StoreKey],
		keys[bridgemoduletypes.MemStoreKey],
		app.GetSubspace(bridgemoduletypes.ModuleName),
		app.BankKeeper,
		app.SudoKeeper,
		app.RelayersKeeper,
	)
	bridgeModule := bridgemodule.NewAppModule(appCodec, app.BridgeKeeper, app.AccountKeeper, app.BankKeeper)

	rmintrewardModule := rmintrewardmodule.NewAppModule(appCodec, app.RmintrewardKeeper, app.AccountKeeper, app.BankKeeper)

	app.RdexKeeper = *rdexmodulekeeper.NewKeeper(
		appCodec,
		keys[rdexmoduletypes.StoreKey],
		keys[rdexmoduletypes.MemStoreKey],
		app.GetSubspace(rdexmoduletypes.ModuleName),
		app.BankKeeper,
		app.SudoKeeper,
	)
	rdexModule := rdexmodule.NewAppModule(appCodec, app.RdexKeeper, app.AccountKeeper, app.BankKeeper)

	app.MiningKeeper = *miningmodulekeeper.NewKeeper(
		appCodec,
		keys[miningmoduletypes.StoreKey],
		keys[miningmoduletypes.MemStoreKey],
		app.GetSubspace(miningmoduletypes.ModuleName),
		app.SudoKeeper,
		app.BankKeeper,
		app.RdexKeeper,
	)
	miningModule := miningmodule.NewAppModule(appCodec, app.MiningKeeper, app.AccountKeeper, app.BankKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, ibcTransferIBCModule)

	ibcRouter.AddRoute(ledgertypes.ModuleName, icaControllerIBCModule)
	ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerIBCModule)
	ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostIBCModule)
	// this line is used by starport scaffolding # ibc/app/router
	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		gov.NewAppModule(appCodec, &app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName)),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		staking.NewAppModule(appCodec, &app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		upgrade.NewAppModule(&app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		ibcTransferModule,
		icaModule,
		sudo.NewAppModule(appCodec, app.SudoKeeper),
		relayers.NewAppModule(appCodec, app.RelayersKeeper),
		ledger.NewAppModule(appCodec, app.LedgerKeeper),
		rvalidatorModule,
		claimModule,
		rvote.NewAppModule(appCodec, app.RvoteKeeper),

		rstakingModule,
		bridgeModule,
		rmintrewardModule,
		rbankModule,
		rdexModule,
		miningModule,
		// this line is used by starport scaffolding # stargate/app/appModule
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: rstaking module should happens after mint module and before distribution module, as it will burn minted coins
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		rstakingmoduletypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibchost.ModuleName,
		feegrant.ModuleName,
		ledgertypes.ModuleName,
		genutiltypes.ModuleName,
		paramstypes.ModuleName,
		sudotypes.ModuleName,
		authtypes.ModuleName,
		crisistypes.ModuleName,
		vestingtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		relayerstypes.ModuleName,
		rvotetypes.ModuleName,
		bridgemoduletypes.ModuleName,
		rmintrewardmoduletypes.ModuleName,
		rbankmoduletypes.ModuleName,
		rdexmoduletypes.ModuleName,
		miningmoduletypes.ModuleName,
		rvalidatormoduletypes.ModuleName,
		claimmoduletypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		rstakingmoduletypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		ibchost.ModuleName,
		feegrant.ModuleName,
		ledgertypes.ModuleName,
		genutiltypes.ModuleName,
		paramstypes.ModuleName,
		sudotypes.ModuleName,
		authtypes.ModuleName,
		vestingtypes.ModuleName,
		banktypes.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		relayerstypes.ModuleName,
		rvotetypes.ModuleName,
		bridgemoduletypes.ModuleName,
		rmintrewardmoduletypes.ModuleName,
		rbankmoduletypes.ModuleName,
		rdexmoduletypes.ModuleName,
		miningmoduletypes.ModuleName,
		rvalidatormoduletypes.ModuleName,
		claimmoduletypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	// NOTE: rstaking module must occur after auth/bank/mint moduels so that coinToBeBurned can be set rightly and must
	// before staking module so that hooks can work rightly.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		minttypes.ModuleName,
		rstakingmoduletypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		upgradetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		sudotypes.ModuleName,
		relayerstypes.ModuleName,
		ledgertypes.ModuleName,
		rvotetypes.ModuleName,
		bridgemoduletypes.ModuleName,
		rmintrewardmoduletypes.ModuleName,
		rbankmoduletypes.ModuleName,
		rdexmoduletypes.ModuleName,
		miningmoduletypes.ModuleName,
		rvalidatormoduletypes.ModuleName,
		claimmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)

	// configurator
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)
	app.setupUpgradeHandlers()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  app.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedIBCTransferKeeper = scopedIBCTransferKeeper

	app.ScopedICAControllerKeeper = scopedICAControllerKeeper
	app.ScopedICAHostKeeper = scopedICAHostKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	// authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	// ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register app's OpenAPI routes.
	apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
	apiSvr.Router.HandleFunc("/", openapiconsole.Handler(Name, "/static/openapi.yml"))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypesv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)

	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)

	paramsKeeper.Subspace(sudotypes.ModuleName)
	paramsKeeper.Subspace(relayerstypes.ModuleName)
	paramsKeeper.Subspace(ledgertypes.ModuleName)
	paramsKeeper.Subspace(rvotetypes.ModuleName)
	paramsKeeper.Subspace(rstakingmoduletypes.ModuleName)
	paramsKeeper.Subspace(bridgemoduletypes.ModuleName)
	paramsKeeper.Subspace(rmintrewardmoduletypes.ModuleName)
	paramsKeeper.Subspace(rbankmoduletypes.ModuleName)
	paramsKeeper.Subspace(rdexmoduletypes.ModuleName)
	paramsKeeper.Subspace(miningmoduletypes.ModuleName)
	paramsKeeper.Subspace(rvalidatormoduletypes.ModuleName)
	paramsKeeper.Subspace(claimmoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}
