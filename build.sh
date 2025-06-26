#!/bin/bash
# This script is used to build a zip and/or a docker image of a plugin.
# The script takes in optional arguments:
# --zip: build only the zip file
# --image: build only the docker image
# --upload: both the zip and image will be built and uploaded zip to the release server
# --override-registry-url: override the Registry URL to which the image is uploaded, zip file will still contain original URL
# --override-registry-org: override the Organization in which this image is uploaded, zip file will still contain original Organization
# If no argument is passed, both the zip and image will be built

read_properties(){
  # Remove the tmp directory and create it again
  rm -rf tmp 2>/dev/null
  mkdir tmp 2>/dev/null

  # Remove all the carriage returns
  sed 's/\r$//' project.properties > tmp/project.properties

  # Read project properties from project.properties file and set them as variables
  . ./tmp/project.properties

  # Remove project.properties from tmp
  rm tmp/project.properties
}

build_zip(){

  # Copy the resources directory contents to tmp
  cp -R resources/. tmp/

  if [ "$(uname)" = "Darwin" ]; then
    echo "Detected MAC OS X platform"

    sed -i '' 's/@project.name@/'"$PLUGIN"'/g' tmp/plugin-version.properties
    sed -i '' 's/@project.version@/'"$VERSION"'/g' tmp/plugin-version.properties
    if [ -s tmp/type-definitions.xml ]; then
      sed -i '' 's/@project.name@/'"$PLUGIN"'/g' tmp/type-definitions.xml
      sed -i '' 's/@project.version@/'"$VERSION"'/g' tmp/type-definitions.xml
      sed -i '' 's/@registry.url@/'"$REGISTRY_URL"'/g' tmp/type-definitions.xml
      sed -i '' 's/@registry.org@/'"$REGISTRY_ORG"'/g' tmp/type-definitions.xml
    fi
    if [ -s tmp/type-definitions.yaml ]; then
      sed -i '' 's/@project.name@/'"$PLUGIN"'/g' tmp/type-definitions.yaml
      sed -i '' 's/@project.version@/'"$VERSION"'/g' tmp/type-definitions.yaml
      sed -i '' 's/@registry.url@/'"$REGISTRY_URL"'/g' tmp/type-definitions.yaml
      sed -i '' 's/@registry.org@/'"$REGISTRY_ORG"'/g' tmp/type-definitions.yaml
    fi

  elif [ "$(expr substr $(uname -s) 1 5)" = "Linux" ]; then
    echo "Detected GNU/Linux platform"

    sed -i.bak 's/@project.name@/'"$PLUGIN"'/g' tmp/plugin-version.properties
    sed -i.bak 's/@project.version@/'"$VERSION"'/g' tmp/plugin-version.properties
    if [ -s tmp/type-definitions.xml ]; then
      sed -i.bak 's/@project.name@/'"$PLUGIN"'/g' tmp/type-definitions.xml
      sed -i.bak 's/@project.version@/'"$VERSION"'/g' tmp/type-definitions.xml
      sed -i.bak 's/@registry.url@/'"$REGISTRY_URL"'/g' tmp/type-definitions.xml
      sed -i.bak 's/@registry.org@/'"$REGISTRY_ORG"'/g' tmp/type-definitions.xml
      rm tmp/type-definitions.xml.bak
    fi
    if [ -s tmp/type-definitions.yaml ]; then
      sed -i.bak 's/@project.name@/'"$PLUGIN"'/g' tmp/type-definitions.yaml
      sed -i.bak 's/@project.version@/'"$VERSION"'/g' tmp/type-definitions.yaml
      sed -i.bak 's/@registry.url@/'"$REGISTRY_URL"'/g' tmp/type-definitions.yaml
      sed -i.bak 's/@registry.org@/'"$REGISTRY_ORG"'/g' tmp/type-definitions.yaml
      rm tmp/type-definitions.yaml.bak
    fi
    rm tmp/plugin-version.properties.bak
  fi

  # Create the build directory and remove any previously created zip file
  mkdir build 2>/dev/null
  rm -f "build/$PLUGIN-$VERSION.zip" 2>/dev/null

  # Create a zip file from the contents of the tmp directory and place it in the build directory
  cd tmp && zip -r "../build/$PLUGIN-$VERSION.zip" . && cd ..
  echo "Build completed: $PLUGIN-$VERSION.zip"
  # Remove the tmp directory
  rm -rf tmp
}

