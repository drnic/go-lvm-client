VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "precise64"
  config.vm.box_url = "http://dl.dropbox.com/u/1537815/precise64.box"

  config.vm.define "lvm-0", primary: true do |v|
    v.vm.network "private_network", ip: "192.168.69.100"
    v.vm.hostname = "lvm-0"
    v.vm.provision "shell" do |sh|
      sh.path = "bin/provision"
    end
  end

end
