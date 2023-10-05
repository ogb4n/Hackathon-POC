echo "Installing tailwindcss"
if command -v npm &> /dev/null
then
    npm i > /dev/null
    npx tailwindcss -i ./src/css/styles.css -o ./src/css/output.css
    echo "Tailwindcss installed and built"
else
    echo "npm could not be found"
    echo "Please install npm and try again"
    exit
fi