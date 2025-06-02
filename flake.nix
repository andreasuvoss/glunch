{
  description = "Flake for the glunch application";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      systems = [
        "x86_64-linux"
        "x86_64-darwin"
        "aarch64-linux"
        "aarch64-darwin"
      ];
      eachSystem = nixpkgs.lib.genAttrs systems;
      pkgsFor = eachSystem (
        system:
        import nixpkgs {
          inherit system;
        }
      );
    in
    {
      packages = eachSystem (system: {
        default = pkgsFor.${system}.buildGoModule {
          pname = "glunch";
          version = "0.0.3";
          src = ./.;
          vendorHash = null;

          meta = {
            description = "Get the lunch menu at work";
            mainProgram = "glunch";
          };
        };
      });
    };
}
