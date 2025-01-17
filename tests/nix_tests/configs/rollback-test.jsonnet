local config = import 'default.jsonnet';

config {
  'aizel_9002-1'+: {
    validators: super.validators[0:1] + [{
      name: 'fullnode',
    }],
    'app-config'+: {
      'api'+: {
        'enable': true,
      },
      'grpc'+: {
        'enable': true,
      },
    },
  },
}
