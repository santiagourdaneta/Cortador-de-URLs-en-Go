import http from 'k6/http';
import { check, sleep } from 'k6';

export default function () {
  const url = 'http://localhost:8080/shorten';
  const payload = JSON.stringify({ url: 'http://ejemplo.com' });
  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);

  check(res, {
    'status is 200 or 429': (r) => r.status === 200 || r.status === 429,
  });

  sleep(1);
}