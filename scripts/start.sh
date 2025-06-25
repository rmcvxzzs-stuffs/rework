#!/bin/bash

while true; do
    echo "=== Horizon Server ===" # menu
    echo "PLEASE DO NOTE: DO NOT RUN THE SERVER, CODE IS ABSENT AS I DID NOT CODED IT BEFORE. THIS IS JUST A TESTING." # lulz
    echo "1) Start server" # would execute "docker compose up"
    echo "2) Exit" # exits
    read -p "Choose an option [1-2]: " choice

    case "$choice" in
        1)
            echo "Starting server with docker compose..."
            docker compose up
            break
            ;;
        2)
            echo "Exiting..."
            exit 0
            ;;
        *)
            echo "Invalid option. Please enter 1 or 2."
            ;;
    esac
done
