from random import choice
from string import ascii_uppercase as asciiup
from faker import Faker
from integration_tests.clients.api_client import ApiClient

fake = Faker()


def create_airport():
    name_ = "".join(choice(asciiup) for _ in range(3))
    city = fake.city()
    data = f'{{"name": "{name_}", "location": "{city}"}}'
    airport = ApiClient().post_airport(data).json()
    return airport.get("id")


def create_5_airports():
    airports = list()
    for i in range(5):
        name_ = "".join(choice(asciiup) for _ in range(3))
        city = fake.city()
        data = f'{{"name": "{name_}", "location": "{city}"}}'
        airport = ApiClient().post_airport(data).json()
        airports.append(airport)
    return airports
