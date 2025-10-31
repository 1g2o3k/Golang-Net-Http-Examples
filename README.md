# Donate Buy me a Coffe:

https://buymeacoffee.com/gokalp

# Go User Directory

A simple web application built with Go's net/http package that displays a directory of users with search functionality. This project showcases basic web development in Go, including routing, HTML templating, and client-side JavaScript for interactivity.

## Features

- **User Directory**: Displays 12 users in a responsive grid layout on the main page.
- **Search Functionality**: Real-time search to filter users by name using JavaScript.
- **Profile Pages**: Each user has a dedicated profile page showing their name, surname, and age.
- **Responsive Design**: Modern UI built with Tailwind CSS, ensuring compatibility across devices.
- **Icons and Fonts**: Incorporates Font Awesome icons and Courier Prime monospace font for a clean look.
- **Hover Effects**: Smooth animations and transitions for an enhanced user experience.

## Installation

1. Ensure you have Go installed on your system (version 1.16 or later recommended).
2. Clone the repository:
   ```bash
   git clone https://github.com/1g2o3k/Golang-Net-Http-Examples.git
   cd Golang-Net-Http-Examples
   ```
3. Run the application:
   ```bash
   go run *.go
   ```
4. Open your browser and navigate to `http://localhost:8080`.

## Usage

- **Main Page**: View the grid of users. Use the search box at the top to filter users by typing their name.
- **Profile View**: Click on any user card to navigate to their individual profile page.
- **Navigation**: Use the "Back to Main Page" button on profile pages to return to the directory.

## Project Structure

- `main.go`: Main application file that initializes the server, sets up routes, and calls user handler functions.
- `user1.go` to `user12.go`: Individual files defining structs and handler functions for each user's profile page.
- `templates/index.html`: HTML template for the main user directory page.
- `templates/user1.html` to `templates/user12.html`: HTML templates for individual user profile pages.

## Technologies Used

- **Go**: Backend server using the standard `net/http` package.
- **HTML Templates**: Go's `html/template` package for server-side rendering.
- **Tailwind CSS**: Utility-first CSS framework for styling.
- **Font Awesome**: Icon library for visual elements.
- **Flowbite**: Additional UI components and JavaScript utilities.
- **JavaScript**: Client-side scripting for search functionality.

