# ChoreChamps

ChoreChamps is a score keeping application for managing kids' chores. It is designed to run as a Docker container and listens on port 8282. The application is built using the Go programming language and uses a MySQL database. It also integrates with Twilio for WhatsApp notifications and Amazon Echo devices for voice interaction.

## Features

- Configuration section to set up kids' profiles and rules with associated scores
- Support for both positive and negative scores
- Automatic score reset for each kid every Sunday
- WhatsApp notifications to each kid when their score is updated
- Integration with Amazon Echo devices for voice interaction

## Application Structure
ChoreChamps/
├── main.go
├── config/
│ └── config.go
├── model/
│ ├── kid.go
│ └── rule.go
├── handler/
│ ├── kid_handler.go
│ └── rule_handler.go
├── middleware/
│ └── auth_middleware.go
└── util/
└── whatsapp.go


## Prerequisites

- Docker
- Docker Compose
- Twilio Account (for WhatsApp integration)

## Build and Deploy

1. Clone the ChoreChamps repository:
`git clone https://github.com/yourusername/ChoreChamps.git`
`cd ChoreChamps`

2. Update the `docker-compose.yml` file with your Twilio credentials (Account SID, Auth Token, and Phone Number).

3. Build and run the application using Docker Compose:
`docker-compose up -d`

The application will be accessible at `http://localhost:8282`.

## Usage

1. Set up kids' profiles and rules using the configuration section.

2. Use the application's API to add, update, or remove kids and rules.

3. The application will automatically reset scores every Sunday and send WhatsApp notifications to each kid with their updated score.

4. Use Amazon Echo devices to interact with the application using voice commands. (Requires creating an Alexa skill)

## Contributing

Please submit pull requests for bug fixes or improvements. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)