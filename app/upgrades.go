package app

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stafihub/stafihub/app/upgrades/v030"
	"github.com/stafihub/stafihub/app/upgrades/v040"
)

func (app *App) setupUpgradeHandlers() {
	// v030 upgrade handler
	app.UpgradeKeeper.SetUpgradeHandler(
		v030.UpgradeName,
		v030.CreateUpgradeHandler(app.mm, app.configurator),
	)

	// v040 upgrade handler
	app.UpgradeKeeper.SetUpgradeHandler(
		v040.UpgradeName,
		v040.CreateUpgradeHandler(app.mm, app.configurator),
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Errorf("failed to ReadUpgradeInfoFromDisk, err: %w", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	var storeUpgrades *storetypes.StoreUpgrades

	// upgrade store case
	switch upgradeInfo.Name {
	}

	if storeUpgrades != nil {
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	}
}
