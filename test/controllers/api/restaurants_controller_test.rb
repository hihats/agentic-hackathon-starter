require "test_helper"

class Api::RestaurantsControllerTest < ActionDispatch::IntegrationTest
  test "should get nearby restaurants" do
    get "/api/restaurants/nearby"
    assert_response :success

    json_response = JSON.parse(response.body)
    assert json_response.key?("restaurants")
    assert json_response["restaurants"].is_a?(Array)
  end

  test "should get nearby restaurants with genre filter" do
    get "/api/restaurants/nearby", params: { genre: "中華" }
    assert_response :success

    json_response = JSON.parse(response.body)
    assert json_response.key?("restaurants")

    # Check if filtered restaurants contain the genre
    filtered_restaurants = json_response["restaurants"]
    filtered_restaurants.each do |restaurant|
      assert restaurant["genre"].include?("中華") if restaurant["genre"].present?
    end
  end

  test "should get available genres" do
    get "/api/restaurants/genres"
    assert_response :success

    json_response = JSON.parse(response.body)
    assert json_response.key?("genres")
    assert json_response["genres"].is_a?(Array)
  end

  test "should handle service errors gracefully" do
    # Mock the service to raise an error
    service_mock = mock()
    service_mock.expects(:fetch_nearby_restaurants).raises(StandardError.new("Test error"))
    TabelogScraperService.expects(:new).returns(service_mock)

    get "/api/restaurants/nearby"
    assert_response :service_unavailable

    json_response = JSON.parse(response.body)
    assert json_response.key?("error")
  end
end
