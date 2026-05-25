# frozen_string_literal: true

source 'https://rubygems.org'

ruby '4.0.4'

gem 'bootsnap', require: false
gem 'importmap-rails'
gem 'propshaft'
gem 'puma', '>= 5.0'
gem 'rails', '8.1.3'
gem 'solid_cable'
gem 'solid_queue'
gem 'sqlite3', '>= 2.1'
gem 'stimulus-rails'
gem 'tailwindcss-rails', '>= 3.3'
gem 'turbo-rails'

gem 'amatch'
gem 'faraday'
gem 'faraday-retry'
gem 'service_actor'

group :development, :test do
  gem 'debug', platforms: %i[mri windows], require: 'debug/prelude'
  gem 'dotenv-rails'
  gem 'rspec-rails'
  gem 'rubocop', require: false
  gem 'rubocop-performance', require: false
  gem 'rubocop-rails', require: false
  gem 'rubocop-rspec', require: false
  gem 'rubocop-rspec_rails', require: false
end

group :development do
  gem 'web-console'
end
