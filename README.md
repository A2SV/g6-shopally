# ShopAlly

ShopAlly is a multi-platform shopping assistant application, featuring a backend (Go), web frontend (Next.js), and mobile app (Flutter). It helps users compare products, manage alerts, and streamline their shopping experience.

## Project Structure

- **shopally-backend/**: Go backend API, business logic, and integrations.
- **shopally-web/**: Next.js web frontend for user interaction.
- **shopallymobile/**: Flutter mobile app for Android/iOS.
- **LICENSE**: Project license.
- **README.md**: Project documentation.

## Getting Started

### Backend (Go)

1. Navigate to `shopally-backend/`.
2. Install dependencies:
	```sh
	go mod tidy
	```
3. Run the API server:
	```sh
	go run cmd/api/main.go
	```
4. For worker services:
	```sh
	go run worker/main.go
	```

### Web (Next.js)

1. Navigate to `shopally-web/`.
2. Install dependencies:
	```sh
	npm install
	```
3. Start the development server:
	```sh
	npm run dev
	```

### Mobile (Flutter)

1. Navigate to `shopallymobile/`.
2. Install dependencies:
	```sh
	flutter pub get
	```
3. Run the app:
	```sh
	flutter run
	```

## Features

- Product comparison
- Alert management
- Push notifications
- Multi-platform support (Web, Android, iOS)

## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes.
4. Push to the branch.
5. Open a pull request.

## License

This project is licensed under the terms of the LICENSE file.