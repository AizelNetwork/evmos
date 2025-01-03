from pathlib import Path

import pytest

from .network import create_snapshots_dir, setup_custom_aizel
from .utils import EVMOS_6DEC_CHAIN_ID, evm6dec_config, memiavl_config, wait_for_block


@pytest.fixture(scope="module")
def custom_aizel(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp")
    yield from setup_custom_aizel(
        path,
        26260,
        Path(__file__).parent / "configs/discard-abci-resp.jsonnet",
    )


@pytest.fixture(scope="module")
def custom_aizel_6dec(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp-6dec")
    yield from setup_custom_aizel(
        path,
        46860,
        evm6dec_config(path, "discard-abci-resp"),
        chain_id=EVMOS_6DEC_CHAIN_ID,
    )


@pytest.fixture(scope="module")
def custom_aizel_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp-rocksdb")
    yield from setup_custom_aizel(
        path,
        26810,
        memiavl_config(path, "discard-abci-resp"),
        post_init=create_snapshots_dir,
        chain_binary="aizeld-rocksdb",
    )


@pytest.fixture(scope="module", params=["aizel", "aizel-6dec", "aizel-rocksdb"])
def aizel_cluster(request, custom_aizel, custom_aizel_6dec, custom_aizel_rocksdb):
    """
    run on aizel and
    aizel built with rocksdb (memIAVL + versionDB)
    """
    provider = request.param
    if provider == "aizel":
        yield custom_aizel
    elif provider == "aizel-6dec":
        yield custom_aizel_6dec
    elif provider == "aizel-rocksdb":
        yield custom_aizel_rocksdb

    else:
        raise NotImplementedError


def test_gas_eth_tx(aizel_cluster):
    """
    When node does not persist ABCI responses
    eth_gasPrice should return an error instead of crashing
    """
    wait_for_block(aizel_cluster.cosmos_cli(), 3)
    try:
        aizel_cluster.w3.eth.gas_price  # pylint: disable=pointless-statement
        raise Exception(  # pylint: disable=broad-exception-raised
            "This query should have failed"
        )
    except Exception as error:
        assert "block result not found" in error.args[0]["message"]
