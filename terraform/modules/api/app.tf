resource "heroku_app" "default" {
  name   = var.app_name
  region = "us"
  stack  = "container"

  config_vars = {
    GOVERSION = "go1.17"
  }

  buildpacks = [
    "heroku/go"
  ]
}

# Build code & release to the app
resource "heroku_build" "resume-be" {
  app_id = heroku_app.default.id

  source {
    url     = "https://github.com/albertopformoso/resume-be/archive/refs/tags/${var.build_version}.tar.gz"
    version = "${var.build_version}"
  }
}
