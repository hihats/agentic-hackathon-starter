class Api::RestaurantsController < ApplicationController
  def nearby
    service = TabelogScraperService.new
    genre_filter = params[:genre]
    restaurants = service.fetch_nearby_restaurants(genre_filter: genre_filter)

    render json: {
      restaurants: restaurants.map { |r|
        {
          name: r[:name],
          genre: r[:genre],
          imageUrl: r[:image_url],
          rating: r[:rating],
          url: r[:url]
        }
      }
    }
  rescue StandardError => e
    Rails.logger.error("Restaurant API Error: #{e.message}")
    render json: { error: "店舗情報の取得に失敗しました" }, status: :service_unavailable
  end

  def genres
    service = TabelogScraperService.new
    genres = service.fetch_available_genres

    render json: { genres: genres }
  rescue StandardError => e
    Rails.logger.error("Genres API Error: #{e.message}")
    render json: { error: "ジャンル情報の取得に失敗しました" }, status: :service_unavailable
  end
end
