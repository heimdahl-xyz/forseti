# Forseti: Open Source Transfer Collector for Heimdahl Platform

Forseti is an open-source tool designed to collect and index real-time token transfers from Heimdahl **onchain data
aggregator platform**. It utilizes Docker Compose for easy setup and operation.

## Features

* **Zero configuration** Setup your own data consumer in minutes
* **Multi-Chain Support:** Collects token transfers from EVM, Solana, and Tron blockchains.
* **Real-Time Data Indexing:** Indexes transfers into a PostgreSQL database.
* **Docker Compose Setup:** Simplifies deployment and management.
* **REST API for Real-Time Queries:** Provides REST API endpoints for querying real-time transfer data.
* **Extensible Design:** Designed to allow future expansion to other data types, such as swaps and events.
* **Heimdahl Platform Integration:** Designed to seamlessly integrate with the Heimdahl onchain data aggregator
  platform.

## Getting Started

### Prerequisites

* VPS with static IP (any VPS Cloud provider, Digital Ocean, EC2, GCP...)
* Docker
* Docker Compose
* Heimdahl.xyz API Key

### Installation and Running

1. Clone the repository:

   ```bash
   git clone [https://github.com/yourusername/forseti.git](https://www.google.com/search?q=https://github.com/yourusername/forseti.git)
   cd forseti
   ```

2. Start the services using Docker Compose:

   ```bash
   docker compose up -d
   ```

   This command will build and start the PostgreSQL database and the transfer collector service.

3. Run migrations

      ```bash
      make migrate-up  
    ```

4. Access the API:

    * The REST API will be available at `http://localhost:9009`.

### Configuration

The configuration is handled via environment variables within the `docker-compose.yml`
and .env.example files. 
