class Gitso < Formula
  desc      "Transparent Git proxy that organizes GitHub clones"
  homepage  "https://github.com/d3Lap1ace/gitso"
  url       "https://github.com/d3Lap1ace/gitso/archive/refs/tags/v0.1.5.tar.gz"
  sha256    "5dcd6f8d791e2d6750ae33c10cb8f58f00bda3e77c04fdb8bc00f70b9c6036db"
  license   "MIT"

  livecheck do
    url :stable          # check the same tag source used by :stable
    strategy :github_latest
  end

  depends_on "go" => :build
  depends_on "git"

  def install
    ENV["CGO_ENABLED"] = "0"
    system "go", "build", *std_go_args(
      ldflags:  %w[-s -w],
      output:   bin/"gitso",
    ), "."
  end

  test do
    assert_match "git version", shell_output("#{bin}/gitso --version")
  end
end
