#
# Rakefile for agent-payload
#

protoc_binary="protoc"
protoc_version="3.5.1"
gogo_dir="/tmp/gogo"

namespace :codegen do

  task :install_protoc do
    if `bash -c "protoc --version"` != "libprotoc ${protoc_version}"
      protoc_binary="/tmp/protoc#{protoc_version}"
      sh <<-EOF
        /bin/bash <<BASH
        if [ ! -f #{protoc_binary} ] ; then
          echo "Downloading protoc #{protoc_version}"
          cd /tmp
          if [ "$(uname -s)" = "Darwin" ] ; then
            curl -OL https://github.com/google/protobuf/releases/download/v#{protoc_version}/protoc-#{protoc_version}-osx-x86_64.zip
          else
            curl -OL https://github.com/google/protobuf/releases/download/v#{protoc_version}/protoc-#{protoc_version}-linux-x86_64.zip
          fi
          unzip protoc-#{protoc_version}*.zip
          mv bin/protoc #{protoc_binary}
        fi
BASH
      EOF
    end
  end

  task :protoc => ['install_protoc'] do
    sh <<-EOF
      /bin/bash <<BASH
      set -euo pipefail

      rm -rf #{gogo_dir}
      rm -rf /tmp/gogo-bin-*

      mkdir -p #{gogo_dir}/src/github.com/gogo
      git clone https://github.com/gogo/protobuf.git #{gogo_dir}/src/github.com/gogo/protobuf

      # Install v1.0.0
      pushd #{gogo_dir}/src/github.com/gogo/protobuf
      git checkout v1.0.0
      GOBIN=/tmp/gogo-bin-v1.0.0 GOPATH=#{gogo_dir} make clean install

      popd

      echo "Generating logs proto"
      PATH=/tmp/gogo-bin-v1.0.0 #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofast_out=$GOPATH/src --java_out=java proto/logs/agent_logs_payload.proto

      echo "Generating network-devices proto"
      PATH=/tmp/gogo-bin-v1.0.0 #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofast_out=$GOPATH/src proto/network-devices/agent.proto

      echo "Generating metrics proto (go)"
      PATH=/tmp/gogo-bin-v1.0.0 #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofast_out=$GOPATH/src proto/metrics/agent_payload.proto

      echo "Generating metrics proto (python)"
      #{protoc_binary} --proto_path=#{gogo_dir}/src:$GOPATH/src:./proto/metrics --python_out=python agent_payload.proto

      # Install the specific tag that the process-agent needs
      pushd #{gogo_dir}/src/github.com/gogo/protobuf
      git checkout d76fbc1373015ced59b43ac267f28d546b955683
      GOBIN=/tmp/gogo-bin-d76fbc1373015ced59b43ac267f28d546b955683 GOPATH=#{gogo_dir} make clean install

      popd

      echo "Generating process proto"
      PATH=/tmp/gogo-bin-d76fbc1373015ced59b43ac267f28d546b955683 #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofaster_out=$GOPATH/src proto/process/agent.proto
BASH
    EOF
  end

  desc 'Run all code generators.'
  multitask :all => [:protoc]

end

desc "Setup dependencies"
task :deps do
  system("dep ensure -v")
end

desc "Run tests"
task :test do
  cmd = "go list ./... | grep -v vendor | xargs go test -v "
  sh cmd
end

desc "Run all code generation."
task :codegen => ['codegen:all']

desc "Run all protobuf code generation."
task :protobuf => ['codegen:protoc']

task :default => [:deps, :test, :codegen]