# Tollama

Tollama is a terminal-based user interface (TUI) application that allows users to interact with the Ollama API. Users can list available models, select and download a model, and chat with the selected model.

## Features

- List available models from the Ollama API
- Select and download a model
- Chat with the selected model

## Installation

1. **Install Ollama:**

   ```bash
   sudo pacman -S ollama
   ```

2. **Clone the repository:**

   ```bash
   git clone https://github.com/Tollama/Tollama.git
   cd Tollama
   ```

3. **Run the application using the provided bash script:**

   ```bash
   ./tollama.sh
   ```

   The `tollama.sh` script will:
   - Install Ollama
   - Start the Ollama server in the background
   - Run the Tollama application

## Usage

1. **Start the Ollama server:**

   The `tollama.sh` script will automatically start the Ollama server. If you need to start it manually, use:

   ```bash
   ollama serve &
   ```

2. **Run the Tollama application:**

   The `tollama.sh` script will automatically run the Tollama application. If you need to run it manually, use:

   ```bash
   go run src/*.go
   ```

3. **Navigate the TUI:**
   - Use the arrow keys to navigate the menu.
   - Press `Enter` to select an option.
   - Press `q` to quit the application.

## Known Issues

- **Models are not fetched yet:** There is a known bug where the models are not being fetched from the Ollama API. This issue is currently being investigated.

## Compatibility

- **Arch Linux:** This project currently works on Arch Linux unless you install Ollama differently. If you are using a different distribution or operating system, you can still adapt the instructions accordingly.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License.

