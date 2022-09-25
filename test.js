import http from 'k6/http';

export default function () {
    const url = 'http://host.docker.internal:3000/api/project';
    const payload = JSON.stringify({
        name: 'Test'
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}