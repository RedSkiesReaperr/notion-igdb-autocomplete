source "https://rubygems.org"

ruby "4.0.4"

gem "rails", "8.1.3"
gem "propshaft"
gem "sqlite3", ">= 2.1"
gem "puma", ">= 5.0"
gem "importmap-rails"
gem "turbo-rails"
gem "stimulus-rails"
gem "tailwindcss-rails", ">= 3.3"
gem "solid_queue"
gem "solid_cable"
gem "bootsnap", require: false

gem "faraday"
gem "faraday-retry"
gem "amatch"
gem "service_actor"

group :development, :test do
  gem "dotenv-rails"
  gem "rspec-rails"
  gem "debug", platforms: %i[ mri windows ], require: "debug/prelude"
end

group :development do
  gem "web-console"
end
