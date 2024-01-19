package app

import (
	v01 "coss/app/upgrades/v01"

	vaultmodulekeeper "coss/x/vault/keeper"

	"github.com/cosmos/cosmos-sdk/types/module"
)

func (app *App) setupUpgradeHandlers(configurator module.Configurator, keeper vaultmodulekeeper.Keeper) {
	app.UpgradeKeeper.SetUpgradeHandler(v01.UpgradeName, v01.CreateUpgradeHandler(app.mm, configurator))

	//  upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	//  if err != nil {
	//   panic(fmt.Errorf("failed to read upgrade info from disk: %w", err))
	//  }

	//  if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
	//   return
	//  }

	//  var storeUpgrades *storetypes.StoreUpgrades

	//	if storeUpgrades != nil {
	//	 app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	//	}
}
