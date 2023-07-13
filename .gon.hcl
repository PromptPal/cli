# gon.hcl
#
# The path follows a pattern
# ./dist/BUILD-ID_TARGET/BINARY-NAME
# source = ["./dist/${artifact}.dmg"]
source = ["./dist/pp-mac_darwin_amd64_v1/promptpal", "./dist/pp-mac_darwin_arm64/promptpal"]
bundle_id = "com.annatarhe.promptpal.cli"

apple_id {
  username = "iamhele1994@gmail.com"
  password = "@env:AC_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Le He"
}