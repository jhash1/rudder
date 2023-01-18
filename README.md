Helm Security Scanner
A tool written in Go that scans Helm releases and finds security vulnerabilities.

Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

Prerequisites
Go (1.15+)
Helm (3+)
Kubectl (1.18+)
Installation
Clone the repository
Copy code
git clone https://github.com/YOUR-USERNAME/rudder.git
Build the binary
Copy code
cd rudder
go build
Add the binary to your PATH
Copy code
export PATH=$PATH:$PWD
Usage
To scan a Helm release, run the following command:

Copy code
rudder -n NAMESPACE
Where NAMESPACE is the namespace in which the release is installed, and RELEASE_NAME is the name of the release.

For more information, see the usage documentation.

Contributing
We welcome contributions to the Helm Security Scanner. Please see the contributing guidelines for more information.

License
This project is licensed under the Apache 2.0 License - see the LICENSE file for details.
