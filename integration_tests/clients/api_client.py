import requests
from integration_tests.settings import URL, HEADERS


class ApiClient:

    def get_all(self, path="/v1/airports"):
        response = requests.get(
            url=f'{URL}{path}',
            headers=HEADERS,
        )
        response.raise_for_status()
        return response

    def get_limited_list(self, params, path="/v1/airports?limit={lim}"):
        path = path.format(**params)
        response = requests.get(
            url=f'{URL}{path}',
            headers=HEADERS,
        )
        response.raise_for_status()
        return response

    def get_by_id(self, params, path="/v1/airports/{id_air}"):
        path = path.format(**params)
        response = requests.get(
            url=f'{URL}{path}',
            headers=HEADERS,
        )
        response.raise_for_status()
        return response

    def post_airport(self, data, path="/v1/airports"):
        response = requests.post(
            url=f'{URL}{path}',
            headers=HEADERS,
            data=data
        )
        response.raise_for_status()
        return response

    def delete_by_id(self, params, path="/api/v1/airports?airportId={id_air}"):
        path = path.format(**params)
        response = requests.delete(
            url=f'{URL}{path}',
            headers=HEADERS,
        )
        response.raise_for_status()
        return response
