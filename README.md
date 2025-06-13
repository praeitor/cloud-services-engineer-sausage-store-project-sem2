# Sausage Store — Магазин колбасок
![Deploy Workflow](https://github.com/praeitor/cloud-services-engineer-sausage-store-project-sem2/actions/workflows/deploy.yaml/badge.svg?branch=main)

Учебный проект в рамках курса DevOps-мастера (2 семестр), реализованный с применением современных инструментов контейнеризации, CI/CD и мониторинга.

## Стек технологий

- **Backend:** Spring Boot (Java)
- **Frontend:** Vue.js
- **Reporting:** Backend-report на Python
- **База данных:** PostgreSQL, MongoDB
- **CI/CD:** GitHub Actions + Helm + Docker
- **Kubernetes:** Деплой в кластер с Helm-чартами
- **Хранилище секретов:** Vault (under construction)

---

## Развертывание

### 1. Требования

- Установлен `kubectl` и доступ к кластеру
- Helm v3.14+
- Доступ к Nexus (ChartMuseum) и Docker Hub

### 2. CI/CD через GitHub Actions

```yaml
on:
  push:
    branches:
      - main
```

- Сборка образов и пуш в DockerHub
- Версионирование Helm-чарта и отправка в Nexus
- Автоматический деплой в Kubernetes

### 3. Деплой вручную

```bash
helm upgrade --install sausage-store chartmuseum/sausage-store-chart \\
  --namespace <namespace> \\
  --set backend.image.repository=<dockerhub>/sausage-backend \\
  --set backend.image.tag=latest \\
  --set frontend.image.repository=<dockerhub>/sausage-frontend \\
  --set frontend.image.tag=latest \\
  --set backendreport.image.repository=<dockerhub>/sausage-backend-report \\
  --set backendreport.image.tag=latest
```

---

## Структура проекта

```
sausage-store/
├── backend/
├── frontend/
├── backend-report/
├── sausage-store-chart/
│   ├── charts/
│   ├── templates/
│   └── values.yaml
├── .github/workflows/
│   └── deploy.yaml
```

---

## Secrets (GitHub)

- `DOCKER_USER`
- `DOCKER_PASSWORD`
- `NEXUS_HELM_REPO`
- `NEXUS_HELM_REPO_USER`
- `NEXUS_HELM_REPO_PASSWORD`
- `KUBE_CONFIG`
- `VAULT_HOST`
- `VAULT_TOKEN`

---

## Helm-чарт

Единый чарт `sausage-store-chart` управляет:

- `backend`
- `frontend`
- `backend-report`
- `infra` (PostgreSQL, MongoDB)

---

## Тестирование

- Запуск юнит-тестов на стадии `build`
- Проверка успешного деплоя в кластер

---

## Авторы

- Денис Булгаков