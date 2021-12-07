import allure
from hamcrest import assert_that, equal_to, not_none
from integration_tests.clients.api_client import ApiClient


class TestApiMethodPost:

    @allure.title('Can create a new airport')
    def test_post_airport(self):
        data = '{"name": "CDG", "location": "Charles-de-Gaulle"}'

        air = ApiClient().post_airport(data).json()

        assert_that(actual=air,
                    matcher=not_none(),
                    reason=f"Post returned none value but shouldn't. Response: {air}")

    @allure.title("Test method return correct airport's name")
    def test_post_airport_and_check_name(self):
        data = '{"name": "CDG", "location": "Charles-de-Gaulle"}'

        new_ap = ApiClient().post_airport(data).json()

        data = {"id_air": new_ap.get("id")}
        new_ap_vals = ApiClient().get_by_id(params=data).json()
        assert_that(actual=new_ap_vals.get("value").get("name"),
                    matcher=equal_to("CDG"),
                    reason=f'Name should be "CDG" but was'
                           f'{new_ap_vals.get("value").get("name")}')

    @allure.title("Test method return correct airport's location")
    def test_post_airport_and_check_location(self):
        data = '{"name": "CDG", "location": "Charles-de-Gaulle"}'

        new_ap = ApiClient().post_airport(data).json()

        data = {"id_air": new_ap.get("id")}
        new_ap_vals = ApiClient().get_by_id(params=data).json()
        assert_that(actual=new_ap_vals.get("value").get("location"),
                    matcher=equal_to("Charles-de-Gaulle"),
                    reason=f"Location should be 'Charles-de-Gaulle'"
                           f"but was {new_ap_vals.get('value').get('location')}")

    @allure.title("Airport's name should be equal to 3")
    def test_length_of_posted_name_equal_to_3(self):
        data = '{"name": "ORL", "location": "Orly"}'

        new_ap = ApiClient().post_airport(data).json()

        data = {"id_air": new_ap.get("id")}
        new_ap_vals = ApiClient().get_by_id(params=data).json()
        assert_that(actual=len(new_ap_vals.get("value").get("name")),
                    matcher=equal_to(3),
                    reason=f"Name should be if length 3 but was"
                           f"{len(new_ap_vals.get('value').get('name'))}")
