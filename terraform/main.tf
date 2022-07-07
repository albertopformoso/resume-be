module "api" {
  source        = "./modules/api"
  app_name      = var.app_name
  build_version = var.build_version
}

module "addons" {
  source = "./modules/addons"
  app_id = module.api.heroku_app_id
}