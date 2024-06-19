# Distributed File System

## Project Description
A distributed file system implemented in Go using gRPC. The system includes a master server for metadata management, chunk servers for data storage, and clients for data operations.

## Requirements
- Metadata management by the master server.
- Data storage and replication by chunk servers.
- Client operations for reading and writing data.
- Fault tolerance and consistency mechanisms.

## Initial Design
- **Master Server**: Manages file metadata, chunk mappings, and handles client requests.
- **Chunk Servers**: Store data chunks, handle read/write requests, and send heartbeats to the master.
- **Clients**: Communicate with the master server to obtain chunk locations and perform data operations.
- **Replication and Consistency**: Ensure data redundancy and consistency using replication and versioning.
