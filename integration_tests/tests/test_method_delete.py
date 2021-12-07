import allure
from hamcrest import assert_that, is_in
from integration_tests.clients.api_client import ApiClient
from integration_tests.helpers.base_helpers import create_airport


class TestMethodDelete:

    @allure.title('Can delete airport by id')
    def test_delete_airport_success(self):
        ap_id = create_airport()
        data = {"id_air": ap_id}

        del_ap = ApiClient().delete_by_id(params=data)
        print("\nDeleted airport ID:", del_ap.json())

        assert_that(actual=del_ap.status_code,
                    matcher=is_in([200, 201]))
