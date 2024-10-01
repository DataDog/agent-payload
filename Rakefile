# Rakefile for agent-payload

# where all toolchains (protobuf compilers, protobuf plugins,etc) are stored
toolchain_dir=File.join(File.dirname(__FILE__), "toolchains")
# binaries
toolchain_bin_dir=File.join(toolchain_dir, "bin")

# include path for the legacy and main protobuf compilers respectively
toolchain_include_dir=File.join(toolchain_dir, "include", "proto")

protoc_version="21.12"
protoc_binary=File.join(toolchain_bin_dir, "protoc" + protoc_version)

gogo_tag = "v1.3.2"
gogo_dir=File.join(toolchain_dir,  "gogo")
gogo_include = "#{toolchain_dir}/gogo/src:#{toolchain_dir}/gogo/src/github.com/gogo/protobuf/protobuf"
gogo_bin = File.join(toolchain_bin_dir, "gogo-bin-#{gogo_tag}")
protoc_jsonschema_version="73d5723"

### toolchains is meant to store a cache of all binary dependencies needed to build the agent-payload.
### this rakefile will download those dependencies on the fly if needed.
### the structure of toolchains is as follows:
### /toolchains
###     /toolchains/bin -- contains any binaries used for building
###         /toolchains/bin/protoc21.12 -- the new protobuf compiler
###         /toolchains/bin/protoc-gen-go -- the go code generator for protoc
###         /toolchains/bin/protoc-gen-XXX -- other protobuf generators (vtproto, java, etc.)
###     /toolchains/include -- contains any protobuf library files for common types (like proto.Duration)
###         /toolchains/includes/proto -- protobuf libraries bundled with the new protobuf compiler
### =>  /toolchains/gogo -- temp directory used to build the gogo_faster generator

