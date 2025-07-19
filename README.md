# 🌍 IP Country Logger

A backend service in Golang that logs website visits based on users' IP addresses. It converts each IP to a country and city using a public geolocation API, stores the visit with a tag and timestamp in a PostgreSQL database, and provides REST APIs to query analytics like total visits per country or tag.

---

## 🚀 Features

- Accepts user IP + website tag and logs the visit
- Fetches country and city using `ip-api.com`
- Stores visit data (IP, country, city, tag, timestamp)
- Provides REST APIs to:
  - Add new visit
  - View all visits (with optional filters)
  - Get visit summaries by country and tag

---

## 🛠️ Tech Stack

- **Language**: Golang
- **Framework**: Gin (for REST APIs)
- **Database**: PostgreSQL / SQLite (for local)
- **ORM**: GORM
- **Geolocation API**: ip-api.com
- **Env config**: godotenv

---

## 📁 Folder Structure

