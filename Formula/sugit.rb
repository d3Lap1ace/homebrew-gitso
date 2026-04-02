class Sugit < Formula
  desc      "Clone GitHub repo into ~/Documents/GitHub.com/<owner>/<repo>"
  homepage  "https://github.com/d3Lap1ace/sugit"
  url       "https://github.com/d3Lap1ace/sugit/archive/refs/tags/v0.1.4.tar.gz"
  sha256    "PUT_SOURCE_TARBALL_SHA256_HERE"
  license   "MIT"

  livecheck do
    url :stable          # check the same tag source used by :stable
    strategy :github_latest
  end

  depends_on "go" => :build

  def install
    # 与 goreleaser 的 ldflags 保持一致
    commit = Utils.git_short_head(length: 7)
    date    = Time.now.utc.strftime("%Y-%m-%dT%H:%M:%SZ")

    ldflags = %W[
      -s -w
      -X sugit/cmd.version=#{version}
      -X sugit/cmd.commit=#{commit}
      -X sugit/cmd.date=#{date}
    ]

    ENV["CGO_ENABLED"] = "0"
    system "go", "build", *std_go_args(
      ldflags:  ldflags,
      output:   bin/"sugit",
    ), "./cmd/sugit"
  end

  test do
    assert_match "Clone GitHub repo", shell_output("#{bin}/sugit --help")
  end
end
