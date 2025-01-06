local config = import 'default.jsonnet';

config {
  'aizel_9002-1'+: {
    config+: {
      storage: {
        discard_abci_responses: true,
      },
    },
  },
}
