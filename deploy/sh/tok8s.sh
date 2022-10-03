#!/bin/bash
set -eu

PROJECT_NAME='go_project'
SSH_IP='10.0.16.9'
SSH_USER='zhengqilin'
UPLOAD_DIR="/home/zhengqilin/deploy/${PROJECT_NAME}"
DEPLOY_FILE="${UPLOAD_DIR}/deploy.yaml"

echo "[tok8s.1] rm -rf old deploy.yaml!"
ssh ${SSH_USER}@${SSH_IP} "rm -rf ${DEPLOY_FILE}" || true
echo "[tok8s.2] mkdir"
ssh ${SSH_USER}@${SSH_IP} "mkdir -p ${UPLOAD_DIR}"

echo "[tok8s.3]"
scp -r deploy/deploy.yaml ${SSH_USER}@${SSH_IP}:${DEPLOY_FILE}

echo "[tok8s.4]"
ssh ${SSH_USER}@${SSH_IP} "kubectl apply -f ${DEPLOY_FILE}"
echo "[tok8s.5] ${SSH_USER}@${SSH_IP} 'kubectl apply -f ${DEPLOY_FILE}'"