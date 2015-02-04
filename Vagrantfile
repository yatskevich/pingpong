# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box = "yatskevich/service-discovery"

  config.vm.network :forwarded_port, guest: 2375, host: 2375, auto_correct: true

  config.vm.provider "virtualbox" do |vb|
    vb.gui = false

    vb.customize ["modifyvm", :id, "--memory", "512"]
  end

  config.vm.synced_folder ".", "/vagrant"

  config.vm.define "node-one" do |node| 
    node.vm.network "private_network", ip: "192.168.33.10"
  end
  config.vm.define "node-two" do |node| 
    node.vm.network "private_network", ip: "192.168.33.11"
  end
  config.vm.define "node-three" do |node| 
    node.vm.network "private_network", ip: "192.168.33.12"
  end

end
