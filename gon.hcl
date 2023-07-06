# gon.hcl
#
# The path follows a pattern
# ./dist/BUILD-ID_TARGET/BINARY-NAME
# source = ["./dist/${artifact}.dmg"]
source = ["./dist/promptpal.dmg"]
bundle_id = "com.annatarhe.promptpal.cli"

apple_id {
  username = "iamhele1994@gmail.com"
  password = "@env:AC_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Annatar He"
}