package predeploys

import "github.com/ethereum/go-ethereum/common"

// TODO - we should get a single toml yaml or json file source of truth in @eth-optimism/bedrock package
// This needs to be kept in sync with @eth-optimism/contracts-ts/wagmi.config.ts which also specifies this
// To improve robustness and maintainability contracts-bedrock should export all addresses
const (
	L2ToL1MessagePasser           = "0x4200000000000000000000000000000000000016"
	DeployerWhitelist             = "0x4200000000000000000000000000000000000002"
	WETH9                         = "0x4200000000000000000000000000000000000006"
	L2CrossDomainMessenger        = "0x4200000000000000000000000000000000000007"
	L2StandardBridge              = "0x4200000000000000000000000000000000000010"
	SequencerFeeVault             = "0x4200000000000000000000000000000000000011"
	OptimismMintableERC20Factory  = "0x4200000000000000000000000000000000000012"
	L1BlockNumber                 = "0x4200000000000000000000000000000000000013"
	GasPriceOracle                = "0x420000000000000000000000000000000000000F"
	L1Block                       = "0x4200000000000000000000000000000000000015"
	GovernanceToken               = "0x4200000000000000000000000000000000000042"
	LegacyMessagePasser           = "0x4200000000000000000000000000000000000000"
	L2ERC721Bridge                = "0x4200000000000000000000000000000000000014"
	OptimismMintableERC721Factory = "0x4200000000000000000000000000000000000017"
	ProxyAdmin                    = "0x4200000000000000000000000000000000000018"
	BaseFeeVault                  = "0x4200000000000000000000000000000000000019"
	L1FeeVault                    = "0x420000000000000000000000000000000000001a"
	DebankL2Register              = "0x42000000000000000000000000000000000000bb"
	DebankMintBurnManager         = "0x42000000000000000000000000000000000000ba"
	SchemaRegistry                = "0x4200000000000000000000000000000000000020"
	EAS                           = "0x4200000000000000000000000000000000000021"
	Create2Deployer               = "0x13b0D85CcB8bf860b6b79AF3029fCA081AE9beF2"
)

var (
	L2ToL1MessagePasserAddr           = common.HexToAddress(L2ToL1MessagePasser)
	DeployerWhitelistAddr             = common.HexToAddress(DeployerWhitelist)
	WETH9Addr                         = common.HexToAddress(WETH9)
	L2CrossDomainMessengerAddr        = common.HexToAddress(L2CrossDomainMessenger)
	L2StandardBridgeAddr              = common.HexToAddress(L2StandardBridge)
	SequencerFeeVaultAddr             = common.HexToAddress(SequencerFeeVault)
	OptimismMintableERC20FactoryAddr  = common.HexToAddress(OptimismMintableERC20Factory)
	L1BlockNumberAddr                 = common.HexToAddress(L1BlockNumber)
	GasPriceOracleAddr                = common.HexToAddress(GasPriceOracle)
	L1BlockAddr                       = common.HexToAddress(L1Block)
	GovernanceTokenAddr               = common.HexToAddress(GovernanceToken)
	LegacyMessagePasserAddr           = common.HexToAddress(LegacyMessagePasser)
	L2ERC721BridgeAddr                = common.HexToAddress(L2ERC721Bridge)
	OptimismMintableERC721FactoryAddr = common.HexToAddress(OptimismMintableERC721Factory)
	ProxyAdminAddr                    = common.HexToAddress(ProxyAdmin)
	BaseFeeVaultAddr                  = common.HexToAddress(BaseFeeVault)
	L1FeeVaultAddr                    = common.HexToAddress(L1FeeVault)
	DebankL2RegisterAddr              = common.HexToAddress(DebankL2Register)
	DebankMintBurnManagerAddr         = common.HexToAddress(DebankMintBurnManager)
	SchemaRegistryAddr                = common.HexToAddress(SchemaRegistry)
	EASAddr                           = common.HexToAddress(EAS)
	Create2DeployerAddr               = common.HexToAddress(Create2Deployer)

	Predeploys          = make(map[string]*Predeploy)
	PredeploysByAddress = make(map[common.Address]*Predeploy)
)

func init() {
	Predeploys["DebankL2Register"] = &Predeploy{Address: DebankL2RegisterAddr}
	Predeploys["DebankMintBurnManager"] = &Predeploy{Address: DebankMintBurnManagerAddr}
	Predeploys["L2ToL1MessagePasser"] = &Predeploy{Address: L2ToL1MessagePasserAddr}
	Predeploys["DeployerWhitelist"] = &Predeploy{Address: DeployerWhitelistAddr}
	Predeploys["WETH9"] = &Predeploy{Address: WETH9Addr, ProxyDisabled: true}
	Predeploys["L2CrossDomainMessenger"] = &Predeploy{Address: L2CrossDomainMessengerAddr}
	Predeploys["L2StandardBridge"] = &Predeploy{Address: L2StandardBridgeAddr}
	Predeploys["SequencerFeeVault"] = &Predeploy{Address: SequencerFeeVaultAddr}
	Predeploys["OptimismMintableERC20Factory"] = &Predeploy{Address: OptimismMintableERC20FactoryAddr}
	Predeploys["L1BlockNumber"] = &Predeploy{Address: L1BlockNumberAddr}
	Predeploys["GasPriceOracle"] = &Predeploy{Address: GasPriceOracleAddr}
	Predeploys["L1Block"] = &Predeploy{Address: L1BlockAddr}
	Predeploys["GovernanceToken"] = &Predeploy{
		Address:       GovernanceTokenAddr,
		ProxyDisabled: true,
		Enabled: func(config DeployConfig) bool {
			return config.GovernanceEnabled()
		},
	}
	Predeploys["LegacyMessagePasser"] = &Predeploy{Address: LegacyMessagePasserAddr}
	Predeploys["L2ERC721Bridge"] = &Predeploy{Address: L2ERC721BridgeAddr}
	Predeploys["OptimismMintableERC721Factory"] = &Predeploy{Address: OptimismMintableERC721FactoryAddr}
	Predeploys["ProxyAdmin"] = &Predeploy{Address: ProxyAdminAddr}
	Predeploys["BaseFeeVault"] = &Predeploy{Address: BaseFeeVaultAddr}
	Predeploys["L1FeeVault"] = &Predeploy{Address: L1FeeVaultAddr}
	Predeploys["SchemaRegistry"] = &Predeploy{Address: SchemaRegistryAddr}
	Predeploys["EAS"] = &Predeploy{Address: EASAddr}
	Predeploys["Create2Deployer"] = &Predeploy{
		Address:       Create2DeployerAddr,
		ProxyDisabled: true,
		Enabled: func(config DeployConfig) bool {
			canyonTime := config.CanyonTime(0)
			return canyonTime != nil && *canyonTime == 0
		},
	}

	for _, predeploy := range Predeploys {
		PredeploysByAddress[predeploy.Address] = predeploy
	}
}
