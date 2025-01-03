import pytest
from web3 import Web3

from .network import setup_aizel, setup_aizel_6dec, setup_aizel_rocksdb
from .utils import (
    ADDRS,
    KEYS,
    derive_new_account,
    send_transaction,
    w3_wait_for_new_blocks,
)


# start a brand new chain for this test
@pytest.fixture(scope="module")
def custom_aizel(tmp_path_factory):
    path = tmp_path_factory.mktemp("account")
    yield from setup_aizel(path, 26700, long_timeout_commit=True)


@pytest.fixture(scope="module")
def custom_aizel_6dec(tmp_path_factory):
    path = tmp_path_factory.mktemp("account-6dec")
    yield from setup_aizel_6dec(
        path,
        46777,
    )


# ATM rocksdb build is not supported for sdkv0.50
# This is due to cronos dependencies (versionDB, memIAVL)
@pytest.fixture(scope="module")
def custom_aizel_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("account-rocksdb")
    yield from setup_aizel_rocksdb(
        path,
        26777,
    )


@pytest.fixture(
    scope="module", params=["aizel", "aizel-ws", "aizel-6dec", "aizel-rocksdb", "geth"]
)
def cluster(request, custom_aizel, custom_aizel_6dec, custom_aizel_rocksdb, geth):
    """
    run on aizel, aizel websocket,
    aizel built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "aizel":
        yield custom_aizel
    elif provider == "aizel-ws":
        aizel_ws = custom_aizel.copy()
        aizel_ws.use_websocket()
        yield aizel_ws
    elif provider == "aizel-6dec":
        yield custom_aizel_6dec
    elif provider == "aizel-rocksdb":
        yield custom_aizel_rocksdb
    elif provider == "geth":
        yield geth
    else:
        raise NotImplementedError


def test_get_transaction_count(cluster):
    w3: Web3 = cluster.w3
    blk = hex(w3.eth.block_number)
    sender = ADDRS["validator"]

    receiver = derive_new_account().address
    n0 = w3.eth.get_transaction_count(receiver, blk)
    # ensure transaction send in new block
    w3_wait_for_new_blocks(w3, 1, sleep=0.1)
    receipt = send_transaction(
        w3,
        {
            "from": sender,
            "to": receiver,
            "value": 1000,
        },
        KEYS["validator"],
    )
    assert receipt.status == 1
    [n1, n2] = [w3.eth.get_transaction_count(receiver, b) for b in [blk, "latest"]]
    assert n0 == n1
    assert n0 == n2


def test_query_future_blk(cluster):
    w3: Web3 = cluster.w3
    acc = derive_new_account(2).address
    current = w3.eth.block_number
    future = current + 1000
    with pytest.raises(ValueError) as exc:
        w3.eth.get_transaction_count(acc, hex(future))
    print(acc, str(exc))
    assert "-32000" in str(exc)
