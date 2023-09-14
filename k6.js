import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';

export const options = {
    vus: 100,
    duration: '1m',
    rps: 8000,
    summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(95)', 'p(99)']
  };

const client = new grpc.Client();
client.load(['proto'], 'simple.proto');

export default () => {
    client.connect('127.0.0.1:8080', {
      plaintext: true
    });
  
    const data = { message: 'Bert' };
    const response = client.invoke('Echo/Do', data);
    check(response, {
      'status is OK': (r) => r && r.status === grpc.StatusOK,
    });

    console.log(JSON.stringify(response.message));
    client.close();
  };