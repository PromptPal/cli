#!/bin/bash

AMD_DIST=./dist/pp-mac_darwin_amd64_v1/promptpal
ARM_DIST=./dist/pp-mac_darwin_arm64/promptpal

if [ -f "$ARM_DIST" && -f "$AMD_DIST" ]; then
	gon .gon.hcl
else 
	echo "waiting for build to finish..."
fi