build_image(){
  SET_REGISTRY_URL="$REGISTRY_URL"
  SET_REGISTRY_ORG="$REGISTRY_ORG"
  if [ -n "$OVERRIDE_REGISTRY_URL" ]; then
    echo "Overriding image URL for image registry upload: $OVERRIDE_REGISTRY_URL, zip file will still contain original URL: $REGISTRY_URL"
    SET_REGISTRY_URL="$OVERRIDE_REGISTRY_URL"
  fi
  if [ -n "$OVERRIDE_REGISTRY_ORG" ]; then
    echo "Overriding image Organization for image registry upload: $OVERRIDE_REGISTRY_ORG, zip file will still contain original URL: $REGISTRY_ORG"
    SET_REGISTRY_ORG="$OVERRIDE_REGISTRY_ORG"
  fi
  # Set build date
  BUILD_DATE=$(date +%Y-%m-%d)

  # Build docker image and push to registry
  FULL_TAG="$SET_REGISTRY_URL/$SET_REGISTRY_ORG/$PLUGIN:$VERSION"
  if docker build --build-arg VERSION="$VERSION" --build-arg BUILD_DATE="$BUILD_DATE" --tag "$FULL_TAG" .; then
    if docker image push "$FULL_TAG"; then
      echo "Build and push completed: $FULL_TAG"
    else
      echo "Push failed for $FULL_TAG"
    fi
  else
    echo "Build failed for $FULL_TAG"
  fi
}

upload_zip(){
  # upload the zip to the release server
  chmod +x ./xlw
  ./xlw plugin release install --file build/$PLUGIN-$VERSION.zip --config .xebialabs/config.yaml
}

parse_arguments() {
  local state="none"
  local current_flag=""
  while [[ $# -gt 0 ]]; do
    case $state in
      none)
        case $1 in
          --zip)
            BUILD_ZIP=true
            ;;
          --image)
            BUILD_IMAGE=true
            ;;
          --upload)
            UPLOAD_ZIP=true
            ;;
          --override-registry-url|--override-registry-org)
            current_flag="$1"
            state="read-value"
            ;;
          *)
            echo "Unknown option: $1"
            exit 1
            ;;
        esac
        ;;
      read-value)
        if [[ $1 != --* && -n $1 ]]; then
          case $current_flag in
            --override-registry-url)
              OVERRIDE_REGISTRY_URL="$1"
              ;;
            --override-registry-org)
              OVERRIDE_REGISTRY_ORG="$1"
              ;;
          esac
          state="none"
        else
          echo "Error: $current_flag requires a value."
          exit 1
        fi
        ;;
    esac
    shift
  done

  # Final validation for missing value
  if [[ $state == "read-value" ]]; then
    echo "Error: $current_flag requires a value."
    exit 1
  fi
}

# Initialize variables
BUILD_ZIP=false
BUILD_IMAGE=false
UPLOAD_ZIP=false
REGISTRY_URL=""
REGISTRY_ORG=""
OVERRIDE_REGISTRY_URL=""
OVERRIDE_REGISTRY_ORG=""

# Call the function with all arguments
read_properties
parse_arguments "$@"


if [ "$BUILD_ZIP" = true ]; then
  echo "Building zip..."
  build_zip
elif [ "$BUILD_IMAGE" = true ]; then
  echo "Building image..."
  build_image
elif [ "$UPLOAD_ZIP" = true ]; then
  echo "Building zip, image and uploading zip..."
  build_zip
  build_image
  upload_zip
else
  echo "Building zip and image..."
  build_zip
  build_image
fi
