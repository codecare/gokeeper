gokeeper is a command line password vault written in go

sample output
  
      please enter command or ?
      f zoom
      filtered entries for 'zoom': 1 
      0  ------------------------------------------
      name:                [zoom.us]
      description:         []
      login:               [xyz@login.name]
      passwordAlg:         [PBKDF2:16383:256:SHA512:AES:GCM]
    
      selected: zoom.us
      please enter command or ?


use at your own risk

Apache v2 License


you can install it with go get -v github.com/codecare/gokeeper/cmd/gokeeper

this will install an executable in your default path for go executables $GOPATH/bin (default GOPATH=$HOME/go)
