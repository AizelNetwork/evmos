from pathlib import Path

import pytest
from web3 import Web3

from .network import create_snapshots_dir, setup_custom_aizel
from .utils import (
    ADDRS,
    EVMOS_6DEC_CHAIN_ID,
    KEYS,
    eth_to_bech32,
    evm6dec_config,
    memiavl_config,
    send_transaction,
    wait_for_fn,
)


@pytest.fixture(scope="module")
def custom_aizel(tmp_path_factory):
    yield from setup_custom_aizel(
        tmp_path_factory.mktemp("zero-fee"),
        26900,
        Path(__file__).parent / "configs/zero-fee.jsonnet",
    )


@pytest.fixture(scope="module")
def custom_aizel_6dec(tmp_path_factory):
    """
    Setup an aizel chain with
    an evm denom with 6 decimals
    """
    path = tmp_path_factory.mktemp("zero-fee-6dec")
    yield from setup_custom_aizel(
        path, 46900, evm6dec_config(path, "zero-fee"), chain_id=EVMOS_6DEC_CHAIN_ID
    )


@pytest.fixture(scope="module")
def custom_aizel_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("zero-fee-rocksdb")
    yield from setup_custom_aizel(
        path,
        26810,
        memiavl_config(path, "zero-fee"),
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


def test_cosmos_tx(aizel_cluster):
    """
    test basic cosmos transaction works with zero fees
    """
    cli = aizel_cluster.cosmos_cli()
    denom = cli.evm_denom()
    sender = eth_to_bech32(ADDRS["signer1"])
    receiver = eth_to_bech32(ADDRS["signer2"])
    amt = 1000

    old_src_balance = cli.balance(sender, denom)
    old_dst_balance = cli.balance(receiver, denom)

    tx = cli.transfer(
        sender,
        receiver,
        f"{amt}{denom}",
        gas_prices=f"0{denom}",
        generate_only=True,
    )

    tx = cli.sign_tx_json(tx, sender, max_priority_price=0)

    rsp = cli.broadcast_tx_json(tx, broadcast_mode="sync")
    assert rsp["code"] == 0, rsp["raw_log"]

    new_dst_balance = 0

    def check_balance_change():
        nonlocal new_dst_balance
        new_dst_balance = cli.balance(receiver, denom)
        return old_dst_balance != new_dst_balance

    wait_for_fn("balance change", check_balance_change)
    assert old_dst_balance + amt == new_dst_balance
    new_src_balance = cli.balance(sender, denom)
    # no fees paid, so sender balance should be
    # initial_balance - amount_sent
    assert old_src_balance - amt == new_src_balance


def test_eth_tx(aizel_cluster):
    """
    test basic Ethereum transaction works with zero fees
    """
    w3: Web3 = aizel_cluster.w3

    sender = ADDRS["signer1"]
    receiver = ADDRS["signer2"]
    amt = int(1e18)

    old_src_balance = w3.eth.get_balance(sender)
    old_dst_balance = w3.eth.get_balance(receiver)

    receipt = send_transaction(
        w3,
        {
            "from": sender,
            "to": receiver,
            "value": amt,
            "gasPrice": 0,
        },
        KEYS["signer1"],
    )
    assert receipt.status == 1

    new_dst_balance = 0

    def check_balance_change():
        nonlocal new_dst_balance
        new_dst_balance = w3.eth.get_balance(receiver)
        return old_dst_balance != new_dst_balance

    wait_for_fn("balance change", check_balance_change)
    assert old_dst_balance + amt == new_dst_balance
    new_src_balance = w3.eth.get_balance(sender)
    # no fees paid, so sender balance should be
    # initial_balance - amount_sent
    assert old_src_balance - amt == new_src_balance