namespace :codegen do
  task :clean do
      sh "rm -rf #{gogo_dir}"
      sh "rm -rf #{toolchain_dir}"
  end

  task :setup_gogo => ['install_protoc_all'] do
    if ! Dir.exist?(gogo_dir) then
      directory "#{gogo_dir}/src/github.com/gogo"
      sh "git clone https://github.com/gogo/protobuf.git #{gogo_dir}/src/github.com/gogo/protobuf"
    else
      puts "gogo already cloned into #{gogo_dir}"
    end

    Dir.chdir("#{gogo_dir}/src/github.com/gogo/protobuf") do
      sh "git checkout #{gogo_tag}"
      sh "PATH=$PATH:/tmp GOBIN=#{toolchain_bin_dir} GOPATH=#{gogo_dir} make clean install"
    end
  end


  task 'install_protoc_all' => ['install_protoc']

  task :install_protoc  do
    protoversion = `bash -c "#{protoc_binary} --version"`.strip
    if  protoversion == "libprotoc 3.21.12"
      next # bail out of task
    end
    puts "got #{protoversion}"


    puts "Downloading protoc #{protoc_version}"
    Dir.chdir("/tmp") do
      if /Darwin/.match(`uname -s`)
        sh "curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v#{protoc_version}/protoc-#{protoc_version}-osx-x86_64.zip"
      else

        sh "curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v#{protoc_version}/protoc-#{protoc_version}-linux-x86_64.zip"
      end

      sh "mkdir -p #{toolchain_bin_dir}"
      sh "rm -rf include/" # make sure there is no stale include
      sh "unzip -o protoc-#{protoc_version}*.zip"
      sh "mv bin/protoc #{protoc_binary}"

      sh "mkdir -p #{toolchain_include_dir}"
      sh "rm -rf #{toolchain_include_dir}/**"
      sh "mv include/* #{toolchain_include_dir}"
    end
  end

  task :protoc => ['install_protoc', 'setup_gogo'] do

    puts "installing go generators"
    sh "GOPATH=#{toolchain_dir} go install github.com/leeavital/protoc-gen-gostreamer@v0.1.0"
    sh "GOPATH=#{toolchain_dir} go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1"
    sh "GOPATH=#{toolchain_dir} go install github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema@#{protoc_jsonschema_version}"
    sh "GOPATH=#{toolchain_dir} go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.5.0"

    puts "Generating logs proto"
    sh "PATH=#{toolchain_bin_dir} #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:. --gogofast_out=$GOPATH/src proto/logs/agent_logs_payload.proto"

    puts "Generating metrics proto"
    sh "PATH=#{toolchain_bin_dir} #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_include}:#{toolchain_include_dir}:. --gogofast_out=$GOPATH/src proto/metrics/agent_payload.proto"

    puts "generatoring process proto"
    sh "PATH=#{toolchain_bin_dir} #{protoc_binary} --proto_path=#{toolchain_include_dir}:#{gogo_include}:. --gogofaster_out=$GOPATH/src proto/process/*.proto"
    sh "PATH=#{toolchain_bin_dir} #{protoc_binary} --proto_path=$GOPATH/src:#{gogo_dir}/src:.  --gostreamer_out=$GOPATH/src proto/process/*.proto"
    sh "mv v5/process/proto/process/*.go process"

    puts "generating contlcycle proto"
    sh "PATH=#{toolchain_bin_dir} #{protoc_binary} --proto_path=#{toolchain_include_dir}:. --go_out=$GOPATH/src proto/contlcycle/contlcycle.proto"
    sh "PATH=#{toolchain_bin_dir}  #{protoc_binary} --proto_path=#{toolchain_include_dir}:. --go_out=$GOPATH/src proto/contimage/contimage.proto"

    puts "generating autoscaling proto"
    sh "PATH=#{toolchain_bin_dir} #{protoc_binary} --proto_path=$GOPATH/src:#{toolchain_include_dir}:. --go_out=$GOPATH/src --jsonschema_out=type_names_with_no_package:jsonschema proto/autoscaling/kubernetes/autoscaling.proto"
    sh "PATH=#{toolchain_bin_dir} #{protoc_binary} --proto_path=#{toolchain_include_dir}:. --go_out=$GOPATH/src proto/deps/github.com/CycloneDX/specification/schema/bom-1.4.proto"

    puts "generatoring sbom proto"
    sh "PATH=#{toolchain_bin_dir}  #{protoc_binary} --proto_path=#{toolchain_include_dir}:. --go_out=$GOPATH/src proto/sbom/sbom.proto"



    sh <<~EOD.tr("\n", "")
      PATH=#{toolchain_bin_dir} #{protoc_binary} --proto_path=$GOPATH/src:.
        --java_out=java
        --go_out=$GOPATH/src
        --go-vtproto_out=$GOPATH/src
        --go-vtproto_opt=features=pool+marshal+unmarshal+size
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.SecDump
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.ProcessActivityNode
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.FileActivityNode
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.FileInfo
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.ProcessInfo
        proto/cws/dumpsv1/activity_dump.proto
    EOD

    sh "cp -r v5/* ."
    sh "rm -rf v5"
  end

  desc 'Run all code generators.'
  multitask :all => [:protoc]

end

task :gimme do
    go_version = "1.21.9"

    if (`which gimme`; $?.success?)
      sh "gimme #{go_version}"
      results = `. ~/.gimme/envs/go#{go_version}.env > /dev/null; echo $GOOS; echo $GOARCH; echo $GOROOT; echo $PATH`.split("\n")
      ENV['GOOS'] = results[0]
      ENV['GOARCH'] = results[1]
      ENV['GOROOT'] = results[2]
      ENV['PATH'] = results[3]
      puts "The code generated with versions other than the pinned version #{go_version} might cause conflict"
    elsif (`which go`; $?.success?)
      puts "using local go toolchain, install gimme to use pinned go version"
    else
      puts "You need either gimme or go in your path."
      exit 1
  end
end

desc "Setup dependencies"
task :deps do
  system("go mod tidy")
end

desc "Run tests"
task :test do
  cmd = "go list ./... | grep -v vendor | xargs go test -v "
  sh cmd
end

task :fuzz do
  cmd = "go list ./... | grep -v vendor | xargs -n 1 go test -v -fuzztime=20s -fuzz '.*' "
  sh cmd
end

desc "Run all code generation."
task :codegen => [:gimme, 'codegen:all']

desc "Run all protobuf code generation."
task :protobuf => [:gimme, 'codegen:protoc']

task :default => [:deps, :test, :codegen]
