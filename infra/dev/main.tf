module "notifications" {
  source = "../modules/main"

  resend_api_key    = var.resend_api_key
  expo_access_token = var.expo_access_token
}
