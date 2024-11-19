# Blog Application 

## Table of Contents
- [About](#about)
- [How To Build](#how-to-build)
- [Documentation](#documentation)

## About 

A lightweight and dynamic blog application built using **Golang**, **Templ** for templating, and **HTMX** for seamless interactivity without JavaScript frameworks.

## Features

- **Responsive Design**: Designed to work on various devices.
- **Dynamic Interactions**: Powered by HTMX for modern, AJAX-style requests.
- **Templated Views**: Leverages Templ for server-side rendering with clean, type-safe templates.
- **CRUD Functionality**: Create, read, update, and delete blog posts.
- **Golang Backend**: Fast and reliable server written in Go.


## How To Build

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- **Golang** (v1.23+)
- **Templ** (v0.2.793+)
- **Docker** (v26.1.4+)
- **Air** (v1.61.1+) - for hot reloading when developing


### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/deej-tsn/blog-go.git
   cd blog-app
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Configure your environment:

   - Copy the `.env.example` file and rename it to `.env`.
   - Modify the `.env` file with your database credentials and other configurations.

4. Add Markdown files to data/posts/ folder

5. Start the application:

   ```bash
   go run main.go
   ```

6. Open the application in your browser:

   ```
   http://localhost:8080
   ```

### Project Structure

```plaintext

blog-go/
├── cmd/
├── data
│   └── posts/
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── helper
│   │   └── render.go
│   ├── models
│   │   └── models.go
│   └── routes
│       └── controller.go
├── Makefile
├── README.md
└── web
    ├── components/
    └── public
        ├── assets/
        ├── css/
        └── scripts/

```

### Technologies Used

- **Golang**: For backend development.
- **Templ**: For type-safe and concise server-side templates.
- **HTMX**: To add modern AJAX-like interactions with minimal JavaScript.
- **Docker** : 

### Usage

1. Access the homepage to view all blog posts.
2. Click on post to be view the full file.

### Development

To make changes to the application:

1. Modify templates in the `components/` directory.
2. Update routes in the `routes/` directories.
3. Use `htmx` attributes in templates to add interactivity without extra JavaScript.


## Documentation 




### Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m "Add feature name"`.
4. Push to the branch: `git push origin feature-name`.
5. Open a Pull Request.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Made with ❤️ using Golang, Templ, and HTMX.
```
