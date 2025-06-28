{
  description = "Zenta is a lightweight command-line tool designed for developers and terminal users who want to cultivate mindfulness, reduce distractions, and maintain awareness throughout their workday. Inspired by Zen philosophy and Stoic practice, zenta brings calm, clarity, and presence into the world of deep work.";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs";

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f system);
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = import nixpkgs { inherit system; };
        in
        {
          default = pkgs.buildGoModule {
            pname = "zenta";
            version = "0.3.5";

            src = ./.;

            vendorHash = null;

            meta = with pkgs.lib; {
              description = "Zenta is a lightweight command-line tool designed for developers and terminal users who want to cultivate mindfulness, reduce distractions, and maintain awareness throughout their workday. Inspired by Zen philosophy and Stoic practice, zenta brings calm, clarity, and presence into the world of deep work.";
              homepage = "https://github.com/e6a5/zenta";
              license = licenses.mit;
              maintainers = [ maintainers.e6a5 ];
            };
          };
        }
      );

      devShells = forAllSystems (system:
        let pkgs = import nixpkgs { inherit system; };
        in {
          default = pkgs.mkShell {
            buildInputs = [ pkgs.go pkgs.git ];
          };
        }
      );
    };
}
