import allure
from hamcrest import assert_that, equal_to, empty
from integration_tests.clients.api_client import ApiClient
from integration_tests.helpers.base_helpers import create_5_airports


class TestApiMethodList:

    def test_parameterless_request_return_empty_list(self):
        airport_list = ApiClient().get_all().json().get("airports")
        assert_that(actual=airport_list,
                    matcher=empty(),
                    reason=f"Should return an empty list but was {airport_list}")

    @allure.title('Test get airport data by id')
    def test_get_list_of_5_airports(self):
        create_5_airports()
        data = {"lim": 5}

        airport_list = ApiClient().get_limited_list(params=data).json().get("airports")
        print("\nList:", airport_list, sep="\n")

        assert_that(actual=len(airport_list),
                    matcher=equal_to(5),
                    reason=f"Number of airport should be equal to 5 but was {len(airport_list)}")
