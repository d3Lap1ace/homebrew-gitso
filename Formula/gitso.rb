class Gitso < Formula
  desc      "Clone GitHub repo into ~/Code/GitHub.com/<owner>/<repo>"
  homepage  "https://github.com/d3Lap1ace/gitso"
  url       "https://github.com/d3Lap1ace/gitso/archive/refs/tags/v0.1.5.tar.gz"
  sha256    "5dcd6f8d791e2d6750ae33c10cb8f58f00bda3e77c04fdb8bc00f70b9c6036db"
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
      -X gitso/cmd.version=#{version}
      -X gitso/cmd.commit=#{commit}
      -X gitso/cmd.date=#{date}
    ]

    ENV["CGO_ENABLED"] = "0"
    system "go", "build", *std_go_args(
      ldflags:  ldflags,
      output:   bin/"gitso",
    ), "."
  end

  test do
    assert_match "Clone GitHub repo", shell_output("#{bin}/gitso --help")
  end
end
