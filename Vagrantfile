Vagrant.configure("2") do |config|
  # operating system for the VM
  config.vm.box = "ubuntu-xenial64"

  # ssh settings
  config.ssh.username = "vagrant"
  config.ssh.private_key_path = ["~/.ssh/server_key", "~/.vagrant.d/insecure_private_key"]
  config.ssh.insert_key = false

  # vm provider
  config.vm.provider "virtualbox" do |vb|
    # Correct this error Stderr: VBoxManage.exe: error: RawFile#0 failed to create the raw output 
    vb.customize [ "modifyvm", :id, "--uartmode1", "disconnected" ]

    # Customize the number CPUS on the VM:
    vb.cpus = "1"

    # Customize the amount of memory on the VM:
    vb.memory = "1024"
  end

  #synchronize folders between host and guest machine
  config.vm.synced_folder '.', '/home/vagrant/goprojects/src'

  # upload public key into the machine
  config.vm.provision "file", source: "~/.ssh/server_key.pub", destination: "~/.ssh/authorized_keys"

  # configure the development server
  config.vm.define "development" do |development|
    development.vm.hostname = "development"
    development.vm.network "private_network", ip: "192.168.255.8"
    development.vm.provision :ansible do |ansible|
      ansible.inventory_path = "configuration/hosts"
      ansible.compatibility_mode = "2.0"
      ansible.verbose = "vvvv"
      ansible.raw_arguments  = ["--private-key=~/.ssh/server_key"]
      ansible.playbook = "configuration/server.yml"
    end
    development.vm.provision "shell", inline: <<-SHELL
      echo 'Running Serverspec Tests...'
      cd /home/vagrant/goprojects/src/configuration/serverspec
      rake -v -t
    SHELL
  end 
end
