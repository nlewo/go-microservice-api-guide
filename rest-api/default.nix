{ pkgs ? import (import ./nix/sources.nix).nixpkgs {} }:

let
  app = pkgs.buildGoPackage {
    name = "go-microservice-api-guide";
    goPackagePath = "github.com/johan-lejdung/go-microservice-api-guide/rest-api";
    src = ./.;
  };
  
in pkgs.dockerTools.buildImage {
  name = "rest-api";
  config = {
    CMD = [ "${app}/bin/rest-api" ];
  };
}
