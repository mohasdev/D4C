# D4C

Malicious version of Geth for fuzzing the devp2p protocols of ethereum.


## Getting Started


<details>
  <summary>Setup a private network with Kurtosis</summary>
  </br>
  
To start fuzzing the ethereum network you will first need a private network. If you don't already have one, follow the instructions below.

First install kurtosis and docker by following the instructions on [kurtosis docs](https://docs.kurtosis.com/install/)

Once this is done, create a "network_params.json" configuration file.
Here's an example configuration file to launch a network with two nodes: Geth/lighthouse and Besu/lighthouse 

  ### network_params.json
  ```json
  {
    "participants": [
        {
            "el_client_type": "geth",
            "el_client_image": "ethereum/client-go:latest",
            "cl_client_type": "lighthouse",
            "cl_client_image": "sigp/lighthouse:latest",
            "count": 1
        },
        {
            "el_client_type": "besu",
            "el_client_image": "hyperledger/besu:develop",
            "cl_client_type": "lighthouse",
            "cl_client_image": "sigp/lighthouse:latest",
            "count": 1
        }
    ],
    "launch_additional_services": false
}
  ```
  

You can view all the options for the configuration file (useful if you want to choose or add other implementations or activate service/monitoring tools) here :

https://github.com/kurtosis-tech/eth2-package#configuration


Run the command : 

``` 
kurtosis run --enclave myTestnet github.com/kurtosis-tech/eth2-package "$(cat ./network_params.json)"
```

You can replace "myTestnet" with the name of your choice and replace "./network_params.json" with the path and name of your configuration file.

After running the command and installation is done, your private network should be available and running in the background.

For more information and more commands check the [Kurtosis docs](https://docs.kurtosis.com/)

</details>



<details>
  <summary> Launch fuzz test with the devp2p cli</summary>
  
</br>

If you don't already have GO installed, go to [GO website](https://go.dev/doc/install) and follow the installation instructions.


- ### Build D4C on Linux and Mac

Start by cloning the repo and go to the root of the project and run the command :

```
make all
```

- ### Build D4C on Windows

Coming soon...

</br>

Once you have built the project you can now run fuzz tests by going to the root of the project and running devp2p binaries followed by the appropriate command:


```
./build/bin/devp2p 
```

You can find out more about the commands available to launch different fuzz tests, as well as the options, in the Command list section and explanations of the fuzzers in the Fuzzers section. 

</details> 
  





## Command list :


<details>
  <summary>Wrong Version field ping</summary>
  </br>
  - Usage : Sends ping to a node with a wrong version field
  
  ```
  ./build/bin/devp2p discv4 wrong-version-ping <node> <fuzzer-name> <run> <string-to-mutate>
  ```
  Args information :
  
  - node : enode adress of the node you want to ping
  - fuzzer-name : name of the fuzzer you want to use
  - run : indicate the number of fuzz test you want to run
  - string-to-mutate : required if you want to use the mutation-fuzzer, enter the string that you want to mutate.
  
  Available fuzzers :
  - random-fuzzer
  - mutation-fuzzer
  - string-fuzzer
  
</details>

<details>
  <summary>Wrong To field ping</summary>
  </br>
  - Usage : Sends ping to a node with a wrong To field
  
  ```
  ./build/bin/devp2p discv4 wrong-to-ping <node> <fuzzer-name> <run> <string-to-mutate>
  ```
  Args information :
  
  - node : enode adress of the node you want to ping
  - fuzzer-name : name of the fuzzer you want to use
  - run : indicate the number of fuzz test you want to run
  - string-to-mutate : required if you want to use the mutation-fuzzer, enter the string that you want to mutate.
  
  Available fuzzers :
  - random-fuzzer
  - mutation-fuzzer
  - string-fuzzer
</details>


<details>
  <summary>Wrong From field ping</summary>
  </br>
  - Usage : Sends ping to a node with a wrong From field
  
  ```
  ./build/bin/devp2p discv4 wrong-from-ping <node> <fuzzer-name> <run> <string-to-mutate>
  ```
  Args information :
  
  - node : enode adress of the node you want to ping
  - fuzzer-name : name of the fuzzer you want to use
  - run : indicate the number of fuzz test you want to run
  - string-to-mutate : required if you want to use the mutation-fuzzer, enter the string that you want to mutate.
  
  Available fuzzers :
  - random-fuzzer
  - mutation-fuzzer
  - string-fuzzer
</details>

<details>
  <summary>Extra Data ping</summary>
  </br>
  - Usage : Sends ping to a node with fuzzed extra data
  
  ```
  ./build/bin/devp2p discv4 extra-data-ping <node> <fuzzer-name> <run> <string-to-mutate>
  ```
  Args information :
  
  - node : enode adress of the node you want to ping
  - fuzzer-name : name of the fuzzer you want to use
  - run : indicate the number of fuzz test you want to run
  - string-to-mutate : required if you want to use the mutation-fuzzer, enter the string that you want to mutate.
  
  Available fuzzers :
  - random-fuzzer
  - mutation-fuzzer
  - string-fuzzer
  
  Notes :
  - The command send two extra data fields by default, the choice of the number of additional fields will be added soon.
</details>

<details>
  <summary>Wrong From field and Extra Data ping</summary>
  </br>
  - Usage : Sends ping to a node with fuzzed extra data and a wrong From field
  
  ```
  ./build/bin/devp2p discv4 wrong-from-extra-data-ping <node> <fuzzer-name> <run> <string-to-mutate>
  ```
  Args information :
  
  - node : enode adress of the node you want to ping
  - fuzzer-name : name of the fuzzer you want to use
  - run : indicate the number of fuzz test you want to run
  - string-to-mutate : required if you want to use the mutation-fuzzer, enter the string that you want to mutate.
  
  Available fuzzers :
  - random-fuzzer
  - mutation-fuzzer
  - string-fuzzer
  
  Notes :
  - The command send two extra data fields by default, the choice of the number of additional fields will be added soon.
</details>

<details>
  <summary>Rlpx Wrong Auth Version ping</summary>
  </br>
  - Usage : Sends a rlpx ping to a node with a auth wrong version field
  
  ```
  ./build/bin/devp2p rlpx wrong-version-ping <node> <fuzzer-name> <run> <string-to-mutate>
  ```
  Args information :
  
  - node : enode adress of the node you want to ping
  - fuzzer-name : name of the fuzzer you want to use
  - run : indicate the number of fuzz test you want to run
  - string-to-mutate : required if you want to use the mutation-fuzzer, enter the string that you want to mutate.
  
  Available fuzzers :
  - random-fuzzer
  - mutation-fuzzer
  - string-fuzzer
  
</details>

<details>
  <summary>Rlpx Wrong Resp Version ping</summary>
  </br>
  - Usage : Sends a rlpx ping to a node with a resp wrong version field
  
  ```
  ./build/bin/devp2p rlpx wrong-resp-version-ping <node> <fuzzer-name> <run> <string-to-mutate>
  ```
  Args information :
  
  - node : enode adress of the node you want to ping
  - fuzzer-name : name of the fuzzer you want to use
  - run : indicate the number of fuzz test you want to run
  - string-to-mutate : required if you want to use the mutation-fuzzer, enter the string that you want to mutate.
  
  Available fuzzers :
  - random-fuzzer
  - mutation-fuzzer
  - string-fuzzer
  
</details>


## Fuzzers


### Random fuzzer 
Generate a random string, default values are :
  - MinLength: 10,
  - MaxLength: 20,
  - CharStart: 32,
  - CharRange: 32

### Mutation fuzzer
Mutate a string with three different mutations :

- By inserting a random character
- By deleting a random character
- By flipping a random character

The mutations are choosed randomly, default values are :

- MinMutation: 2,
- MaxMutation: 10

### String fuzzer
Mutate a string with 15 different mutations

The mutations are choosed randomly, default values are :

- MinMutation: 2,
- MaxMutation: 10

### Number fuzzer
Mutate a number with 10 different mutations

The mutations are choosed randomly, default values are :

- MinMutation: 2,
- MaxMutation: 10

### Address fuzzer
Mutate a IPv4 adress with 5 different mutations

The mutations are choosed randomly, default values are :

- MinMutation: 2,
- MaxMutation: 10



### Changing defaults values with the CLI is not available for the moment