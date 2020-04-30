#!/bin/bash

readonly CHART_NAME=emitter
readonly CHART_DIR=./deploy/helm

CONSUL_ADDR=${CONSUL_ADDR:=127.0.0.1:8500}
ENV=${ENV:=dev}
VERSION=${VERSION:=`git describe --abbrev=0`-`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD`}
GCP_PROJECT=${GCP_PROJECT:=videocoin-network}

function log {
  local readonly level="$1"
  local readonly message="$2"
  local readonly timestamp=$(date +"%Y-%m-%d %H:%M:%S")
  >&2 echo -e "${timestamp} [${level}] [$SCRIPT_NAME] ${message}"
}

function log_info {
  local readonly message="$1"
  log "INFO" "$message"
}

function log_warn {
  local readonly message="$1"
  log "WARN" "$message"
}

function log_error {
  local readonly message="$1"
  log "ERROR" "$message"
}

function update_deps() {
    log_info "Syncing dependencies..."
    helm dependencies update --kube-context ${KUBE_CONTEXT} ${CHART_DIR}
}

function has_jq {
  [ -n "$(command -v jq)" ]
}

function has_consul {
  [ -n "$(command -v consul)" ]
}

function has_helm {
  [ -n "$(command -v helm)" ]
}

function get_vars() {
    log_info "Getting variables..."
    readonly KUBE_CONTEXT=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/common/kube_context`
    readonly ACCOUNTS_RPC_ADDR=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/vars/accountsRpcAddr`
    readonly STREAM_MANAGER_CONTRACT_ADDR=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/vars/streamManagerContractAddr`
    readonly STAKING_MANAGER_CONTRACT_ADDR=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/vars/stakingManagerContractAddr`
    readonly FAUCET_URL=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/vars/faucetUrl`

    readonly SYMPHONY_ADDR=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/symphonyAddr`
    readonly OAUTH_CLIENT_ID=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/oauthClientId`
    readonly RPC_KEY=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/rpcKey`
    readonly FAUCET_KEY=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/faucetKey`
    readonly MANAGER_KEY=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/managerKey`
    readonly MANAGER_SECRET=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/managerSecret`
    readonly CLIENT_SECRET=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/clientSecret`
    readonly VALIDATOR_KEY=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/validatorKey`
    readonly VALIDATOR_SECRET=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/validatorSecret`
    readonly MQ_URI=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/mqUri`
    readonly SENTRY_DSN=`consul kv get -http-addr=${CONSUL_ADDR} config/${ENV}/services/${CHART_NAME}/secrets/sentryDsn`
}

function deploy() {
    log_info "Deploying ${CHART_NAME} version ${VERSION}"
    helm upgrade \
        --kube-context "${KUBE_CONTEXT}" \
        --install \
        --set image.repository="gcr.io/${GCP_PROJECT}/${CHART_NAME}" \
        --set image.tag="${VERSION}" \
        --set config.accountsRpcAddr="${ACCOUNTS_RPC_ADDR}" \
        --set config.streamManagerContractAddr="${STREAM_MANAGER_CONTRACT_ADDR}" \
        --set config.stakingManagerContractAddr="${STAKING_MANAGER_CONTRACT_ADDR}" \
        --set secrets.symphonyAddr="${SYMPHONY_ADDR}" \
        --set secrets.oauthClientId="${OAUTH_CLIENT_ID}" \
        --set secrets.rpcKey="${RPC_KEY}" \
        --set secrets.faucetKey="${FAUCET_KEY}" \
        --set secrets.mqUri="${MQ_URI}" \
        --set secrets.managerKey="${MANAGER_KEY}" \
        --set secrets.managerSecret="${MANAGER_SECRET}" \
        --set secrets.clientSecret="${CLIENT_SECRET}" \
        --set secrets.validatorKey="${VALIDATOR_KEY}" \
        --set secrets.validatorSecret="${VALIDATOR_SECRET}" \
        --set secrets.sentryDsn="${SENTRY_DSN}" \
        --wait ${CHART_NAME} ${CHART_DIR}
}

if ! $(has_jq); then
    log_error "Could not find jq"
    exit 1
fi

if ! $(has_consul); then
    log_error "Could not find consul"
    exit 1
fi

if ! $(has_helm); then
    log_error "Could not find helm"
    exit 1
fi

get_vars
update_deps
deploy

exit $?