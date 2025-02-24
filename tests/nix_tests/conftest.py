import pytest

from .network import setup_aizel, setup_aizel_6dec, setup_aizel_rocksdb, setup_geth


@pytest.fixture(scope="session")
def aizel(tmp_path_factory):
    path = tmp_path_factory.mktemp("aizel")
    yield from setup_aizel(path, 26650)


@pytest.fixture(scope="session")
def aizel_6dec(tmp_path_factory):
    path = tmp_path_factory.mktemp("aizel-6dec")
    yield from setup_aizel_6dec(path, 46650)


@pytest.fixture(scope="session")
def aizel_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("aizel-rocksdb")
    yield from setup_aizel_rocksdb(path, 20650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["aizel", "aizel-ws", "aizel-6dec"])
def aizel_rpc_ws(request, aizel, aizel_6dec):
    """
    run on both aizel and aizel websocket
    """
    provider = request.param
    if provider == "aizel":
        yield aizel
    elif provider == "aizel-ws":
        aizel_ws = aizel.copy()
        aizel_ws.use_websocket()
        yield aizel_ws
    elif provider == "aizel-6dec":
        yield aizel_6dec
    else:
        raise NotImplementedError


@pytest.fixture(
    scope="module", params=["aizel", "aizel-ws", "aizel-6dec", "aizel-rocksdb", "geth"]
)
def cluster(request, aizel, aizel_6dec, aizel_rocksdb, geth):
    """
    run on aizel, aizel websocket,
    aizel built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "aizel":
        yield aizel
    elif provider == "aizel-ws":
        aizel_ws = aizel.copy()
        aizel_ws.use_websocket()
        yield aizel_ws
    elif provider == "aizel-6dec":
        yield aizel_6dec
    elif provider == "geth":
        yield geth
    elif provider == "aizel-rocksdb":
        yield aizel_rocksdb
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["aizel", "aizel-6dec", "aizel-rocksdb"])
def aizel_cluster(request, aizel, aizel_6dec, aizel_rocksdb):
    """
    run on aizel default build &
    aizel with rocksdb build and memIAVL + versionDB
    """
    provider = request.param
    if provider == "aizel":
        yield aizel
    elif provider == "aizel-6dec":
        yield aizel_6dec
    elif provider == "aizel-rocksdb":
        yield aizel_rocksdb
    else:
        raise NotImplementedError
