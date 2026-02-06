#!/bin/bash

# Integration Test Runner Script
# This script helps run integration tests with proper environment setup

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Athenahealth Integration Test Runner ===${NC}\n"

# Check if required environment variables are set
if [ -z "$ATHENA_PRACTICE_ID" ] || [ -z "$ATHENA_API_KEY" ] || [ -z "$ATHENA_API_SECRET" ]; then
    echo -e "${RED}Error: Required environment variables not set${NC}"
    echo ""
    echo "Please set the following environment variables:"
    echo "  export ATHENA_PRACTICE_ID=your-practice-id"
    echo "  export ATHENA_API_KEY=your-api-key"
    echo "  export ATHENA_API_SECRET=your-api-secret"
    echo ""
    echo "Optional variables for Risk Contract tests:"
    echo "  export ATHENA_TEST_RISK_CONTRACT_ID=123"
    echo "  export ATHENA_TEST_RISK_CONTRACT_NAME=\"Contract Name\""
    echo ""
    exit 1
fi

echo -e "${GREEN}Environment configuration:${NC}"
echo "  ATHENA_PRACTICE_ID: ${ATHENA_PRACTICE_ID}"
echo "  ATHENA_API_KEY: ${ATHENA_API_KEY:0:10}..."
echo "  ATHENA_API_SECRET: ${ATHENA_API_SECRET:0:10}..."
echo ""

if [ -n "$ATHENA_TEST_RISK_CONTRACT_ID" ]; then
    echo "  ATHENA_TEST_RISK_CONTRACT_ID: ${ATHENA_TEST_RISK_CONTRACT_ID}"
fi

if [ -n "$ATHENA_TEST_RISK_CONTRACT_NAME" ]; then
    echo "  ATHENA_TEST_RISK_CONTRACT_NAME: ${ATHENA_TEST_RISK_CONTRACT_NAME}"
fi

echo ""

# Enable integration tests
export ATHENA_RUN_INTEGRATION_TESTS=true

# Determine which tests to run
TEST_PATTERN="${1:-TestIntegration}"

echo -e "${YELLOW}Running tests matching: ${TEST_PATTERN}${NC}\n"

# Run the tests
cd "$(dirname "$0")"
go test -v -run "${TEST_PATTERN}" ./athenahealth

echo ""
echo -e "${GREEN}✓ Tests completed${NC}"
