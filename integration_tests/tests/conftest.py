from random import randint
import pytest

from integration_tests.clients.api_client import ApiClient


@pytest.fixture
def get_random_airport():
    number = randint(1, 5)
    data = {"id_air": number}
    rand_ap = ApiClient().get_by_id(params=data).json()
    return rand_ap
