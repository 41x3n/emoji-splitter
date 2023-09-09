# Emoji Splitter

Emoji Splitter is a small Golang web application that takes one or many emojis as input and returns them broken down into their constituent emojis. It's a simple yet fun tool to analyze and split emojis, allowing you to see their individual components.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Installation

To run Emoji Splitter locally or deploy it to your own server, follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/41x3n/emoji-splitter.git
   cd emoji-splitter
   ```

2. Build and run the server:

   ```sh
   go run cmd/server/main.go
   ```

3. The server will start and be available at `http://localhost:8080`. You can now use the API endpoints described below.

## Usage

Emoji Splitter provides a simple API for splitting emojis into their constituent parts. You can use either a GET or POST request to interact with the API.

## API Endpoints

- **GET /split**: Split emojis using a query parameter.
  - **Example**:
    `GET https://emoji-splitter.azurewebsites.net/split?emoji=<YOUR_EMOJI>`
  
- **POST /split**: Split emojis using a JSON payload.
  - **Example**:
    `POST https://emoji-splitter.azurewebsites.net/split`
    ```json
    {
        "emoji": "<YOUR_EMOJI>"
    }
    ```

## Contributing

We welcome contributions from the community to make Emoji Splitter even better. If you'd like to contribute, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure the code is well-documented.
4. Write tests for new features or modifications.
5. Submit a pull request with a clear description of your changes.

## License

Emoji Splitter is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
