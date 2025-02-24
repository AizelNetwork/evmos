package ledger_test

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdktestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/AizelNetwork/evmos/v20/app"
	"github.com/AizelNetwork/evmos/v20/crypto/hd"
	"github.com/AizelNetwork/evmos/v20/tests/integration/ledger/mocks"
	utiltx "github.com/AizelNetwork/evmos/v20/testutil/tx"
	"github.com/AizelNetwork/evmos/v20/utils"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cometbft/cometbft/version"
	"github.com/stretchr/testify/suite"

	clientkeys "github.com/AizelNetwork/evmos/v20/client/keys"
	aizelkeyring "github.com/AizelNetwork/evmos/v20/crypto/keyring"
	feemarkettypes "github.com/AizelNetwork/evmos/v20/x/feemarket/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmversion "github.com/cometbft/cometbft/proto/tendermint/version"
	rpcclientmock "github.com/cometbft/cometbft/rpc/client/mock"
	cosmosledger "github.com/cosmos/cosmos-sdk/crypto/ledger"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/gomega"
)

var s *LedgerTestSuite

type LedgerTestSuite struct {
	suite.Suite

	app *app.Aizel
	ctx sdk.Context

	ledger       *mocks.SECP256K1
	accRetriever *mocks.AccountRetriever

	accAddr sdk.AccAddress

	privKey types.PrivKey
	pubKey  types.PubKey
}

func TestLedger(t *testing.T) {
	s = new(LedgerTestSuite)
	suite.Run(t, s)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Aizeld Suite")
}

func (suite *LedgerTestSuite) SetupTest() {
	var (
		err     error
		ethAddr common.Address
	)

	suite.ledger = mocks.NewSECP256K1(s.T())

	ethAddr, s.privKey = utiltx.NewAddrKey()

	s.Require().NoError(err)
	suite.pubKey = s.privKey.PubKey()

	suite.accAddr = sdk.AccAddress(ethAddr.Bytes())
}

func (suite *LedgerTestSuite) SetupAizelApp() {
	consAddress := sdk.ConsAddress(utiltx.GenerateAddress().Bytes())

	// init app
	chainID := utils.MainnetChainID + "-1"
	suite.app = app.Setup(false, feemarkettypes.DefaultGenesisState(), chainID)
	suite.ctx = suite.app.BaseApp.NewContextLegacy(false, tmproto.Header{
		Height:          1,
		ChainID:         chainID,
		Time:            time.Now().UTC(),
		ProposerAddress: consAddress.Bytes(),

		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		LastBlockId: tmproto.BlockID{
			Hash: tmhash.Sum([]byte("block_id")),
			PartSetHeader: tmproto.PartSetHeader{
				Total: 11,
				Hash:  tmhash.Sum([]byte("partset_header")),
			},
		},
		AppHash:            tmhash.Sum([]byte("app")),
		DataHash:           tmhash.Sum([]byte("data")),
		EvidenceHash:       tmhash.Sum([]byte("evidence")),
		ValidatorsHash:     tmhash.Sum([]byte("validators")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators")),
		ConsensusHash:      tmhash.Sum([]byte("consensus")),
		LastResultsHash:    tmhash.Sum([]byte("last_result")),
	})
}

func (suite *LedgerTestSuite) NewKeyringAndCtxs(krHome string, input io.Reader, encCfg sdktestutil.TestEncodingConfig) (keyring.Keyring, client.Context, context.Context) {
	kr, err := keyring.New(
		sdk.KeyringServiceName(),
		keyring.BackendTest,
		krHome,
		input,
		encCfg.Codec,
		s.MockKeyringOption(),
	)
	s.Require().NoError(err)
	s.accRetriever = mocks.NewAccountRetriever(s.T())

	initClientCtx := client.Context{}.
		WithCodec(encCfg.Codec).
		// NOTE: cmd.Execute() panics without account retriever
		WithAccountRetriever(s.accRetriever).
		WithTxConfig(encCfg.TxConfig).
		WithLedgerHasProtobuf(true).
		WithUseLedger(true).
		WithKeyring(kr).
		WithClient(mocks.MockCometRPC{Client: rpcclientmock.Client{}}).
		WithChainID(utils.TestnetChainID + "-13").
		WithSignModeStr(flags.SignModeLegacyAminoJSON)

	srvCtx := server.NewDefaultContext()
	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &initClientCtx)
	ctx = context.WithValue(ctx, server.ServerContextKey, srvCtx)

	return kr, initClientCtx, ctx
}

func (suite *LedgerTestSuite) aizelAddKeyCmd() *cobra.Command {
	cmd := keys.AddKeyCommand()

	algoFlag := cmd.Flag(flags.FlagKeyType)
	algoFlag.DefValue = string(hd.EthSecp256k1Type)

	err := algoFlag.Value.Set(string(hd.EthSecp256k1Type))
	suite.Require().NoError(err)

	cmd.Flags().AddFlagSet(keys.Commands().PersistentFlags())

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		clientCtx := client.GetClientContextFromCmd(cmd).WithKeyringOptions(hd.EthSecp256k1Option())
		clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
		if err != nil {
			return err
		}
		buf := bufio.NewReader(clientCtx.Input)
		return clientkeys.RunAddCmd(clientCtx, cmd, args, buf)
	}
	return cmd
}

func (suite *LedgerTestSuite) MockKeyringOption() keyring.Option {
	return func(options *keyring.Options) {
		options.SupportedAlgos = aizelkeyring.SupportedAlgorithms
		options.SupportedAlgosLedger = aizelkeyring.SupportedAlgorithmsLedger
		options.LedgerDerivation = func() (cosmosledger.SECP256K1, error) { return suite.ledger, nil }
		options.LedgerCreateKey = aizelkeyring.CreatePubkey
		options.LedgerAppName = aizelkeyring.AppName
		options.LedgerSigSkipDERConv = aizelkeyring.SkipDERConversion
	}
}

func (suite *LedgerTestSuite) FormatFlag(flag string) string {
	return fmt.Sprintf("--%s", flag)
}
