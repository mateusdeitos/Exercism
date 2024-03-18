# if .env not found alert
if [ ! -f .env ]; then
  echo "Please create a .env file and set the TOKEN variable. See .env.example for an example."
  exit 1
fi

# exports the variables from .env
export $(grep -v '^#' .env | xargs -d '\n')

# run the docker build command
docker build . -t $IMAGE_NAME --build-arg TOKEN=$TOKEN
