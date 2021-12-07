from random import randint
import allure
from hamcrest import assert_that, equal_to, not_none
from integration_tests.clients.api_client import ApiClient


class TestApiMethodGet:

    @allure.title('Test get airport data by id')
    def test_get_airport_by_id(self):
        data = '{"name": "AAA", "location": "Test location of AAA"}'
        new_ap = ApiClient().post_airport(data).json()
        new_id = {"id_air": new_ap.get("id")}

        airport = ApiClient().get_by_id(params=new_id).json()
        print("\nRequested airport:", airport.get("value"))

        assert_that(actual=airport,
                    matcher=not_none(),
                    reason='')

    def test_should_return_correct_name(self):
        data = '{"name": "BBB", "location": "Test location of BBB"}'
        new_ap = ApiClient().post_airport(data).json()
        new_id = {"id_air": new_ap.get("id")}

        airport = ApiClient().get_by_id(params=new_id).json()
        print("\nRequested airport:", airport.get("value"))

        assert_that(actual=airport.get("value").get("name"),
                    matcher=equal_to("BBB"),
                    reason=f'Name of airport should be "BBB" but was'
                           f'{airport.get("value").get("name")}')

    def test_should_return_correct_location(self):
        data = '{"name": "CCC", "location": "Test location of CCC"}'
        new_ap = ApiClient().post_airport(data).json()
        new_id = {"id_air": new_ap.get("id")}

        airport = ApiClient().get_by_id(params=new_id).json()
        print("\nRequested airport:", airport.get("value"))

        assert_that(actual=airport.get("value").get("location"),
                    matcher=equal_to("Test location of CCC"),
                    reason=f'Name of airport should be "Test location of CCC" but was'
                           f'{airport.get("value").get("location")}')
